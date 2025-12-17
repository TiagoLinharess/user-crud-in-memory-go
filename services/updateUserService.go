package services

import (
	"encoding/json"
	"net/http"
	"userCrud/models/application"
	"userCrud/models/response"
	"userCrud/models/user"
	"userCrud/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func HandleUpdateUser(db *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")
		id, err := uuid.Parse(idString)

		if err != nil {
			utils.SendJSON(
				w,
				response.ResponseError{Error: "id is invalid"},
				http.StatusBadRequest,
			)
			return
		}

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

		u, id, err = db.Update(u, id)

		if err != nil {
			utils.SendJSON(
				w,
				response.ResponseError{Error: err.Error()},
				http.StatusNotFound,
			)
			return
		}

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
