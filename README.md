# gorm-model
基于gorm的一键生成关联mysql数据表模型结构体 

### 安装到项目中
```shell script
go get -u github.com/MQEnergy/gorm-model
```

### 方法

#### 1、生成全部数据表对应的model 
##### GenerateAllModel(db *gorm.DB, dbName, mDir, prefix string) 
```
db: gorm连接实例
dbName: 数据库名称
mDir: 模型存储目录
prefix: 去除的表前缀名称，不去除传空，去除出前缀名 如：gin
```

#### 2、生成单个数据表对应的model 
##### GenerateSingleModel(db *gorm.DB, dbName, tbName, mDir, prefix string) 
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
  -p string
        数据表前缀 如: gin_ (default "gin_")
  -tb string
        模型名称 如：初始化所有（all）单个数据表就填写表名（如：gin_admin） (default "all")
```
按照以上参数可自定义生成数据表模型