package controllers

import (
	commentRepository "mygram/modules/v1/comments/interfaces/repositories"
	commentUseCase "mygram/modules/v1/comments/usecases"

	"gorm.io/gorm"
)

type CommentController struct {
	CommentUseCase *commentUseCase.CommentUseCase
}

func NewCommentController(db *gorm.DB) *CommentController {
	repo := commentRepository.NewCommentRepository(db)
	cu := commentUseCase.NewCommentUseCase(repo)
	return &CommentController{
		CommentUseCase: cu,
	}
}
