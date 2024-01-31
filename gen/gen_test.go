package gen

import (
	"fmt"
	"testing"

	"github.com/liyonge-cm/go-api-cli/config"
)

func TestGen(t *testing.T) {
	// 读取配置文件
	config.LoadConfig("../config/config.yml")

	s := NewGenServer()
	s.ConnDB()

	// s.RemaneTableFileName = func(name string) string {
	// 	return "A" + name
	// }

	// 获取表字段
	s.GetTableFields()

	// 生成model
	fmt.Println(s.modelPath)
	s.GenModel()

	// 生成API
	s.GenApi()

}
