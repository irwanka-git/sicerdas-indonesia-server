// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameRefKelasMinatSma = "ref_kelas_minat_sma"

// RefKelasMinatSma mapped from table <ref_kelas_minat_sma>
type RefKelasMinatSma struct {
	ID              int16  `gorm:"column:id;primaryKey" json:"id"`
	Kelas           string `gorm:"column:kelas" json:"kelas"`
	KeteranganMinat string `gorm:"column:keterangan_minat" json:"keterangan_minat"`
	MapelUtama      string `gorm:"column:mapel_utama" json:"mapel_utama"`
	MapelTambahan   string `gorm:"column:mapel_tambahan" json:"mapel_tambahan"`
}

// TableName RefKelasMinatSma's table name
func (*RefKelasMinatSma) TableName() string {
	return TableNameRefKelasMinatSma
}
