package main

import (
	"fmt"

	h "github.com/solympe/Golang_Training/pkg/pMemento/history"
	p "github.com/solympe/Golang_Training/pkg/pMemento/palette"
)

func main() {
	pallete := p.NewPalette()
	history := h.NewHistory()
	pallete.SetHistory(history)

	pallete.SetColor("red")
	pallete.SetColor("green")

	fmt.Println(pallete.GetColor())
	pallete.SetPreviousColor()
	fmt.Println(pallete.GetColor())

}