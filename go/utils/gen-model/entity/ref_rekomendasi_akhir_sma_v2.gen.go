// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameRefRekomendasiAkhirSmaV2 = "ref_rekomendasi_akhir_sma_v2"

// RefRekomendasiAkhirSmaV2 mapped from table <ref_rekomendasi_akhir_sma_v2>
type RefRekomendasiAkhirSmaV2 struct {
	ID                  int32  `gorm:"column:id;primaryKey" json:"id"`
	RekomMinat          string `gorm:"column:rekom_minat" json:"rekom_minat"`
	RekomSikapPelajaran string `gorm:"column:rekom_sikap_pelajaran" json:"rekom_sikap_pelajaran"`
	RekomAkhir          string `gorm:"column:rekom_akhir" json:"rekom_akhir"`
}

// TableName RefRekomendasiAkhirSmaV2's table name
func (*RefRekomendasiAkhirSmaV2) TableName() string {
	return TableNameRefRekomendasiAkhirSmaV2
}
