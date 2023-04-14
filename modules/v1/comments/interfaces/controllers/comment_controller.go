package controllers

import (
	"log"
	api "mygram/pkg/api_response"
	error "mygram/pkg/http-error"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cc *CommentController) GetAllComments(c *gin.Context) {
	idPhotos := c.Param("id_photos")
	comments, err := cc.CommentUseCase.GetAllComments(idPhotos)
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
