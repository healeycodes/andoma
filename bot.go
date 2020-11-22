package main

import (
	"fmt"
	"math"

	"github.com/healeycodes/chess-bot/tables"
	"github.com/notnil/chess"
)

func main() {
	fen, _ := chess.FEN("3k4/8/8/8/r4P2/4P3/r7/4K3 b - - 0 1")
	game := chess.NewGame(fen)
	depth := 10
	isPlayer := true

	// for true {
	// 	BestMove(game, depth, isPlayer)
	// 	if game.Outcome() != "*" {
	// 		fmt.Println(game.Outcome())
	// 		break
	// 	}
	// }
	fmt.Println(BestMove(game, depth, isPlayer))
}

// BestMove returns the strongest next move.
func BestMove(game *chess.Game, depth int, isPlayer bool) *chess.Move {
	bestValue := -math.MaxInt32
	bestMove := &chess.Move{}
	for _, move := range game.ValidMoves() {
		clone := game.Clone()
		clone.Move(move)
		value := minimax(clone, depth, !isPlayer)
		if value >= bestValue {
			bestValue = value
			bestMove = move
		}
	}

	fmt.Println(game.Position().Board().Draw())
	game.Move(bestMove)
	return bestMove
}

func minimax(game *chess.Game, depth int, isPlayer bool) int {
	if isPlayer {
		return maxi(game, depth-1, -math.MaxInt32, math.MaxInt32)
	}
	return mini(game, depth-1, math.MaxInt32, -math.MaxInt32)
}

func maxi(game *chess.Game, depth int, alpha int, beta int) int {
	if depth == 0 {
		return -BoardValue(game)
	}

	value := -math.MaxInt32
	for _, move := range game.ValidMoves() {
		clone := game.Clone()
		clone.Move(move)
		value = max(value, mini(clone, depth-1, alpha, beta))
		alpha = max(alpha, value)
		if alpha >= beta {
			break
		}
	}
	return value
}

func mini(game *chess.Game, depth int, alpha int, beta int) int {
	if depth == 0 {
		return -BoardValue(game)
	}

	value := math.MaxInt32
	for _, move := range game.ValidMoves() {
		clone := game.Clone()
		clone.Move(move)
		value = min(value, maxi(clone, depth-1, alpha, beta))
		beta = min(beta, value)
		if beta <= alpha {
			break
		}
	}
	return value
}

// BoardValue evaluates the board's value
func BoardValue(game *chess.Game) int {
	if game.Position().Status() == chess.Checkmate {
		return math.MaxInt32
	}

	sum := 0
	for square, piece := range game.Position().Board().SquareMap() {
		sum += tables.Evaluate(int(square), piece)
	}
	return sum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
