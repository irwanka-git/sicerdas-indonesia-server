// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameInfoCerda = "info_cerdas"

// InfoCerda mapped from table <info_cerdas>
type InfoCerda struct {
	IDInfo    int32     `gorm:"column:id_info;primaryKey;autoIncrement:true" json:"id_info"`
	Judul     string    `gorm:"column:judul" json:"judul"`
	Isi       string    `gorm:"column:isi" json:"isi"`
	Gambar    string    `gorm:"column:gambar" json:"gambar"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UUID      string    `gorm:"column:uuid" json:"uuid"`
	URL       string    `gorm:"column:url" json:"url"`
}

// TableName InfoCerda's table name
func (*InfoCerda) TableName() string {
	return TableNameInfoCerda
}