// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameRefKecerdasanMajemuk = "ref_kecerdasan_majemuk"

// RefKecerdasanMajemuk mapped from table <ref_kecerdasan_majemuk>
type RefKecerdasanMajemuk struct {
	No             string `gorm:"column:no;primaryKey" json:"no"`
	NamaKecerdasan string `gorm:"column:nama_kecerdasan" json:"nama_kecerdasan"`
	NamaKecil      string `gorm:"column:nama_kecil" json:"nama_kecil"`
	Icon           string `gorm:"column:icon" json:"icon"`
}

// TableName RefKecerdasanMajemuk's table name
func (*RefKecerdasanMajemuk) TableName() string {
	return TableNameRefKecerdasanMajemuk
}
