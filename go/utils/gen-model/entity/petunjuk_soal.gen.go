// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNamePetunjukSoal = "petunjuk_soal"

// PetunjukSoal mapped from table <petunjuk_soal>
type PetunjukSoal struct {
	IDPetunjuk  int32  `gorm:"column:id_petunjuk;primaryKey;autoIncrement:true" json:"id_petunjuk"`
	IsiPetunjuk string `gorm:"column:isi_petunjuk" json:"isi_petunjuk"`
	UUID        string `gorm:"column:uuid" json:"uuid"`
	Keterangan  string `gorm:"column:keterangan" json:"keterangan"`
}

// TableName PetunjukSoal's table name
func (*PetunjukSoal) TableName() string {
	return TableNamePetunjukSoal
}