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

// @Summary Register User
// @Description Register User
// @Tags Users
// @Accept json
// @Produce json
// @Param input body domain.RegisterUserInput true "Register User Input"
// @Success 200 {object} api.Response
// @Failure 400 {object} api.Response
// @Failure 500 {object} api.Response
// @Router /api/v1/users/register [post]
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

// @Summary Login User
// @Description Login User
// @Tags Users
// @Accept json
// @Produce json
// @Param input body domain.LoginUserInput true "Login User Input"
// @Success 200 {object} api.Response
// @Failure 400 {object} api.Response
// @Failure 500 {object} api.Response
// @Router /api/v1/users/login [post]
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
		if error.IsSame(err, error.ErrEmailNotFound) || error.IsSame(err, error.ErrDataLoginNotFound) {
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

// @Summary Get All Social Media
// @Description Get All Social Media
// @Tags Social Media
// @Accept json
// @Produce json
// @Success 200 {object} api.Response
// @Failure 500 {object} api.Response
// @Router /api/v1/media [get]
func (uc *UserController) GetAllSocialMedia(c *gin.Context) {
	media, err := uc.UserUseCase.AllSocialMedia()
	if err != nil {
		log.Println(err)
		resp := api.APIResponse("Get Social Media Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	if len(media) == 0 {
		resp := api.APIResponse("Social Media Not Found", http.StatusOK, "success", nil)
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := api.APIResponse("Get Social Media Success", http.StatusOK, "success", media)
	c.JSON(http.StatusOK, resp)
}

// @Summary Get One Social Media
// @Description Get One Social Media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param id path string true "Social Media ID"
// @Success 200 {object} api.Response
// @Failure 404 {object} api.Response
// @Failure 500 {object} api.Response
// @Router /api/v1/media/{id} [get]
func (uc *UserController) GetOneSocialMedia(c *gin.Context) {
	id := c.Param("id")
	media, err := uc.UserUseCase.OneSocialMedia(id)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrDataNotFound) {
			errMessage := api.SetError("Social Media Not Found!")
			resp := api.APIResponse("Get Social Media Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Get Social Media Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Get Social Media Success", http.StatusOK, "success", media)
	c.JSON(http.StatusOK, resp)
}

// @Summary Create Social Media
// @Description Create Social Media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param name body domain.InsertSocialMedia true "Create Social Media"
// @Success 200 {object} api.Response
// @Failure 400 {object} api.Response
// @Failure 500 {object} api.Response
// @Router /api/v1/media [post]
func (uc *UserController) CreateSocialMedia(c *gin.Context) {
	var input domain.InsertSocialMedia
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
			resp := api.APIResponse("Create Social Media Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Create Social Media Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	currentUser := c.MustGet("currentUser").(domain.User)
	err := uc.UserUseCase.CheckSocialMedia(currentUser.ID)
	if err != nil {
		if error.IsSame(err, error.ErrSocialMediaAlreadyExist) {
			errorMessage := api.SetError("User Already Have Social Media")
			resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
	}

	socialmedia, err := uc.UserUseCase.CreateSocialMedia(input, currentUser.ID)
	if err != nil {
		log.Println(err)
		resp := api.APIResponse("Create Social Media Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Create Social Media Success", http.StatusOK, "success", socialmedia)
	c.JSON(http.StatusOK, resp)
}

// @Summary Update Social Media
// @Description Update Social Media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param id path string true "Social Media ID"
// @Param name body domain.UpdateSocialMedia true "Update Social Media"
// @Success 200 {object} api.Response
// @Failure 400 {object} api.Response
// @Failure 404 {object} api.Response
// @Failure 500 {object} api.Response
// @Router /api/v1/media/{id} [put]
func (uc *UserController) UpdateSocialMedia(c *gin.Context) {
	id := c.Param("id")
	var input domain.UpdateSocialMedia
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
			resp := api.APIResponse("Update Social Media Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Update Social Media Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if input.Name == "" && input.Social_media_url == "" {
		errorMessage := api.SetError("Name and Social Media Url Cannot Be Empty")
		resp := api.APIResponse("Update Social Media Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	currentUser := c.MustGet("currentUser").(domain.User)
	socialmedia, err := uc.UserUseCase.UpdateSocialMedia(input, id, currentUser.ID)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrSocialMediaNotFound) {
			errorMessage := api.SetError("Social Media Not Found!")
			resp := api.APIResponse("Update Social Media Failed", http.StatusNotFound, "error", errorMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Update Social Media Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := api.APIResponse("Update Social Media Success", http.StatusOK, "success", socialmedia)
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete Social Media
// @Description Delete Social Media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param id path string true "Social Media ID"
// @Success 200 {object} api.Response
// @Failure 404 {object} api.Response
// @Failure 500 {object} api.Response
// @Router /api/v1/media/{id} [delete]
func (uc *UserController) DeleteSocialMedia(c *gin.Context) {
	id := c.Param("id")
	currentUser := c.MustGet("currentUser").(domain.User)
	err := uc.UserUseCase.DeleteSocialMedia(id, currentUser.ID)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrSocialMediaNotFound) {
			errorMessage := api.SetError("Social Media Not Found!")
			resp := api.APIResponse("Delete Social Media Failed", http.StatusNotFound, "error", errorMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Delete Social Media Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Delete Social Media Success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, resp)
}
