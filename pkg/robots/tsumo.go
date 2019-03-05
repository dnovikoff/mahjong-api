package robots

import (
	proto_game "github.com/dnovikoff/mahjong-api/genproto/public/game"
)

// Tsumo is a Simple robot, rejecting all incomming suggest
// This should be a terminal decider of any robot
type Tsumo struct {
}

func (t *Tsumo) Request(req *proto_game.Server) *proto_game.Client {
	if req.Suggest == nil || req.Suggest.Canceled {
		return nil
	}
	return &proto_game.Client{
		SuggestId: req.Suggest.SuggestId,
		OneofClient: &proto_game.Client_Cancel{
			Cancel: true,
		},
	}
}
