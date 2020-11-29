package communication

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/healeycodes/chess-bot/movegeneration"
	"github.com/notnil/chess"
)

type match struct {
	fen string
}

var initialFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

// Listen starts listening for UCI commands and acts on them
func Listen(reader *bufio.Reader) {
	currentMatch := &match{fen: initialFen}

	for true {
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		println(">>>", text)

		quit := Commands(text, currentMatch)
		if quit {
			break
		}
	}
}

// Commands accepts a command and the current match
func Commands(text string, currentMatch *match) bool {

	if text == "uci" {
		fmt.Println("id name Andoma") // Andrew/Roma -> And/oma
		fmt.Println("id author Andrew Healey & Roma Parramore")
		fmt.Println("uciok")
		return false
	}

	if text == "isready" {
		fmt.Println("readyok")
		return false
	}

	if text == "ucinewgame" {
		return false
	}

	if strings.Contains(text, "position fen") {
		if len(strings.Split(text, "")) > 2 {
			currentMatch.fen = strings.Split(text, "")[2]
		}
		return false
	}

	if text == "go" {
		fen, _ := chess.FEN(currentMatch.fen)
		game := chess.NewGame(fen)
		depth := 30 // TODO: accept this as an option/config value somewhere

		fmt.Println("bestmove", movegeneration.BestMove(game, depth))
		return false
	}

	if text == "quit" {
		return true
	}

	return false
}
