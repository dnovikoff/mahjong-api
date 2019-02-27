package robots

import (
	proto_api "github.com/dnovikoff/mahjong-api/genproto/api"
)

type DebugDecorator struct {
	robot Robot
	debug func(*proto_api.Server, *proto_api.Client)
}

func NewDebugDecorator(robot Robot, f func(*proto_api.Server, *proto_api.Client)) *DebugDecorator {
	return &DebugDecorator{robot, f}
}

func (d *DebugDecorator) Request(req *proto_api.Server) *proto_api.Client {
	res := d.robot.Request(req)
	d.debug(req, res)
	return res
}
