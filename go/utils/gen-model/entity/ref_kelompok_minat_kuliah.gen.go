// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameRefKelompokMinatKuliah = "ref_kelompok_minat_kuliah"

// RefKelompokMinatKuliah mapped from table <ref_kelompok_minat_kuliah>
type RefKelompokMinatKuliah struct {
	ID              int16  `gorm:"column:id;primaryKey" json:"id"`
	Kelompok        string `gorm:"column:kelompok" json:"kelompok"`
	KeteranganMinat string `gorm:"column:keterangan_minat" json:"keterangan_minat"`
}

// TableName RefKelompokMinatKuliah's table name
func (*RefKelompokMinatKuliah) TableName() string {
	return TableNameRefKelompokMinatKuliah
}
