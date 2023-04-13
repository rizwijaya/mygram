package repository

import (
	"gorm.io/gorm"
)

type RepositoryPresenter interface {
}

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}