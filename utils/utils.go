package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// DefaultTrimChars are the characters which are stripped by Trim* functions in default.
	DefaultTrimChars = string([]byte{
		'\t', // Tab.
		'\v', // Vertical tab.
		'\n', // New line (line feed).
		'\r', // Carriage return.
		'\f', // New page.
		' ',  // Ordinary space.
		0x00, // NUL-byte.
		0x85, // Delete.
		0xA0, // Non-breaking space.
	})
)

type iString interface {
	String() string
}

type iError interface {
	Error() string
}

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

func Trim(str string, characterMask ...string) string {
	trimChars := DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.Trim(str, trimChars)
}

func ContainsI(str, substr string) bool {
	return PosI(str, substr) != -1
}

func PosI(haystack, needle string, startOffset ...int) int {
	length := len(haystack)
	offset := 0
	if len(startOffset) > 0 {
		offset = startOffset[0]
	}
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(strings.ToLower(haystack[offset:]), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// String converts `any` to string.
// It's most commonly used converting function.
func String(any interface{}) string {
	if any == nil {
		return ""
	}
	switch value := any.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	default:
		// Empty checks.
		if value == nil {
			return ""
		}
		if f, ok := value.(iString); ok {
			// If the variable implements the String() interface,
			// then use that interface to perform the conversion
			return f.String()
		}
		if f, ok := value.(iError); ok {
			// If the variable implements the Error() interface,
			// then use that interface to perform the conversion
			return f.Error()
		}
		// Reflect checks.
		var (
			rv   = reflect.ValueOf(value)
			kind = rv.Kind()
		)
		switch kind {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return ""
			}
		case reflect.String:
			return rv.String()
		}
		if kind == reflect.Ptr {
			return String(rv.Elem().Interface())
		}
		// Finally, we use json.Marshal to convert.
		if jsonContent, err := json.Marshal(value); err != nil {
			return fmt.Sprint(value)
		} else {
			return string(jsonContent)
		}
	}
}
