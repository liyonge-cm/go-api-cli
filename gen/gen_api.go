package gen

import (
	"fmt"
	"os"
	"path"
	"strings"

	"go-cli/gen/templates"
	"go-cli/utils"
)

func (s *GenServer) GenApi() error {
	err := utils.CreateDir(s.apiPath)
	if err != nil {
		return err
	}
	err = s.generatorApis()
	if err != nil {
		return err
	}
	return nil
}

func (s *GenServer) generatorApis() error {
	for _, table := range s.tableInfos {
		tableName := table.TableFileName
		groupPath := path.Join(s.apiPath, tableName)
		err := utils.CreateDir(groupPath)
		if err != nil {
			return err
		}
		err = s.generatorApiGroup(groupPath, table)
		if err != nil {
			return err
		}
		_ = s.importApiRouter(s.apiPath, s.cfg.Frame.PrjName, table.TableFileName)

	}
	return nil
}

func (s *GenServer) generatorApiGroup(groupPath string, table *TableInfo) (err error) {
	var (
		prjName     = s.cfg.Frame.PrjName
		packageName = table.TableFileName
		modelName   = table.TableModelName
		params      = []string{}

		createFuncName  = utils.ToCamel(fmt.Sprintf("create_%v", packageName))
		readFuncName    = utils.ToCamel(fmt.Sprintf("get_%v", packageName))
		readAllFuncName = utils.ToCamel(fmt.Sprintf("get_%v_list", packageName))
		updateFuncName  = utils.ToCamel(fmt.Sprintf("update_%v", packageName))
		deleteFuncName  = utils.ToCamel(fmt.Sprintf("delete_%v", packageName))
	)

	for _, v := range table.Columns {
		if v.Field != "id" && v.Field != "status" && v.Field != "created_at" && v.Field != "updated_at" {
			params = append(params, v.Name)
		}
	}
	_ = s.generatorCreateApi(groupPath, prjName, packageName, modelName, createFuncName, params)
	_ = s.generatorReadApi(groupPath, prjName, packageName, modelName, readFuncName)
	_ = s.generatorReadAllApi(groupPath, prjName, packageName, modelName, readAllFuncName)
	_ = s.generatorUpdateApi(groupPath, prjName, packageName, modelName, updateFuncName, params)
	_ = s.generatorDeleteApi(groupPath, prjName, packageName, modelName, deleteFuncName)
	_ = s.generatorApiRouter(groupPath, prjName, packageName, packageName, createFuncName, readFuncName, readAllFuncName, updateFuncName, deleteFuncName)

	return err
}

func (s *GenServer) generatorCreateApi(groupPath, prjName, packageName, modelName, funcName string, params []string) error {
	fileName := path.Join(groupPath, packageName+"_create.go")
	content := templates.GetCreateContent(prjName, packageName, modelName, funcName, params)
	if s.isJsonCamel {
		content = utils.JsonToCamel(content)
	}
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *GenServer) generatorReadApi(groupPath, prjName, packageName, modelName, funcName string) error {
	fileName := path.Join(groupPath, packageName+"_read.go")
	content := templates.GetReadContent(prjName, packageName, modelName, funcName)
	if s.isJsonCamel {
		content = utils.JsonToCamel(content)
	}
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *GenServer) generatorReadAllApi(groupPath, prjName, packageName, modelName, funcName string) error {
	fileName := path.Join(groupPath, packageName+"_read_all.go")
	content := templates.GetReadAllContent(prjName, packageName, modelName, funcName)
	if s.isJsonCamel {
		content = utils.JsonToCamel(content)
	}
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *GenServer) generatorUpdateApi(groupPath, prjName, packageName, modelName, funcName string, params []string) error {
	fileName := path.Join(groupPath, packageName+"_update.go")
	content := templates.GetUpdateContent(prjName, packageName, modelName, funcName, params)
	if s.isJsonCamel {
		content = utils.JsonToCamel(content)
	}
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *GenServer) generatorDeleteApi(groupPath, prjName, packageName, modelName, funcName string) error {
	fileName := path.Join(groupPath, packageName+"_delete.go")
	content := templates.GetDeleteContent(prjName, packageName, modelName, funcName)
	if s.isJsonCamel {
		content = utils.JsonToCamel(content)
	}
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *GenServer) generatorApiRouter(groupPath, prjName, packageName, groupRouter, createFuncName, readFuncName, readAllFuncName, updateFuncName, deleteFuncName string) error {
	fileName := path.Join(groupPath, packageName+".go")
	content := templates.GetRouterContent(prjName, packageName, groupRouter, createFuncName, readFuncName, readAllFuncName, updateFuncName, deleteFuncName)
	if s.isJsonCamel {
		content = utils.JsonToCamel(content)
	}
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *GenServer) importApiRouter(filePath, prjName, groupName string) error {
	apiPath := path.Join(filePath, "apis.go")

	datab, err := os.ReadFile(apiPath)
	if err != nil {
		return err
	}
	content := string(datab)
	newImportContent := fmt.Sprintf(`"%v/service/apis/%v"`, prjName, groupName)
	if !strings.Contains(content, newImportContent) {
		content += fmt.Sprintf(`
import _ %v`, newImportContent)
	}

	err = utils.SaveFile(apiPath, []byte(content))
	if err != nil {
		fmt.Println("SaveFile err", err.Error())
		return err
	}
	return nil
}
