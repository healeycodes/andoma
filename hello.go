package main

import (
	"fmt"
	"math"

	"github.com/notnil/chess"
	"github.com/healeycodes/chess-bot/tables"
)

func main() {
	fen, _ := chess.FEN("4kbnr/8/8/8/8/8/PPPPPPPP/RNBQKBNR w KQk - 0 1")
	game := chess.NewGame(fen)
	depth := 10
	isPlayer := true

	fmt.Println(game.Position().Board().Draw())

	for true {
		bestMove(game, depth, isPlayer)
		if game.Outcome() != "*" {
			fmt.Println(game.Outcome())
			break
		}
	}
}

func bestMove(game *chess.Game, depth int, isPlayer bool) {
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

	game.Move(bestMove)
	fmt.Println(game.Position().Board().Draw())
	// fmt.Println(bestValue, "best value")
	// fmt.Println(bestMove, "best move")
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

func BoardValue(game *chess.Game) int {
	if game.Position().Status() == chess.Checkmate {
		return math.MaxInt32
	}

	board := game.Position().Board()
	pieceValue := map[chess.PieceType]int{chess.Pawn: 100, chess.Bishop: 330, chess.King: 20000, chess.Knight: 320, chess.Queen: 900, chess.Rook: 500}
	sum := 0

	for i, piece := range board.SquareMap() {
		fmt.Println(tables.Evaluate(i, chess.PieceType))
		sum += pieceValue[piece.Type()]
	}
	return sum
}

// FirstMove provides one valid move
func FirstMove() *chess.Move {
	game := chess.NewGame()
	moves := game.ValidMoves()
	game.Move(moves[0])
	return moves[0] // b1a3
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
