package controller

import (
	"os"
)

type controller struct{}

func init() {
	os.Setenv("TZ", "Asia/Jakarta")
}
