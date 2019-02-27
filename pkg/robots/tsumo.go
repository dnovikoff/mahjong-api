package robots

import (
	proto_api "github.com/dnovikoff/mahjong-api/genproto/api"
)

// Tsumo is a Simple robot, rejecting all incomming suggest
// This should be a terminal decider of any robot
type Tsumo struct {
}

func (t *Tsumo) Request(req *proto_api.Server) *proto_api.Client {
	if req.Suggest == nil || req.Suggest.Canceled {
		return nil
	}
	return &proto_api.Client{
		SuggestId: req.Suggest.SuggestId,
		OneofClient: &proto_api.Client_Cancel{
			Cancel: true,
		},
	}
}
