package robots

import (
	"github.com/dnovikoff/tempai-core/compact"
	"github.com/dnovikoff/tempai-core/tile"

	proto_api "github.com/dnovikoff/mahjong-api/genproto/api"
	"github.com/dnovikoff/mahjong-api/pkg/convert"
)

// Zaichik is a robot, that is gready to declare kans
type Zaichik struct {
	Robot
}

func (z *Zaichik) Request(req *proto_api.Server) *proto_api.Client {
	res := z.Robot.Request(req)
	if req.GetSuggest().GetKan() && req.GetDrop() != nil {
		i := convert.Instance(req.GetDrop().Instance)
		cc := compact.NewInstances()
		cc.SetCount(i.Tile(), 4)
		cc.Remove(i)
		res.OneofClient = &proto_api.Client_Call{
			convert.ProtoInstances(cc.Instances()),
		}
	} else if req.GetSuggest().GetClosedKanMask() != 0 && req.GetTake() != nil {
		convert.Mask(req.GetSuggest().GetClosedKanMask()).Each(func(t tile.Tile) bool {
			cc := compact.NewInstances()
			cc.SetCount(t, 4)
			res.OneofClient = &proto_api.Client_Call{
				convert.ProtoInstances(cc.Instances()),
			}
			return false
		})

	}
	return res
}
