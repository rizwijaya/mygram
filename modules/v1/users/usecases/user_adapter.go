package usecases

import (
	userRepository "mygram/modules/v1/users/interfaces/repositories"
)

type UserAdapter interface {
}

type UserUseCase struct {
	repoUser *userRepository.Repository
}

func NewUserUseCase(repoUser *userRepository.Repository) *UserUseCase {
	return &UserUseCase{repoUser}
}
