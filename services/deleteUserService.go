package services

import (
	"net/http"
	"userCrud/models/application"
	"userCrud/models/response"
	"userCrud/models/user"
	"userCrud/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func HandleDeleteUser(db *application.Application) http.HandlerFunc {
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

		u, id, err := db.Delete(id)

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
