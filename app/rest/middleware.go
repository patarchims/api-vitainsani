package rest

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"vincentcoreapi/helper"
	"vincentcoreapi/modules/telegram"
	"vincentcoreapi/modules/user"
	"vincentcoreapi/modules/user/dto"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://rsmethodist.vincentcore.co.id:28444")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")

		if c.Request.Method == "OPTIONS" {
			c.Writer.Write([]byte("allowed"))
			return
		}

		c.Next()
	}
}

var SecretKey = "V1nc3nC0R3_CO_ID"

// GenerateTokenPair
func GenerateTokenPair(users user.ApiUser) (map[string]string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	// ENDCODE STRING ID USER
	var encodedString = base64.StdEncoding.EncodeToString([]byte(users.ID))
	claims["id"] = encodedString
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}

	// REFRESH TOKEN NOT USE
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["id"] = users.ID
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	_, err = refreshToken.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"token": t,
		//"refresh_token": rt,
	}, nil
}

func JwtVerifyFiber(c *fiber.Ctx) error {
	// func(ctx *fiber.Ctx)

	return c.JSON(helper.APIResponseFailure("Token invalid", http.StatusCreated))
}

func JwtVerify() gin.HandlerFunc {
	return func(c *gin.Context) {

		type requestHeader struct {
			Username string `header:"x-username" binding:"required"`
			Token    string `header:"x-token" binding:"required"`
		}

		r := new(requestHeader)
		c.ShouldBindHeader(&r)
		data, _ := json.Marshal(r)

		_, err := jwt.Parse(r.Token, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err != nil {
			er := errors.New("Token expired")
			response := helper.APIResponseFailure(er.Error(), http.StatusCreated)
			c.AbortWithStatusJSON(http.StatusCreated, response)
			telegram.RunFailureMessage("Verify Token", response, c, data)
			return
		}
		c.Next()
	}
}

// Parse token
func ParseToken(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil
	}
	return token.Claims.(jwt.MapClaims)
}

func JWTVeifyHandler(Logging *logrus.Logger) fiber.Handler {

	return func(c *fiber.Ctx) error {

		var token = c.Get("x-token")
		var userName = c.Get("x-username")

		var datas = dto.RequestToken{Username: userName, Token: token}

		data2, _ := json.Marshal(datas)

		if token == "" {
			Logging.Error("Token Expired")
			response := helper.APIResponseFailure("Token harus diisi", http.StatusCreated)
			telegram.RunFailureMessageFiber("Verify Token", response, c, data2)
			return c.Status(fiber.StatusCreated).JSON(response)
		}

		_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

		if err != nil {
			Logging.Error("Token Expired")
			response := helper.APIResponseFailure("Token Expired", http.StatusCreated)
			telegram.RunFailureMessageFiber("Verify Token", response, c, data2)
			return c.Status(fiber.StatusCreated).JSON(response)
		}

		return c.Next()
	}

}
