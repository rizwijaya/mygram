package repository

import (
	"mygram/modules/v1/comments/domain"

	"gorm.io/gorm/clause"
)

func (r *Repository) FindAllComments(idPhotos string, idUser int) ([]domain.Comment, error) {
	var comments []domain.Comment
	err := r.db.Where("photo_id = ? AND user_id = ?", idPhotos, idUser).Find(&comments).Error
	return comments, err
}

func (r *Repository) FindCommentById(id string) (domain.Comment, error) {
	var comment domain.Comment
	err := r.db.Where("id = ?", id).First(&comment).Error
	return comment, err
}

func (r *Repository) SaveComment(comment domain.Comment) (domain.Comment, error) {
	err := r.db.Create(&comment).Error
	return comment, err
}

func (r *Repository) UpdateComment(comment domain.Comment, id string) (domain.Comment, error) {
	err := r.db.Model(&comment).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&comment).Error
	return comment, err
}

func (r *Repository) FindPhotoById(id int) (domain.Photo, error) {
	var photo domain.Photo
	err := r.db.Preload("User").Preload("Comments").Where("id = ?", id).First(&photo).Error
	return photo, err
}
