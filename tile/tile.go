package tile

import "sort"

type Tile interface {
	PrintCard()
	PrintCardOrigin()
	Number() int
}

type Tiles []Tile

type KindTile interface {
	ReturnTilesInterfaceArray()[]Tile
	ReturnTilesIntArray()[]int
	RemainTiles()int
	Tiles2Analyzer(tiles []Tile)Analyzer
}

func (tiles Tiles) Len() int { return len(tiles) }
func (tiles Tiles) Swap(i, j int){ tiles[i], tiles[j] = tiles[j], tiles[i] }
func (tiles Tiles) Less(i, j int) bool {
	return tiles[i].Number() < tiles[j].Number()
}

func Sort(tiles Tiles)  {
	sort.Sort(tiles)
}

type Analyzer struct {
	Million []Tile
	Stick []Tile
	Plate []Tile
	Wind []Tile
	Dragon []Tile
}
