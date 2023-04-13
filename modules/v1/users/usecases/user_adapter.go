package usecases

import (
	"mygram/modules/v1/users/domain"
	userRepository "mygram/modules/v1/users/interfaces/repositories"
)

type UserAdapter interface {
	RegisterUser(input domain.RegisterUserInput) (domain.User, error)
	LoginUser(input domain.LoginUserInput) (domain.User, error)
}

type UserUseCase struct {
	repoUser *userRepository.Repository
}

func NewUserUseCase(repoUser *userRepository.Repository) *UserUseCase {
	return &UserUseCase{repoUser}
}
