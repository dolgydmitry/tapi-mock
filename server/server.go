package server

import (
	"errors"
	"fmt"
	"log"
	"tapi-controller/controllers"
	"tapi-controller/token"
	"tapi-controller/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config   utils.Config
	db       *utils.InMemoryDb
	engine   *gin.Engine
	tapiCtrl controllers.TapiCtrl
	token    token.TokenMaker
}

func InitServer(config utils.Config, db *utils.InMemoryDb) (Server, error) {
	log.Println("Initilize server")
	keys, err := utils.LoadKeys(config)
	if err != nil {
		message := "cannot create tokenMaker obj"
		log.Println(message)
		return Server{}, errors.New(message)
	}
	server := Server{
		config:   config,
		db:       db,
		engine:   gin.Default(),
		tapiCtrl: new(controllers.TapiCtrlInMemDB),
		token:    &token.JwtMaker{PrivateKey: keys.PrivateKey, PublicKey: keys.PublicKey},
	}
	return server, nil
}

func (server *Server) Start() error {
	log.Println("Add routes")
	server.AddRoute()
	host := fmt.Sprintf("%s:%d", server.config.Address, server.config.Port)
	log.Printf("run server on the host: %s", host)
	return server.engine.Run(host)
}

func (server *Server) AddRoute() {
	server.engine.POST("/login/", server.Login)

	authRoutes := server.engine.Group("/").Use(authMiddleware(server.token))

	authRoutes.POST("/tapi-common:context/tapi-connectivity:connectivity-context/", server.CreateConnSrvRoute)
	authRoutes.GET("/tapi-common:context/tapi-connectivity:connectivity-context/", server.GetAllConnSrvRoute)
	authRoutes.GET("/tapi-common:context/tapi-connectivity:connectivity-context/:uuid", server.GetConnSrvRoute)
	authRoutes.PATCH("/tapi-common:context/tapi-connectivity:connectivity-context/:uuid", server.UpdateConSrvRoute)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
