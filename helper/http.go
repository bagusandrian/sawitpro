package helper

import (
	"encoding/json"
	"net/http"
)

func SendValidationErrorResponse(w http.ResponseWriter, statusCode int, errors []string) {
	SendJSONResponse(w, statusCode, map[string]interface{}{"error": "Validation failed", "details": errors})
}

func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
