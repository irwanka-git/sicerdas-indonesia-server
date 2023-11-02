// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSkorKejiwaanDewasa = "skor_kejiwaan_dewasa"

// SkorKejiwaanDewasa mapped from table <skor_kejiwaan_dewasa>
type SkorKejiwaanDewasa struct {
	IDUser           int32 `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz           int32 `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	SkorDepresi      int32 `gorm:"column:skor_depresi" json:"skor_depresi"`
	SkorKecemasan    int32 `gorm:"column:skor_kecemasan" json:"skor_kecemasan"`
	SkorManipulatif  int32 `gorm:"column:skor_manipulatif" json:"skor_manipulatif"`
	SkorParanoid     int32 `gorm:"column:skor_paranoid" json:"skor_paranoid"`
	SkorPsikopat     int32 `gorm:"column:skor_psikopat" json:"skor_psikopat"`
	SkorScizopernia  int32 `gorm:"column:skor_scizopernia" json:"skor_scizopernia"`
	SkorHisteria     int32 `gorm:"column:skor_histeria" json:"skor_histeria"`
	SkorHipokratis   int32 `gorm:"column:skor_hipokratis" json:"skor_hipokratis"`
	SkorHipomania    int32 `gorm:"column:skor_hipomania" json:"skor_hipomania"`
	SkorImpulsif     int32 `gorm:"column:skor_impulsif" json:"skor_impulsif"`
	NilaiDepresi     int32 `gorm:"column:nilai_depresi" json:"nilai_depresi"`
	NilaiKecemasan   int32 `gorm:"column:nilai_kecemasan" json:"nilai_kecemasan"`
	NilaiManipulatif int32 `gorm:"column:nilai_manipulatif" json:"nilai_manipulatif"`
	NilaiParanoid    int32 `gorm:"column:nilai_paranoid" json:"nilai_paranoid"`
	NilaiPsikopat    int32 `gorm:"column:nilai_psikopat" json:"nilai_psikopat"`
	NilaiScizopernia int32 `gorm:"column:nilai_scizopernia" json:"nilai_scizopernia"`
	NilaiHisteria    int32 `gorm:"column:nilai_histeria" json:"nilai_histeria"`
	NilaiHipokratis  int32 `gorm:"column:nilai_hipokratis" json:"nilai_hipokratis"`
	NilaiHipomania   int32 `gorm:"column:nilai_hipomania" json:"nilai_hipomania"`
	NilaiImpulsif    int32 `gorm:"column:nilai_impulsif" json:"nilai_impulsif"`
}

// TableName SkorKejiwaanDewasa's table name
func (*SkorKejiwaanDewasa) TableName() string {
	return TableNameSkorKejiwaanDewasa
}
