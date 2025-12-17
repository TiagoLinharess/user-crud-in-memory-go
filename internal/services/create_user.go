package services

import (
	"encoding/json"
	"net/http"
	"userCrud/internal/models"
	"userCrud/internal/utils"
)

func HandleSaveUser(db *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body models.UserBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			utils.SendJSON(
				w,
				models.ResponseError{Error: "invalid body"},
				http.StatusUnprocessableEntity,
			)
			return
		}

		u := models.User(body)

		u, id := db.Insert(u)

		response := models.UserResponse{
			Id:        id.String(),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Biography: u.Biography,
		}

		utils.SendJSON(
			w,
			response,
			http.StatusCreated,
		)
	}
}
