package dameng

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/migrator"
	"gorm.io/gorm/schema"
)

var (
	regFullDataType = regexp.MustCompile(`\D*(\d+)\D?`)
)

type Migrator struct {
	migrator.Migrator
	Dialector
}

// AutoMigrate 自动迁移模型为表结构
//
//	// 迁移并设置单个表注释
//	db.Set("gorm:table_comments", "用户信息表").AutoMigrate(&User{})
//
//	// 迁移并设置多个表注释
//	db.Set("gorm:table_comments", []string{"用户信息表", "公司信息表"}).AutoMigrate(&User{}, &Company{})
func (m Migrator) AutoMigrate(dst ...interface{}) error {
	if err := m.Migrator.AutoMigrate(dst...); err != nil {
		return err
	}
	if tableComments, ok := m.DB.Get("gorm:table_comments"); ok {
		var comments []string
		switch c := tableComments.(type) {
		case string:
			comments = []string{c}
		case []string:
			comments = c
		default:
			return nil
		}
		for i := 0; i < len(dst) && i < len(comments); i++ {
			value := dst[i]
			tx := m.DB.Session(&gorm.Session{})
			comment := strings.ReplaceAll(comments[i], "'", "''")
			if err := m.RunWithValue(value, func(stmt *gorm.Statement) error {
				return tx.Exec(fmt.Sprintf("COMMENT ON TABLE ? IS '%s'", comment), m.CurrentTable(stmt)).Error
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m Migrator) CurrentDatabase() (name string) {
	_ = m.DB.Raw("SELECT SYS_CONTEXT('USERENV', 'CURRENT_SCHEMA');").Row().Scan(&name)
	return
}

func (m Migrator) FullDataTypeOf(field *schema.Field) clause.Expr {
	expr := m.Migrator.FullDataTypeOf(field)

	// https://eco.dameng.com/community/question/2966fa8cdb97f9444dd80afb78d7c9a6
	//if value, ok := field.TagSettings["COMMENT"]; ok {
	//	expr.SQL += " COMMENT " + m.Dialector.Explain("?", value)
	//}

	return expr
}

func (m Migrator) GetTypeAliases(databaseTypeName string) []string {
	// super
	return m.Migrator.GetTypeAliases(databaseTypeName)
}

func (m Migrator) CreateTable(values ...interface{}) error {
	// 将`gorm:"default:true"`转为`gorm:"default:1"`
	for _, value := range values {
		if err := m.RunWithValue(value, func(stmt *gorm.Statement) error {
			for _, v := range stmt.Schema.Fields {
				if v.HasDefaultValue {
					if vv, ok := v.DefaultValueInterface.(bool); ok {
						if vv {
							v.DefaultValueInterface = int64(1)
						} else {
							v.DefaultValueInterface = int64(0)
						}
					}
				}
			}
			return nil
		}); err != nil {
			return err
		}
	}
	// super
	return m.Migrator.CreateTable(values...)
}

//goland:noinspection SqlNoDataSourceInspection
func (m Migrator) DropTable(values ...interface{}) error {
	values = m.ReorderModels(values, false)
	for i := len(values) - 1; i >= 0; i-- {
		tx := m.DB.Session(&gorm.Session{})
		if err := m.RunWithValue(values[i], func(stmt *gorm.Statement) error {
			return tx.Exec("DROP TABLE IF EXISTS ? CASCADE", m.CurrentTable(stmt)).Error
		}); err != nil {
			return err
		}
	}
	return nil
}

func (m Migrator) HasTable(value interface{}) bool {
	tableSql := `SELECT /*+ MAX_OPT_N_TABLES(5) */ COUNT(TABS.NAME) FROM
(SELECT ID, PID FROM SYS.SYSOBJECTS WHERE TYPE$ = 'SCH' AND NAME = ?) SCHEMAS,
(SELECT ID, SCHID, NAME FROM SYS.SYSOBJECTS WHERE
NAME = ? AND TYPE$ = 'SCHOBJ' AND SUBTYPE$ IN ('UTAB', 'STAB', 'VIEW', 'SYNOM')
AND ((SUBTYPE$ ='UTAB' AND CAST((INFO3 & 0x00FF & 0x003F) AS INT) not in (9, 27, 29, 25, 12, 7, 21, 23, 18, 5))
OR SUBTYPE$ in ('STAB', 'VIEW', 'SYNOM'))) TABS
WHERE TABS.SCHID = SCHEMAS.ID AND SF_CHECK_PRIV_OPT(UID(), CURRENT_USERTYPE(), TABS.ID, SCHEMAS.PID, -1, TABS.ID) = 1;`

	var count int64
	_ = m.RunWithValue(value, func(stmt *gorm.Statement) error {
		return m.DB.Raw(tableSql, m.CurrentDatabase(), stmt.Table).Row().Scan(&count)
	})
	return count > 0
}

func (m Migrator) RenameTable(oldName, newName interface{}) error {
	// super
	return m.Migrator.RenameTable(oldName, newName)
}

func (m Migrator) GetTables() (tableList []string, err error) {
	tableSql := `SELECT /*+ MAX_OPT_N_TABLES(5) */ TABS.NAME FROM
(SELECT ID, PID FROM SYS.SYSOBJECTS WHERE TYPE$ = 'SCH' AND NAME = ?) SCHEMAS,
(SELECT ID, SCHID, NAME FROM SYS.SYSOBJECTS WHERE TYPE$ = 'SCHOBJ' AND SUBTYPE$ IN ('UTAB', 'STAB', 'VIEW', 'SYNOM')
AND ((SUBTYPE$ ='UTAB' AND CAST((INFO3 & 0x00FF & 0x003F) AS INT) not in (9, 27, 29, 25, 12, 7, 21, 23, 18, 5))
OR SUBTYPE$ in ('STAB', 'VIEW', 'SYNOM'))) TABS
WHERE TABS.SCHID = SCHEMAS.ID AND SF_CHECK_PRIV_OPT(UID(), CURRENT_USERTYPE(), TABS.ID, SCHEMAS.PID, -1, TABS.ID) = 1;`

	err = m.DB.Raw(tableSql, m.CurrentDatabase()).Scan(&tableList).Error
	return
}

func (m Migrator) AddColumn(dst interface{}, field string) error {
	// super
	return m.Migrator.AddColumn(dst, field)
}

func (m Migrator) DropColumn(dst interface{}, field string) error {
	// super
	return m.Migrator.DropColumn(dst, field)
}

//goland:noinspection SqlNoDataSourceInspection
func (m Migrator) AlterColumn(value interface{}, field string) error {
	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		if field := stmt.Schema.LookUpField(field); field != nil {
			return m.DB.Exec(
				"ALTER TABLE ? MODIFY ? ?",
				clause.Table{Name: stmt.Table},
				clause.Column{Name: field.DBName},
				m.FullDataTypeOf(field),
			).Error
		}
		return fmt.Errorf("failed to look up field with name: %s", field)
	})
}

// containsUnique: 原本是否就有UNIQUE字段
func (m Migrator) alterColumn(value interface{}, field string, containsUnique bool) error {
	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		if field := stmt.Schema.LookUpField(field); field != nil {
			typeof := m.FullDataTypeOf(field)
			// 如果列原本就有UNIQUE，且修改后仍有UNIQUE，则在MODIFY COLUMN时不再添加UNIQUE字段
			// 这样也不会影响 有UNIQUE -> 无UNIQUE、无UNIQUE -> 有UNIQUE、无UNIQUE -> 无UNIQUE的情况
			if containsUnique && field.Unique {
				typeof.SQL = strings.Replace(typeof.SQL, " UNIQUE", "", 1)
			}
			return m.DB.Exec(
				"ALTER TABLE ? MODIFY ? ?",
				clause.Table{Name: stmt.Table},
				clause.Column{Name: field.DBName},
				typeof,
			).Error
		}
		return fmt.Errorf("failed to look up field with name: %s", field)
	})
}

