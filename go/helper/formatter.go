package helper

import (
	"fmt"
	"time"
)

var months = [...]string{
	"Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "November", "Desember",
}

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
	DDMMYYYYIndo   = "02 January 2006"
)

func StringTimeYMDHIS(waktu time.Time) string {
	return waktu.Format(DDMMYYYYhhmmss)
}

func StringTimeTglIndo(t time.Time) string {
	return fmt.Sprintf("%02d %s %v", t.Day(), months[t.Month()-1], t.Year())
}
