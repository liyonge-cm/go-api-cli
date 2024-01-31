package gen_frame

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/liyonge-cm/go-api-cli/utils"
)

func TestGen(t *testing.T) {
	s := NewGenFrameService(&GenFrameConfig{
		OutPath:  "../../",
		PrjName:  "prj_aiee_demo",
		JsonCase: "",
	})
	err := s.GenFrame()
	if err != nil {
		fmt.Println("gen err,", err.Error())
	}
}

func TestGenJson(t *testing.T) {
	contentPath := "../prj/service/apis/apis.go"
	datab, err := os.ReadFile(contentPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	content := string(datab)

	re := regexp.MustCompile(`json:"(.*?)"`)
	matchAll := re.FindAllStringSubmatch(content, 1000)
	for _, match := range matchAll {
		if len(match) > 1 {
			jsonContent := match[1]
			camelCase := utils.ToCamel(jsonContent)
			content = strings.ReplaceAll(content, match[0], fmt.Sprintf(`json:"%v"`, camelCase))
		}
	}
	fmt.Println(content)
}
