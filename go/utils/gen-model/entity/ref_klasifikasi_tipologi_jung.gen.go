// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameRefKlasifikasiTipologiJung = "ref_klasifikasi_tipologi_jung"

// RefKlasifikasiTipologiJung mapped from table <ref_klasifikasi_tipologi_jung>
type RefKlasifikasiTipologiJung struct {
	ID            int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Kode          string `gorm:"column:kode;not null;default:0" json:"kode"`
	Klasifikasi   string `gorm:"column:klasifikasi;not null;default:0" json:"klasifikasi"`
	FieldSkoringA string `gorm:"column:field_skoring_a" json:"field_skoring_a"`
	FieldSkoringB string `gorm:"column:field_skoring_b" json:"field_skoring_b"`
}

// TableName RefKlasifikasiTipologiJung's table name
func (*RefKlasifikasiTipologiJung) TableName() string {
	return TableNameRefKlasifikasiTipologiJung
}
