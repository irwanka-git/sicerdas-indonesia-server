// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameRefRekomendasiSikapPelajaran = "ref_rekomendasi_sikap_pelajaran"

// RefRekomendasiSikapPelajaran mapped from table <ref_rekomendasi_sikap_pelajaran>
type RefRekomendasiSikapPelajaran struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Perbedaan   int32  `gorm:"column:perbedaan;not null;comment:PERBEDAAN = (MAT+FIS+BIO ) - (EKO+SEJ+GEO)" json:"perbedaan"` // PERBEDAAN = (MAT+FIS+BIO ) - (EKO+SEJ+GEO)
	Rekomendasi string `gorm:"column:rekomendasi;not null;default:0" json:"rekomendasi"`
}

// TableName RefRekomendasiSikapPelajaran's table name
func (*RefRekomendasiSikapPelajaran) TableName() string {
	return TableNameRefRekomendasiSikapPelajaran
}