package convert

import (
	"sort"

	"github.com/dnovikoff/tempai-core/base"
	"github.com/dnovikoff/tempai-core/compact"
	"github.com/dnovikoff/tempai-core/hand/calc"
	"github.com/dnovikoff/tempai-core/score"
	"github.com/dnovikoff/tempai-core/tile"
	"github.com/dnovikoff/tempai-core/yaku"

	proto_base "github.com/dnovikoff/mahjong-api/genproto/public/base"
)

func ProtoIndex(i int) int64 {
	return int64(i) + 1
}

var limit = map[proto_base.Limit]yaku.Limit{
	proto_base.Limit_MANGAN:    yaku.LimitMangan,
	proto_base.Limit_HANEMAN:   yaku.LimitHaneman,
	proto_base.Limit_BAIMAN:    yaku.LimitBaiman,
	proto_base.Limit_SANBAIMAN: yaku.LimitSanbaiman,
	proto_base.Limit_YAKUMAN:   yaku.LimitYakuman,
}

var protoLimit = func() map[yaku.Limit]proto_base.Limit {
	x := make(map[yaku.Limit]proto_base.Limit, len(limit))
	for k, v := range limit {
		x[v] = k
	}
	return x
}()

func Limit(x proto_base.Limit) yaku.Limit {
	return limit[x]
}

func ProtoLimit(x yaku.Limit) proto_base.Limit {
	return protoLimit[x]
}

func ProtoWind(w base.Wind) proto_base.Wind {
	switch w {
	case base.WindEast:
		return proto_base.Wind_EAST
	case base.WindSouth:
		return proto_base.Wind_SOUTH
	case base.WindWest:
		return proto_base.Wind_WEST
	case base.WindNorth:
		return proto_base.Wind_NORTH
	}
	return proto_base.Wind_WIND_UNSPECIFIED
}

func ProtoInstance(input tile.Instance) int64 {
	return int64(input)
}

func Instance(input int64) tile.Instance {
	return tile.Instance(input)
}

func ProtoMoney(x score.Money) int64 {
	return int64(x)
}

func Instances(input []int64) tile.Instances {
	ret := make(tile.Instances, len(input))
	for k, v := range input {
		ret[k] = Instance(v)
	}
	return ret
}

func Hand(x *proto_base.Instances) compact.Instances {
	if x == nil {
		return nil
	}
	r := compact.NewInstances()
	for _, v := range x.Instances {
		r.Set(Instance(v))
	}
	return r
}

func ProtoInstances(input tile.Instances) *proto_base.Instances {
	return &proto_base.Instances{
		Instances: protoInstancesInts(input),
	}
}

func protoInstancesInts(input tile.Instances) []int64 {
	result := make([]int64, len(input))
	for k, v := range input {
		result[k] = ProtoInstance(v)
	}
	return result
}

func Money(x []int64) []score.Money {
	r := make([]score.Money, len(x))
	for k, v := range x {
		r[k] = score.Money(v)
	}
	return r
}

func Mask(x int64) compact.Tiles {
	return compact.Tiles(x)
}

func ProtoMask(x compact.Tiles) int64 {
	return int64(x)
}

func Meld(x *proto_base.Meld) calc.Meld {
	called := Instance(x.CalledInstance).Tile()
	// Only possible for closed kan
	if x.Opponent == proto_base.Opponent_SELF {
		return calc.Kan(called)
	}
	ln := len(x.GetHand().GetInstances())
	if x.UpgradeInstance != 0 || ln == 3 {
		return calc.Open(calc.Kan(called))
	}
	if ln != 2 {
		return nil
	}
	instances := x.GetHand().GetInstances()
	t1 := Instance(instances[0]).Tile()
	t2 := Instance(instances[1]).Tile()
	var m calc.Meld
	switch t2 - t1 {
	case 0:
		m = calc.Pon(called)
	case 2:
		m = calc.Chi(called)
	case 1:
		if called < t1 {
			m = calc.Chi(called)
		} else {
			m = calc.Chi(t1)
		}
	}
	if m == nil {
		return nil
	}
	return calc.Open(m)
}

func Melds(x []*proto_base.Meld) calc.Melds {
	out := make(calc.Melds, len(x))
	for k, v := range x {
		out[k] = Meld(v)
	}
	return out
}

// TODO: generate mappings
func ProtoYaku(y yaku.Yaku) proto_base.Yaku {
	return proto_base.Yaku(y)
}

func ProtoYakuman(y yaku.Yakuman) proto_base.Yakuman {
	return proto_base.Yakuman(y)
}

func ProtoYakumans(y yaku.Yakumans, rules score.Rules) []*proto_base.YakumanValue {
	x := make([]*proto_base.YakumanValue, 0, len(y))
	for _, v := range y {
		m := int64(1)
		if rules.IsDoubleYakuman(v) {
			m = 2
		}
		x = append(x, &proto_base.YakumanValue{
			Yakuman:    ProtoYakuman(v),
			Multiplier: m})
	}
	// TODO: correct sorting
	sort.Slice(x, func(i, j int) bool {
		return x[i].Yakuman < x[j].Yakuman
	})
	return x
}

func ProtoYakus(y yaku.YakuSet) []*proto_base.YakuValue {
	x := make([]*proto_base.YakuValue, 0, len(y))
	for k, v := range y {
		x = append(x, &proto_base.YakuValue{
			Yaku: ProtoYaku(k),
			Han:  int64(v)})
	}
	// TODO: correct sorting
	sort.Slice(x, func(i, j int) bool {
		return x[i].Yaku < x[j].Yaku
	})
	return x
}
