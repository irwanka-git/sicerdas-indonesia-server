// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSkorModeBelajar = "skor_mode_belajar"

// SkorModeBelajar mapped from table <skor_mode_belajar>
type SkorModeBelajar struct {
	IDUser        int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz        int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	IDModeBelajar int32  `gorm:"column:id_mode_belajar;primaryKey" json:"id_mode_belajar"`
	Prioritas1    string `gorm:"column:prioritas_1" json:"prioritas_1"`
	Prioritas2    string `gorm:"column:prioritas_2" json:"prioritas_2"`
	Prioritas3    string `gorm:"column:prioritas_3" json:"prioritas_3"`
	Prioritas4    string `gorm:"column:prioritas_4" json:"prioritas_4"`
	Prioritas5    string `gorm:"column:prioritas_5" json:"prioritas_5"`
}

// TableName SkorModeBelajar's table name
func (*SkorModeBelajar) TableName() string {
	return TableNameSkorModeBelajar
}
