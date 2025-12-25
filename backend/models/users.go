package models

import "time"

type Users struct {
	ID        int64     `gorm:"primary_key;auto_increment" column:"id"`
	Email     string    `gorm:"column:email"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
