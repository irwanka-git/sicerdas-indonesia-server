// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSkorKognitifPmk = "skor_kognitif_pmk"

// SkorKognitifPmk mapped from table <skor_kognitif_pmk>
type SkorKognitifPmk struct {
	ID            int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	IDUser        int32   `gorm:"column:id_user" json:"id_user"`
	IDQuiz        int32   `gorm:"column:id_quiz" json:"id_quiz"`
	TpaIu         int32   `gorm:"column:tpa_iu" json:"tpa_iu"`
	TpaPv         int32   `gorm:"column:tpa_pv" json:"tpa_pv"`
	TpaPk         int32   `gorm:"column:tpa_pk" json:"tpa_pk"`
	TpaPa         int32   `gorm:"column:tpa_pa" json:"tpa_pa"`
	TpaPs         int32   `gorm:"column:tpa_ps" json:"tpa_ps"`
	TpaPm         int32   `gorm:"column:tpa_pm" json:"tpa_pm"`
	TpaKt         int32   `gorm:"column:tpa_kt" json:"tpa_kt"`
	TpaIq         int32   `gorm:"column:tpa_iq" json:"tpa_iq"`
	SkorIq        float32 `gorm:"column:skor_iq" json:"skor_iq"`
	KlasifikasiPv string  `gorm:"column:klasifikasi_pv" json:"klasifikasi_pv"`
	KlasifikasiPk string  `gorm:"column:klasifikasi_pk" json:"klasifikasi_pk"`
	KlasifikasiPa string  `gorm:"column:klasifikasi_pa" json:"klasifikasi_pa"`
	KlasifikasiPs string  `gorm:"column:klasifikasi_ps" json:"klasifikasi_ps"`
	KlasifikasiPm string  `gorm:"column:klasifikasi_pm" json:"klasifikasi_pm"`
	KlasifikasiKt string  `gorm:"column:klasifikasi_kt" json:"klasifikasi_kt"`
	KlasifikasiIq string  `gorm:"column:klasifikasi_iq" json:"klasifikasi_iq"`
	KlasifikasiIu string  `gorm:"column:klasifikasi_iu" json:"klasifikasi_iu"`
}

// TableName SkorKognitifPmk's table name
func (*SkorKognitifPmk) TableName() string {
	return TableNameSkorKognitifPmk
}
