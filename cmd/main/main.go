package main

import (
	"fmt"
	"tapi-controller/utils"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewColorFigure("TAPI controller simulator", "", "yellow", true)
	myFigure.Print()

	inMemoryDB := utils.InitMemoryDb()
	fmt.Println(inMemoryDB)

}
