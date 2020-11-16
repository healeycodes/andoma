package main

import (
	"fmt"

	"github.com/notnil/chess"
)

func main() {
	firstMove := FirstMove()
	fmt.Println("Hello, world.", firstMove)
}

// FirstMove provides one valid move
func FirstMove() *chess.Move {
	game := chess.NewGame()
	moves := game.ValidMoves()
	game.Move(moves[0])
	return moves[0] // b1a3
}
