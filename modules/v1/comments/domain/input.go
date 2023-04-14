package domain

type InsertComment struct {
	UserID  int    `json:"_" gorm:"column:user_id"`
	PhotoID int    `json:"photo_id" gorm:"column:photo_id" binding:"required,number"`
	Message string `json:"message" gorm:"column:message" binding:"required"`
}
