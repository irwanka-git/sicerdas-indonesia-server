// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID             int32     `gorm:"column:id;primaryKey;default:nextval('mysequence_id_user'" json:"id"`
	Username       string    `gorm:"column:username" json:"username"`
	Password       string    `gorm:"column:password" json:"password"`
	NamaPengguna   string    `gorm:"column:nama_pengguna" json:"nama_pengguna"`
	JenisKelamin   string    `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	Organisasi     string    `gorm:"column:organisasi" json:"organisasi"`
	UnitOrganisasi string    `gorm:"column:unit_organisasi" json:"unit_organisasi"`
	Alamat         string    `gorm:"column:alamat" json:"alamat"`
	Email          string    `gorm:"column:email" json:"email"`
	Telp           string    `gorm:"column:telp" json:"telp"`
	KopBiro        string    `gorm:"column:kop_biro" json:"kop_biro"`
	RememberToken  string    `gorm:"column:remember_token" json:"remember_token"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
	UUID           string    `gorm:"column:uuid" json:"uuid"`
	Avatar         string    `gorm:"column:avatar" json:"avatar"`
	CreateBy       int32     `gorm:"column:create_by;comment:id USER yg buat akun" json:"create_by"` // id USER yg buat akun
	KunciUser      int32     `gorm:"column:kunci_user" json:"kunci_user"`
	BlokUser       int32     `gorm:"column:blok_user" json:"blok_user"`
	TokenUpload    string    `gorm:"column:token_upload" json:"token_upload"`
	DeviceID       string    `gorm:"column:device_id;comment:Device ID Login" json:"device_id"`  // Device ID Login
	CoverBiro      string    `gorm:"column:cover_biro;comment:Cover BIRO PDF" json:"cover_biro"` // Cover BIRO PDF
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}