// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameStatusCronjob = "status_cronjob"

// StatusCronjob mapped from table <status_cronjob>
type StatusCronjob struct {
	ID     int32 `gorm:"column:id;primaryKey" json:"id"`
	Status int16 `gorm:"column:status" json:"status"`
}

// TableName StatusCronjob's table name
func (*StatusCronjob) TableName() string {
	return TableNameStatusCronjob
}