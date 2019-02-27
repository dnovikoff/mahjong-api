package rules

import (
	"github.com/dnovikoff/tempai-core/tile"

	proto_base "github.com/dnovikoff/mahjong-api/genproto/base"
	proto_rules "github.com/dnovikoff/mahjong-api/genproto/rules"
)

func Fishing() *proto_rules.Ruleset {
	result := FishingRed()
	result.Yaku.AkaDora = nil
	result.Id = "Fishing"
	return result
}

func FishingRed() *proto_rules.Ruleset {
	return &proto_rules.Ruleset{
		Id: "FishingRed",
		Scoring: &proto_rules.Scoring{
			ManganRound:    false,
			KazoeYakuman:   true,
			YakumanSum:     true,
			HonbaMoney:     100,
			DoubleYakumans: []proto_base.Yakuman{proto_base.Yakuman_DAISUUSHI},
		},
		Yaku: &proto_rules.Yaku{
			OpenTanyao:         true,
			HaiteiFromLiveOnly: true,
			Ura:                true,
			Ipatsu:             true,
			AkaDora:            redInstances(tile.Man5, tile.Pin5, tile.Sou5),
			Renhou:             proto_base.Limit_MANGAN,
			RinshanFu:          true, // TODO: check
			GreenRequired:      false,
		},
		Game: &proto_rules.Game{
			NumberOfPlayers:          4,
			ChiAllowed:               true,
			AgariYame:                proto_rules.AgariYame_IMPLICIT_CONTINUE,
			ChiShift:                 proto_rules.Shifting_SHIFTING_FORBIDDEN_STRICT,
			EndByBancrocity:          true,
			Nagashi:                  true,
			Atamahane:                false,
			HonbaPayedToAll:          false,
			RiichiReturnOnMultiron:   false,
			SpeedChi:                 false,
			SayOnClick:               false,
			ShouldHaveMoneyForRiichi: true,
			Atodzuke:                 proto_rules.Atodzuke_ATODZUKE_ALLOWED,
			StartMoney:               25000,
			EndReduceMoney:           25000,
			OkaMoney:                 0,
			MinWinMoney:              30000,
			MaxLastWind:              proto_base.Wind_WEST,
			LastWind:                 proto_base.Wind_SOUTH,
			Uma:                      SimpleUma(25000, 10000, -10000, -25000),
			Draw: &proto_rules.Draw{
				Winds:   true,
				Kokushi: true,
				Kans:    true,
				Riichi:  true,
				Ron3:    true,
			},
			KanDoraOpen:             proto_rules.KanDoraOpen_AFTER_ACTION,
			UmaShare:                true,
			RiichiSticksGoesToFirst: true,
			Pao: &proto_rules.Pao{
				Winds:   true,
				Dragons: true,
				Kans:    false,
				Rinshan: false,
			},
		},
	}
}
