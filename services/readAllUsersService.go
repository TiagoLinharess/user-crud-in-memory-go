package services

import (
	"net/http"
	"userCrud/models/application"
	"userCrud/models/user"
	"userCrud/utils"
)

func HandleGetUsers(db *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := db.FindAll()
		users := make([]user.UserResponse, 0, len(data))
		for id, u := range data {
			u := user.UserResponse{
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
