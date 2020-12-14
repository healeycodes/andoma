package movegeneration

import (
	"math"

	"github.com/healeycodes/chess-bot/tables"
	"github.com/notnil/chess"
)

// BestMove returns the best move available
func BestMove(depth int, game *chess.Game, isMaximisingPlayer bool) *chess.Move {
	bestMove := -math.MaxInt32
	bestMoveFound := &chess.Move{}

	for _, move := range game.ValidMoves() {
		child := game.Clone()
		child.Move(move)
		value := minimax(depth-1, child, -math.MaxInt32, math.MaxInt32, !isMaximisingPlayer)
		if value >= bestMove {
			bestMove = value
			bestMoveFound = move
		}
	}
	return bestMoveFound
}

func minimax(depth int, game *chess.Game, alpha int, beta int, isMaximisingPlayer bool) int {
	if depth == 0 || game.Position().Status() == chess.Checkmate {
		return -tables.EvaluateBoard(game)
	}

	if isMaximisingPlayer {
		bestMove := -math.MaxInt32
		for _, move := range game.ValidMoves() {
			child := game.Clone()
			child.Move(move)
			bestMove := max(bestMove, minimax(depth-1, game, alpha, beta, !isMaximisingPlayer))
			alpha = max(alpha, bestMove)
			if beta <= alpha {
				return bestMove
			}
		}
		return bestMove
	}

	bestMove := math.MaxInt32
	for _, move := range game.ValidMoves() {
		child := game.Clone()
		child.Move(move)
		bestMove := min(bestMove, minimax(depth-1, game, alpha, beta, !isMaximisingPlayer))
		beta := min(beta, bestMove)
		if beta <= alpha {
			return bestMove
		}
	}
	return bestMove
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

func copy(game *chess.Game) *chess.Game {
	g := &chess.Game{}
	g.tagPairs = game.TagPairs()
	g.moves = game.Moves()
	g.positions = game.Positions()
	g.pos = game.pos
	g.outcome = game.outcome
	g.method = game.method
}
