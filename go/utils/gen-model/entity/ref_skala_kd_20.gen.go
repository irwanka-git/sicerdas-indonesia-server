// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameRefSkalaKd20 = "ref_skala_kd_20"

// RefSkalaKd20 mapped from table <ref_skala_kd_20>
type RefSkalaKd20 struct {
	ID          int32  `gorm:"column:id;primaryKey" json:"id"`
	Skor        int32  `gorm:"column:skor;not null" json:"skor"`
	Klasifikasi string `gorm:"column:klasifikasi;not null" json:"klasifikasi"`
}

// TableName RefSkalaKd20's table name
func (*RefSkalaKd20) TableName() string {
	return TableNameRefSkalaKd20
}
