package middlewares

import (
	// "bytes"
	// "encoding/json"
	// "fmt"
	// "io"
	"context"
	"net/http"

	"AuthInGo/dto"
	"AuthInGo/utils"
)

// func RequestValidator(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		bodyBytes, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			fmt.Println("Error reading body:", err)
// 			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
// 			return
// 		}

// 		var payload map[string]interface{}
// 		if err := json.Unmarshal(bodyBytes, &payload); err != nil {
// 			fmt.Println("Error parsing JSON in middleware:", err)
// 			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid JSON", err)
// 			return
// 		}

// 		fmt.Println("Payload in middleware:", payload)

// 		if len(payload) == 0 {
// 			fmt.Println("Empty payload received")
// 			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Request body cannot be empty", fmt.Errorf("empty payload"))
// 			return
// 		}

// 		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

// 		next.ServeHTTP(w, r)
// 	})
// }


func UserLoginRequestValidator(next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        var payload dto.LoginUserRequestDTO

        // Read and decode the JSON body into the payload
        if err := utils.ReadJsonBody(r, &payload); err != nil {
            utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
            return
        }

        // Validate the payload using the Validator instance
        if err := utils.Validator.Struct(payload); err != nil {
            utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
            return
        }

		req_context := r.Context()

		ctx := context.WithValue(req_context, "login_payload", payload)

        next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
    })
}

func UserCreateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateUserRequestDTO							
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}	
		if err := utils.Validator.Struct(payload); err != nil {		
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
			return
		}

		req_context := r.Context()

		ctx := context.WithValue(req_context, "create_payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}
