package main

import (
	"fmt"

	"github.com/notnil/chess"
)

func main() {
	fmt.Println("Hello, world.")

	game := chess.NewGame()
	moves := game.ValidMoves()
	game.Move(moves[0])
	fmt.Println(moves[0])
}
