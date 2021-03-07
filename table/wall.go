package table

import (
	"go-SocketMahjong/tile"
	ri_chi "go-SocketMahjong/tile/ri-chi"
	"math/rand"
)



type RichiWall struct {
	Tiles []ri_chi.Richi
}

func (rw *RichiWall) Init()[]tile.Tile {
	richi := ri_chi.Richi{}
	rw.Tiles = richi.ReturnTilesSelfArray()
	var t []tile.Tile
	for _, i2 := range rw.Tiles {
		t = append(t,&i2)
	}
	return t
}

func (rw *RichiWall) SpecialRule() bool {
	return false
}

//洗牌
func (rw *RichiWall)Shuffle()[]tile.Tile  {
	//Fisher-Yates随机置乱
	for i := len(rw.Tiles) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		rw.Tiles[i], rw.Tiles[num] = rw.Tiles[num], rw.Tiles[i]
	}
	var t []tile.Tile
	for _, i2 := range rw.Tiles {
		t = append(t,&i2)
	}
	return t
}

func (rw *RichiWall)Draw()(tile tile.Tile)  {
	////获取n张牌
	//for i:=0;i<num;i++{
	//	//c.CheckShuffle()
	//	tile = append(re,c.MainDeck[len(c.MainDeck)-1])
	//	c.MainDeck = c.MainDeck[:len(c.MainDeck)-1]
	//}
	//return re
	return tile
}

func CreateWallRiChi()  {
	
}
