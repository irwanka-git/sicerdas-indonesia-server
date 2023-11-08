package helper

import (
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GetArrFilename(direktory string) []string {
	var list []string
	files, err := os.ReadDir(direktory)
	if err != nil {

		return list
	}

	for _, file := range files {
		if !file.IsDir() {
			list = append(list, direktory+"/"+file.Name())
		}
	}
	return list
}