func (m Migrator) MigrateColumn(dst interface{}, field *schema.Field, columnType gorm.ColumnType) error {
	// super
	// return m.Migrator.MigrateColumn(dst, field, columnType)
	// bug629968 不再使用父类默认的MigrateColumn函数，主要修改：
	// 添加了containsUnique参数和最后的调用从AlterColumn改为alterColumn

	// found, smart migrate
	fullDataType := strings.TrimSpace(strings.ToLower(m.DB.Migrator().FullDataTypeOf(field).SQL))
	realDataType := strings.ToLower(columnType.DatabaseTypeName())
	containsUnique := false
	if unique, ok := columnType.Unique(); ok {
		containsUnique = unique
	}

	var (
		alterColumn bool
		isSameType  = fullDataType == realDataType
	)

	if !field.PrimaryKey {
		// check type
		if !strings.HasPrefix(fullDataType, realDataType) {
			// check type aliases
			aliases := m.DB.Migrator().GetTypeAliases(realDataType)
			for _, alias := range aliases {
				if strings.HasPrefix(fullDataType, alias) {
					isSameType = true
					break
				}
			}

			if !isSameType {
				alterColumn = true
			}
		}
	}

	if !isSameType {
		// check size
		if length, ok := columnType.Length(); length != int64(field.Size) {
			if length > 0 && field.Size > 0 {
				alterColumn = true
			} else {
				// has size in data type and not equal
				// Since the following code is frequently called in the for loop, reg optimization is needed here
				matches2 := regFullDataType.FindAllStringSubmatch(fullDataType, -1)
				if !field.PrimaryKey &&
					(len(matches2) == 1 && matches2[0][1] != fmt.Sprint(length) && ok) {
					alterColumn = true
				}
			}
		}

		// check precision
		if precision, _, ok := columnType.DecimalSize(); ok && int64(field.Precision) != precision {
			if regexp.MustCompile(fmt.Sprintf("[^0-9]%d[^0-9]", field.Precision)).MatchString(m.Migrator.DataTypeOf(field)) {
				alterColumn = true
			}
		}
	}

	// check nullable
	if nullable, ok := columnType.Nullable(); ok && nullable == field.NotNull {
		// not primary key & database is nullable
		if !field.PrimaryKey && nullable {
			alterColumn = true
		}
	}

	// check unique
	if unique, ok := columnType.Unique(); ok && unique != field.Unique {
		// not primary key
		if !field.PrimaryKey {
			alterColumn = true
		}
	}

	// check default value
	if !field.PrimaryKey {
		currentDefaultNotNull := field.HasDefaultValue && (field.DefaultValueInterface != nil || !strings.EqualFold(field.DefaultValue, "NULL"))
		dv, dvNotNull := columnType.DefaultValue()
		if dvNotNull && !currentDefaultNotNull {
			// defalut value -> null
			alterColumn = true
		} else if !dvNotNull && currentDefaultNotNull {
			// null -> default value
			alterColumn = true
		} else if (field.GORMDataType != schema.Time && dv != field.DefaultValue) ||
			(field.GORMDataType == schema.Time && !strings.EqualFold(strings.TrimSuffix(dv, "()"), strings.TrimSuffix(field.DefaultValue, "()"))) {
			// default value not equal
			// not both null
			if currentDefaultNotNull || dvNotNull {
				alterColumn = true
			}
		}
	}

	// check comment
	if comment, ok := columnType.Comment(); ok && comment != field.Comment {
		// not primary key
		if !field.PrimaryKey {
			alterColumn = true
		}
	}

	if alterColumn && !field.IgnoreMigration {
		return m.DB.Migrator().(Migrator).alterColumn(dst, field.DBName, containsUnique)
	}

	return nil
}

