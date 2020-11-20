package main

import (
	"fmt"
	"math"

	"github.com/notnil/chess"
)

func main() {
	fen, _ := chess.FEN("rnbqkbnr/pppp1ppp/8/4p3/3PP3/8/PPP2PPP/RNBQKBNR b KQkq - 0 2")
	game := chess.NewGame(fen)
	depth := 2
	isPlayer := true
	bestMove(depth, game, isPlayer)
}

func bestMove(depth int, game *chess.Game, isPlayer bool) {
	moves := game.ValidMoves()

	bestValue := -math.MaxInt32
	bestMove := moves[0]
	fmt.Println(game.Position().Board().Draw())

	for _, move := range moves {
		clone := game.Clone()
		clone.Move(move)
		value := minimax(depth, clone, !isPlayer)
		if value >= bestValue {
			bestValue = value
			bestMove = move
		}
	}
	game.Move(bestMove)
	fmt.Println(game.Position().Board().Draw())
	fmt.Println(bestValue, "best value")
	fmt.Println(bestMove, "best move")
}

func minimax(depth int, game *chess.Game, isPlayer bool) int {

	if isPlayer {
		return maxi(depth-1, game)
	}

	return mini(depth-1, game)

}

func maxi(depth int, game *chess.Game) int {
	if depth == 0 {
		return -BoardValue(game.Position().Board())
	}

	moves := game.ValidMoves()
	max := -math.MaxInt32

	for _, move := range moves {
		clone := game.Clone()
		clone.Move(move)

		score := mini(depth-1, clone)
		if score > max {
			max = score
		}
	}
	// fmt.Println(max, "score, max in maxi")
	return max
}

func mini(depth int, game *chess.Game) int {
	if depth == 0 {
		return -BoardValue(game.Position().Board())
	}

	moves := game.ValidMoves()
	min := math.MaxInt32

	for _, move := range moves {
		clone := game.Clone()
		clone.Move(move)

		score := maxi(depth-1, clone)
		if score < min {
			min = score
		}
	}
	// fmt.Println(min, "score, min in mini")
	return min
}

func BoardValue(board *chess.Board) int {
	pieceValue := map[chess.PieceType]int{chess.Pawn: 100, chess.Bishop: 350, chess.King: 10000, chess.Knight: 350, chess.Queen: 1000, chess.Rook: 525}
	sum := 0

	for _, piece := range board.SquareMap() {
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
