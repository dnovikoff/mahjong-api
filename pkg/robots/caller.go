package robots

import (
	"github.com/dnovikoff/tempai-core/tile"

	proto_api "github.com/dnovikoff/mahjong-api/genproto/api"
	"github.com/dnovikoff/mahjong-api/pkg/convert"
)

type CallFunc func(*Caller, *proto_api.Server, *proto_api.Client) bool

type Caller struct {
	Tracker
	Sequence []CallFunc
}

func NewCaller(s ...CallFunc) *Caller {
	return &Caller{NewTracker(), s}
}

func Left(c *Caller, req *proto_api.Server, res *proto_api.Client) bool {
	if !req.GetSuggest().GetChiLeft() {
		return false
	}
	t := c.tile(req)
	return c.call(res, t+1, t+2)
}

func Center(c *Caller, req *proto_api.Server, res *proto_api.Client) bool {
	if !req.GetSuggest().GetChiCenter() {
		return false
	}
	t := c.tile(req)
	return c.call(res, t-1, t+1)
}

func Right(c *Caller, req *proto_api.Server, res *proto_api.Client) bool {
	if !req.GetSuggest().GetChiRight() {
		return false
	}
	t := c.tile(req)
	return c.call(res, t-2, t-1)
}

func Pon(c *Caller, req *proto_api.Server, res *proto_api.Client) bool {
	if !req.GetSuggest().GetPon() {
		return false
	}
	t := c.tile(req)
	return c.call(res, t, t)
}

func Kan(c *Caller, req *proto_api.Server, res *proto_api.Client) bool {
	if !req.GetSuggest().GetKan() {
		return false
	}
	t := c.tile(req)
	return c.call(res, t, t, t)
}

func (c *Caller) tile(req *proto_api.Server) tile.Tile {
	return convert.Instance(req.GetDrop().GetInstance()).Tile()
}

func (c *Caller) Request(req *proto_api.Server) *proto_api.Client {
	res := c.Tracker.Request(req)
	d := req.GetDrop()
	if d == nil || req.Suggest == nil || c.ClientIndex == d.WhoIndex {
		return res
	}
	for _, f := range c.Sequence {
		if f(c, req, res) {
			return res
		}
	}
	return res
}

func (c *Caller) call(res *proto_api.Client, tiles ...tile.Tile) bool {
	cp := c.Hand.Clone()
	var result tile.Instances
	types := make(map[tile.Type]bool, 4)
	for _, v := range tiles {
		first := cp.GetMask(v).First()
		if first == tile.InstanceNull {
			return false
		}
		cp.Remove(first)
		types[first.Tile().Type()] = true
		result = append(result, first)
	}
	if len(types) != 1 {
		return false
	}
	res.OneofClient = &proto_api.Client_Call{
		convert.ProtoInstances(result),
	}
	return true
}
