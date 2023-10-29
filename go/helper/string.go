package helper

import (
	"strings"

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

func GetPilihanFromTagDot(data string, index int) string {
	arr := strings.Split(data, ":")
	if len(arr) >= index {
		return arr[index-1]
	}
	return ""
}
