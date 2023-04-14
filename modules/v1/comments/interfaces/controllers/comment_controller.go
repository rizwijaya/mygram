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
