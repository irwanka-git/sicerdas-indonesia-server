// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSoalKejiwaanDewasa = "soal_kejiwaan_dewasa"

// SoalKejiwaanDewasa mapped from table <soal_kejiwaan_dewasa>
type SoalKejiwaanDewasa struct {
	IDSoal  int16  `gorm:"column:id_soal;primaryKey" json:"id_soal"`
	IDModel int16  `gorm:"column:id_model" json:"id_model"`
	Unsur   string `gorm:"column:unsur" json:"unsur"`
	Urutan  int16  `gorm:"column:urutan" json:"urutan"`
}

// TableName SoalKejiwaanDewasa's table name
func (*SoalKejiwaanDewasa) TableName() string {
	return TableNameSoalKejiwaanDewasa
}