// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSkorKesehatanMental = "skor_kesehatan_mental"

// SkorKesehatanMental mapped from table <skor_kesehatan_mental>
type SkorKesehatanMental struct {
	IDUser               int32 `gorm:"column:id_user;not null" json:"id_user"`
	IDQuiz               int32 `gorm:"column:id_quiz;not null" json:"id_quiz"`
	SkorStressKehidupan  int32 `gorm:"column:skor_stress_kehidupan" json:"skor_stress_kehidupan"`
	SkorOverThinking     int32 `gorm:"column:skor_over_thinking" json:"skor_over_thinking"`
	SkorAdiksiMedsos     int32 `gorm:"column:skor_adiksi_medsos" json:"skor_adiksi_medsos"`
	SkorImpulsiveBuying  int32 `gorm:"column:skor_impulsive_buying" json:"skor_impulsive_buying"`
	SkorGangguanMood     int32 `gorm:"column:skor_gangguan_mood" json:"skor_gangguan_mood"`
	SkorGangguanMakan    int32 `gorm:"column:skor_gangguan_makan" json:"skor_gangguan_makan"`
	SkorPenampilanTubuh  int32 `gorm:"column:skor_penampilan_tubuh" json:"skor_penampilan_tubuh"`
	SkorKecemasanBicara  int32 `gorm:"column:skor_kecemasan_bicara" json:"skor_kecemasan_bicara"`
	SkorPanicAttack      int32 `gorm:"column:skor_panic_attack" json:"skor_panic_attack"`
	SkorBipolarDisorder  int32 `gorm:"column:skor_bipolar_disorder" json:"skor_bipolar_disorder"`
	SkorAdiksiZat        int32 `gorm:"column:skor_adiksi_zat" json:"skor_adiksi_zat"`
	NilaiStressKehidupan int32 `gorm:"column:nilai_stress_kehidupan" json:"nilai_stress_kehidupan"`
	NilaiOverThinking    int32 `gorm:"column:nilai_over_thinking" json:"nilai_over_thinking"`
	NilaiAdiksiMedsos    int32 `gorm:"column:nilai_adiksi_medsos" json:"nilai_adiksi_medsos"`
	NilaiImpulsiveBuying int32 `gorm:"column:nilai_impulsive_buying" json:"nilai_impulsive_buying"`
	NilaiGangguanMood    int32 `gorm:"column:nilai_gangguan_mood" json:"nilai_gangguan_mood"`
	NilaiGangguanMakan   int32 `gorm:"column:nilai_gangguan_makan" json:"nilai_gangguan_makan"`
	NilaiPenampilanTubuh int32 `gorm:"column:nilai_penampilan_tubuh" json:"nilai_penampilan_tubuh"`
	NilaiKecemasanBicara int32 `gorm:"column:nilai_kecemasan_bicara" json:"nilai_kecemasan_bicara"`
	NilaiPanicAttack     int32 `gorm:"column:nilai_panic_attack" json:"nilai_panic_attack"`
	NilaiBipolarDisorder int32 `gorm:"column:nilai_bipolar_disorder" json:"nilai_bipolar_disorder"`
	NilaiAdiksiZat       int32 `gorm:"column:nilai_adiksi_zat" json:"nilai_adiksi_zat"`
}

// TableName SkorKesehatanMental's table name
func (*SkorKesehatanMental) TableName() string {
	return TableNameSkorKesehatanMental
}