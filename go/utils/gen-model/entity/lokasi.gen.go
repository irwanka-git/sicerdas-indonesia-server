// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameLokasi = "lokasi"

// Lokasi mapped from table <lokasi>
type Lokasi struct {
	IDLokasi      int16  `gorm:"column:id_lokasi;primaryKey;autoIncrement:true" json:"id_lokasi"`
	NamaLokasi    string `gorm:"column:nama_lokasi" json:"nama_lokasi"`
	UUID          string `gorm:"column:uuid" json:"uuid"`
	KodeProvinsi  string `gorm:"column:kode_provinsi" json:"kode_provinsi"`
	KodeKabupaten string `gorm:"column:kode_kabupaten" json:"kode_kabupaten"`
}

// TableName Lokasi's table name
func (*Lokasi) TableName() string {
	return TableNameLokasi
}
