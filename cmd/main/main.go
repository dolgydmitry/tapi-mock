package main

import (
	"log"
	"tapi-controller/server"
	"tapi-controller/utils"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewColorFigure("TAPI simulator", "", "yellow", true)
	myFigure.Print()

	inMemoryDB := utils.InitMemoryDb()
	config, err := utils.LoadConfig()
	if err != nil {
		log.Panicf("cannot load config, err: %x", err)
	}
	server, err := server.InitServer(config, inMemoryDB)
	if err != nil {
		log.Panicf("cannot initilaize server, err: %x", err)
	}
	err = server.Start()
	if err != nil {
		log.Panicf("cannot run server, err: %x", err)
	}

}
