package ri_chi

import (
	"fmt"
	"go-SocketMahjong/tile"
)

const (
	lizhi = iota
	daunyaojiu

)



func (r Richi)Tiles2Analyzer(tiles []tile.Tile)(analyzer tile.Analyzer)  {
	tile.Sort(tiles)
	var buffer int
	var finish int
	for key, value := range tiles {
		if value.Number() >= 36{
			analyzer.Million = append(tiles[0:key])
			buffer = key
			finish = 1
			continue
		}
		if value.Number() >= 72&&finish == 1{
			analyzer.Plate = append(tiles[buffer:key])
			buffer = key
			finish = 2
			continue
		}
		if value.Number() >= 108&&finish == 2{
			analyzer.Stick = append(tiles[buffer:key])
			buffer = key
			finish = 3
			continue
		}
		if value.Number() >= 125&&finish == 3{
			analyzer.Wind = append(tiles[buffer:key])
			buffer = key
			finish = 4
			continue
		}
		if finish == 4{
			analyzer.Plate = append(tiles[buffer:key])
		}
	}
	fmt.Println(analyzer)
	return analyzer
}