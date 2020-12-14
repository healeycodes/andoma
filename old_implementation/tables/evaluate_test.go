package tables

import (
	"math"
	"testing"

	"github.com/notnil/chess"
)

func TestTwoEvaluations(t *testing.T) {
	game := chess.NewGame()
	evaluation := EvaluateBoard(game)

	fen, _ := chess.FEN("3r4/8/1R4pk/1P3p1p/3bn2P/3R2P1/6K1/3B4 b - - 0 1")
	gameTwo := chess.NewGame(fen)
	evaluationTwo := EvaluateBoard(gameTwo)

	if evaluation == evaluationTwo {
		t.Errorf("Different boards should have a different weighting %v %v", evaluation, evaluationTwo)
	}
}

func TestCheckmate(t *testing.T) {
	fen, _ := chess.FEN("3k4/8/8/8/8/3qqq2/P7/4K3 w - - 0 1")
	game := chess.NewGame(fen)
	evaluation := EvaluateBoard(game)

	if evaluation != math.MaxInt32 {
		t.Errorf("A checkmate state should be weighted at the maximum value %v", evaluation)
	}
}
