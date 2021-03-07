package table

import "go-SocketMahjong/tile"

type Wall interface {
	//初始化，返回牌山
	Init()
	//洗牌，返回牌山
	Shuffle()
	//玩家抽
	Draw()tile.Tile
	//特殊规则 如翻牌等
	//后续可能需要重构为
	SpecialRule()bool
	
}

type RichiWall struct {
	
}

func (rw *RichiWall) Init() {
	

}

func (rw *RichiWall) SpecialRule() bool {
	return false
}

func (rw *RichiWall)Shuffle()  {
	
}

func (rw *RichiWall)Draw()(tile tile.Tile)  {
	
	return 
}

func CreateWallRiChi()  {
	
}
