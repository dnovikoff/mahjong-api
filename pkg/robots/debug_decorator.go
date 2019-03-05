package robots

import (
	proto_game "github.com/dnovikoff/mahjong-api/genproto/public/game"
)

type DebugDecorator struct {
	robot Robot
	debug func(*proto_game.Server, *proto_game.Client)
}

func NewDebugDecorator(robot Robot, f func(*proto_game.Server, *proto_game.Client)) *DebugDecorator {
	return &DebugDecorator{robot, f}
}

func (d *DebugDecorator) Request(req *proto_game.Server) *proto_game.Client {
	res := d.robot.Request(req)
	d.debug(req, res)
	return res
}
