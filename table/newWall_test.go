package table

import "testing"

func TestWall_Init(t *testing.T) {
	var w Wall
	w.Init()
	w.PrintWall()
	w.Shuffle()
	w.PrintWall()
}
