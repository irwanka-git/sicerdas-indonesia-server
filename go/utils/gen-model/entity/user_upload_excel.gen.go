// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameUserUploadExcel = "user_upload_excel"

// UserUploadExcel mapped from table <user_upload_excel>
type UserUploadExcel struct {
	IDUserUpload   int32  `gorm:"column:id_user_upload;primaryKey;autoIncrement:true" json:"id_user_upload"`
	Token          string `gorm:"column:token" json:"token"`
	Username       string `gorm:"column:username" json:"username"`
	NamaPengguna   string `gorm:"column:nama_pengguna" json:"nama_pengguna"`
	JenisKelamin   string `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	Organisasi     string `gorm:"column:organisasi" json:"organisasi"`
	UnitOrganisasi string `gorm:"column:unit_organisasi" json:"unit_organisasi"`
	Password       string `gorm:"column:password" json:"password"`
	Email          string `gorm:"column:email" json:"email"`
	Telp           string `gorm:"column:telp" json:"telp"`
	Valid          int32  `gorm:"column:valid;default:1" json:"valid"`
}

// TableName UserUploadExcel's table name
func (*UserUploadExcel) TableName() string {
	return TableNameUserUploadExcel
}
