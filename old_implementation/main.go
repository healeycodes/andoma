package main

import (
	"bufio"
	"os"

	"github.com/healeycodes/chess-bot/communication"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	communication.Listen(reader)
}
