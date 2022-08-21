package main

import (
	"flag"
	"fmt"
	gomodel "github.com/MQEnergy/gorm-model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	dbName  string
	tbName  string
	mDir    string // 模型存储目录
	prefix  string // 数据表前缀
	dsn     string // 数据库连接信息
	isCover bool   // 是否覆盖
)

func init() {
	flag.StringVar(&dbName, "db", "gin_framework", "数据库名称 如：gin_framework")
	flag.StringVar(&tbName, "tb", "all", "模型名称 如：初始化所有（all）单个数据表就填写表名（如：gin_admin）")
	flag.StringVar(&prefix, "p", "gin_", "数据表前缀 如: gin_")
	flag.StringVar(&dsn, "dsn", "root:123456@tcp(127.0.0.1:3306)/gin_framework?charset=utf8mb4&parseTime=True&loc=Local", "数据库连接信息 如：root:123456@tcp(127.0.0.1:3306)/gin_framework?charset=utf8mb4&parseTime=True&loc=Local")
	flag.StringVar(&mDir, "dir", "./models", "模型存储目录 如：./models（存入在当前执行命令所在目录，支持多级目录）")
	flag.BoolVar(&isCover, "ic", false, "是否覆盖原生成的模型结构体 true:覆盖 false:不覆盖")
}

func main() {
	flag.Parse()
	db, err := newMysql()
	if err != nil {
		fmt.Println(err)
		return
	}
	modelConfig := &gomodel.ModelConfig{
		DB:      db,
		DbName:  dbName,
		MDir:    mDir,
		Prefix:  prefix,
		IsCover: isCover,
	}
	if tbName == "all" {
		strs, errs := gomodel.GenerateAllModel(modelConfig)
		for i, str := range strs {
			if errs[i] != nil {
				fmt.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", str))
			} else {
				fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", str))
			}
		}
	} else {
		str, err := gomodel.GenerateSingleModel(modelConfig, tbName)
		if err != nil {
			fmt.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", str))
			return
		}
		fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", str))
	}
}

func newMysql() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the database, please check the MySQL configuration information first,the error details are:" + err.Error())
	}
	return db, nil
}
