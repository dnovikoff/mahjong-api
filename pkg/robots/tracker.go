package robots

import (
	"github.com/dnovikoff/tempai-core/compact"
	"github.com/dnovikoff/tempai-core/tile"

	proto_base "github.com/dnovikoff/mahjong-api/genproto/public/base"
	proto_game "github.com/dnovikoff/mahjong-api/genproto/public/game"
	proto_log "github.com/dnovikoff/mahjong-api/genproto/public/log"
	"github.com/dnovikoff/mahjong-api/pkg/convert"
)

// Tracker is a robot, tracking the hand state
// State is reset on a new round
// Use this robot as a helper struct, to implement others.
// Do not forget to pass every request to Tracker.
type Tracker struct {
	Tsumo
	ClientIndex int64
	Visible     compact.Instances
	Melds       []*proto_base.Meld
	Hand        compact.Instances
}

func NewTracker() Tracker {
	return Tracker{
		Visible: compact.NewInstances(),
		Hand:    compact.NewInstances(),
	}
}

func (t *Tracker) Request(req *proto_game.Server) *proto_game.Client {
	if start := req.GetGameStart(); start != nil {
		t.ClientIndex = start.ClientIndex
	}
	x := t.Tsumo.Request(req)
	t.gameStart(req.GetGameStart())
	t.rountStart(req.GetRoundStart())
	t.drop(req.GetDrop())
	t.declare(req.GetDeclare())
	t.take(req.GetTake())
	t.indicator(req.GetIndicator())
	return x
}

func (t *Tracker) gameStart(e *proto_game.GameStartEvent) {
	if e == nil {
		return
	}
	t.ClientIndex = e.ClientIndex
}

func (t *Tracker) rountStart(e *proto_log.RoundInfo) {
	if e == nil {
		return
	}
	t.Visible = compact.NewInstances()
	hands := e.GetWall().GetHands()
	t.Hand = convert.Hand(hands[t.ClientIndex-1])
	t.Visible.Merge(t.Hand)
}

func (t *Tracker) drop(e *proto_log.DropEvent) {
	if e == nil {
		return
	}
	i := convert.Instance(e.GetInstance())
	if i == tile.InstanceNull {
		return
	}
	t.Visible.Set(i)
	t.Hand.Remove(i)
}

func (t *Tracker) declare(e *proto_log.DeclareEvent) {
	if e == nil {
		return
	}
	withMeldTiles(e.Meld, func(i tile.Instance) {
		t.Visible.Set(i)
		t.Hand.Remove(i)
	})
	if t.ClientIndex == e.WhoIndex {
		t.Melds = append(t.Melds, e.Meld)
	}
}

func (t *Tracker) take(e *proto_log.TakeEvent) {
	if e == nil {
		return
	}
	i := convert.Instance(e.Instance)
	t.Visible.Set(i)
	if e.WhoIndex == t.ClientIndex {
		t.Hand.Set(i)
	}
}

func (t *Tracker) indicator(e *proto_log.IndicatorEvent) {
	if e == nil {
		return
	}
	t.Visible.Set(convert.Instance(e.Instance))
}

func withMeldTiles(x *proto_base.Meld, f func(x tile.Instance)) {
	set := func(x int64) {
		if x == 0 {
			return
		}
		f(convert.Instance(x))
	}

	set(x.CalledInstance)
	set(x.UpgradeInstance)
	for _, v := range x.GetHand().GetInstances() {
		set(v)
	}
}
