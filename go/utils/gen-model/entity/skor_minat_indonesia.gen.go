// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSkorMinatIndonesium = "skor_minat_indonesia"

// SkorMinatIndonesium mapped from table <skor_minat_indonesia>
type SkorMinatIndonesium struct {
	IDUser        int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz        int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	TmiIlmuAlam   int32  `gorm:"column:tmi_ilmu_alam" json:"tmi_ilmu_alam"`
	TmiIlmuSosial int32  `gorm:"column:tmi_ilmu_sosial" json:"tmi_ilmu_sosial"`
	TmiRentang    int32  `gorm:"column:tmi_rentang" json:"tmi_rentang"`
	RekomTmi      string `gorm:"column:rekom_tmi" json:"rekom_tmi"`
}

// TableName SkorMinatIndonesium's table name
func (*SkorMinatIndonesium) TableName() string {
	return TableNameSkorMinatIndonesium
}
