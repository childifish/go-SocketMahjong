package ground

import (
	"go-SocketMahjong/player"
	"go-SocketMahjong/table"
)

type Room struct {
	Players []player.Player
	Table table.Wall
}
