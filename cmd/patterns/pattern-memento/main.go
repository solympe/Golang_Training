package main

import (
	"fmt"

	h "github.com/solympe/Golang_Training/pkg/patterns/pattern-memento/history"
	p "github.com/solympe/Golang_Training/pkg/patterns/pattern-memento/palette"
)

func main() {
	pallete := p.NewPaletteChanger()
	history := h.NewHistoryService()
	pallete.SetHistory(history)

	pallete.SetColor("red")
	pallete.SetColor("green")
	pallete.SetColor("yellow")

	pallete.SetPreviousColor()
	fmt.Println(pallete.GetColor())

	pallete.SetPreviousColor()
	fmt.Println(pallete.GetColor())

	pallete.SetPreviousColor()
	fmt.Println(pallete.GetColor())

	pallete.SetPreviousColor()
	fmt.Println(pallete.GetColor())
}
