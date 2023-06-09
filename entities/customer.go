package entities

import "time"

type Customer struct {
	Id         uint   `gorm:"primary_key"`
	First_name string `gorm:"column:first_name"`
	Last_name  string `gorm:"column:last_name"`
	Email      string `gorm:"column:email"`
	Avatar     string `gorm:"column:avatar"`
	Created_at time.Time
	Updated_at time.Time
}
