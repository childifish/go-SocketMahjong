package table

import (
	"fmt"
	"go-SocketMahjong/tile"
	"go-SocketMahjong/tile/ri-chi"
	"math/rand"
)

type Wall struct {
	MainTiles []tile.Tile
}

func (w *Wall) Init()[]tile.Tile {
	Type := ChooseType()
	w.MainTiles = Type.ReturnTilesInterfaceArray()
	fmt.Println(w.MainTiles)
	return w.MainTiles
}

func (w *Wall) PrintWall() {
	flag := 0
	fmt.Println("\n=====牌山=====")
	for i, i2 := range w.MainTiles {
		i2.PrintCardOrigin()
		if (i+1)%36 == 0{
			fmt.Println("")
		}
		if i+14 <len(w.MainTiles)&&flag==0{
			fmt.Println("\n======宝牌===")
			flag = 1
		}
		if i+4 <len(w.MainTiles)&&flag==1{
			fmt.Println("\n======岭上===")
			flag = 2
		}
	}

}

func (w *Wall) BreakWall (){

}

func (w *Wall) Shuffle() {
	for i := len(w.MainTiles) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		w.MainTiles[i], w.MainTiles[num] = w.MainTiles[num], w.MainTiles[i]
	}
}


//现在先做立直麻将
func ChooseType()tile.KindTile  {
	return &ri_chi.Richi{}
}