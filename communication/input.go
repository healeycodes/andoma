package communication

import (
	"bufio"
	"fmt"
	"os"
)

func Listen() {
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	}

	// fen, _ := chess.FEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	// game := chess.NewGame(fen)
	// depth := 100

	// fmt.Println(BestMove(game, depth))
}
