// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameStatusSkoring = "status_skoring"

// StatusSkoring mapped from table <status_skoring>
type StatusSkoring struct {
	ID      int32  `gorm:"column:id;primaryKey" json:"id"`
	Status  int16  `gorm:"column:status" json:"status"`
	Mulai   string `gorm:"column:mulai" json:"mulai"`
	Selesai string `gorm:"column:selesai" json:"selesai"`
	Jumlah  int32  `gorm:"column:jumlah" json:"jumlah"`
}

// TableName StatusSkoring's table name
func (*StatusSkoring) TableName() string {
	return TableNameStatusSkoring
}
