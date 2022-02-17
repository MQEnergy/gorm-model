# gorm-model
一键生成关联数据表的gorm model 

### 安装到项目中
```shell script
go get -u github.com/MQEnergy/gorm-model
```

### 方法

#### 1、生成全部数据表对应的model 
##### GenerateAllModel(db *gorm.DB, dbName string) 
```go
// 调用方式
package main
import gorm_model "github.com/MQEnergy/gorm-model"
func main() {
  // dbName 是数据库名称
  gorm_model.GenerateAllModel(*gorm.DB, dbName)
}
```

#### 2、生成单个数据表对应的model 
##### GenerateSingleModel(db *gorm.DB, tbName string, table Table) 
```go
package main
import gorm_model "github.com/MQEnergy/gorm-model"

func main() {
	// dbName 数据库名称 tbName 数据表名称
    table = gorm_model.GetSingleTable(*gorm.DB, dbName, tbName)
    // tbName 是数据表名称
    err := gorm_model.GenerateSingleModel(*gorm.DB, tbName, table)
    if err != nil {
        panic(err)
    }
}
```
#### 3、获取单个表信息和字段信息 
##### GetSingleTable(db *gorm.DB, dbName string, tbName string) 
#### 4、根据表获取表字段 
##### GetFieldsByTable(db *gorm.DB, tbName string) 
#### 5、获取数据库中所有表信息和字段信息 
##### GetAllTables(db *gorm.DB, dbName string) 

