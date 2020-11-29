package communication

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestUCICommunication(t *testing.T) {
	// Mock stdin and capture stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Send two UCI messages
	messages := "uci\nquit\n"
	reader := bufio.NewReader(strings.NewReader(messages))
	Listen(reader)

	w.Close()
	out, _ := ioutil.ReadAll(r)

	s := string(out)

	// Does the engine respond that it's ready?
	if !strings.Contains(s, "uciok") {
		t.Errorf("It didn't respond to UCI messages %v", s)
	}
}
