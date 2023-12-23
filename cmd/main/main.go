package main

import (
	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewColorFigure("TAPI controller", "", "yellow", true)
	myFigure.Print()

}
