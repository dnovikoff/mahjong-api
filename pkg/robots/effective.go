package robots

import (
	"github.com/dnovikoff/tempai-core/hand/calc"
	"github.com/dnovikoff/tempai-core/hand/effective"
	"github.com/dnovikoff/tempai-core/tile"

	proto_api "github.com/dnovikoff/mahjong-api/genproto/api"
	"github.com/dnovikoff/mahjong-api/pkg/convert"
)

// Effective is a robot, that uses uke-ire to build its hand.
// Declares riichi when ready.
// Note, that this robot does not know how to win. See SettingsDecorator.
type Effective struct {
	Tracker
}

func NewEffective() *Effective {
	return &Effective{NewTracker()}
}

func (t *Effective) Request(req *proto_api.Server) *proto_api.Client {
	// Note, that the tracker must be informed of every request.
	// Still the result could be replace with any of yours.
	r := t.Tracker.Request(req)
	e := req.GetTake()
	if e == nil || req.Suggest == nil || e.WhoIndex != t.ClientIndex {
		return r
	}
	// fmt.Println("Hand debug: ", t.Hand.Instances().String(), t.Hand.CountBits())
	res := effective.Calculate(t.Hand,
		calc.Declared(convert.Melds(t.Melds)),
		calc.Used(t.Visible),
	)
	best := res.Sorted(t.Visible).Best()
	bestTile := tile.InstanceNull
	t.Hand.GetMask(best.Tile).Each(func(x tile.Instance) bool {
		bestTile = x
		return true
	})
	dm := convert.Mask(req.Suggest.DropMask)
	rm := convert.Mask(req.Suggest.RiichiMask)
	if dm.IsEmpty() {
		return r
	}
	r.OneofClient = &proto_api.Client_Drop{
		&proto_api.ClientDrop{
			Instance: convert.ProtoInstance(bestTile),
			Riichi:   rm.Check(bestTile.Tile()),
		},
	}
	return r
}
