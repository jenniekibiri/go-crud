package models
 import (
	"time"
 )
 type Person struct {
	ID int32 `gorm:"primarykey;column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Address string `gorm:"column:address" json:"address"`
	Phone uint `gorm:"column:phone" json:"phone"`
	CreatedAt  time.Time `gorm:"colum:createdAt;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;autoUpdateTime" json:"updatedAt"`
 }

 func(Person) TableName() string{
	return "person"
 }