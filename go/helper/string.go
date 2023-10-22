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
