package rules

import (
	proto_base "github.com/dnovikoff/mahjong-api/genproto/public/base"
	proto_rules "github.com/dnovikoff/mahjong-api/genproto/public/rules"
)

func JPMLA() *proto_rules.Ruleset {
	return &proto_rules.Ruleset{
		Id: "JPMLA",
		Scoring: &proto_rules.Scoring{
			ManganRound:  false,
			KazoeYakuman: false,
			YakumanSum:   false,
			HonbaMoney:   100,
		},
		Yaku: &proto_rules.Yaku{
			OpenTanyao:         true,
			HaiteiFromLiveOnly: true,
			Ura:                false,
			Ipatsu:             false,
			AkaDora:            nil,
			Renhou:             proto_base.Limit_MANGAN,
			RinshanFu:          false,
			GreenRequired:      true,
		},
		Game: &proto_rules.Game{
			NumberOfPlayers:          4,
			ChiAllowed:               true,
			AgariYame:                proto_rules.AgariYame_IMPLICIT_CONTINUE, // TODO: check
			ChiShift:                 proto_rules.Shifting_SHIFTING_FORBIDDEN_SOFT,
			EndByBancrocity:          false,
			Nagashi:                  false,
			Atamahane:                true,
			HonbaPayedToAll:          false, // TODO: check
			RiichiReturnOnMultiron:   false,
			SpeedChi:                 true,  // TODO: check
			SayOnClick:               true,  // TODO: choose
			ShouldHaveMoneyForRiichi: false, // TODO: check
			Atodzuke:                 proto_rules.Atodzuke_ATODZUKE_ALLOWED,
			StartMoney:               30000,
			EndReduceMoney:           30000,
			OkaMoney:                 0,
			MinWinMoney:              30000, // TODO: check
			MaxLastWind:              proto_base.Wind_SOUTH,
			LastWind:                 proto_base.Wind_SOUTH,
			Uma: &proto_rules.ComplexUma{
				Minus1Money:  []int64{8000, 3000, 1000, -12000},
				Plus1Money:   []int64{12000, -1000, -3000, -8000},
				DefaultMoney: []int64{8000, 4000, -4000, -8000},
			},
			Draw: &proto_rules.Draw{
				Winds:   true,
				Kokushi: true,
				Kans:    true,
				Riichi:  true,
				Ron3:    true,
			},
			KanDoraOpen:             proto_rules.KanDoraOpen_DONT_OPEN,
			UmaShare:                true, // TODO: check
			RiichiSticksGoesToFirst: false,
			Pao: &proto_rules.Pao{
				Winds:   true,
				Dragons: true,
				Kans:    true,
				Rinshan: false,
			},
		},
	}
}
