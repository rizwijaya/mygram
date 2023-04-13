package controllers

import (
	"errors"
	"log"
	"mygram/modules/v1/users/domain"
	api "mygram/pkg/api_response"
	error "mygram/pkg/http-error"
	jwt "mygram/pkg/token"
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
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	//Insert Data to Database
	newUser, err := uc.UserUseCase.RegisterUser(input)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrUsernameAlreadyExist) {
			errorMessage := api.SetError("username already exist")
			resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, resp)
			return
		} else if error.IsSame(err, error.ErrEmailAlreadyExist) {
			errorMessage := api.SetError("email already exist")
			resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		resp := api.APIResponse("Register Account Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
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

func (uc *UserController) Login(c *gin.Context) {
	//Check Input Users and Validation
	var input domain.LoginUserInput
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
			resp := api.APIResponse("Login Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Login Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user, err := uc.UserUseCase.LoginUser(input)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrEmailNotFound) {
			errorMessage := api.SetError("email/password is wrong")
			resp := api.APIResponse("Login Failed", http.StatusNotFound, "error", errorMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		if error.IsSame(err, error.ErrUsernameNotFound) {
			errorMessage := api.SetError("username/password is wrong")
			resp := api.APIResponse("Login Failed", http.StatusNotFound, "error", errorMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Login Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	//Generate JWT Token
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		log.Println(err)
		resp := api.APIResponse("Login Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	userResponse := api.SetUserResponse(user, token)
	resp := api.APIResponse("Login Success", http.StatusOK, "success", userResponse)
	c.JSON(http.StatusOK, resp)
}

func (uc *UserController) GetAllSocialMedia(c *gin.Context) {
	media, err := uc.UserUseCase.AllSocialMedia()
	if err != nil {
		log.Println(err)
		resp := api.APIResponse("Get Social Media Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	if len(media) == 0 {
		resp := api.APIResponse("Social Media Not Found!", http.StatusOK, "success", nil)
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := api.APIResponse("Get Social Media Success", http.StatusOK, "success", media)
	c.JSON(http.StatusOK, resp)
}
