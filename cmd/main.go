package main

import (
	"go-SocketMahjong/table"
	"math/rand"
	"time"
)

func main()  {
	var w table.Wall
	w.Init()
	w.PrintWall()
	w.Shuffle()
	w.PrintWall()
}

func init() {
	rand.Seed(time.Now().Unix())
}