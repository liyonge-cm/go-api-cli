package templates

import (
	"fmt"
	"strings"
)

func getSysParam() (id, data, allData, count string) {
	id = "Id	int	`json:\"id\"`"
	data = "Data	*model.$model_name	`json:\"data\"`"
	allData = "Data	[]*model.$model_name	`json:\"data\"`"
	count = "Count int64	`json:\"count\"`"

	return

}

func getParamsContent(params []string) string {
	content := ""
	for _, v := range params {
		content += fmt.Sprintf(`
		%v:	req.Data.%v,`, v, v)
	}
	return content
}

func GetCreateContent(prjName, packageName, modelName, funcName string, params []string) string {
	content := createContent
	content = strings.ReplaceAll(content, "$package_name", packageName)
	content = strings.ReplaceAll(content, "$prj_name", prjName)
	content = strings.ReplaceAll(content, "$func_name", funcName)
	content = strings.ReplaceAll(content, "$model_name", modelName)
	content = strings.ReplaceAll(content, "$params_content", getParamsContent(params))
	return content
}

func GetDeleteContent(prjName, packageName, modelName, funcName string) string {
	content := deleteContent

	id, _, _, _ := getSysParam()
	content = strings.ReplaceAll(content, "$id_request", id)

	content = strings.ReplaceAll(content, "$package_name", packageName)
	content = strings.ReplaceAll(content, "$prj_name", prjName)
	content = strings.ReplaceAll(content, "$func_name", funcName)
	content = strings.ReplaceAll(content, "$model_name", modelName)

	return content
}

func GetReadAllContent(prjName, packageName, modelName, funcName string) string {
	content := readAllContent

	_, _, allData, count := getSysParam()
	content = strings.ReplaceAll(content, "$all_data_response", allData)
	content = strings.ReplaceAll(content, "$count_response", count)

	content = strings.ReplaceAll(content, "$package_name", packageName)
	content = strings.ReplaceAll(content, "$prj_name", prjName)
	content = strings.ReplaceAll(content, "$func_name", funcName)
	content = strings.ReplaceAll(content, "$model_name", modelName)
	return content
}

func GetReadContent(prjName, packageName, modelName, funcName string) string {
	content := readContent

	id, data, _, count := getSysParam()
	content = strings.ReplaceAll(content, "$id_request", id)
	content = strings.ReplaceAll(content, "$data_response", data)
	content = strings.ReplaceAll(content, "$count_response", count)

	content = strings.ReplaceAll(content, "$package_name", packageName)
	content = strings.ReplaceAll(content, "$prj_name", prjName)
	content = strings.ReplaceAll(content, "$func_name", funcName)
	content = strings.ReplaceAll(content, "$model_name", modelName)
	return content
}

func GetUpdateContent(prjName, packageName, modelName, funcName string, params []string) string {
	content := updateContent
	content = strings.ReplaceAll(content, "$package_name", packageName)
	content = strings.ReplaceAll(content, "$prj_name", prjName)
	content = strings.ReplaceAll(content, "$func_name", funcName)
	content = strings.ReplaceAll(content, "$model_name", modelName)
	content = strings.ReplaceAll(content, "$params_content", getParamsContent(params))
	return content
}

func GetRouterContent(prjName, packageName, groupRouter, createFuncName, readFuncName, readAllFuncName, updateFuncName, deleteFuncName string) string {
	content := routerContent
	content = strings.ReplaceAll(content, "$package_name", packageName)
	content = strings.ReplaceAll(content, "$prj_name", prjName)
	content = strings.ReplaceAll(content, "$group_router", groupRouter)
	content = strings.ReplaceAll(content, "$create_func_name", createFuncName)
	content = strings.ReplaceAll(content, "$read_func_name", readFuncName)
	content = strings.ReplaceAll(content, "$read_all_func_name", readAllFuncName)
	content = strings.ReplaceAll(content, "$update_func_name", updateFuncName)
	content = strings.ReplaceAll(content, "$delete_func_name", deleteFuncName)
	return content
}
