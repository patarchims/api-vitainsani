package handler

import (
	"encoding/json"
	"net/http"
	"vincentcoreapi/app/rest"
	"vincentcoreapi/helper"
	"vincentcoreapi/modules/telegram"
	"vincentcoreapi/modules/user/entity"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase    entity.UserUseCase
	UserRepository entity.UserRepository
}

func (uh *UserHandler) Login(c *gin.Context) {
	type requestHeader struct {
		Username string `header:"x-username" binding:"required"`
		Password string `header:"x-password" binding:"required"`
	}

	r := new(requestHeader)
	err := c.ShouldBindHeader(&r)

	data, _ := json.Marshal(r)
	if err != nil {
		response := helper.APIResponseFailure("Username atau Password Tidak Sesuai", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		telegram.RunFailureMessage("GET TOKEN", response, c, data)
		return
	}

	user, exist := uh.UserRepository.GetByUser(c, r.Username)
	if !exist {
		response := helper.APIResponseFailure("Username atau Password Tidak Sesuai", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		telegram.RunFailureMessage("GET TOKEN", response, c, data)
		return
	}

	if user.Password != r.Password {
		response := helper.APIResponseFailure("Username atau Password Tidak Sesuai", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		telegram.RunFailureMessage("GET TOKEN", response, c, data)
		return
	}

	m, _ := rest.GenerateTokenPair(user)
	response := helper.APIResponse("Ok", http.StatusOK, "Ok", m)
	c.JSON(http.StatusOK, response)
	telegram.RunSuccessMessage("GET TOKEN", response, c, data)
}
