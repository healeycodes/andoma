package main

import (
	"fmt"
	"testing"

	"github.com/notnil/chess"
)

func TestWhiteWillAvoidValueAndCheckmate(t *testing.T) {
	// White should avoid taking a pawn and instead seek the
	// checkmate that is two moves away. Black cannot avoid this.
	fen, _ := chess.FEN("k7/8/5ppQ/8/8/8/8/1R2K3 w - - 0 1")
	game := chess.NewGame(fen)
	move := BestMove(game, 10, true)

	notation := fmt.Sprintf("%v", move)
	if notation != "h6h7" {
		t.Errorf("It returned the wrong move %v", notation)
	}
}

func TestBlackWillAvoidValueAndCheckmate(t *testing.T) {
	// Black should avoid taking a pawn and instead seek the
	// checkmate that is two moves away. White cannot avoid this.
	fen, _ := chess.FEN("3k4/8/8/8/5P2/1r2P3/r7/4K3 b - - 0 1")
	game := chess.NewGame(fen)
	move := BestMove(game, 10, true)

	notation := fmt.Sprintf("%v", move)
	if notation != "b3b1" {
		t.Errorf("It returned the wrong move %v", notation)
	}
}
