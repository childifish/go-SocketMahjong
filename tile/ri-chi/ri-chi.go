package ri_chi

import (
	"errors"
	"fmt"
	"go-SocketMahjong/tile"
	"strconv"
)

const AllTileNum = 136

//配牌 条s筒p万m
const (
	Million = 0 	//万"m"
	Plate = 1		//筒"p"
	Stick = 2		//条"s"
	Dragon = 3		//三元"d"//风"w"
)

//字牌 东南西北中发白
const (
	East = 0
	South = 1
	West = 2
	North = 3
	Red = 4
	Green = 5
	White = 6
)

type Richi struct {
	Num int
	Flower int
}

func (r *Richi)ReturnTilesIntArray()[]int  {
	var TilesArray []int
	TilesArray = make([]int,AllTileNum)
	for i:=0;i<AllTileNum;i++{
		TilesArray[i] = i+1
	}
	return TilesArray
}

func (r *Richi)ReturnTilesInterfaceArray()[]tile.Tile  {
	var TilesArray []Richi
	var t []tile.Tile
	TilesArray = make([]Richi,AllTileNum)
	for i:=0;i<AllTileNum;i++{
		TilesArray[i].Flower = i+1
		TilesArray[i].Num = i%4+1
		t = append(t, TilesArray[i])
	}
	return t
}

func (r *Richi)ReturnTilesSelfArray()[]Richi  {
	var TilesArray []Richi
	TilesArray = make([]Richi,AllTileNum)
	for i:=0;i<AllTileNum;i++{
		TilesArray[i].Flower = i+1
		TilesArray[i].Num = i%4+1
	}
	return TilesArray
}

func (r Richi)PrintCard()  {
	fmt.Printf("%d-%s\n",r.Num,Analyze(r.Flower))
}

func (r Richi)PrintCardOrigin()  {
	fmt.Printf("%d-%s  ",r.Num,Analyze(r.Flower))
}

func Analyze(tileNum int)(chara string)  {
	switch (tileNum-1)/36{
	case Million:
		chara = analyzeNum((tileNum-1)/4)+"万"
		return chara
	case Plate:
		chara = analyzeNum((tileNum-1)/4)+"筒"
		return chara
	case Stick:
		chara = analyzeNum((tileNum-1)/4)+"条"
		return chara
	case Dragon:
		chara = analyzeDragon(tileNum)
		return chara
	default:
		fmt.Println("err in Analyze:",errors.New("too big num"))
		return ""
	}
}

func analyzeNum(num int)string  {
	return strconv.Itoa(num%9+1)
}

func analyzeDragon(num int)string  {
	switch (num - 109)/4 {
	case East:
		return "东"
	case South:
		return "南"
	case West:
		return "西"
	case North:
		return "北"
	case Red:
		return "中"
	case Green:
		return "发"
	case White:
		return "白"
	default:
		fmt.Println("err in Analyze:",errors.New("too big num"))
		return ""
	}
}