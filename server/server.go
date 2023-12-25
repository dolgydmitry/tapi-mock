package server

import (
	"fmt"
	"log"
	"tapi-controller/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config utils.Config
	db     utils.InMemoryDb
	engine *gin.Engine
}

func InitServer(config utils.Config, db utils.InMemoryDb) (Server, error) {
	log.Println("Initilize server")
	server := Server{
		config: config,
		db:     db,
		engine: gin.Default(),
	}
	// add route in this place

	return server, nil
}

func (server *Server) Start() error {
	host := fmt.Sprintf("%s:%d", server.config.Address, server.config.Port)
	log.Printf("run server on the host: %s", host)
	return server.engine.Run(host)
}
