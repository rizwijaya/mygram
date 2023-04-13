package controllers

import (
	"errors"
	"log"
	"mygram/modules/v1/users/domain"
	api "mygram/pkg/api_response"
	error "mygram/pkg/http-error"
	"mygram/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (uc *UserController) Register(c *gin.Context) {
	//Check Input Users and Validation
	var input domain.RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		var verification validator.ValidationErrors
		if errors.As(err, &verification) {
			result := make([]error.Form, len(verification))
			for i, val := range verification {
				result[i] = error.Form{
					Field:   val.Field(),
					Message: error.FormValidationError(val),
				}
			}
			resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}

		resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", "Input Invalid")
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	//Insert Data to Database
	newUser, err := uc.UserUseCase.RegisterUser(input)
	if err != nil {
		log.Println(err)
		var res string
		if err.Error() == error.ErrUsernameAlreadyExist {
			res = "Username Already Exist"
		}
		if err.Error() == error.ErrEmailAlreadyExist {
			res = "Email Already Exist"
		}
		resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", res)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	//Generate JWT Token
	token, err := jwt.GenerateToken(newUser.ID)
	if err != nil {
		log.Println(err)
		resp := api.APIResponse("Register Account Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	user := api.SetUserResponse(newUser, token)
	resp := api.APIResponse("Register Account Success", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, resp)
}