func (m Migrator) HasColumn(value interface{}, field string) bool {
	columnSql := `SELECT /*+ MAX_OPT_N_TABLES(5) */ COUNT(DISTINCT COLS.NAME) FROM
(SELECT ID FROM SYS.SYSOBJECTS WHERE TYPE$ = 'SCH' AND NAME = ?) SCHS,
(SELECT ID, SCHID FROM SYS.SYSOBJECTS WHERE TYPE$ = 'SCHOBJ' AND SUBTYPE$ IN ('UTAB', 'STAB', 'VIEW') AND NAME = ?) TABS,
(SELECT NAME, ID FROM SYS.SYSCOLUMNS WHERE NAME = ?) COLS
WHERE TABS.ID = COLS.ID AND SCHS.ID = TABS.SCHID;`

	var count int64
	_ = m.RunWithValue(value, func(stmt *gorm.Statement) error {
		return m.DB.Raw(columnSql, m.CurrentDatabase(), stmt.Table, field).Row().Scan(&count)
	})
	return count > 0
}

func (m Migrator) RenameColumn(value interface{}, oldName, newName string) error {
	// super
	return m.Migrator.RenameColumn(value, oldName, newName)
}

func (m Migrator) ColumnTypes(dst interface{}) ([]gorm.ColumnType, error) {
	columnTypes := make([]gorm.ColumnType, 0)
	execErr := m.RunWithValue(dst, func(stmt *gorm.Statement) error {
		var (
			currentDatabase = m.CurrentDatabase()
			table           = stmt.Table
			columnTypeSQL   = `SELECT /*+ MAX_OPT_N_TABLES(5) */ COLS.NAME, COLS.DEFVAL FROM
(SELECT ID FROM SYS.SYSOBJECTS WHERE TYPE$ = 'SCH' AND NAME = ?) SCHS,
(SELECT ID, SCHID FROM SYS.SYSOBJECTS WHERE TYPE$ = 'SCHOBJ' AND SUBTYPE$ IN ('UTAB', 'STAB', 'VIEW') AND NAME = ?) TABS,
SYS.SYSCOLUMNS COLS
WHERE TABS.ID=COLS.ID AND SCHS.ID = TABS.SCHID`
			columnConsSQL = `SELECT /*+ MAX_OPT_N_TABLES(5) */ COLS.NAME, LNNVL(CONS.TYPE$!='P'), LNNVL(CONS.TYPE$!='U') FROM
(SELECT ID FROM SYS.SYSOBJECTS WHERE TYPE$ = 'SCH' AND NAME = ?) SCHS,
(SELECT ID, SCHID FROM SYS.SYSOBJECTS WHERE TYPE$ = 'SCHOBJ' AND SUBTYPE$ IN ('UTAB', 'STAB', 'VIEW') AND NAME = ?) TABS,
SYS.SYSCOLUMNS COLS,
SYS.SYSCONS CONS,
SYS.SYSINDEXES INDS
WHERE SCHS.ID=TABS.SCHID AND TABS.ID=COLS.ID AND COLS.ID=CONS.TABLEID and CONS.INDEXID=INDS.ID and SF_COL_IS_IDX_KEY(INDS.KEYNUM, INDS.KEYINFO, COLS.COLID)=1`
			consMap   = make(map[string][][]bool)
			rows, err = m.DB.Session(&gorm.Session{}).Table(stmt.Table).Limit(1).Rows()
		)

		if err != nil {
			return err
		}

		rawColumnTypes, err := rows.ColumnTypes()

		if err := rows.Close(); err != nil {
			return err
		}

		// 约束
		cons, consErr := m.DB.Table(table).Raw(columnConsSQL, currentDatabase, table).Rows()
		if consErr != nil {
			return consErr
		}
		defer func(cons *sql.Rows) {
			_ = cons.Close()
		}(cons)

		for cons.Next() {
			var (
				colName string
				values  = make([]bool, 2)
			)
			if scanErr := cons.Scan(&colName, &values[0], &values[1]); scanErr != nil {
				return scanErr
			}

			if consMap[colName] == nil {
				consMap[colName] = *new([][]bool)
			}
			consMap[colName] = append(consMap[colName], values)
		}

		// 列信息
		columns, rowErr := m.DB.Table(table).Raw(columnTypeSQL, currentDatabase, table).Rows()
		if rowErr != nil {
			return rowErr
		}
		defer func(columns *sql.Rows) {
			_ = columns.Close()
		}(columns)

		for columns.Next() {
			var (
				column migrator.ColumnType
				values = []interface{}{
					&column.NameValue, &column.DefaultValueValue,
				}
			)

			if scanErr := columns.Scan(values...); scanErr != nil {
				return scanErr
			}

			column.DefaultValueValue.String = strings.Trim(column.DefaultValueValue.String, "'")

			// 设置列的主键和值唯一信息
			for key, value := range consMap {
				if key == column.NameValue.String {
					for _, con := range value {
						if con[0] {
							column.PrimaryKeyValue.Bool = true
							column.PrimaryKeyValue.Valid = true
						}
						if con[1] {
							column.UniqueValue.Bool = true
							column.UniqueValue.Valid = true
						}
					}
					break
				}
			}

			// 设置默认的列信息（来自go驱动标准类型sql.ColumnType）
			for _, c := range rawColumnTypes {
				if c.Name() == column.NameValue.String {
					column.SQLColumnType = c
					break
				}
			}

			columnTypes = append(columnTypes, column)
		}

		return nil
	})

	return columnTypes, execErr
}

