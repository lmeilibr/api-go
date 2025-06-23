package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response struct {
	Error string `json:"message,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func SendJSON(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err)
		SendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err = w.Write(data); err != nil {
		slog.Error("failed to write response", "error", err)
		return
	}

}
