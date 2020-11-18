package main

import (
	"fmt"

	"github.com/notnil/chess"
)

func main() {
	fen, _ := chess.FEN("rnbqkbnr/pppp1ppp/8/4p3/5P2/4P3/PPPP2PP/RNBQKBNR b KQkq - 0 2")
	game := chess.NewGame(fen)
	moves := game.ValidMoves()

	bestValue := BoardValue(game.Position().Board())
	bestMove := moves[0]

	for _, move := range moves {
		clone := game.Clone()
		clone.Move(move)
		boardValue := BoardValue(clone.Position().Board())
		if boardValue < bestValue {
			bestValue = boardValue
			bestMove = move
		}
	}
	fmt.Println(bestValue, "best value")
	fmt.Println(bestMove, "best move")
}

func BoardValue(board *chess.Board) int {
	pieceValue := map[chess.PieceType]int{chess.Pawn: 1, chess.Bishop: 0, chess.King: 0, chess.Knight: 0, chess.Queen: 0, chess.Rook: 0}
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
