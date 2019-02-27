package robots

import (
	proto_api "github.com/dnovikoff/mahjong-api/genproto/api"
)

type Robot interface {
	Request(*proto_api.Server) *proto_api.Client
}
