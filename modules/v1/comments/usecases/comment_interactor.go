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

func (cu *CommentUseCase) UpdateComment(idComments string, input domain.UpdateComment, idUser int) (domain.Comment, error) {
	if input.PhotoID != 0 {
		//Validate Photo Exist
		photo, err := cu.repoComment.FindPhotoById(input.PhotoID)
		if err != nil {
			return domain.Comment{}, errorHandling.ErrPhotoNotFound
		}

		if photo.ID == 0 {
			return domain.Comment{}, errorHandling.ErrPhotoNotFound
		}
	}
	//Check Comment Exist
	comment, err := cu.GetCommentById(idComments)
	if err != nil {
		return domain.Comment{}, err
	}

	if comment.ID == 0 || comment.UserID != idUser {
		return domain.Comment{}, errorHandling.ErrCommentNotFound
	}

	updateComment := domain.Comment{
		PhotoID: input.PhotoID,
		Message: input.Message,
	}

	return cu.repoComment.UpdateComment(updateComment, idComments)
}

func (cu *CommentUseCase) DeleteComment(idComment string, idUser int) error {
	//Check Comment Exist
	comment, err := cu.repoComment.FindCommentById(idComment)
	if err != nil {
		if errorHandling.IsSame(err, errorHandling.ErrDataNotFound) {
			return errorHandling.ErrCommentNotFound
		}
		return err
	}

	if comment.ID == 0 || comment.UserID != idUser {
		return errorHandling.ErrCommentNotFound
	}

	return cu.repoComment.DeleteComment(idComment)
}
