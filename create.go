package dameng

import (
	"database/sql"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

func Create(db *gorm.DB) {
	if db.Error != nil {
		return
	}

	if db.Statement.Schema != nil && !db.Statement.Unscoped {
		for _, c := range db.Statement.Schema.CreateClauses {
			db.Statement.AddClause(c)
		}
	}

	var onConflict clause.OnConflict
	var hasConflict bool
	if db.Statement.SQL.String() == "" {
		var (
			values = callbacks.ConvertToCreateValues(db.Statement)
			c      = db.Statement.Clauses["ON CONFLICT"]
		)
		onConflict, hasConflict = c.Expression.(clause.OnConflict)

		if hasConflict {
			if len(db.Statement.Schema.PrimaryFields) > 0 {
				columnsMap := map[string]bool{}
				for _, column := range values.Columns {
					columnsMap[column.Name] = true
				}

				for _, field := range db.Statement.Schema.PrimaryFields {
					if _, ok := columnsMap[field.DBName]; !ok {
						hasConflict = false
					}
				}
			} else {
				hasConflict = false
			}
		}

		if hasConflict {
			MergeCreate(db, onConflict, values)
		} else {
			setIdentityInsert := false

			if db.Statement.Schema != nil {
				if field := db.Statement.Schema.PrioritizedPrimaryField; field != nil && field.AutoIncrement {
					switch db.Statement.ReflectValue.Kind() {
					case reflect.Struct:
						_, isZero := field.ValueOf(db.Statement.Context, db.Statement.ReflectValue)
						setIdentityInsert = !isZero
					case reflect.Slice, reflect.Array:
						for i := 0; i < db.Statement.ReflectValue.Len(); {
							obj := db.Statement.ReflectValue.Index(i)
							if reflect.Indirect(obj).Kind() == reflect.Struct {
								_, isZero := field.ValueOf(db.Statement.Context, db.Statement.ReflectValue.Index(i))
								setIdentityInsert = !isZero
							}
							break
						}
					default:
					}

					if setIdentityInsert && !db.DryRun && db.Error == nil {
						db.Statement.SQL.Reset()
						_, _ = db.Statement.WriteString("SET IDENTITY_INSERT ")
						db.Statement.WriteQuoted(db.Statement.Table)
						_, _ = db.Statement.WriteString(" ON;")
						_, err := db.Statement.ConnPool.ExecContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
						if db.AddError(err) != nil {
							return
						}
						defer func() {
							db.Statement.SQL.Reset()
							_, _ = db.Statement.WriteString("SET IDENTITY_INSERT ")
							db.Statement.WriteQuoted(db.Statement.Table)
							_, _ = db.Statement.WriteString(" OFF;")
							_, _ = db.Statement.ConnPool.ExecContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
						}()
					}
				}
			}

			db.Statement.SQL.Reset()
			db.Statement.AddClauseIfNotExists(clause.Insert{})
			db.Statement.Build("INSERT")
			_ = db.Statement.WriteByte(' ')

			db.Statement.AddClause(values)
			if values, ok := db.Statement.Clauses["VALUES"].Expression.(clause.Values); ok {
				if len(values.Columns) > 0 {
					_ = db.Statement.WriteByte('(')
					for idx, column := range values.Columns {
						if idx > 0 {
							_ = db.Statement.WriteByte(',')
						}
						db.Statement.WriteQuoted(column)
					}
					_ = db.Statement.WriteByte(')')

					//outputInserted(db)

					_, _ = db.Statement.WriteString(" VALUES ")

					for idx, value := range values.Values {
						if idx > 0 {
							_ = db.Statement.WriteByte(',')
						}

						_ = db.Statement.WriteByte('(')
						db.Statement.AddVar(db.Statement, value...)
						_ = db.Statement.WriteByte(')')
					}

					_, _ = db.Statement.WriteString(";")
				} else {
					_, _ = db.Statement.WriteString("DEFAULT VALUES;")
				}
			}
		}
	}

	if !db.DryRun && db.Error == nil {
		var (
			rows           *sql.Rows
			result         sql.Result
			err            error
			updateInsertID bool  // 是否需要更新主键自增列
			insertID       int64 // 主键自增列最新值
		)
		if hasConflict {
			rows, err = db.Statement.ConnPool.QueryContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
			if db.AddError(err) != nil {
				return
			}
			defer func(rows *sql.Rows) {
				_ = rows.Close()
			}(rows)
			if rows.Next() {
				_ = rows.Scan(&insertID)
				if insertID > 0 {
					updateInsertID = true
				}
			}
		} else {
			result, err = db.Statement.ConnPool.ExecContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
			if db.AddError(err) != nil {
				return
			}
			db.RowsAffected, _ = result.RowsAffected()
			if db.RowsAffected != 0 && db.Statement.Schema != nil &&
				db.Statement.Schema.PrioritizedPrimaryField != nil &&
				db.Statement.Schema.PrioritizedPrimaryField.HasDefaultValue {
				insertID, err = result.LastInsertId()
				insertOk := err == nil && insertID > 0
				if !insertOk {
					_ = db.AddError(err)
					return
				}
				updateInsertID = true
			}
		}

		if !updateInsertID {
			return
		}
		// map insert support return increment id
		// https://github.com/go-gorm/gorm/pull/6662
		var pkFieldName = "@id"
		if db.Statement.Schema != nil {
			if db.Statement.Schema.PrioritizedPrimaryField == nil || !db.Statement.Schema.PrioritizedPrimaryField.HasDefaultValue {
				return
			}
			pkFieldName = db.Statement.Schema.PrioritizedPrimaryField.DBName
		}
		// append @id column with value for auto-increment primary key
		// the @id value is correct, when: 1. without setting auto-increment primary key, 2. database AutoIncrementIncrement = 1
		switch values := db.Statement.Dest.(type) {
		case map[string]interface{}:
			values[pkFieldName] = insertID
		case *map[string]interface{}:
			(*values)[pkFieldName] = insertID
		case []map[string]interface{}, *[]map[string]interface{}:
			mapValues, ok := values.([]map[string]interface{})
			if !ok {
				if v, ok := values.(*[]map[string]interface{}); ok {
					if *v != nil {
						mapValues = *v
					}
				}
			}
			// if config.LastInsertIDReversed {
			insertID -= int64(len(mapValues)-1) * schema.DefaultAutoIncrementIncrement
			// }
			for _, mapValue := range mapValues {
				if mapValue != nil {
					mapValue[pkFieldName] = insertID
				}
				insertID += schema.DefaultAutoIncrementIncrement
			}
		default:
			switch db.Statement.ReflectValue.Kind() {
			case reflect.Slice, reflect.Array:
				//if config.LastInsertIDReversed {
				for i := db.Statement.ReflectValue.Len() - 1; i >= 0; i-- {
					rv := db.Statement.ReflectValue.Index(i)
					if reflect.Indirect(rv).Kind() != reflect.Struct {
						break
					}

					_, isZero := db.Statement.Schema.PrioritizedPrimaryField.ValueOf(db.Statement.Context, rv)
					if isZero {
						_ = db.AddError(db.Statement.Schema.PrioritizedPrimaryField.Set(db.Statement.Context, rv, insertID))
						insertID -= db.Statement.Schema.PrioritizedPrimaryField.AutoIncrementIncrement
					}
				}
				//} else {
				//	for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
				//		rv := db.Statement.ReflectValue.Index(i)
				//		if reflect.Indirect(rv).Kind() != reflect.Struct {
				//			break
				//		}
				//
				//		if _, isZero := db.Statement.Schema.PrioritizedPrimaryField.ValueOf(db.Statement.Context, rv); isZero {
				//			db.AddError(db.Statement.Schema.PrioritizedPrimaryField.Set(db.Statement.Context, rv, insertID))
				//			insertID += db.Statement.Schema.PrioritizedPrimaryField.AutoIncrementIncrement
				//		}
				//	}
				//}
			case reflect.Struct:
				_, isZero := db.Statement.Schema.PrioritizedPrimaryField.ValueOf(db.Statement.Context, db.Statement.ReflectValue)
				if isZero {
					_ = db.AddError(db.Statement.Schema.PrioritizedPrimaryField.Set(db.Statement.Context, db.Statement.ReflectValue, insertID))
				}
			default:
			}
		}
	}
}

func MergeCreate(db *gorm.DB, onConflict clause.OnConflict, values clause.Values) {
	_, _ = db.Statement.WriteString("MERGE INTO ")
	db.Statement.WriteQuoted(db.Statement.Table)
	_, _ = db.Statement.WriteString(" USING (")
	for idx, value := range values.Values {
		if idx > 0 {
			_, _ = db.Statement.WriteString(" UNION ALL ")
		}

		_, _ = db.Statement.WriteString("SELECT ")
		db.Statement.AddVar(db.Statement, value...)
		_, _ = db.Statement.WriteString(" FROM DUAL")
	}

	_, _ = db.Statement.WriteString(`) AS "excluded" (`)
	for idx, column := range values.Columns {
		if idx > 0 {
			_ = db.Statement.WriteByte(',')
		}
		db.Statement.WriteQuoted(column.Name)
	}
	_, _ = db.Statement.WriteString(") ON ")

	var where clause.Where
	var whereFields []string
	if len(onConflict.Columns) > 0 {
		for _, column := range onConflict.Columns {
			whereFields = append(whereFields, column.Name)
		}
	} else {
		for _, field := range db.Statement.Schema.PrimaryFields {
			whereFields = append(whereFields, field.DBName)
		}
	}
	for _, field := range whereFields {
		where.Exprs = append(where.Exprs, clause.Eq{
			Column: clause.Column{Table: db.Statement.Table, Name: field},
			Value:  clause.Column{Table: "excluded", Name: field},
		})
	}
	where.Build(db.Statement)

	if len(onConflict.DoUpdates) > 0 {
		// 将UPDATE子句中出现在关联条件中的列去除（即上面的ON子句），否则会报错：-4064:不能更新关联条件中的列
		var withoutOnColumns = make([]clause.Assignment, 0, len(onConflict.DoUpdates))
	a:
		for _, assignment := range onConflict.DoUpdates {
			for _, field := range whereFields {
				if assignment.Column.Name == field {
					continue a
				}
			}
			withoutOnColumns = append(withoutOnColumns, assignment)
		}
		onConflict.DoUpdates = withoutOnColumns
		if len(onConflict.DoUpdates) > 0 {
			_, _ = db.Statement.WriteString(" WHEN MATCHED THEN UPDATE SET ")
			onConflict.DoUpdates.Build(db.Statement)
		}
	}

	_, _ = db.Statement.WriteString(" WHEN NOT MATCHED THEN INSERT (")

	written := false
	for _, column := range values.Columns {
		if db.Statement.Schema.PrioritizedPrimaryField == nil || !db.Statement.Schema.PrioritizedPrimaryField.AutoIncrement || db.Statement.Schema.PrioritizedPrimaryField.DBName != column.Name {
			if written {
				_ = db.Statement.WriteByte(',')
			}
			written = true
			db.Statement.WriteQuoted(column.Name)
		}
	}

	_, _ = db.Statement.WriteString(") VALUES (")

	written = false
	for _, column := range values.Columns {
		if db.Statement.Schema.PrioritizedPrimaryField == nil || !db.Statement.Schema.PrioritizedPrimaryField.AutoIncrement || db.Statement.Schema.PrioritizedPrimaryField.DBName != column.Name {
			if written {
				_ = db.Statement.WriteByte(',')
			}
			written = true
			db.Statement.WriteQuoted(clause.Column{
				Table: "excluded",
				Name:  column.Name,
			})
		}
	}

	_, _ = db.Statement.WriteString(")")
	//outputInserted(db)
	_, _ = db.Statement.WriteString(";")

	// merge into 语句插入的记录，无法通过LastInsertID获取
	if db.Statement.Schema.PrioritizedPrimaryField != nil && db.Statement.Schema.PrioritizedPrimaryField.AutoIncrement {
		_, _ = db.Statement.WriteString("SELECT ")
		db.Statement.WriteQuoted(db.Statement.Schema.PrioritizedPrimaryField.DBName)
		_, _ = db.Statement.WriteString(" FROM ")
		db.Statement.WriteQuoted(db.Statement.Table)
		_, _ = db.Statement.WriteString(" ORDER BY ")
		db.Statement.WriteQuoted(db.Statement.Schema.PrioritizedPrimaryField.DBName)
		_, _ = db.Statement.WriteString(" DESC LIMIT 1;")
	}
}
