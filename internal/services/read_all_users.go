package services

import (
	"net/http"
	"userCrud/internal/models"
	"userCrud/internal/utils"
)

func HandleGetUsers(db *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := db.FindAll()
		users := make([]models.UserResponse, 0, len(data))
		for id, u := range data {
			u := models.UserResponse{
				Id:        id.String(),
				FirstName: u.FirstName,
				LastName:  u.LastName,
				Biography: u.Biography,
			}
			users = append(users, u)
		}

		utils.SendJSON(
			w,
			users,
			http.StatusOK,
		)
	}
}
