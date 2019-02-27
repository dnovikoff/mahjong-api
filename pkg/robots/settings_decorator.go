package robots

import (
	"github.com/golang/protobuf/proto"

	proto_api "github.com/dnovikoff/mahjong-api/genproto/api"
)

func DefaultSettings() *proto_api.Settings {
	return &proto_api.Settings{
		AutoTempai:  true,
		AutoEndGame: true,
	}
}

type SuggestFunc func(*proto_api.Suggest)

func NoChi(s *proto_api.Suggest) {
	NoLeft(s)
	NoRight(s)
	NoCenter(s)
}

func NoLeft(s *proto_api.Suggest) {
	s.ChiLeft = false
}

func NoRight(s *proto_api.Suggest) {
	s.ChiRight = false
}

func NoCenter(s *proto_api.Suggest) {
	s.ChiCenter = false
}

func NoPon(s *proto_api.Suggest) {
	s.Pon = false
}

func NoKan(s *proto_api.Suggest) {
	s.Kan = false
}

type Settings struct {
	proto_api.Settings
	Remove []SuggestFunc
}

// SettingsDecorator is a decorator, that could accept wins.
// Also could protect other implementations from making particular calls.
type SettingsDecorator struct {
	Robot
	Settings Settings
}

func NewSettingsDecorator(c Robot, f ...SuggestFunc) *SettingsDecorator {
	return &SettingsDecorator{c, Settings{Remove: f}}
}

func (d *SettingsDecorator) Update(x *proto_api.Settings) {
	d.Settings.Settings = *x
}

func ApplySettings(req *proto_api.Server, settings *proto_api.Settings) *proto_api.Client {
	x := applySettings(req, settings)
	if x != nil {
		x.SuggestId = req.GetSuggest().GetSuggestId()
	}
	return x
}

func applySettings(req *proto_api.Server, settings *proto_api.Settings) *proto_api.Client {
	if req.Suggest == nil {
		return nil
	}
	s := req.GetSuggest()
	switch {
	case settings.GetAutoWin() && s.GetWin():
		return &proto_api.Client{
			OneofClient: &proto_api.Client_Win{true},
		}
	case settings.GetDropTsumo() && s.GetDropMask() != 0,
		settings.GetAutoTempai() && s.GetNoten(),
		settings.GetAutoEndGame() && s.GetContinueGame():
		return &proto_api.Client{
			OneofClient: &proto_api.Client_Cancel{true},
		}
	}
	return nil
}

func (d *SettingsDecorator) Request(req *proto_api.Server) *proto_api.Client {
	if req.Suggest == nil || req.GetSuggest().GetCanceled() {
		return d.Robot.Request(req)
	}
	applyResult := ApplySettings(req, &d.Settings.Settings)
	if applyResult != nil {
		d.Robot.Request(req)
		return applyResult
	}
	cReq := proto.Clone(req).(*proto_api.Server)
	for _, v := range d.Settings.Remove {
		v(cReq.Suggest)
	}
	return d.Robot.Request(cReq)
}
