package responses

import (
	"encoding/json"
	"github.com/Xurliman/auth-service/internal/constants"
	"github.com/Xurliman/auth-service/internal/server/app/models"
)

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func UserListTransformer(user []*models.User) (rows []UserResponse) {
	for _, row := range user {
		var mapResponse UserResponse
		jsonResponse, _ := json.Marshal(row)
		_ = json.Unmarshal(jsonResponse, &mapResponse)

		mapResponse = UserResponse{
			Id:       row.Id.String(),
			Username: row.Username,
			Name:     row.Name,
			Email:    row.Email,
		}
		rows = append(rows, mapResponse)
	}
	return rows
}

type UserDetailResponse struct {
	Id              string `json:"id"`
	Username        string `json:"username"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	IsEmailVerified bool   `json:"is_email_verified"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func UserDetailTransformer(user models.User) (row UserDetailResponse) {
	return UserDetailResponse{
		Id:              user.Id.String(),
		Username:        user.Username,
		Name:            user.Name,
		Email:           user.Email,
		IsEmailVerified: user.IsEmailVerified,
		CreatedAt:       user.CreatedAt.Format(constants.TimestampFormat),
		UpdatedAt:       user.UpdatedAt.Format(constants.TimestampFormat),
	}
}
