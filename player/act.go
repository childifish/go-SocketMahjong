package player

import (
	"go-SocketMahjong/table"
	"go-SocketMahjong/tile"
)

type Player struct {
	Position int
	Point int
	HandCard []tile.Tile
	Wall *table.Wall
}

func (p *Player)DrawCard(num int)  {
	tiles, bool := p.Wall.DrawTiles(num)
	if bool{

	}
	p.HandCard = append(p.HandCard,tiles...)
	p.Sort()
}

func (p *Player)PlayCard(i int)  {
	p.HandCard = append(p.HandCard[:i], p.HandCard[i+1:]...)
	p.Sort()
}

func (p *Player)AddPoint(point int)  {
	p.Point += point
}

func (p *Player)Sort()  {
	tile.Sort(p.HandCard)
}