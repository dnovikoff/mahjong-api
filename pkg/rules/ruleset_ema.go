package rules

import (
	proto_base "github.com/dnovikoff/mahjong-api/genproto/public/base"
	proto_rules "github.com/dnovikoff/mahjong-api/genproto/public/rules"
)

func EMA() *proto_rules.Ruleset {
	return &proto_rules.Ruleset{
		Id: "EMA",
		Scoring: &proto_rules.Scoring{
			ManganRound:  false,
			KazoeYakuman: false,
			YakumanSum:   false,
			HonbaMoney:   100,
		},
		Yaku: &proto_rules.Yaku{
			OpenTanyao:         true,
			HaiteiFromLiveOnly: true,
			Ura:                true,
			Ipatsu:             true,
			AkaDora:            nil,
			Renhou:             proto_base.Limit_MANGAN,
			RinshanFu:          true,
			GreenRequired:      false,
		},
		Game: &proto_rules.Game{
			NumberOfPlayers:          4,
			ChiAllowed:               true,
			AgariYame:                proto_rules.AgariYame_IMPLICIT_CONTINUE,
			ChiShift:                 proto_rules.Shifting_SHIFTING_FORBIDDEN_STRICT,
			EndByBancrocity:          false,
			Nagashi:                  false,
			Atamahane:                false,
			HonbaPayedToAll:          true,
			RiichiReturnOnMultiron:   true,
			SpeedChi:                 false,
			SayOnClick:               true,
			ShouldHaveMoneyForRiichi: false,
			Atodzuke:                 proto_rules.Atodzuke_ATODZUKE_ALLOWED,
			StartMoney:               30000,
			EndReduceMoney:           30000,
			OkaMoney:                 0,
			MinWinMoney:              0,
			MaxLastWind:              proto_base.Wind_SOUTH,
			LastWind:                 proto_base.Wind_SOUTH,
			Uma:                      SimpleUma(15000, 5000, -5000, -15000),
			Draw:                     nil,
			KanDoraOpen:              proto_rules.KanDoraOpen_INSTANT,
			UmaShare:                 true,
			RiichiSticksGoesToFirst:  true,
			Pao: &proto_rules.Pao{
				Winds:   true,
				Dragons: true,
				Kans:    false,
				Rinshan: false,
			},
		},
	}
}
