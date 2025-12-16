package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"userCrud/models/response"
)

func SendJSON[T any](w http.ResponseWriter, resp T, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to mershal json data", "error", err)
		SendJSON(
			w,
			response.ResponseError{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}
}
