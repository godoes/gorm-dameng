# GORM DM8 Driver

基于达梦数据库官方 [Go 驱动源码](https://download.dameng.com/eco/adapter/resource/go/go-20230418.zip)
二次开发整理的开箱即用的 [GORM](https://gorm.io/zh_CN/) 达梦数据库驱动，无需单独复制驱动源码到项目中。
获取达梦最新官方 Go 驱动请访问 <https://eco.dameng.com/download/>。

## 最低要求

- go 1.18：<https://go.dev/dl/>
- gorm v2：<https://github.com/go-gorm/gorm>
- 达梦数据库 DM8：<https://eco.dameng.com/download/>

## 快速上手

### 安装

```shell
go get -d github.com/godoes/gorm-dameng
```

### 用例

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/godoes/gorm-dameng"
	"gorm.io/gorm"
)

func main() {
	options := map[string]string{
		"schema":         "SYSDBA",
		"appName":        "GORM 连接达梦数据库示例",
		"connectTimeout": "30000",
	}

	// dm://user:password@host:port?schema=SYSDBA[&...]
	dsn := dameng.BuildUrl("user", "password", "127.0.0.1", 5236, options)
	// VARCHAR 类型大小为字符长度
	//db, err := gorm.Open(dameng.New(dameng.Config{DSN: dsn, VarcharSizeIsCharLength: true}))
	// VARCHAR 类型大小为字节长度（默认）
	db, err := gorm.Open(dameng.Open(dsn), &gorm.Config{})
	if err != nil {
		// panic error or log error info
	}

	// do somethings
	var versionInfo []map[string]interface{}
	db.Table("SYS.V$VERSION").Find(&versionInfo)
	if err := db.Error; err == nil {
		versionBytes, _ := json.MarshalIndent(versionInfo, "", "  ")
		fmt.Printf("达梦数据库版本信息：\n%s\n", versionBytes)
	}
}

/****************** 控制台输出内容 *****************

达梦数据库版本信息：
[
  {
    "BANNER": "DM Database Server 64 V8"
  },
  {
    "BANNER": "DB Version: 0x7000c"
  },
  {
    "BANNER": "03134284094-20230927-******-*****"
  }
]

*************************************************/
```

## 参考文档

### 达梦技术文档

- 产品手册：<https://eco.dameng.com/document/dm/zh-cn/pm/>
- 常见问题：<https://eco.dameng.com/document/dm/zh-cn/faq/>
- GO 数据库接口：<https://eco.dameng.com/document/dm/zh-cn/app-dev/go_go.html>
- DM Go 编程指南：<https://eco.dameng.com/document/dm/zh-cn/pm/go-rogramming-guide.html>
- SQL 开发指南：<https://eco.dameng.com/document/dm/zh-cn/sql-dev/>

### GORM 文档

- 文档首页：<https://gorm.io/zh_CN/docs/>
- 声明模型：<https://gorm.io/zh_CN/docs/models.html>
- 连接到数据库：<https://gorm.io/zh_CN/docs/connecting_to_the_database.html>
