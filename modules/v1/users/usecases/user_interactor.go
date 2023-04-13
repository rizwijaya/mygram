package usecases

import (
	"mygram/modules/v1/users/domain"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserUseCase) RegisterUser(input domain.RegisterUserInput) (domain.User, error) {
	user := domain.User{}
	user.UserName = input.Username
	user.Email = input.Email
	user.Age = input.Age

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	newUser, err := u.repoUser.Save(user)
	if err != nil {
		return user, err
	}

	return newUser, nil
}
