// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameRefSkalaSikapPelajaran = "ref_skala_sikap_pelajaran"

// RefSkalaSikapPelajaran mapped from table <ref_skala_sikap_pelajaran>
type RefSkalaSikapPelajaran struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Skor        int32  `gorm:"column:skor" json:"skor"`
	Klasifikasi string `gorm:"column:klasifikasi" json:"klasifikasi"`
}

// TableName RefSkalaSikapPelajaran's table name
func (*RefSkalaSikapPelajaran) TableName() string {
	return TableNameRefSkalaSikapPelajaran
}