func (m Migrator) CreateView(name string, option gorm.ViewOption) error {
	// super, not support
	return m.Migrator.CreateView(name, option)
}

func (m Migrator) DropView(name string) error {
	// super, not support
	return m.Migrator.DropView(name)
}

func (m Migrator) CreateConstraint(dst interface{}, name string) error {
	// super
	return m.Migrator.CreateConstraint(dst, name)
}

func (m Migrator) DropConstraint(value interface{}, name string) error {
	// super
	return m.Migrator.DropConstraint(value, name)
}

func (m Migrator) HasConstraint(value interface{}, name string) bool {
	conSql := `select count(CON_OBJ.NAME) from
(select ID from SYSOBJECTS where TYPE$='SCH' and NAME = ?) SCH_OBJ, 
(select ID, SCHID from SYSOBJECTS where TYPE$='SCHOBJ' and SUBTYPE$ like '_TAB') TAB_OBJ, 
(select ID, NAME from SYSOBJECTS where SUBTYPE$ = 'CONS' and NAME=?) CON_OBJ,
SYSCONS CONS
where CON_OBJ.ID=CONS.ID and TAB_OBJ.ID=CONS.TABLEID and TAB_OBJ.SCHID=SCH_OBJ.ID;`

	var count int64
	_ = m.RunWithValue(value, func(stmt *gorm.Statement) error {
		constraint, _ := m.GuessConstraintInterfaceAndTable(stmt, name)
		switch c := constraint.(type) {
		case *schema.Constraint:
			name = c.Name
		case *schema.CheckConstraint:
			name = c.Name
		default:
		}
		return m.DB.Raw(conSql, m.CurrentDatabase(), name).Row().Scan(&count)
	})
	return count > 0
}

