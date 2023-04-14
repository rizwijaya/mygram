package usecases

import (
	"mygram/modules/v1/comments/domain"
	commentRepository "mygram/modules/v1/comments/interfaces/repositories"
)

type CommentAdapter interface {
	GetAllComments(idPhotos string) ([]domain.Comment, error)
	GetCommentById(id string) (domain.Comment, error)
}

type CommentUseCase struct {
	repoComment *commentRepository.Repository
}

func NewCommentUseCase(repoComment *commentRepository.Repository) *CommentUseCase {
	return &CommentUseCase{repoComment}
}
