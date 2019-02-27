package rules

import (
	proto_rules "github.com/dnovikoff/mahjong-api/genproto/rules"
	"github.com/dnovikoff/mahjong-api/pkg/convert"
	"github.com/dnovikoff/tempai-core/score"
	"github.com/dnovikoff/tempai-core/tile"
	"github.com/dnovikoff/tempai-core/yaku"
)

type rulesWrapper struct {
	*proto_rules.Ruleset
}

var _ yaku.Rules = rulesWrapper{}
var _ score.Rules = rulesWrapper{}

type Wrapper interface {
	yaku.Rules
	score.Rules
}

func Wrap(x *proto_rules.Ruleset) Wrapper {
	return rulesWrapper{x}
}

// score.Rules
func (w rulesWrapper) ManganRound() bool {
	return w.Scoring.ManganRound
}

func (w rulesWrapper) KazoeYakuman() bool {
	return w.Scoring.KazoeYakuman
}

func (w rulesWrapper) IsDoubleYakuman(y yaku.Yakuman) bool {
	py := convert.ProtoYakuman(y)
	for _, v := range w.Scoring.DoubleYakumans {
		if py == v {
			return true
		}
	}
	return false
}

func (w rulesWrapper) YakumanSum() bool {
	return w.Scoring.YakumanSum
}

func (w rulesWrapper) Honba() score.Money {
	return score.Money(w.Scoring.HonbaMoney)
}

// yaku.Rules
func (w rulesWrapper) OpenTanyao() bool {
	return w.Yaku.OpenTanyao
}

func (w rulesWrapper) CheckAka(i tile.Instance) bool {
	p := convert.ProtoInstance(i)
	for _, v := range w.Yaku.GetAkaDora().GetInstances() {
		if p == v {
			return true
		}
	}
	return false
}

func (w rulesWrapper) Renhou() yaku.Limit {
	return convert.Limit(w.Yaku.Renhou)
}

func (w rulesWrapper) HaiteiFromLiveOnly() bool {
	return w.Yaku.HaiteiFromLiveOnly
}

func (w rulesWrapper) Ura() bool {
	return w.Yaku.Ura
}

func (w rulesWrapper) Ipatsu() bool {
	return w.Yaku.Ipatsu
}

func (w rulesWrapper) GreenRequired() bool {
	return w.Yaku.GreenRequired
}

func (w rulesWrapper) RinshanFu() bool {
	return w.Yaku.RinshanFu
}
