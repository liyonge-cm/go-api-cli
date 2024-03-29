package gen

import (
	"fmt"
	"os"
	"path"

	"github.com/liyonge-cm/go-api-cli/utils"
)

func (s *GenServer) GenModel() error {
	err := utils.CreateDir(s.modelPath)
	if err != nil {
		return err
	}
	err = s.generatorModels()
	if err != nil {
		return err
	}
	return nil
}

func (s *GenServer) generatorModels() error {
	for _, table := range s.tableInfos {
		filePath := path.Join(s.modelPath, table.TableFileName+".go")
		content, err := s.genFieldsContent(table.Table, table.TableModelName, table.Columns, s.isJsonCamel)
		if err != nil {
			return err
		}
		err = os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *GenServer) genFieldsContent(tableName, tableModelName string, fields []*ColumnInfo, isJsonCamel bool) (content string, err error) {
	content = fmt.Sprintf(`package %v

type %v struct{`, s.modelPkgName, tableModelName)
	for _, field := range fields {
		jsonCase := field.Field
		if isJsonCamel {
			jsonCase = field.Name
		}
		pcontent := fmt.Sprintf(`
	%v	%v	`, field.Name, field.Type) + "`json:" + `"` + jsonCase + `"`
		pcontent += "` // " + field.Comment
		content += pcontent
	}
	content += `
}
`
	content += fmt.Sprintf(`
func (%v) TableName() string {
	return "%v"
}
`, tableModelName, tableName)

	return
}
