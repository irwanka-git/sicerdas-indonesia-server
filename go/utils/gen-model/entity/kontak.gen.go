// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameKontak = "kontak"

// Kontak mapped from table <kontak>
type Kontak struct {
	IDKontak int32  `gorm:"column:id_kontak;primaryKey;autoIncrement:true" json:"id_kontak"`
	Telepon  string `gorm:"column:telepon" json:"telepon"`
	Email    string `gorm:"column:email" json:"email"`
	Website  string `gorm:"column:website" json:"website"`
	UUID     string `gorm:"column:uuid" json:"uuid"`
	WaMe     string `gorm:"column:wa_me" json:"wa_me"`
}

// TableName Kontak's table name
func (*Kontak) TableName() string {
	return TableNameKontak
}
