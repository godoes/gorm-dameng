package dameng

import (
	"log"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/godoes/gorm-dameng/dm8"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dmUsername string
	dmPassword string

	dmHost string
	dmPort int

	dmSchema string

	waitOnce sync.Once
)

func init() {
	dmUsername = os.Getenv("DM_USERNAME")
	if dmUsername == "" {
		dmUsername = "SYSDBA"
	}
	dmPassword = os.Getenv("DM_PASSWORD")
	if dmPassword == "" {
		dmPassword = "SYSDBA"
	}

	dmHost = os.Getenv("DM_HOST")
	if dmHost == "" {
		dmHost = "localhost"
	}
	dmPort, _ = strconv.Atoi(os.Getenv("DM_PORT"))
	if dmPort == 0 {
		dmPort = int(dm8.DEFAULT_PORT)
	}

	dmSchema = os.Getenv("DM_SCHEMA")
	if dmSchema == "" {
		dmSchema = "SYSDBA"
	}

	log.Printf("数据库连接信息：\nUser: \t\t%s\nPassword: \t%s\nHOST: \t\t%s\nPORT: \t\t%d\n",
		dmUsername, dmPassword, dmHost, dmPort)
}

func testWaitInit() {
	waitOnce.Do(func() {
		if wait := os.Getenv("WAIT_MIN"); wait != "" {
			if min, e := strconv.Atoi(wait); e == nil {
				log.Println("wait for dm database initialization to complete...")
				time.Sleep(time.Duration(min) * time.Minute)
			}
		}
	})
}

type Product struct {
	gorm.Model
	Code  string
	Price uint

	Remark  string `gorm:"column:remark;size:0"`
	Remark1 string `gorm:"column:remark_1;size:-1"`
	Remark2 string `gorm:"column:remark_2;size:32768"`
}

