// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameRefSkoringKecerdasanMajemuk = "ref_skoring_kecerdasan_majemuk"

// RefSkoringKecerdasanMajemuk mapped from table <ref_skoring_kecerdasan_majemuk>
type RefSkoringKecerdasanMajemuk struct {
	IDSkoringKecerdasanMajemuk int32  `gorm:"column:id_skoring_kecerdasan_majemuk;primaryKey;autoIncrement:true" json:"id_skoring_kecerdasan_majemuk"`
	IDQuiz                     int32  `gorm:"column:id_quiz;comment:ID_QUIZ" json:"id_quiz"`                    // ID_QUIZ
	IDUser                     int32  `gorm:"column:id_user;comment:ID_USER" json:"id_user"`                    // ID_USER
	No                         string `gorm:"column:no;comment:NO SEKOLAH DINAS (ref_sekolah_dinas)" json:"no"` // NO SEKOLAH DINAS (ref_sekolah_dinas)
	B1                         int32  `gorm:"column:b1" json:"b1"`
	B2                         int32  `gorm:"column:b2" json:"b2"`
	B3                         int32  `gorm:"column:b3" json:"b3"`
	B4                         int32  `gorm:"column:b4" json:"b4"`
	B5                         int32  `gorm:"column:b5" json:"b5"`
	B6                         int32  `gorm:"column:b6" json:"b6"`
	B7                         int32  `gorm:"column:b7" json:"b7"`
	B8                         int32  `gorm:"column:b8" json:"b8"`
	B9                         int32  `gorm:"column:b9" json:"b9"`
	Total                      int32  `gorm:"column:total" json:"total"`
	Rangking                   int32  `gorm:"column:rangking" json:"rangking"`
}

// TableName RefSkoringKecerdasanMajemuk's table name
func (*RefSkoringKecerdasanMajemuk) TableName() string {
	return TableNameRefSkoringKecerdasanMajemuk
}
