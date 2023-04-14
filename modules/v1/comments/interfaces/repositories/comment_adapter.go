package repository

import (
	"mygram/modules/v1/comments/domain"

	"gorm.io/gorm"
)

type RepositoryPresenter interface {
	FindAllComments(idPhotos string, idUser int) ([]domain.Comment, error)
	FindCommentById(id string) (domain.Comment, error)
	SaveComment(comment domain.Comment) (domain.Comment, error)
	UpdateComment(comment domain.Comment, id string) (domain.Comment, error)
	DeleteComment(id string) error
	FindPhotoById(id int) (domain.Photo, error)
}

type Repository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
