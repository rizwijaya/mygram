package usecases

import (
	"mygram/modules/v1/users/domain"
	userRepository "mygram/modules/v1/users/interfaces/repositories"
)

type UserAdapter interface {
	RegisterUser(input domain.RegisterUserInput) (domain.User, error)
	LoginUser(input domain.LoginUserInput) (domain.User, error)
	GetUserByID(id int) (domain.User, error)
	AllSocialMedia() ([]domain.SocialMedia, error)
	OneSocialMedia(id string) (domain.SocialMedia, error)
	CreateSocialMedia(input domain.InsertSocialMedia, id int) (domain.CreatedSocialMedia, error)
	CheckSocialMedia(id int) error
}

type UserUseCase struct {
	repoUser *userRepository.Repository
}

func NewUserUseCase(repoUser *userRepository.Repository) *UserUseCase {
	return &UserUseCase{repoUser}
}
