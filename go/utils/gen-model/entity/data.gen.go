// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameDatum = "data"

// Datum mapped from table <data>
type Datum struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Nama       string `gorm:"column:nama" json:"nama"`
	Keterangan string `gorm:"column:keterangan" json:"keterangan"`
}

// TableName Datum's table name
func (*Datum) TableName() string {
	return TableNameDatum
}
