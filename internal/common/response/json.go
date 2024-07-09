package response

import (
	"encoding/json"
	"net/http"
)

type ResponseBody struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, status int, data any) error {
	return JSONWithHeader(w, status, data, nil)
}

func JSONWithHeader(w http.ResponseWriter, status int, data any, headers http.Header) error {
	jsm, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return err
	}

	jsm = append(jsm, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsm)

	return nil
}
