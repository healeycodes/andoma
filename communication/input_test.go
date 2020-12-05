package communication

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestUCIReady(t *testing.T) {
	// Mock stdin and capture stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Send 2 UCI messages
	messages := "uci\nquit\n"
	reader := bufio.NewReader(strings.NewReader(messages))
	Listen(reader)

	w.Close()
	out, _ := ioutil.ReadAll(r)

	s := string(out)

	// Does the engine respond that it's ready?
	// The first messages will engine information
	reply := strings.Split(s, "\n")
	lastMsg := reply[len(reply)-2]
	if lastMsg != "uciok" {
		t.Errorf("It didn't respond to UCI messages %v", lastMsg)
	}
}

func TestUCIFEN(t *testing.T) {
	// Mock stdin and capture stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Send 4 UCI messages
	messages := "uci\nposition fen 3r4/8/1R4pk/1P3p1p/3bn2P/3R2P1/6K1/3B4 b - - 0 1\ngo\nquit"
	reader := bufio.NewReader(strings.NewReader(messages))
	Listen(reader)

	w.Close()
	out, _ := ioutil.ReadAll(r)

	s := string(out)

	// Does the engine reply with the obvious move (a free rook)
	reply := strings.Split(s, "\n")
	lastMsg := reply[len(reply)-2]
	if lastMsg != "bestmove d4b6" { //
		t.Errorf("It didn't reply with the correct move %v", lastMsg)
	}
}

func TestUCIMoves(t *testing.T) {
	// Mock stdin and capture stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Send 4 UCI messages
	messages := "uci\nposition startpos e2e4\ngo\nquit"
	reader := bufio.NewReader(strings.NewReader(messages))
	Listen(reader)

	w.Close()
	out, _ := ioutil.ReadAll(r)

	s := string(out)

	// Does the engine reply with any move?
	reply := strings.Split(s, "\n")
	lastMsg := reply[len(reply)-2]
	if !strings.Contains(lastMsg, "bestmove") && len(lastMsg) > 9 { //
		t.Errorf("It didn't reply with the correct move %v", lastMsg)
	}
}
