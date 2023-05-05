package helper

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Data interface{} `json:"response"`
	Meta Meta        `json:"metadata"`
}

type FailureResponse struct {
	Meta Meta `json:"metadata"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	// Status  string `json:"status"`
}

func APIResponseFailure(message string, code int) FailureResponse {
	meta := Meta{
		Message: message,
		Code:    code,
	}

	jsonResponse := FailureResponse{
		Meta: meta,
	}

	return jsonResponse
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		// Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

// Hash - Hash password using Bcrypt
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckHash - Check if password and hash password is valid
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
