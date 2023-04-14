package usecases

import (
	"mygram/modules/v1/comments/domain"
	errorHandling "mygram/pkg/http-error"
)

func (cu *CommentUseCase) GetAllComments(idPhotos string, idUser int) ([]domain.Comment, error) {
	comment, err := cu.repoComment.FindAllComments(idPhotos, idUser)
	if err != nil {
		return nil, err
	}

	if len(comment) == 0 {
		return nil, errorHandling.ErrDataNotFound
	}

	return comment, nil
}

func (cu *CommentUseCase) GetCommentById(id string) (domain.Comment, error) {
	comment, err := cu.repoComment.FindCommentById(id)
	if err != nil {
		return domain.Comment{}, err
	}

	if comment.ID == 0 {
		return domain.Comment{}, errorHandling.ErrDataNotFound
	}

	return comment, nil
}

func (cu *CommentUseCase) CreateComment(input domain.InsertComment) (domain.Comment, error) {
	//Validate Photo Exist
	photo, err := cu.repoComment.FindPhotoById(input.PhotoID)
	if err != nil {
		return domain.Comment{}, errorHandling.ErrPhotoNotFound
	}

	if photo.ID == 0 {
		return domain.Comment{}, errorHandling.ErrPhotoNotFound
	}

	comment := domain.Comment{
		PhotoID: input.PhotoID,
		UserID:  input.UserID,
		Message: input.Message,
	}

	return cu.repoComment.SaveComment(comment)
}
