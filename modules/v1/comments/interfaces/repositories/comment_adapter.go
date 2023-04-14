package repository

import (
	"mygram/modules/v1/comments/domain"

	"gorm.io/gorm"
)

type RepositoryPresenter interface {
	FindAllComments(idPhotos string) ([]domain.Comment, error)
	FindCommentById(id string) (domain.Comment, error)
}

type Repository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
