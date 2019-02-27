package rules

import (
	rules_proto "github.com/dnovikoff/mahjong-api/genproto/rules"
)

func All() []*rules_proto.Ruleset {
	return []*rules_proto.Ruleset{
		WRC(),
		EMA(),
		JPMLA(),
		JPMLB(),
		TenhouRed(),
		Fishing(),
		FishingRed(),
	}
}
