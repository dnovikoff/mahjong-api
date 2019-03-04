package rules

import (
	proto_base "github.com/dnovikoff/mahjong-api/genproto/public/base"
	proto_rules "github.com/dnovikoff/mahjong-api/genproto/public/rules"
)

func JPMLB() *proto_rules.Ruleset {
	return &proto_rules.Ruleset{
		Id: "JPMLB",
		Scoring: &proto_rules.Scoring{
			ManganRound:  true,
			KazoeYakuman: false,
			YakumanSum:   false,
			HonbaMoney:   100,
		},
		Yaku: &proto_rules.Yaku{
			OpenTanyao:         true,
			HaiteiFromLiveOnly: true,
			Ura:                true,
			Ipatsu:             true,
			Renhou:             proto_base.Limit_MANGAN,
			RinshanFu:          false, // like in A ruleset
			GreenRequired:      false,
		},
		Game: &proto_rules.Game{
			NumberOfPlayers:          4,
			ChiAllowed:               true,
			AgariYame:                proto_rules.AgariYame_IMPLICIT_CONTINUE,
			ChiShift:                 proto_rules.Shifting_SHIFTING_FORBIDDEN_STRICT,
			EndByBancrocity:          false,
			Nagashi:                  false,
			Atamahane:                true,
			HonbaPayedToAll:          false,
			RiichiReturnOnMultiron:   false,
			SpeedChi:                 true,
			SayOnClick:               true,
			ShouldHaveMoneyForRiichi: false,
			Atodzuke:                 proto_rules.Atodzuke_ATODZUKE_ALLOWED,
			StartMoney:               30000,
			EndReduceMoney:           30000,
			OkaMoney:                 0,
			MinWinMoney:              30000,
			MaxLastWind:              proto_base.Wind_SOUTH,
			LastWind:                 proto_base.Wind_SOUTH,
			Uma:                      SimpleUma(15000, 5000, -5000, -15000),
			Draw: &proto_rules.Draw{
				Winds:   true,
				Kokushi: true,
				Kans:    true,
				Riichi:  true,
				Ron3:    true,
			},
			KanDoraOpen:             proto_rules.KanDoraOpen_AFTER_ACTION, // TODO: check (mb instant)
			UmaShare:                true,
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
