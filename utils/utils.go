package utils

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

func CreateDir(dir string) error {
	_, err := os.Stat(dir) //os.Stat获取文件信息
	if err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		return err
	}
	return nil
}

func SaveFile(filePath string, content []byte) error {
	dir := path.Dir(filePath)
	_, err := os.Stat(dir) //os.Stat获取文件信息
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Println("os.MkdirAll err", err.Error())
			return err
		}
	} else if err != nil {
		fmt.Println("os.Stat err", err.Error())
		return err
	}

	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		fmt.Println("os.WriteFile err", err.Error())
		return err
	}
	return err
}

// 将字符串按下划线转成驼峰字符串
func ToCamel(c string) string {
	strSlice := strings.Split(c, "_")
	for i := range strSlice {
		strSlice[i] = strings.Title(strSlice[i])
	}
	camelStr := strings.Join(strSlice, "")
	return camelStr
}

func JsonToCamel(content string) string {
	re := regexp.MustCompile(`json:"(.*?)"`)
	matchAll := re.FindAllStringSubmatch(content, 1000)
	for _, match := range matchAll {
		if len(match) > 1 {
			jsonContent := match[1]
			camelCase := ToCamel(jsonContent)
			content = strings.ReplaceAll(content, match[0], fmt.Sprintf(`json:"%v"`, camelCase))
		}
	}
	return content
}
