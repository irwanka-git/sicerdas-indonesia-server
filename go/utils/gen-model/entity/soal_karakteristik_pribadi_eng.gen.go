// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSoalKarakteristikPribadiEng = "soal_karakteristik_pribadi_eng"

// SoalKarakteristikPribadiEng mapped from table <soal_karakteristik_pribadi_eng>
type SoalKarakteristikPribadiEng struct {
	IDSoal     int32  `gorm:"column:id_soal;primaryKey;autoIncrement:true" json:"id_soal"`
	IDKomponen int32  `gorm:"column:id_komponen" json:"id_komponen"`
	Urutan     int32  `gorm:"column:urutan" json:"urutan"`
	Pernyataan string `gorm:"column:pernyataan" json:"pernyataan"`
	UUID       string `gorm:"column:uuid" json:"uuid"`
	Pilihan1   string `gorm:"column:pilihan_1;default:'Tidak Sesuai'" json:"pilihan_1"`
	Pilihan2   string `gorm:"column:pilihan_2;default:'Agak Tidak Sesuai'" json:"pilihan_2"`
	Pilihan3   string `gorm:"column:pilihan_3;default:'Agak Sesuai'" json:"pilihan_3"`
	Pilihan4   string `gorm:"column:pilihan_4;default:'Sesuai'" json:"pilihan_4"`
}

// TableName SoalKarakteristikPribadiEng's table name
func (*SoalKarakteristikPribadiEng) TableName() string {
	return TableNameSoalKarakteristikPribadiEng
}