func TestGormConnExample(t *testing.T) {
	testWaitInit()
	options := map[string]string{
		"schema":         dmSchema,
		"appName":        "GORM 连接操作达梦数据库测试",
		"connectTimeout": "30000",
	}

	dsn := BuildUrl(dmUsername, dmPassword, dmHost, dmPort, options)
	t.Logf("连接地址： %s", dsn)

	// 参考链接： https://eco.dameng.com/document/dm/zh-cn/pm/go-rogramming-guide.html#11.8%20ORM%20%E6%96%B9%E8%A8%80%E5%8C%85
	dialector := New(Config{DSN: dsn, VarcharSizeIsCharLength: true})
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.New(log.Default(), logger.Config{LogLevel: logger.Info}),
	})
	if err != nil {
		t.Fatalf("连接数据库 [%s@%s:%d] 失败：%v", dmUsername, dmHost, dmPort, err)
	} else {
		t.Logf("连接数据库 [%s@%s:%d] 成功！", dmUsername, dmHost, dmPort)
	}

	// 迁移 schema
	if err = db.AutoMigrate(&Product{}); err != nil {
		t.Errorf("迁移表结构失败：%v", err)
	} else {
		t.Logf("迁移表结构成功！")
	}
	// 测试重新迁移 schema
	if err = db.AutoMigrate(&Product{}); err != nil {
		t.Errorf("迁移表结构失败：%v", err)
	} else {
		t.Logf("迁移表结构成功！")
	}

	// Create
	data := Product{Code: "D42", Price: 100, Remark1: "VARCHAR", Remark2: "CLOB"}
	db.Create(&data)
	if err = db.Error; err != nil {
		t.Errorf("创建数据失败：%v", err)
	} else {
		t.Logf("创建数据成功！数据 ID：%d", data.ID)
	}
	// Create - 批量创建 map 型数据
	list := []map[string]any{
		{"code": "M42", "price": 200, "remark": "map1"},
		{"code": "N42", "price": 200, "remark": "map2"},
	}
	var listIDs []any
	db.Model(&Product{}).Create(list)
	if err = db.Error; err != nil {
		t.Errorf("批量创建 map 型数据失败：%v", err)
	} else {
		listIDs = make([]any, len(list))
		for i, item := range list {
			listIDs[i] = item["id"]
		}
		t.Logf("批量创建 map 型数据成功！数据 IDs：%+v", listIDs)
	}

	// Read
	var product Product
	// 根据整型主键查找
	if err = db.First(&product, data.ID).Error; err != nil {
		t.Errorf("根据 ID 获取数据失败：%v", err)
	} else {
		t.Logf("根据 ID 获取数据成功！\n%+v", product)
	}

	// 若 创建数据库 初始化参数 时选中了“字符串比较大小写敏感”，则查询时 SQL 字段名要加双引号，
	// 可通过 SELECT CASE_SENSITIVE() 查询数据库是否启用了大小写敏感，若未启用则引号可加可不加。
	// 参考链接： https://eco.dameng.com/community/article/a9a9fab6fc1b86483e82317d1ccc1acb
	// 查找 code 字段值为 D42 的记录
	if err = db.First(&product, `"code" = ?`, "D42").Error; err != nil {
		t.Errorf("获取数据失败：%v", err)
	} else {
		t.Logf("获取数据数据成功！\n%+v", product)
	}

	// Update - 将 product 的 price 更新为 200
	if err = db.Model(&product).Update("Price", 200).Error; err != nil {
		t.Errorf("根据列名更新单个字段失败：%v", err)
	} else {
		t.Log("根据列名更新单个字段成功！")
	}
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	if err = db.Error; err != nil {
		t.Errorf("根据结构体更新非零字段失败：%v", err)
	} else {
		t.Log("根据结构体更新非零字段成功！")
	}
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	if err = db.Error; err != nil {
		t.Errorf("根据 map 更新指定字段失败：%v", err)
	} else {
		t.Log("根据 map 更新指定字段成功！")
	}

	// Delete - 删除 product
	db.Delete(&product, data.ID)
	if err = db.Error; err != nil {
		t.Errorf("删除数据失败：%v", err)
	} else {
		t.Log("删除数据成功！")
	}
	// Delete - 批量删除
	db.Delete(&Product{}, "id IN (?)", listIDs)
	if err = db.Error; err != nil {
		t.Errorf("批量删除数据失败：%v", err)
	} else {
		t.Log("批量删除数据成功！")
	}

	//goland:noinspection SqlNoDataSourceInspection
	db.Exec(`DROP table "products"`)
	if err = db.Error; err != nil {
		t.Errorf("删除表结构失败：%v", err)
	} else {
		t.Log("删除表结构成功！")
	}
}

/******************** TestGormConnExample 测试结果 ********************

数据库连接信息：
User:           SYSDBA
Password:       SYSDBA
HOST:           127.0.0.1
PORT:           5236
=== RUN   TestGormConnExample
    dameng_test.go:65: 连接地址： dm://SYSDBA:SYSDBA@127.0.0.1:5236?appName=GORM+%E8%BF%9E%E6%8E%A5%E6%93%8D%E4%BD%9C%E8%BE%BE%E6%A2%A6%E6%95%B0%E6%8D%AE%E5%BA%93%E6%B5%8B%E8%AF%95&connectTimeout=30000&schema=SYSDBA
    dameng_test.go:72: 连接数据库 [SYSDBA@127.0.0.1:5236] 成功！
    dameng_test.go:79: 迁移表结构成功！
    dameng_test.go:87: 创建数据成功！
    dameng_test.go:96: 根据 ID 获取数据成功！
    dameng_test.go:106: 获取数据数据成功！
    dameng_test.go:114: 根据列名更新单个字段成功！
    dameng_test.go:121: 根据结构体更新非零字段成功！
    dameng_test.go:127: 根据 map 更新指定字段成功！
    dameng_test.go:135: 删除数据成功！
    dameng_test.go:143: 删除表结构成功！
--- PASS: TestGormConnExample (0.43s)
PASS

**************************************************************/
