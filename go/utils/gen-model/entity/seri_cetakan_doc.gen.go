// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameSeriCetakanDoc = "seri_cetakan_doc"

// SeriCetakanDoc mapped from table <seri_cetakan_doc>
type SeriCetakanDoc struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NoSeri    string    `gorm:"column:no_seri" json:"no_seri"`
	Token     string    `gorm:"column:token" json:"token"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	IDUser    int32     `gorm:"column:id_user" json:"id_user"`
	URL       string    `gorm:"column:url" json:"url"`
	JenisTes  string    `gorm:"column:jenis_tes" json:"jenis_tes"`
	Filename  string    `gorm:"column:filename" json:"filename"`
	Pathname  string    `gorm:"column:pathname" json:"pathname"`
}

// TableName SeriCetakanDoc's table name
func (*SeriCetakanDoc) TableName() string {
	return TableNameSeriCetakanDoc
}
