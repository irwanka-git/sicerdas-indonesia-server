// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameRole = "roles"

// Role mapped from table <roles>
type Role struct {
	IDRole    int64     `gorm:"column:id_role" json:"id_role"`
	NamaRole  string    `gorm:"column:nama_role" json:"nama_role"`
	UUID      string    `gorm:"column:uuid" json:"uuid"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}