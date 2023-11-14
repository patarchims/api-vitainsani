package helper

import (
	"log"
	"net/http"
	"strings"

	humanize "github.com/dustin/go-humanize"
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

func APIResponse(message string, code int, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
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

func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp " + stringValue
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
