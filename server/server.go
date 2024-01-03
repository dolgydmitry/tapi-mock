package server

import (
	"fmt"
	"log"
	"tapi-controller/controllers"
	"tapi-controller/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config   utils.Config
	db       *utils.InMemoryDb
	engine   *gin.Engine
	tapiCtrl controllers.TapiCtrl
}

func InitServer(config utils.Config, db *utils.InMemoryDb) (Server, error) {
	log.Println("Initilize server")
	server := Server{
		config:   config,
		db:       db,
		engine:   gin.Default(),
		tapiCtrl: new(controllers.TapiCtrlInMemDB),
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
	server.engine.POST("/tapi-common:context/tapi-connectivity:connectivity-context/", server.CreateConnSrvRoute)
	server.engine.GET("/tapi-common:context/tapi-connectivity:connectivity-context/", server.GetAllConnSrvRoute)
	server.engine.GET("/tapi-common:context/tapi-connectivity:connectivity-context/:uuid", server.GetConnSrvRoute)
	server.engine.PATCH("/tapi-common:context/tapi-connectivity:connectivity-context/:uuid", server.UpdateConSrvRoute)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
