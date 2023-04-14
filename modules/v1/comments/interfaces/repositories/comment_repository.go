package repository

import "mygram/modules/v1/comments/domain"

func (r *Repository) FindAllComments(idPhotos string) ([]domain.Comment, error) {
	var comments []domain.Comment
	err := r.db.Where("photo_id = ?", idPhotos).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *Repository) FindCommentById(id string) (domain.Comment, error) {
	var comment domain.Comment
	err := r.db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return domain.Comment{}, err
	}
	return comment, nil
}
