package services

import (
	"encoding/json"
	"net/http"
	"userCrud/internal/models"
	"userCrud/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func HandleUpdateUser(db *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")
		id, err := uuid.Parse(idString)

		if err != nil {
			utils.SendJSON(
				w,
				models.ResponseError{Error: "id is invalid"},
				http.StatusBadRequest,
			)
			return
		}

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

		u, id, err = db.Update(u, id)

		if err != nil {
			utils.SendJSON(
				w,
				models.ResponseError{Error: err.Error()},
				http.StatusNotFound,
			)
			return
		}

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
