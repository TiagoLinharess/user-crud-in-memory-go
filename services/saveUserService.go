package services

import (
	"encoding/json"
	"net/http"
	"userCrud/models/application"
	"userCrud/models/response"
	"userCrud/models/user"
	"userCrud/utils"
)

func HandleSaveUser(db *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body user.UserBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			utils.SendJSON(
				w,
				response.ResponseError{Error: "invalid body"},
				http.StatusUnprocessableEntity,
			)
			return
		}

		u := user.User(body)

		u, id := db.Insert(u)

		response := user.UserResponse{
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
