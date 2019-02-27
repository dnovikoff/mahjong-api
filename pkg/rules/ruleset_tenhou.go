package rules

import (
	"github.com/dnovikoff/tempai-core/tile"

	proto_base "github.com/dnovikoff/mahjong-api/genproto/base"
	proto_rules "github.com/dnovikoff/mahjong-api/genproto/rules"
	"github.com/dnovikoff/mahjong-api/pkg/convert"
)

func TenhouRed() *proto_rules.Ruleset {
	return &proto_rules.Ruleset{
		Id: "TenhouRed",
		Scoring: &proto_rules.Scoring{
			ManganRound:  false,
			KazoeYakuman: true,
			YakumanSum:   true,
			HonbaMoney:   100,
		},
		Yaku: &proto_rules.Yaku{
			OpenTanyao:         true,
			HaiteiFromLiveOnly: true,
			Ura:                true,
			Ipatsu:             true,
			AkaDora:            redInstances(tile.Man5, tile.Pin5, tile.Sou5),
			Renhou:             proto_base.Limit_LIMIT_UNSPECIFIED, // not yaku
			RinshanFu:          true,                               // TODO: check
			GreenRequired:      false,                              // TODO: check
		},
		Game: &proto_rules.Game{
			NumberOfPlayers:          4,
			ChiAllowed:               true,
			AgariYame:                proto_rules.AgariYame_IMPLICIT_END, //also TenpaiYame same (not sure about cases with abortive draws)
			ChiShift:                 proto_rules.Shifting_SHIFTING_FORBIDDEN_STRICT,
			EndByBancrocity:          true, //<0
			Nagashi:                  true,
			Atamahane:                false,
			HonbaPayedToAll:          false,
			RiichiReturnOnMultiron:   false,
			SpeedChi:                 false,
			SayOnClick:               false,
			ShouldHaveMoneyForRiichi: true,
			Atodzuke:                 proto_rules.Atodzuke_ATODZUKE_ALLOWED,
			StartMoney:               25000,
			EndReduceMoney:           30000,
			OkaMoney:                 20000,
			MinWinMoney:              30000,
			MaxLastWind:              proto_base.Wind_WEST, // fixed, also sudden death (if someone reach MinWinMoney)
			LastWind:                 proto_base.Wind_SOUTH,
			Uma:                      SimpleUma(20000, 10000, -10000, -20000),
			Draw: &proto_rules.Draw{
				Winds:   true,
				Kokushi: true,
				Kans:    true,
				Riichi:  true,
				Ron3:    true,
			},
			KanDoraOpen:             proto_rules.KanDoraOpen_AFTER_ACTION,
			UmaShare:                false,
			RiichiSticksGoesToFirst: true,
			Pao: &proto_rules.Pao{
				Winds:   true,
				Dragons: true,
				Kans:    false, // TODO: check
				Rinshan: false,
			},
		},
	}
}

func redInstances(input ...tile.Tile) *proto_base.Instances {
	r := make([]int64, len(input))
	for k, v := range input {
		r[k] = convert.ProtoInstance(v.Instance(0))
	}
	return &proto_base.Instances{Instances: r}
}
