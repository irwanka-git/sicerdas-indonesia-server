// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSkorSuasanaKerja = "skor_suasana_kerja"

// SkorSuasanaKerja mapped from table <skor_suasana_kerja>
type SkorSuasanaKerja struct {
	IDUser        int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz        int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	SuasanaKerja1 string `gorm:"column:suasana_kerja1;default:0" json:"suasana_kerja1"`
	SuasanaKerja2 string `gorm:"column:suasana_kerja2;default:0" json:"suasana_kerja2"`
	SuasanaKerja3 string `gorm:"column:suasana_kerja3;default:0" json:"suasana_kerja3"`
}

// TableName SkorSuasanaKerja's table name
func (*SkorSuasanaKerja) TableName() string {
	return TableNameSkorSuasanaKerja
}