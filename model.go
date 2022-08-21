package gomodel

import (
	"errors"
	"gorm.io/gorm"
	"io"
	"os"
	"strings"
)

// Table table
type Table struct {
	TableName    string `gorm:"column:TABLE_NAME"`    // table name
	TableComment string `gorm:"column:TABLE_COMMENT"` // table comment
}

// Field field
type Field struct {
	Field      string `gorm:"column:Field"`
	Type       string `gorm:"column:Type"`
	Collation  string `gorm:"column:Collation"`
	Null       string `gorm:"column:Null"`
	Key        string `gorm:"column:Key"`
	Default    string `gorm:"column:Default"`
	Extra      string `gorm:"column:Extra"`
	Privileges string `gorm:"column:Privileges"`
	Comment    string `gorm:"column:Comment"`
}

type ModelConfig struct {
	DB      *gorm.DB
	DbName  string
	MDir    string
	Prefix  string
	IsCover bool
}

// GenerateAllModel generates the struct for all table
func GenerateAllModel(cfg *ModelConfig) ([]string, []error) {
	var strs []string
	var errs []error
	tables := getAllTables(cfg.DB, cfg.DbName)
	for _, table := range tables {
		str, err := GenerateSingleModel(cfg, table.TableName)
		strs = append(strs, str)
		errs = append(errs, err)
	}
	return strs, errs
}

// GenerateSingleModel generates the structure content of a single Model
func GenerateSingleModel(cfg *ModelConfig, tbName string) (string, error) {
	var fields []Field
	table := getSingleTable(cfg.DB, cfg.DbName, tbName)
	fields = getFieldsByTable(cfg.DB, tbName)

	content, camelTbName, _tbName := parseFieldsByTable(tbName, table.TableComment, fields, cfg.MDir, cfg.Prefix)
	if err := makeMultiDir(cfg.MDir); err != nil {
		return cfg.MDir + " create failed.", err
	}
	var f *os.File
	fileName := cfg.MDir + "/" + _tbName + ".go"
	if cfg.IsCover == false {
		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			return camelTbName + " already existed.", errors.New(camelTbName + " already existed")
		}
	}
	f, err := os.Create(fileName)
	if err != nil {
		return fileName + " create failed.", err
	}
	defer f.Close()
	if _, err := io.WriteString(f, content); err != nil {
		return camelTbName + " generate failed.", err
	}
	return camelTbName + " generate success.", nil
}

// getAllTables get all table information and field information
func getAllTables(db *gorm.DB, dbName string) []Table {
	var tables []Table
	db.Raw(`
			SELECT
				TABLE_NAME,      -- table name
				TABLE_COMMENT    -- table comment
				FROM
				INFORMATION_SCHEMA.TABLES
				WHERE TABLE_SCHEMA = "` + dbName + `" -- db name
		`).Scan(&tables)
	return tables
}

// getSingleTable get individual table information and field information
func getSingleTable(db *gorm.DB, dbName string, tbName string) Table {
	var table Table
	db.Raw(`
			SELECT
				TABLE_NAME,      -- table name
				TABLE_COMMENT    -- table comment
				FROM
				INFORMATION_SCHEMA.TABLES
				WHERE TABLE_SCHEMA = "` + dbName + `" AND TABLE_NAME = "` + tbName + `"
		`).Find(&table)
	return table
}

// parseFieldsByTable converts the type of a data table field
func parseFieldsByTable(tbName, tbComment string, fields []Field, mDir, prefix string) (string, string, string) {
	pkgNameArr := strings.Split(mDir, "/")
	content := "package " + pkgNameArr[len(pkgNameArr)-1] + "\n\n"
	_tbName := tbName
	// whether to remove the table prefix
	if prefix != "" {
		_tbName = strings.TrimPrefix(tbName, prefix)
	}
	camelTbName := CamelCase(_tbName)
	content += "var " + camelTbName + "TbName = \"" + tbName + "\"\n\n"
	if len(tbComment) > 0 {
		content += "// " + camelTbName + " " + tbComment + "\n"
	}
	content += "type " + camelTbName + " struct {\n"
	for _, val := range fields {
		// 生成字段
		columnField := CamelCase(val.Field)
		columnJson := "`gorm:\""
		if val.Key == "PRI" {
			columnJson += "primaryKey;"
		}
		if val.Extra == "auto_increment" {
			columnJson += "autoIncrement;"
		}
		columnJson += "column:" + val.Field + ";type:" + val.Type + ";"
		if val.Default != "" {
			columnJson += "default:" + val.Default + ";"
		}
		if val.Null == "NO" {
			columnJson += "NOT NULL;"
		} else {
			columnJson += "NULL;"
		}
		if val.Comment != "" {
			columnJson += "comment:" + val.Comment
		}
		columnJson += "\" json:\"" + val.Field + "\"`"
		columnType := parseFieldTypeByTable(val.Type)
		columnComment := ""
		if len(val.Comment) > 0 {
			columnComment = "// " + val.Comment
		}
		content += "    " + columnField + " " + columnType + " " + columnJson + " " + columnComment + "\n"
	}
	content += "}"
	return content, camelTbName, _tbName
}

// getFieldsByTable get the table fields from the table
func getFieldsByTable(db *gorm.DB, tbName string) []Field {
	var field []Field
	db.Raw(`
		SHOW FULL COLUMNS FROM ` + tbName + `
	`).Find(&field)
	return field
}

// parseFieldTypeByTable escape a database field type to a struct
func parseFieldTypeByTable(fieldType string) string {
	typeArr := strings.Split(fieldType, "(")
	switch typeArr[0] {
	case "int", "integer", "mediumint", "bit", "year", "smallint", "int unsigned", "mediumint unsigned", "smallint unsigned":
		return "int"
	case "tinyint":
		return "int8"
	case "tinyint unsigned":
		return "uint8"
	case "bigint":
		return "int64"
	case "bigint unsigned":
		return "uint64"
	case "double", "float", "real", "numeric":
		return "float32"
	case "decimal":
		return "decimal.Decimal"
	case "timestamp", "datetime", "time":
		return "time.Time"
	default:
		return "string"
	}
}

// isPathExist 判断所给路径文件/文件夹是否存在
func isPathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// makeMultiDir 调用os.MkdirAll递归创建文件夹
func makeMultiDir(filePath string) error {
	if !isPathExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}
