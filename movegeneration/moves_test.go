package movegeneration

import (
	"fmt"
	"testing"

	"github.com/notnil/chess"
)

func TestAvoidingValueForCheckmate(t *testing.T) {
	// White should avoid taking a pawn and instead seek the
	// checkmate that is two moves away. Black cannot avoid this.
	fen, _ := chess.FEN("k7/8/5ppQ/8/8/8/8/1R2K3 w - - 0 1")
	game := chess.NewGame(fen)
	game.Move(BestMove(game, 10))
	game.Move(BestMove(game, 10))
	game.Move(BestMove(game, 10))

	if game.Position().Status() != chess.Checkmate {
		t.Errorf("It could not find checkmate")
	}

	// Black should avoid taking a pawn and instead seek the
	// checkmate that is two moves away. White cannot avoid this.
	fen, _ = chess.FEN("3k4/8/8/8/5P2/1r2P3/r7/4K3 b - - 0 1")
	game = chess.NewGame(fen)
	game.Move(BestMove(game, 10))
	game.Move(BestMove(game, 10))
	game.Move(BestMove(game, 10))

	if game.Position().Status() != chess.Checkmate {
		t.Errorf("It could not find checkmate")
	}
}

func TestTradeForValue(t *testing.T) {
	// Black will trade a bishop for a queen
	fen, _ := chess.FEN("rnbqk1nr/p1ppppbp/1p4p1/8/2P5/2Q5/PP1PPPPP/RNB1KBNR b KQkq - 0 1")
	game := chess.NewGame(fen)
	move := BestMove(game, 3)

	notation := fmt.Sprintf("%v", move)
	if notation != "g7c3" {
		t.Errorf("It returned the wrong move %v", notation)
	}

	// White will trade a bishop for a queen
	fen, _ = chess.FEN("rnb1k1nr/p1p1ppb1/1p1p2pp/6q1/2P5/3P4/PPQ1PPPP/RNB1KBNR w KQkq - 0 1")
	game = chess.NewGame(fen)
	move = BestMove(game, 3)

	notation = fmt.Sprintf("%v", move)
	if notation != "c1g5" {
		t.Errorf("It returned the wrong move %v", notation)
	}
}
