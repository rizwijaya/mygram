package controllers

import (
	"errors"
	"log"
	"mygram/modules/v1/comments/domain"
	domainUser "mygram/modules/v1/users/domain"
	api "mygram/pkg/api_response"
	error "mygram/pkg/http-error"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Get All Comments
// @Description Get All Comments
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param id_photos path string true "Id Photos"
// @Success 200 {object} api.Response
// @Failure 404 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 401 {object} api.Response
// @Router /api/v1/comments/{id_photos} [get]
func (cc *CommentController) GetAllComments(c *gin.Context) {
	idPhotos := c.Param("id_photos")
	user := c.MustGet("currentUser").(domainUser.User)
	comments, err := cc.CommentUseCase.GetAllComments(idPhotos, user.ID)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrDataNotFound) {
			errMessage := api.SetError("Comment Not Found!")
			resp := api.APIResponse("Get All Comments Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Get All Comments Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Get All Comments Success", http.StatusOK, "success", comments)
	c.JSON(http.StatusOK, resp)
}

// @Summary Get Comment By Id
// @Description Get Comment By Id
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param id path string true "Id Comment"
// @Success 200 {object} api.Response
// @Failure 404 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 401 {object} api.Response
// @Router /api/v1/comments/id/{id} [get]
func (cc *CommentController) GetCommentById(c *gin.Context) {
	id := c.Param("id")
	comment, err := cc.CommentUseCase.GetCommentById(id)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrDataNotFound) {
			errMessage := api.SetError("Comment Not Found!")
			resp := api.APIResponse("Get Comment By Id Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Get Comment By Id Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Get Comment By Id Success", http.StatusOK, "success", comment)
	c.JSON(http.StatusOK, resp)
}

// @Summary Create Comment
// @Description Create Comment
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param body body domain.InsertComment true "Body"
// @Success 200 {object} api.Response
// @Failure 400 {object} api.Response
// @Failure 404 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 401 {object} api.Response
// @Router /api/v1/comments [post]
func (cc *CommentController) CreateComment(c *gin.Context) {
	var input domain.InsertComment
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		var validation validator.ValidationErrors
		if errors.As(err, &validation) {
			result := make([]error.Form, len(validation))
			for i, v := range validation {
				result[i] = error.Form{
					Field:   v.Field(),
					Message: error.FormValidationError(v),
				}
			}
			resp := api.APIResponse("Create Comment Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Create Comment Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user := c.MustGet("currentUser").(domainUser.User)
	input.UserID = user.ID
	comment, err := cc.CommentUseCase.CreateComment(input)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrPhotoNotFound) {
			errMessage := api.SetError("Cannot Create Comment, Photo Not Found!")
			resp := api.APIResponse("Create Comment Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Create Comment Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Create Comment Success", http.StatusOK, "success", comment)
	c.JSON(http.StatusOK, resp)
}

// @Summary Update Comment
// @Description Update Comment
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param id path string true "Id Comment"
// @Param body body domain.UpdateComment true "Body"
// @Success 200 {object} api.Response
// @Failure 400 {object} api.Response
// @Failure 404 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 401 {object} api.Response
// @Router /api/v1/comments/{id} [put]
func (cc *CommentController) UpdateComment(c *gin.Context) {
	id := c.Param("id")
	var input domain.UpdateComment
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		var validation validator.ValidationErrors
		if errors.As(err, &validation) {
			result := make([]error.Form, len(validation))
			for i, v := range validation {
				result[i] = error.Form{
					Field:   v.Field(),
					Message: error.FormValidationError(v),
				}
			}
			resp := api.APIResponse("Update Comment Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Update Comment Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if input.PhotoID == 0 && input.Message == "" {
		errMessage := api.SetError("Photo ID and Comment Cannot Be Empty!")
		resp := api.APIResponse("Update Comment Failed", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user := c.MustGet("currentUser").(domainUser.User)
	comment, err := cc.CommentUseCase.UpdateComment(id, input, user.ID)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrPhotoNotFound) {
			errMessage := api.SetError("Cannot Update Comment, Photo Not Found!")
			resp := api.APIResponse("Update Comment Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		if error.IsSame(err, error.ErrCommentNotFound) || error.IsSame(err, error.ErrDataNotFound) {
			errMessage := api.SetError("Comment Not Found!")
			resp := api.APIResponse("Update Comment Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Update Comment Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Update Comment Success", http.StatusOK, "success", comment)
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete Comment
// @Description Delete Comment
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param id path string true "Id Comment"
// @Success 200 {object} api.Response
// @Failure 404 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 401 {object} api.Response
// @Router /api/v1/comments/{id} [delete]
func (cc *CommentController) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	user := c.MustGet("currentUser").(domainUser.User)
	err := cc.CommentUseCase.DeleteComment(id, user.ID)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrCommentNotFound) {
			errMessage := api.SetError("Comment Not Found!")
			resp := api.APIResponse("Delete Comment Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Delete Comment Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Delete Comment Success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, resp)
}

// @Summary Get All Photos
// @Description Get All Photos
// @Tags Photos
// @Accept  json
// @Produce  json
// @Success 200 {object} api.Response
// @Failure 404 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 401 {object} api.Response
// @Router /api/v1/photos [get]
func (cc *CommentController) GetAllPhotos(c *gin.Context) {
	photos, err := cc.CommentUseCase.GetAllPhotos()
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrPhotoNotFound) {
			errMessage := api.SetError("Photo Not Found!")
			resp := api.APIResponse("Get All Photos Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Get All Photos Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Get All Photos Success", http.StatusOK, "success", photos)
	c.JSON(http.StatusOK, resp)
}

// @Summary Get Photo By ID
// @Description Get Photo By ID
// @Tags Photos
// @Accept  json
// @Produce  json
// @Param id path string true "Id Photo"
// @Success 200 {object} api.Response
// @Failure 404 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 401 {object} api.Response
// @Router /api/v1/photos/{id} [get]
func (cc *CommentController) GetPhotoById(c *gin.Context) {
	id := c.Param("id")
	user := c.MustGet("currentUser").(domainUser.User)
	photo, err := cc.CommentUseCase.GetPhotoById(id, user.ID)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrPhotoNotFound) {
			errMessage := api.SetError("Photo Not Found!")
			resp := api.APIResponse("Get Photo By ID Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Get Photo By ID Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Get Photo By ID Success", http.StatusOK, "success", photo)
	c.JSON(http.StatusOK, resp)
}

// @Summary Create Photo
// @Description Create Photo
// @Tags Photos
// @Accept  json
// @Produce  json
// @Param input body domain.InsertPhoto true "Input Create Photo"
// @Success 200 {object} api.Response
// @Failure 400 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 401 {object} api.Response
// @Router /api/v1/photos [post]
func (cc *CommentController) CreatePhoto(c *gin.Context) {
	var input domain.InsertPhoto
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		var validation validator.ValidationErrors
		if errors.As(err, &validation) {
			result := make([]error.Form, len(validation))
			for i, v := range validation {
				result[i] = error.Form{
					Field:   v.Field(),
					Message: error.FormValidationError(v),
				}
			}
			resp := api.APIResponse("Create Photo Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Create Photo Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user := c.MustGet("currentUser").(domainUser.User)
	input.UserID = user.ID
	photo, err := cc.CommentUseCase.CreatePhoto(input)
	if err != nil {
		log.Println(err)
		resp := api.APIResponse("Create Photo Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Create Photo Success", http.StatusOK, "success", photo)
	c.JSON(http.StatusOK, resp)
}

// @Summary Update Photo
// @Description Update Photo
// @Tags Photos
// @Accept  json
// @Produce  json
// @Param id path string true "Id Photo"
// @Param input body domain.UpdatePhoto true "Input Update Photo"
// @Success 200 {object} api.Response
// @Failure 400 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 401 {object} api.Response
// @Router /api/v1/photos/{id} [put]
func (cc *CommentController) UpdatePhoto(c *gin.Context) {
	id := c.Param("id")
	var input domain.UpdatePhoto
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		var validation validator.ValidationErrors
		if errors.As(err, &validation) {
			result := make([]error.Form, len(validation))
			for i, v := range validation {
				result[i] = error.Form{
					Field:   v.Field(),
					Message: error.FormValidationError(v),
				}
			}
			resp := api.APIResponse("Update Photo Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Update Photo Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if input.Caption == "" && input.Title == "" && input.Photo_url == "" {
		errMessage := api.SetError("Caption, Title, Photo_url cannot be empty!")
		resp := api.APIResponse("Update Photo Failed", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user := c.MustGet("currentUser").(domainUser.User)
	input.UserID = user.ID
	photo, err := cc.CommentUseCase.UpdatePhoto(id, input)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrPhotoNotFound) {
			errMessage := api.SetError("Photo Not Found!")
			resp := api.APIResponse("Update Photo Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Update Photo Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Update Photo Success", http.StatusOK, "success", photo)
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete Photo
// @Description Delete Photo
// @Tags Photos
// @Accept  json
// @Produce  json
// @Param id path string true "Id Photo"
// @Success 200 {object} api.Response
// @Failure 404 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 401 {object} api.Response
// @Router /api/v1/photos/{id} [delete]
func (cc *CommentController) DeletePhoto(c *gin.Context) {
	id := c.Param("id")
	user := c.MustGet("currentUser").(domainUser.User)
	err := cc.CommentUseCase.DeletePhoto(id, user.ID)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrPhotoNotFound) {
			errMessage := api.SetError("Photo Not Found!")
			resp := api.APIResponse("Delete Photo Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Delete Photo Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Delete Photo Success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, resp)
}
