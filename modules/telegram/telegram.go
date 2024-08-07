package telegram

import (
	"errors"
	"net"
	"net/http"
	"os"
	"strings"
	"vincentcoreapi/helper"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func getIP(r *http.Request) (string, error) {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1", nil
		}
		return ip, nil
	}

	return "", errors.New("IP not found")
}

func RunSuccessMessage(method string, response helper.Response, c *gin.Context, data []byte) {
	ip, _ := getIP(c.Request)
	host := os.Getenv("HOST")
	var message Message
	message.ChatID = -1001867821168
	message.Text = SendMessageTelegram(method, response, string(data), string(c.GetHeader("User-Agent")), ip, host)
	SendMessage(&message)
}

func RunFailureMessage(method string, response helper.FailureResponse, c *gin.Context, data []byte) {
	ip, _ := getIP(c.Request)
	host := os.Getenv("HOST")
	var message Message
	message.ChatID = -1001867821168
	message.Text = SendMessageFailureTelegram(method, response, string(data), string(c.GetHeader("User-Agent")), ip, host)
	SendMessage(&message)
}

func RunFailureMessageFiber(method string, response helper.FailureResponse, c *fiber.Ctx, data []byte) {
	// GET IP
	ip, _ := getIPFromFiber(c)
	host := os.Getenv("HOST")
	var message Message
	message.ChatID = -1001867821168
	message.Text = SendMessageFailureTelegram(method, response, string(data), string(c.Get("User-Agent")), ip, host)
	SendMessage(&message)
}

func RunSuccessMessageFiber(method string, response helper.Response, c *fiber.Ctx, data []byte) {
	ip, _ := getIPFromFiber(c)
	host := os.Getenv("HOST")
	var message Message
	message.ChatID = -1001867821168
	message.Text = SendMessageTelegram(method, response, string(data), string(c.Get("User-Agent")), ip, host)
	SendMessage(&message)
}

func getIPFromFiber(c *fiber.Ctx) (string, error) {
	ips := c.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	ip, _, err := net.SplitHostPort(c.Context().RemoteAddr().Network())
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(ip)

	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1", nil
		}
		return ip, nil
	}

	return "", nil
}
