package repository

import (
	"mygram/modules/v1/users/domain"

	"gorm.io/gorm"
)

type RepositoryPresenter interface {
	SaveUsers(user domain.User) (domain.User, error)
	FindUser(field string, value string) (domain.User, error)
	FindUserByID(id int) (domain.User, error)
	AllSocialMedia() ([]domain.SocialMedia, error)
	FindSocialMediaByID(id string) (domain.SocialMedia, error)
	SaveSocailMedia(socialMedia domain.SocialMedia) (domain.SocialMedia, error)
	FindSocialMediaByUserID(id int) (domain.SocialMedia, error)
	UpdateSocialMedia(socialMedia domain.SocialMedia, id int) (domain.SocialMedia, error)
}

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
