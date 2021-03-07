package tile

type Tile interface {
	PrintCard()
	PrintCardOrigin()
}

type Tiles []Tile

type KindTile interface {
	ReturnTilesInterfaceArray()[]Tile
	ReturnTilesIntArray()[]int
}