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
			return kingEndGame[square] + K
		}
		return kingMidGame[square] + K
	case chess.WhiteQueen:
		return queen[square] + Q
	case chess.WhiteRook:
		return rook[square] + R
	case chess.WhiteBishop:
		return bishop[square] + B
	case chess.WhiteKnight:
		return knight[square] + N
	case chess.WhitePawn:
		return pawn[square] + P
	case chess.BlackKing:
		if endgame {
			return kingEndGame[len(kingEndGame)-square-1] + K
		}
		return kingMidGame[len(kingMidGame)-square-1] + K
	case chess.BlackQueen:
		return queen[len(queen)-square-1] + Q
	case chess.BlackRook:
		return rook[len(rook)-square-1] + R
	case chess.BlackBishop:
		return bishop[len(bishop)-square-1] + B
	case chess.BlackKnight:
		return knight[len(knight)-square-1] + N
	case chess.BlackPawn:
		return pawn[len(pawn)-square-1] + P
	}

	println("Shouldn't be reachable")
	return 0
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