func (m Migrator) CreateIndex(dst interface{}, name string) error {
	// super
	return m.Migrator.CreateIndex(dst, name)
}

func (m Migrator) DropIndex(value interface{}, name string) error {
	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		if idx := stmt.Schema.LookIndex(name); idx != nil {
			name = idx.Name
		}

		return m.DB.Exec("DROP INDEX ?", clause.Column{Name: name}).Error
	})
}

func (m Migrator) HasIndex(value interface{}, name string) bool {
	indexSql := `WITH USERS(ID) AS (SELECT ID FROM SYS.SYSOBJECTS WHERE TYPE$ = 'SCH' AND NAME = ?),
TAB(ID,SCHID) AS (SELECT ID, SCHID FROM SYS.SYSOBJECTS WHERE TYPE$ = 'SCHOBJ' AND SUBTYPE$ = 'UTAB' AND NAME = ?)
SELECT COUNT(DISTINCT INDEX_NAME) FROM (
SELECT /*+ MAX_OPT_N_TABLES(5) */ OBJ_INDS.NAME AS INDEX_NAME FROM USERS, TAB, SYS.SYSINDEXES AS INDS, SYS.SYSCOLUMNS AS COLS,
(SELECT ID, PID, NAME FROM SYS.SYSOBJECTS WHERE SUBTYPE$='INDEX' AND NAME = ?) OBJ_INDS
WHERE TAB.ID =COLS.ID AND TAB.ID =OBJ_INDS.PID AND INDS.ID=OBJ_INDS.ID AND TAB.SCHID= USERS.ID
AND SF_COL_IS_IDX_KEY(INDS.KEYNUM, INDS.KEYINFO, COLS.COLID)=1
UNION SELECT OBJ_INDS.NAME AS INDEX_NAME FROM USERS, TAB, SYSCONTEXTINDEXES AS OBJ_INDS, SYS.SYSCOLUMNS AS COLS
WHERE TAB.ID = COLS.ID AND TAB.ID = OBJ_INDS.TABLEID AND COLS.COLID = OBJ_INDS.COLID AND TAB.SCHID = USERS.ID AND OBJ_INDS.NAME = ?)`

	var count int64
	_ = m.RunWithValue(value, func(stmt *gorm.Statement) error {
		if idx := stmt.Schema.LookIndex(name); idx != nil {
			name = idx.Name
		}
		return m.DB.Raw(indexSql, m.CurrentDatabase(), stmt.Schema.Table, name, name).Row().Scan(&count)
	})
	return count > 0
}

