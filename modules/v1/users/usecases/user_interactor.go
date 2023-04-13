package usecases

import (
	"fmt"
	"mygram/modules/v1/users/domain"

	errorHandling "mygram/pkg/http-error"

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

func (u *UserUseCase) LoginUser(input domain.LoginUserInput) (domain.User, error) {
	var (
		user domain.User
		err  error
	)

	if input.Email != "" {
		user, err = u.repoUser.FindUser("email", input.Email)
		if err != nil {
			if errorHandling.IsSame(err, errorHandling.ErrDataNotFound) {
				return user, errorHandling.ErrEmailNotFound
			}
			return user, err
		}
	} else {
		user, err = u.repoUser.FindUser("username", input.Username)
		if err != nil {
			fmt.Println("Ga ada data username")
			if errorHandling.IsSame(err, errorHandling.ErrDataNotFound) {
				fmt.Println("custome error")
				return user, errorHandling.ErrUsernameNotFound
			}
			return user, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, err
	}

	return user, nil
}
