package helper

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"vincentcoreapi/pkg/constant"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
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
		var message = fmt.Sprintf("%s %v %s", e.Field(), e.Value(), e.Tag())
		errors = append(errors, message)
	}

	return errors
}

func PaylayotHandler(c *gin.Context, errorMessage map[string]any) gin.HandlerFunc {
	return func(c *gin.Context) {
		response := APIResponse(constant.DataGagalDiProses.Error(), http.StatusCreated, errorMessage)
		c.JSON(http.StatusCreated, response)
	}
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

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "sak.kapal@gmail.com"
const CONFIG_AUTH_EMAIL = "sak.kapal@gmail.com"
const CONFIG_AUTH_PASSWORD = "xkngbwsyanoiakqo"

func SendMail(to string, cc string, otp string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", to)
	mailer.SetAddressHeader("Cc", cc, "VERIFIKASI DATA")
	mailer.SetHeader("Subject", "Verifikasi Data")
	body := fmt.Sprintf("<FONT COLOR=BLACK><RIGHT>KODE OTP       : %s </RIGHT></FONT><br> ", otp)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}

const API_KEY_VONAGE = "59bd523e"
const API_SECRET_VONAGE = "l2V9LHi4ZWfNWmeU"

func LayananPoli(strLayanan string) (strChange string) {
	if strLayanan == "1" {
		return "Buka"
	} else {
		return "Tutup"
	}
}

func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
