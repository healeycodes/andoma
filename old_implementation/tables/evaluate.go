package tables

import (
	"math"

	"github.com/notnil/chess"
)

// EvaluateBoard implements the Simplified Evaluation Function by Tomasz Michniewski
// https://www.chessprogramming.org/Simplified_Evaluation_Function
func EvaluateBoard(game *chess.Game) int {
	if game.Position().Status() == chess.Checkmate {
		return math.MaxInt32
	}

	zeroQueens := true
	for _, piece := range game.Position().Board().SquareMap() {
		if piece.Type() == chess.Queen {
			zeroQueens = false
		}
	}

	sum := 0
	for square, piece := range game.Position().Board().SquareMap() {
		// We determine the endgame as neither side having queens
		// In the future, we might use:
		// > Every side which has a queen has additionally no other pieces or one minorpiece maximum
		sum += evaluatePiece(int(square), piece, zeroQueens)
	}

	return sum
}

func evaluatePiece(square int, piece chess.Piece, endgame bool) int {
	var value int

	// Piece weights
	P := 100
	N := 320
	B := 330
	R := 500
	Q := 900
	K := 20000

	// Apply square bonus
	switch piece {
	case chess.WhiteKing:
		if endgame {
			value = kingEndGame[square] + K
		}
		value = kingMidGame[square] + K
	case chess.WhiteQueen:
		value = queen[square] + Q
	case chess.WhiteRook:
		value = rook[square] + R
	case chess.WhiteBishop:
		value = bishop[square] + B
	case chess.WhiteKnight:
		value = knight[square] + N
	case chess.WhitePawn:
		value = pawn[square] + P
	case chess.BlackKing:
		if endgame {
			value = kingEndGame[len(kingEndGame)-square-1] + K
		}
		value = kingMidGame[len(kingMidGame)-square-1] + K
	case chess.BlackQueen:
		value = queen[len(queen)-square-1] + Q
	case chess.BlackRook:
		value = rook[len(rook)-square-1] + R
	case chess.BlackBishop:
		value = bishop[len(bishop)-square-1] + B
	case chess.BlackKnight:
		value = knight[len(knight)-square-1] + N
	case chess.BlackPawn:
		value = pawn[len(pawn)-square-1] + P
	}

	if piece.Color() == 1 {
		return value
	}

	return -value
}

// Tables lifted from:
// https://www.chessprogramming.org/Simplified_Evaluation_Function
// They are from White's POV and should be flipped for Black

var pawn = [...]int{0, 0, 0, 0, 0, 0, 0, 0,
	50, 50, 50, 50, 50, 50, 50, 50,
	10, 10, 20, 30, 30, 20, 10, 10,
	5, 5, 10, 25, 25, 10, 5, 5,
	0, 0, 0, 20, 20, 0, 0, 0,
	5, -5, -10, 0, 0, -10, -5, 5,
	5, 10, 10, -20, -20, 10, 10, 5,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var knight = [...]int{-50, -40, -30, -30, -30, -30, -40, -50,
	-40, -20, 0, 0, 0, 0, -20, -40,
	-30, 0, 10, 15, 15, 10, 0, -30,
	-30, 5, 15, 20, 20, 15, 5, -30,
	-30, 0, 15, 20, 20, 15, 0, -30,
	-30, 5, 10, 15, 15, 10, 5, -30,
	-40, -20, 0, 5, 5, 0, -20, -40,
	-50, -40, -30, -30, -30, -30, -40, -50,
}

var bishop = [...]int{-20, -10, -10, -10, -10, -10, -10, -20,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-10, 0, 5, 10, 10, 5, 0, -10,
	-10, 5, 5, 10, 10, 5, 5, -10,
	-10, 0, 10, 10, 10, 10, 0, -10,
	-10, 10, 10, 10, 10, 10, 10, -10,
	-10, 5, 0, 0, 0, 0, 5, -10,
	-20, -10, -10, -10, -10, -10, -10, -20,
}

var rook = [...]int{0, 0, 0, 0, 0, 0, 0, 0,
	5, 10, 10, 10, 10, 10, 10, 5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	0, 0, 0, 5, 5, 0, 0, 0,
}

var queen = [...]int{-20, -10, -10, -5, -5, -10, -10, -20,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-10, 0, 5, 5, 5, 5, 0, -10,
	-5, 0, 5, 5, 5, 5, 0, -5,
	0, 0, 5, 5, 5, 5, 0, -5,
	-10, 5, 5, 5, 5, 5, 0, -10,
	-10, 0, 5, 0, 0, 0, 0, -10,
	-20, -10, -10, -5, -5, -10, -10, -20,
}

var kingMidGame = [...]int{-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-20, -30, -30, -40, -40, -30, -30, -20,
	-10, -20, -20, -20, -20, -20, -20, -10,
	20, 20, 0, 0, 0, 0, 20, 20,
	20, 30, 10, 0, 0, 10, 30, 20,
}

var kingEndGame = [...]int{-50, -40, -30, -20, -20, -30, -40, -50,
	-30, -20, -10, 0, 0, -10, -20, -30,
	-30, -10, 20, 30, 30, 20, -10, -30,
	-30, -10, 30, 40, 40, 30, -10, -30,
	-30, -10, 30, 40, 40, 30, -10, -30,
	-30, -10, 20, 30, 30, 20, -10, -30,
	-30, -30, 0, 0, 0, 0, -30, -30,
	-50, -30, -30, -30, -30, -30, -30, -50,
}
