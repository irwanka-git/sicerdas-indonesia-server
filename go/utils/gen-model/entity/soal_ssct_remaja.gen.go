// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSoalSsctRemaja = "soal_ssct_remaja"

// SoalSsctRemaja mapped from table <soal_ssct_remaja>
type SoalSsctRemaja struct {
	IDSoal          int32  `gorm:"column:id_soal;primaryKey;default:nextval('serial_id_soal_ssct_remaja'" json:"id_soal"`
	Urutan          int32  `gorm:"column:urutan" json:"urutan"`
	SubjekPenilaian string `gorm:"column:subjek_penilaian" json:"subjek_penilaian"`
	SikapNegatif1   string `gorm:"column:sikap_negatif1" json:"sikap_negatif1"`
	SikapPositif1   string `gorm:"column:sikap_positif1" json:"sikap_positif1"`
	SikapNegatif2   string `gorm:"column:sikap_negatif2" json:"sikap_negatif2"`
	SikapPositif2   string `gorm:"column:sikap_positif2" json:"sikap_positif2"`
	SikapNegatif3   string `gorm:"column:sikap_negatif3" json:"sikap_negatif3"`
	SikapPositif3   string `gorm:"column:sikap_positif3" json:"sikap_positif3"`
	Aspek           string `gorm:"column:aspek" json:"aspek"`
	UUID            string `gorm:"column:uuid" json:"uuid"`
	Komponen        string `gorm:"column:komponen" json:"komponen"`
}

// TableName SoalSsctRemaja's table name
func (*SoalSsctRemaja) TableName() string {
	return TableNameSoalSsctRemaja
}
