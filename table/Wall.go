package table

import (
	"fmt"
	"go-SocketMahjong/tile"
	"go-SocketMahjong/tile/ri-chi"
	"math/rand"
)

type Wall struct {
	MainTiles     []tile.Tile
	//MountainTiles []tile.Tile
	//DoraTiles []tile.Tile
	//UponTiles []tile.Tile
}

//初始化
func (w *Wall) Init()[]tile.Tile {
	Type := ChooseType()
	w.MainTiles = Type.ReturnTilesInterfaceArray()
	fmt.Println(w.MainTiles)
	return w.MainTiles
}

//打印牌山
func (w *Wall) PrintWall() {
	flag := 0
	fmt.Println("\n=====牌山=====")
	for i, i2 := range w.MainTiles {
		i2.PrintCardOrigin()
		if (i+1)%36 == 0{
			fmt.Println("")
		}
		if i+16 >w.len()&&flag==0{
			fmt.Println("\n======宝牌===")
			flag = 1
		}
		if i+6 >w.len()&&flag==1{
			fmt.Println("\n======岭上===")
			flag = 2
		}
	}

}

//抽牌
func (w *Wall)DrawTile()(tile.Tile,bool)  {
	if w.len()<=14{
		return nil,false
	}
	re := w.MainTiles[0]
	w.MainTiles = w.MainTiles[1:]
	return re,true
}

//抽很多张牌
func (w *Wall)DrawTiles(num int)([]tile.Tile,bool)  {
	var re []tile.Tile
	for i:=0;i<num;i++ {
		if w.len() <= 14 {
			//todo:埋点-状态写入

			return nil, false
		}
		re = append(re, w.MainTiles[w.len()-1])
		w.MainTiles = w.MainTiles[:w.len()-1]
	}
	return re,true
}

//返回长度
func (w *Wall)len() int  {
	return len(w.MainTiles)
}

//将牌堆分为几部分//大概是立直麻将专属，不知道该写在哪里
//func (w *Wall) BreakWall (){
//	w.MountainTiles = w.MainTiles[:len(w.MainTiles)-14]
//	w.DoraTiles = w.MainTiles[len(w.MainTiles)-14:len(w.MainTiles)-4]
//	w.UponTiles = w.MainTiles[len(w.MainTiles)-4:]
//}

func (w *Wall) Shuffle() {
	for i := w.len() - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		w.MainTiles[i], w.MainTiles[num] = w.MainTiles[num], w.MainTiles[i]
	}
}


//现在先做立直麻将
func ChooseType()tile.KindTile  {
	return &ri_chi.Richi{}
}