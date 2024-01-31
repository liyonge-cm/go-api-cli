package gen

import (
	"fmt"
	"regexp"

	"github.com/liyonge-cm/go-api-cli/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type TableInfo struct {
	Table          string
	TableFileName  string
	TableModelName string
	Columns        []*ColumnInfo
}
type ColumnInfo struct {
	Field   string
	Name    string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
	Comment string
}

func (s *GenServer) ConnDB() error {
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使用单数表名
		},
	}
	db, err := gorm.Open(mysql.Open(s.dsn()), gormConfig)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *GenServer) dsn() string {
	dsn := s.cfg.MySQL
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dsn.Username,
		dsn.Password,
		dsn.Endpoint,
		dsn.Database,
	)
}

func (s *GenServer) GetTableFields() (err error) {
	tableColumnInfo, err := s.getTableColumnInfo(s.cfg.Api.Tables)
	if err != nil {
		return err
	}

	for table, columnInfo := range tableColumnInfo {
		tableInfo := &TableInfo{
			Table:          table,
			TableFileName:  table,
			TableModelName: utils.ToCamel(table),
			Columns:        columnInfo,
		}
		if s.RemaneTableFileName != nil {
			tableInfo.TableFileName = s.RemaneTableFileName(table)
		}
		if s.RemaneTableModelName != nil {
			tableInfo.TableFileName = s.RemaneTableModelName(table)
		}
		s.tableInfos[table] = tableInfo
	}
	return nil
}

func (s *GenServer) getTableColumnInfo(tables []string) (tableColumnInfo map[string][]*ColumnInfo, err error) {
	tableColumnInfo = map[string][]*ColumnInfo{}
	for _, table := range tables {
		var columns []*ColumnInfo
		s.db.Raw(fmt.Sprintf("SHOW FULL COLUMNS FROM `%v` ", table)).Scan(&columns)
		for _, v := range columns {
			v.Name = utils.ToCamel(v.Field)
			v.Type = s.getFieldType(v.Type)
		}
		tableColumnInfo[table] = columns
	}
	return
}

func (s *GenServer) getFieldType(fieldType string) string {
	re := regexp.MustCompile(`^(\w+)\(`)
	match := re.FindStringSubmatch(fieldType)
	if len(match) > 1 {
		fieldType = match[1]
	}

	fieldTypeMap := map[string]string{
		"varchar": "string",
		"decimal": "float64",
		"bigint":  "int64",
		"tinyint": "int",
	}
	if newType, ok := fieldTypeMap[fieldType]; ok {
		return newType
	}
	return fieldType
}
