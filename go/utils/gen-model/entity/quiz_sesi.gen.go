// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameQuizSesi = "quiz_sesi"

// QuizSesi mapped from table <quiz_sesi>
type QuizSesi struct {
	IDQuiz               int32     `gorm:"column:id_quiz;primaryKey;autoIncrement:true" json:"id_quiz"`
	Token                string    `gorm:"column:token" json:"token"`
	NamaSesi             string    `gorm:"column:nama_sesi" json:"nama_sesi"`
	Lokasi               string    `gorm:"column:lokasi" json:"lokasi"`
	Tanggal              time.Time `gorm:"column:tanggal" json:"tanggal"`
	Open                 int32     `gorm:"column:open;default:1" json:"open"`
	Gambar               string    `gorm:"column:gambar" json:"gambar"`
	SkoringTabel         string    `gorm:"column:skoring_tabel" json:"skoring_tabel"`
	IDQuizTemplate       int32     `gorm:"column:id_quiz_template" json:"id_quiz_template"`
	IDUserBiro           int32     `gorm:"column:id_user_biro;comment:Id User BIRO" json:"id_user_biro"` // Id User BIRO
	Kota                 string    `gorm:"column:kota" json:"kota"`
	NamaAsesor           string    `gorm:"column:nama_asesor" json:"nama_asesor"`
	NomorSipp            string    `gorm:"column:nomor_sipp" json:"nomor_sipp"`
	UUID                 string    `gorm:"column:uuid" json:"uuid"`
	TtdAsesor            string    `gorm:"column:ttd_asesor" json:"ttd_asesor"`
	IDLokasi             int16     `gorm:"column:id_lokasi" json:"id_lokasi"`
	Jenis                string    `gorm:"column:jenis" json:"jenis"`
	FilenameReportZip    string    `gorm:"column:filename_report_zip;default:NULL" json:"filename_report_zip"`
	JSONURL              string    `gorm:"column:json_url" json:"json_url"`
	JSONURLEncrypt       string    `gorm:"column:json_url_encrypt" json:"json_url_encrypt"`
	CoverTemplate        string    `gorm:"column:cover_template;default:'default.pdf'" json:"cover_template"`
	FilenameReportZipDoc string    `gorm:"column:filename_report_zip_doc" json:"filename_report_zip_doc"`
	ModelReport          string    `gorm:"column:model_report" json:"model_report"`
}

// TableName QuizSesi's table name
func (*QuizSesi) TableName() string {
	return TableNameQuizSesi
}
