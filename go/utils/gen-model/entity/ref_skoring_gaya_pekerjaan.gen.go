// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameRefSkoringGayaPekerjaan = "ref_skoring_gaya_pekerjaan"

// RefSkoringGayaPekerjaan mapped from table <ref_skoring_gaya_pekerjaan>
type RefSkoringGayaPekerjaan struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	IDUser      int32  `gorm:"column:id_user" json:"id_user"`
	IDQuiz      int32  `gorm:"column:id_quiz" json:"id_quiz"`
	Kode        string `gorm:"column:kode" json:"kode"`
	Skor        int16  `gorm:"column:skor" json:"skor"`
	Rangking    int16  `gorm:"column:rangking" json:"rangking"`
	Klasifikasi string `gorm:"column:klasifikasi" json:"klasifikasi"`
}

// TableName RefSkoringGayaPekerjaan's table name
func (*RefSkoringGayaPekerjaan) TableName() string {
	return TableNameRefSkoringGayaPekerjaan
}
