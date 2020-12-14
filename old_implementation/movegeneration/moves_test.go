package movegeneration

import (
	"fmt"
	"testing"

	"github.com/notnil/chess"
)

func TestTakeUndefendedRook(t *testing.T) {
	// Black bishop should take a undefended rook
	fen, _ := chess.FEN("3r4/8/1R4pk/1P3p1p/3bn2P/3R2P1/6K1/3B4 b - - 0 1")
	game := chess.NewGame(fen)
	move := BestMove(3, game, true)

	notation := fmt.Sprintf("%v", move)
	if notation != "d4b6" {
		t.Errorf("It returned the wrong move %v", notation)
	}
}

func TestTradeForValue(t *testing.T) {
	// Black will trade a bishop for a queen
	fen, _ := chess.FEN("rnbqk1nr/p1ppppbp/1p4p1/8/2P5/2Q5/PP1PPPPP/RNB1KBNR b KQkq - 0 1")
	game := chess.NewGame(fen)
	move := BestMove(3, game, true)

	notation := fmt.Sprintf("%v", move)
	if notation != "g7c3" {
		t.Errorf("It returned the wrong move %v", notation)
	}
}

func TestTakeQueenOverPawnAdvancement(t *testing.T) {
	// Black should take a free queen
	fen, _ := chess.FEN("1nbqkb2/3ppp2/8/Q1p4p/4PB1N/2NB4/P4KP1/R6R b - - 0 1")
	game := chess.NewGame(fen)
	move := BestMove(3, game, true)

	notation := fmt.Sprintf("%v", move)
	if notation != "d8a5" {
		t.Errorf("It returned the wrong move %v", notation)
	}
}

// TODO: Make the engine aware of fork opportunities
// func TestForkKingAndRook(t *testing.T) {
// 	// White knight should fork the black king/rook
// 	fen, _ := chess.FEN("8/3R1pk1/3N1rp1/3K3p/5P1P/6P1/8/4b3 w - - 0 1")
// 	game := chess.NewGame(fen)
// 	move := BestMove(game, 4)
// 	game.Move(move)
// 	moveTwo := BestMove(game, 4)

// 	// Fork
// 	notation := fmt.Sprintf("%v", move)
// 	if notation != "d5e8" {
// 		t.Errorf("It returned the wrong move %v", notation)
// 	}

// 	// Take rook
// 	notation = fmt.Sprintf("%v", moveTwo)
// 	if notation != "e8f6" {
// 		t.Errorf("It returned the wrong move %v", notation)
// 	}
// }
