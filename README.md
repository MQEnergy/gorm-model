# gorm-model
基于gorm的一键生成关联mysql数据表模型结构体

[![GoDoc](https://pkg.go.dev/badge/github.com/MQEnergy/gorm-model/?status.svg)](https://pkg.go.dev/github.com/MQEnergy/gorm-model)
[![Go Report Card](https://goreportcard.com/badge/github.com/MQEnergy/gorm-model)](https://goreportcard.com/report/github.com/MQEnergy/gorm-model)
[![codebeat badge](https://codebeat.co/badges/a1d6feb8-909f-49a5-8d5c-3600b64bda04)](https://codebeat.co/projects/github-com-mqenergy-gorm-model-main)
[![GitHub license](https://img.shields.io/github/license/MQEnergy/gorm-model)](https://github.com/MQEnergy/gorm-model/blob/main/LICENSE)

### 注意
```
需要注意本组件根据数据表的数据字段有两个特殊格式 time和decimal
需要引入相应组件：
1、time
2、github.com/shopspring/decimal（或其他自定义支持decimal.Decimal方法的组件）
```
### 安装到项目中
```shell script
go get -u github.com/MQEnergy/gorm-model
```

### 方法

#### 1、生成全部数据表对应的model 
##### GenerateAllModel(cfg *ModelConfig) 
```
db: gorm连接实例
dbName: 数据库名称
mDir: 模型存储目录
prefix: 去除的表前缀名称，不去除传空，去除出前缀名 如：gin
```

#### 2、生成单个数据表对应的model 
##### GenerateSingleModel(cfg *ModelConfig, tbName string) 
```
db: gorm连接实例
dbName: 数据库名称
tbName: 数据表名称
mDir: 模型存储目录
prefix: 去除的表前缀名称，不去除传空，去除出前缀名 如：gin
```

### 查看案例examples
```go
go run examples/model.go --help
```
```
 -db string
        数据库名称 如：gin_framework (default "gin_framework")
  -dir string
        模型存储目录 如：./models（存入在当前执行命令所在目录，支持多级目录） (default "./models")
  -dsn string
        数据库连接信息 如：root:123456@tcp(127.0.0.1:3306)/gin_framework?charset=utf8mb4&parseTime=True&loc=Local (default "root:123456@tcp(127.0.0.1:3306)/gin_framework?charset=utf8mb4&parseTime=True&loc=Local")
   -ic
        是否覆盖原生成的模型结构体 true:覆盖 false:不覆盖
  -p string
        数据表前缀 如: gin_ (default "gin_")
  -tb string
        模型名称 如：初始化所有（all）单个数据表就填写表名（如：gin_admin） (default "all")
```
按照以上参数可自定义生成数据表模型

### 注意：
1、`time.Time`类型需要import "time"，需手动加载

2、`decimal.Decimal`类型需要 import "github.com/shopspring/decimal"，需手动安装加载