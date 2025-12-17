package services

import (
	"net/http"
	"userCrud/internal/models"
	"userCrud/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func HandleGetUserById(db *models.Application) http.HandlerFunc {
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

		u, err := db.FindById(id)

		if err != nil {
			utils.SendJSON(
				w,
				models.ResponseError{Error: err.Error()},
				http.StatusNotFound,
			)
			return
		}

		res := models.UserResponse{
			Id:        id.String(),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Biography: u.Biography,
		}

		utils.SendJSON(
			w,
			res,
			http.StatusOK,
		)
	}
}
