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
	match := &match{fen: initialFen}

	for true {
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		println(">>>", text)

		quit := Commands(text, match)
		if quit {
			break
		}
	}
}

// Commands accepts a command and the current match
func Commands(text string, match *match) bool {
	depth := 2 // TODO: accept this as an option/config value somewhere

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

	if strings.Contains(text, "position startpos moves") {
		moves := strings.Split(text, " ")[3:]
		game := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))
		for _, move := range moves {
			err := game.MoveStr(move)
			if err != nil {
				println("Move decode error:")
				println(err)
			}
		}
		match.fen = game.Position().String()
		return false
	}

	if strings.Contains(text, "position fen") {
		if len(strings.Split(text, " ")) > 2 {
			match.fen = strings.Join(strings.Split(text, " ")[2:], " ")
		}
		return false
	}

	if strings.Contains(text[:2], "go") {
		fen, _ := chess.FEN(match.fen)
		game := chess.NewGame(fen)
		fmt.Println("bestmove", movegeneration.BestMove(depth, game, true))
		return false
	}

	if text == "quit" {
		return true
	}

	return false
}
