package api

import "mygram/modules/v1/users/domain"

func SetMessage(message string) Message {
	return Message{"message": message}
}

func SetUserResponse(user domain.User, token string) UserResponse {
	return UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Age:      user.Age,
		Token:    "Bearer " + token,
	}
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}
