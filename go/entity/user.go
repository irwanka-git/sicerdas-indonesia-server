package entity

import "time"

type User struct {
	ID             int32     `json:"id"`
	Username       string    `json:"username"`
	Password       string    `json:"-"`
	NamaPengguna   string    `json:"nama_pengguna"`
	JenisKelamin   string    `json:"jenis_kelamin"`
	Organisasi     string    `json:"organisasi"`
	UnitOrganisasi string    `json:"unit_organisasi"`
	Alamat         string    `json:"alamat"`
	Email          string    `json:"email"`
	Telp           string    `json:"telp"`
	KopBiro        string    `json:"kop_biro"`
	RememberToken  string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	UUID           string    `json:"uuid"`
	Avatar         string    `json:"avatar"`
	CreateBy       int32     `json:"create_by"` // id USER yg buat akun
	KunciUser      int32     `json:"kunci_user"`
	BlokUser       int32     `json:"blok_user"`
	TokenUpload    string    `json:"-"`
	DeviceID       string    `json:"device_id"`  // Device ID Login
	CoverBiro      string    `json:"cover_biro"` // Cover BIRO PDF
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Ref      string `json:"ref"`
}
