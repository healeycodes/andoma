package tables

import (
	"fmt"
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
	switch piece {
	case chess.WhiteKing:
		if endgame {
			return kingEndGame[square]
		}
		return kingMidGame[square]
	case chess.WhiteQueen:
		return queen[square]
	case chess.WhiteRook:
		return rook[square]
	case chess.WhiteBishop:
		return bishop[square]
	case chess.WhiteKnight:
		return knight[square]
	case chess.WhitePawn:
		return pawn[square]
	case chess.BlackKing:
		if endgame {
			return kingEndGame[len(kingEndGame)-square-1]
		}
		return kingMidGame[len(kingMidGame)-square-1]
	case chess.BlackQueen:
		return queen[len(queen)-square-1]
	case chess.BlackRook:
		return rook[len(rook)-square-1]
	case chess.BlackBishop:
		return bishop[len(bishop)-square-1]
	case chess.BlackKnight:
		return knight[len(knight)-square-1]
	case chess.BlackPawn:
		return pawn[len(pawn)-square-1]
	}

	fmt.Println("Shouldn't be reachable")
	return 0
}

// Tables lifted from and are from White's POV
// https://www.chessprogramming.org/Simplified_Evaluation_Function

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
