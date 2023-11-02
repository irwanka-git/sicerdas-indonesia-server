package helper

import (
	"time"
)

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

func StringTimeYMDHIS(waktu time.Time) string {
	return waktu.Format(DDMMYYYYhhmmss)
}
