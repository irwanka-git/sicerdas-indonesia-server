// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameRefSkorKarakterPribadi = "ref_skor_karakter_pribadi"

// RefSkorKarakterPribadi mapped from table <ref_skor_karakter_pribadi>
type RefSkorKarakterPribadi struct {
	ID      int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Respon  string `gorm:"column:respon;not null" json:"respon"`
	Skor    int32  `gorm:"column:skor;not null" json:"skor"`
	Jawaban string `gorm:"column:jawaban" json:"jawaban"`
}

// TableName RefSkorKarakterPribadi's table name
func (*RefSkorKarakterPribadi) TableName() string {
	return TableNameRefSkorKarakterPribadi
}
