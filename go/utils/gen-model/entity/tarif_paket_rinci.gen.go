// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameTarifPaketRinci = "tarif_paket_rinci"

// TarifPaketRinci mapped from table <tarif_paket_rinci>
type TarifPaketRinci struct {
	IDTarifRinci int32  `gorm:"column:id_tarif_rinci;not null" json:"id_tarif_rinci"`
	IDTarif      int64  `gorm:"column:id_tarif" json:"id_tarif"`
	NamaRincian  string `gorm:"column:nama_rincian" json:"nama_rincian"`
	Urutan       int32  `gorm:"column:urutan" json:"urutan"`
	UUID         string `gorm:"column:uuid" json:"uuid"`
}

// TableName TarifPaketRinci's table name
func (*TarifPaketRinci) TableName() string {
	return TableNameTarifPaketRinci
}