package gen

import (
	"fmt"
	"regexp"
	"strings"

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
			// v.Type = s.getFieldType(v.Type)
			t, err := s.checkLocalTypeForField(v.Type)
			if err != nil {
				return nil, err
			}
			v.Type = string(t)
		}
		tableColumnInfo[table] = columns
	}
	return
}

func (s *GenServer) GetFieldType(fieldType string) string {
	re := regexp.MustCompile(`^(\w+)\(`)
	match := re.FindStringSubmatch(fieldType)
	if len(match) > 1 {
		fieldType = match[1]
	}

	fieldTypeMap := map[string]string{
		"varchar":   "string",
		"decimal":   "float64",
		"bigint":    "int64",
		"tinyint":   "int",
		"smallint":  "int",
		"mediumint": "int",
		"float":     "float32",
	}
	if newType, ok := fieldTypeMap[fieldType]; ok {
		return newType
	}
	return fieldType
}

// CheckLocalTypeForField checks and returns corresponding type for given db type.
func (c *GenServer) checkLocalTypeForField(fieldType string) (LocalType, error) {
	var (
		typeName string
	)

	re := regexp.MustCompile(`^(\w+)\(`)
	match := re.FindStringSubmatch(fieldType)
	if len(match) > 1 {
		typeName = match[1]
	} else {
		typeName = fieldType
	}

	typeName = strings.ToLower(typeName)

	switch typeName {
	case
		fieldTypeBinary,
		fieldTypeVarbinary,
		fieldTypeBlob,
		fieldTypeTinyblob,
		fieldTypeMediumblob,
		fieldTypeLongblob:
		return LocalTypeBytes, nil

	case
		fieldTypeInt,
		fieldTypeTinyint,
		fieldTypeSmallInt,
		fieldTypeSmallint,
		fieldTypeMediumInt,
		fieldTypeMediumint,
		fieldTypeSerial:
		if utils.ContainsI(fieldType, "unsigned") {
			return LocalTypeUint, nil
		}
		return LocalTypeInt, nil

	case
		fieldTypeBigInt,
		fieldTypeBigint,
		fieldTypeBigserial:
		if utils.ContainsI(fieldType, "unsigned") {
			return LocalTypeUint64, nil
		}
		return LocalTypeInt64, nil

	case
		fieldTypeReal:
		return LocalTypeFloat32, nil

	case
		fieldTypeDecimal,
		fieldTypeMoney,
		fieldTypeNumeric,
		fieldTypeSmallmoney:
		return LocalTypeString, nil
	case
		fieldTypeFloat,
		fieldTypeDouble:
		return LocalTypeFloat64, nil

	case
		fieldTypeBit:
		// It is suggested using bit(1) as boolean.
		// if typePattern == "1" {
		// 	return LocalTypeBool, nil
		// }
		return LocalTypeInt64Bytes, nil

	case
		fieldTypeBool:
		return LocalTypeBool, nil

	case
		fieldTypeDate:
		return LocalTypeDate, nil

	case
		fieldTypeDatetime,
		fieldTypeTimestamp,
		fieldTypeTimestampz:
		return LocalTypeDatetime, nil

	case
		fieldTypeJson:
		return LocalTypeJson, nil

	case
		fieldTypeJsonb:
		return LocalTypeJsonb, nil

	default:
		// Auto-detect field type, using key match.
		switch {
		case strings.Contains(typeName, "text") || strings.Contains(typeName, "char") || strings.Contains(typeName, "character"):
			return LocalTypeString, nil

		case strings.Contains(typeName, "float") || strings.Contains(typeName, "double") || strings.Contains(typeName, "numeric"):
			return LocalTypeFloat64, nil

		case strings.Contains(typeName, "bool"):
			return LocalTypeBool, nil

		case strings.Contains(typeName, "binary") || strings.Contains(typeName, "blob"):
			return LocalTypeBytes, nil

		case strings.Contains(typeName, "int"):
			if utils.ContainsI(fieldType, "unsigned") {
				return LocalTypeUint, nil
			}
			return LocalTypeInt, nil

		case strings.Contains(typeName, "time"):
			return LocalTypeDatetime, nil

		case strings.Contains(typeName, "date"):
			return LocalTypeDatetime, nil

		default:
			return LocalTypeString, nil
		}
	}
}
