package movegeneration

import (
	"math"

	"github.com/healeycodes/chess-bot/tables"
	"github.com/notnil/chess"
)

// BestMove returns next strongest move
func BestMove(game *chess.Game, depth int) *chess.Move {
	bestValue := -math.MaxInt32
	bestMove := &chess.Move{}
	for _, move := range game.ValidMoves() {
		clone := game.Clone()
		clone.Move(move)
		value := alphabeta(clone, depth, -math.MaxInt32, math.MaxInt32, true)
		if value >= bestValue {
			bestValue = value
			bestMove = move
		}
	}
	return bestMove
}

func alphabeta(game *chess.Game, depth int, alpha int, beta int, maximizingPlayer bool) int {
	if depth == 0 || game.Position().Status() == chess.Checkmate {
		return -tables.EvaluateBoard(game)
	}
	if maximizingPlayer {
		value := -math.MaxInt32
		for _, move := range game.ValidMoves() {
			clone := game.Clone()
			clone.Move(move)
			value = max(value, alphabeta(clone, depth-1, alpha, beta, false))
			alpha := max(alpha, value)
			if alpha >= beta {
				break
			}
		}
		return value
	}

	// We're the minimizing player
	value := math.MaxInt32
	for _, move := range game.ValidMoves() {
		clone := game.Clone()
		clone.Move(move)
		value = min(value, alphabeta(clone, depth-1, alpha, beta, true))
		beta := min(beta, value)
		if beta <= alpha {
			break
		}
	}
	return value
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
