package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"tapi-controller/models"
	"time"

	"github.com/gin-gonic/gin"
)

var ExpiredTime time.Duration = time.Second * 180

func (server *Server) Login(ctx *gin.Context) {
	var data models.LoginData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if data.Username == server.config.Usenrname && data.Password == server.config.Password {
		token, err := server.token.CreateToken(data.Username, ExpiredTime)
		if err != nil {
			message := fmt.Sprintf("unautorisued user %s", data.Username)
			log.Println(message)
			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New(message)))
			return
		}
		log.Printf("user %s is authorised", data.Username)
		ctx.JSON(http.StatusOK, token)
		return
	}
	message := fmt.Sprintf("unautorisued user %s", data.Username)
	log.Println(message)
	ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New(message)))
}
