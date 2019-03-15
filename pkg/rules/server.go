package rules

import (
	"context"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	public_rules "github.com/dnovikoff/mahjong-api/genproto/public/rules"
)

type RulesServer struct {
	rules   []*public_rules.Ruleset
	indexed indexed
}

type indexed map[string]*public_rules.Ruleset

var _ public_rules.RulesServiceServer = &RulesServer{}

func (r *RulesServer) Init(input []*public_rules.Ruleset) {
	r.rules = input
	r.indexed = indexRules(input)
}

func (r *RulesServer) ListRules(context.Context, *public_rules.ListRulesRequest) (*public_rules.ListRulesResponse, error) {
	resp := &public_rules.ListRulesResponse{Rules: r.rules}
	resp = proto.Clone(resp).(*public_rules.ListRulesResponse)
	return resp, nil
}

func (r *RulesServer) GetRule(_ context.Context, req *public_rules.GetRuleRequest) (*public_rules.GetRuleResponse, error) {
	rule, ok := r.indexed[req.RuleId]
	if ok {
		rule = proto.Clone(rule).(*public_rules.Ruleset)
		return &public_rules.GetRuleResponse{Ruleset: rule}, nil
	}
	return nil, status.Errorf(codes.NotFound, "rule '%v' not found", req.RuleId)
}

func indexRules(x []*public_rules.Ruleset) indexed {
	ret := make(indexed, len(x))
	for _, v := range x {
		ret[v.GetId()] = v
	}
	return ret
}
