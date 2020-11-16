package main

import (
	"fmt"
	"testing"
)

func TestFirstMove(t *testing.T) {
	move := FirstMove()

	if fmt.Sprintf("%v", move) != "b1a3" {
		t.Errorf("It returned the wrong first move")
	}
}
