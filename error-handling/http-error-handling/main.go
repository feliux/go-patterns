package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

// APIError is a custom error handler.
type APIError struct {
	StatusCode int `json:"statusCode"`
	Msg        any `json:"msg"`
}

// Error implements the Error interface for APIError.
func (e APIError) Error() string {
	return fmt.Sprintf("statuscode %d with msg %s", e.StatusCode, e.Msg)
}

// NewAPIError creates a new APIError instance.
func NewAPIError(statuscode int, err error) APIError {
	return APIError{
		StatusCode: statuscode,
		Msg:        err.Error(),
	}
}

// InvalidRequestData sends an error with a http.StatusUnprocessableEntity.
func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Msg:        errors,
	}
}

// InvalidJSON sends an error with a http.StatusBadRequest.
func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("Invalid JSON request data"))
}

// CustomData is a custom type for the name.
type CustomData struct {
	Name string `json:"name"`
}

// validate validates that the name contains more than 3 letters.
func (c CustomData) validate() map[string]string {
	errors := make(map[string]string)
	if len(c.Name) < 3 {
		errors["name"] = "name must be at least 3 characters long"
	}
	return errors
}

// CustomHandler handles requests for decoding JSON data into CustomData.
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

// Make is a custom handler.
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

// writeJSON write the response to the client.
func writeJSON(w http.ResponseWriter, statuscode int, v any) error {
	w.WriteHeader(statuscode)
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// func main() {
// 	// this method sends "not implemented" when method is not a POST. Inside if
// 	// curl -X POST http://localhost:3000 -d '{"name": "foo"}'
// 	http.HandleFunc("/", Make(CustomHandler))
// 	http.ListenAndServe(":3000", nil)
// }

func main() {
	// this method sends "method not allowed" when method is not POST. Not inside if
	router := http.NewServeMux()
	router.HandleFunc("POST /", Make(CustomHandler))
	http.ListenAndServe(":3000", router)
}
