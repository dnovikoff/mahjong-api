package rules

import (
	rules_proto "github.com/dnovikoff/mahjong-api/genproto/rules"
)

func SimpleUma(uma ...int64) *rules_proto.ComplexUma {
	return &rules_proto.ComplexUma{
		DefaultMoney: uma,
		Minus1Money:  uma,
		Plus1Money:   uma,
	}
}
