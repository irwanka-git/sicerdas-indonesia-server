// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameRefSkalaKdInformasiUmum = "ref_skala_kd_informasi_umum"

// RefSkalaKdInformasiUmum mapped from table <ref_skala_kd_informasi_umum>
type RefSkalaKdInformasiUmum struct {
	ID          int32  `gorm:"column:id;primaryKey" json:"id"`
	Skor        int32  `gorm:"column:skor;not null" json:"skor"`
	Klasifikasi string `gorm:"column:klasifikasi;not null" json:"klasifikasi"`
}

// TableName RefSkalaKdInformasiUmum's table name
func (*RefSkalaKdInformasiUmum) TableName() string {
	return TableNameRefSkalaKdInformasiUmum
}
