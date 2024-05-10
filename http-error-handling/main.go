package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type APIError struct {
	StatusCode int `json:"statusCode"`
	Msg        any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("statuscode %d with msg %s", e.StatusCode, e.Msg)
}

func NewAPIError(statuscode int, err error) APIError {
	return APIError{
		StatusCode: statuscode,
		Msg:        err.Error(),
	}
}

func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Msg:        errors,
	}
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("Invalid JSON request data"))
}

type CustomData struct {
	Name string `json:"name"`
}

func (c CustomData) validate() map[string]string {
	errors := make(map[string]string)
	if len(c.Name) < 3 {
		errors["name"] = "name must be at least 3 characters long"
	}
	return errors
}

func CustomHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		var customData CustomData
		if err := json.NewDecoder(r.Body).Decode(&customData); err != nil {
			return InvalidJSON()
		}
		defer r.Body.Close()
		if errors := customData.validate(); len(errors) > 0 {
			return InvalidRequestData(errors)
		}
		return writeJSON(w, http.StatusCreated, customData)
	} else {
		return writeJSON(w, http.StatusNotImplemented, nil)
	}
}

func Make(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("Internal server error", "err", err, "path", r.URL.Path)
			if apiERR, ok := err.(APIError); ok {
				writeJSON(w, apiERR.StatusCode, apiERR)
			} else {
				// errResp := map[string]any{
				// 	"statuscode": http.StatusInternalServerError,
				// 	"msg":        "internal server error",
				// }
				writeJSON(w, http.StatusInternalServerError, map[string]any{
					"statuscode": http.StatusInternalServerError,
					"msg":        "internal server error",
				})
			}
		}
	}
}

func writeJSON(w http.ResponseWriter, statuscode int, v any) error {
	w.WriteHeader(statuscode)
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// func main() {
// 	// this method sends "not implemented" when method is not a POST. Inside if
// 	// curl -X POST http://localhost:3000 -d '{"name": "fo"}'
// 	http.HandleFunc("/", Make(CustomHandler))
// 	http.ListenAndServe(":3000", nil)
// }

func main() {
	// this method sends "method not allowed" when method is not POST. Not inside if
	router := http.NewServeMux()
	router.HandleFunc("POST /", Make(CustomHandler))
	http.ListenAndServe(":3000", router)
}
