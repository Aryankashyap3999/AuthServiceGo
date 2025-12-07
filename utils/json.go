package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())	
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}

func WriteJsonSuccessResponse(w http.ResponseWriter, status int,message string, data any) error {
	reponse := map[string]any{}

	reponse["message"] = message
	reponse["data"] = data
	reponse["status"] = "success"

	return WriteJSONResponse(w, status, reponse)
}

func WriteJsonErrorResponse(w http.ResponseWriter, statusCode  int, message string, errorMessage error) error {
	response := map[string]any{}
	response["message"] = message
	response["error"] = errorMessage.Error()
	response["status"] = "error"	
	return WriteJSONResponse(w, statusCode, response)
}	

func ReadJsonBody(r *http.Request, result any) error {
	fmt.Println("Reading JSON body in utils", r.Body)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(result)
}