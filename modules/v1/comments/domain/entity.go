package domain

import "time"

type GormModel struct {
	ID        int        `json:"id" gorm:"column:id"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type Comment struct {
	GormModel
	UserID  int    `json:"user_id" gorm:"column:user_id"`
	PhotoID int    `json:"photo_id" gorm:"column:photo_id"`
	Message string `json:"message" gorm:"column:message"`
}
