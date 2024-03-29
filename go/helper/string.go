package helper

import (
	"strings"
	"unicode"

	"github.com/google/uuid"
)

func GenUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func CleanString(data string) string {
	data = strings.ReplaceAll(data, "`", "")
	data = strings.ReplaceAll(data, "'", "")
	data = strings.ReplaceAll(data, "\"", "")
	data = strings.ReplaceAll(data, "*", "")
	data = strings.ReplaceAll(data, "^", "")
	return data
}

func CleanParameterIDOnly(data string) string {
	data = strings.ReplaceAll(data, " ", "")
	data = strings.ReplaceAll(data, "`", "")
	data = strings.ReplaceAll(data, "'", "")
	data = strings.ReplaceAll(data, "\"", "")
	data = strings.ReplaceAll(data, "*", "")
	data = strings.ReplaceAll(data, "^", "")
	return data
}

func CleanNamaFileOnly(data string) string {
	data = strings.ReplaceAll(data, "  ", " ")
	data = strings.ReplaceAll(data, " ", "-")
	data = strings.ReplaceAll(data, "'", "")
	data = strings.ReplaceAll(data, ".", "")
	data = strings.ReplaceAll(data, "*", "")
	data = strings.ReplaceAll(data, "^", "")
	data = strings.ReplaceAll(data, "`", "")
	data = strings.ReplaceAll(data, ",", "")
	data = strings.ReplaceAll(data, "&", "")
	data = strings.ReplaceAll(data, "@", "")
	data = strings.ReplaceAll(data, "!", "")
	data = strings.ReplaceAll(data, "$", "")
	data = strings.ReplaceAll(data, ":", "")
	data = strings.ReplaceAll(data, "\"", "")
	data = strings.ReplaceAll(data, "?", "")
	data = strings.ReplaceAll(data, ">", "")
	data = strings.ReplaceAll(data, ")", "")
	data = strings.ReplaceAll(data, "(", "")
	data = strings.ReplaceAll(data, "}", "")
	data = strings.ReplaceAll(data, "{", "")
	data = strings.ReplaceAll(data, "=", "")
	data = strings.ReplaceAll(data, "+", "")
	data = strings.ReplaceAll(data, "_", "")
	return data
}

func GetPilihanFromTagDot(data string, index int) string {
	arr := strings.Split(data, ":")
	if len(arr) >= index {
		return arr[index-1]
	}
	return ""
}

func Capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func IsContainsInArrayString(sl []string, name string) bool {
	// iterate over the array and compare given string to each element
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}
