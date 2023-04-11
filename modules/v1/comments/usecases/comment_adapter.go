package usecases

import (
	commentRepository "mygram/modules/v1/comments/interfaces/repositories"
)

type CommentAdapter interface {
}

type CommentUseCase struct {
	repoComment *commentRepository.Repository
}

func NewCommentUseCase(repoComment *commentRepository.Repository) *CommentUseCase {
	return &CommentUseCase{repoComment}
}
