// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSkorDisc = "skor_disc"

// SkorDisc mapped from table <skor_disc>
type SkorDisc struct {
	IDUser       int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz       int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	SkorD        int32  `gorm:"column:skor_d" json:"skor_d"`
	SkorI        int32  `gorm:"column:skor_i" json:"skor_i"`
	SkorS        int32  `gorm:"column:skor_s" json:"skor_s"`
	SkorC        int32  `gorm:"column:skor_c" json:"skor_c"`
	KlasifikasiD string `gorm:"column:klasifikasi_d" json:"klasifikasi_d"`
	KlasifikasiS string `gorm:"column:klasifikasi_s" json:"klasifikasi_s"`
	KlasifikasiI string `gorm:"column:klasifikasi_i" json:"klasifikasi_i"`
	KlasifikasiC string `gorm:"column:klasifikasi_c" json:"klasifikasi_c"`
}

// TableName SkorDisc's table name
func (*SkorDisc) TableName() string {
	return TableNameSkorDisc
}