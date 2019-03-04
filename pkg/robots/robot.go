package robots

import (
	proto_game "github.com/dnovikoff/mahjong-api/genproto/public/game"
)

type Robot interface {
	Request(*proto_game.Server) *proto_game.Client
}
