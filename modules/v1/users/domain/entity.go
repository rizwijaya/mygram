package domain

import "time"

type GormModel struct {
	ID        int        `gorm:"column:id"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

type User struct {
	GormModel
	UserName string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Age      int    `gorm:"column:age"`
}