func (m Migrator) RenameIndex(value interface{}, oldName, newName string) error {
	return m.RunWithValue(value, func(stmt *gorm.Statement) error {
		return m.DB.Exec(
			"ALTER INDEX ? RENAME TO ?",
			clause.Column{Name: oldName}, clause.Column{Name: newName},
		).Error
	})
}

func (m Migrator) GetIndexes(value interface{}) ([]gorm.Index, error) {
	indexSql := `WITH USERS(ID) AS (SELECT ID FROM SYS.SYSOBJECTS WHERE TYPE$ = 'SCH' AND NAME = ?),
TAB(ID,SCHID,NAME) AS (SELECT ID, SCHID, NAME FROM SYS.SYSOBJECTS WHERE TYPE$ = 'SCHOBJ' AND SUBTYPE$ = 'UTAB' AND NAME = ?)
SELECT /*+ MAX_OPT_N_TABLES(5) */ TAB.NAME AS TABLE_NAME, COLS.NAME AS COLUMN_NAME,
OBJ_INDS.NAME AS INDEX_NAME, CASE INDS.ISUNIQUE WHEN 'Y' THEN 0 ELSE 1 END AS NON_UNIQUE,
CASE OBJ_INDS.TYPE$ WHEN 'P' THEN 1 ELSE 0 END AS IS_PRIMARY FROM USERS, TAB, SYS.SYSINDEXES AS INDS, SYS.SYSCOLUMNS AS COLS,
(SELECT INDS.ID, INDS.PID, INDS.NAME, CONS.TYPE$ FROM SYS.SYSOBJECTS AS INDS LEFT JOIN SYS.SYSCONS AS CONS ON CONS.INDEXID=INDS.ID AND SUBTYPE$='INDEX') OBJ_INDS
WHERE TAB.ID =COLS.ID AND TAB.ID =OBJ_INDS.PID AND INDS.ID=OBJ_INDS.ID AND TAB.SCHID=USERS.ID AND SF_COL_IS_IDX_KEY(INDS.KEYNUM, INDS.KEYINFO, COLS.COLID)=1
UNION SELECT TAB.NAME AS TABLE_NAME, COLS.NAME AS COLUMN_NAME, OBJ_INDS.NAME AS INDEX_NAME, 1 AS NON_UNIQUE, 0 AS IS_PRIMARY FROM
USERS, TAB, SYSCONTEXTINDEXES AS OBJ_INDS, SYS.SYSCOLUMNS AS COLS WHERE
TAB.ID = COLS.ID AND TAB.ID = OBJ_INDS.TABLEID AND COLS.COLID = OBJ_INDS.COLID AND TAB.SCHID = USERS.ID;`

	indexes := make([]gorm.Index, 0)
	err := m.RunWithValue(value, func(stmt *gorm.Statement) error {
		result := make([]*Index, 0)
		if scanErr := m.DB.Raw(indexSql, m.CurrentDatabase(), stmt.Table).Scan(&result).Error; scanErr != nil {
			return scanErr
		}
		indexMap := groupByIndexName(result)
		for _, idx := range indexMap {
			tempIdx := &migrator.Index{
				TableName: idx[0].TableName,
				NameValue: idx[0].IndexName,
				PrimaryKeyValue: sql.NullBool{
					Bool:  idx[0].Primary,
					Valid: true,
				},
				UniqueValue: sql.NullBool{
					Bool:  idx[0].NonUnique,
					Valid: true,
				},
			}
			for _, x := range idx {
				tempIdx.ColumnList = append(tempIdx.ColumnList, x.ColumnName)
			}
			indexes = append(indexes, tempIdx)
		}
		return nil
	})
	return indexes, err
}

type Index struct {
	TableName  string `gorm:"column:TABLE_NAME"`
	ColumnName string `gorm:"column:COLUMN_NAME"`
	IndexName  string `gorm:"column:INDEX_NAME"`
	NonUnique  bool   `gorm:"column:NON_UNIQUE"`
	Primary    bool   `gorm:"column:IS_PRIMARY"`
}

func groupByIndexName(indexList []*Index) map[string][]*Index {
	columnIndexMap := make(map[string][]*Index, len(indexList))
	for _, idx := range indexList {
		columnIndexMap[idx.IndexName] = append(columnIndexMap[idx.IndexName], idx)
	}
	return columnIndexMap
}
