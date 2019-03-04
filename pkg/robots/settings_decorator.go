package robots

import (
	"github.com/golang/protobuf/proto"

	proto_game "github.com/dnovikoff/mahjong-api/genproto/public/game"
)

func DefaultSettings() *proto_game.Settings {
	return &proto_game.Settings{
		AutoTempai:  true,
		AutoEndGame: true,
	}
}

type SuggestFunc func(*proto_game.Suggest)

func NoChi(s *proto_game.Suggest) {
	NoLeft(s)
	NoRight(s)
	NoCenter(s)
}

func NoLeft(s *proto_game.Suggest) {
	s.ChiLeft = false
}

func NoRight(s *proto_game.Suggest) {
	s.ChiRight = false
}

func NoCenter(s *proto_game.Suggest) {
	s.ChiCenter = false
}

func NoPon(s *proto_game.Suggest) {
	s.Pon = false
}

func NoKan(s *proto_game.Suggest) {
	s.Kan = false
}

type Settings struct {
	proto_game.Settings
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

func (d *SettingsDecorator) Update(x *proto_game.Settings) {
	d.Settings.Settings = *x
}

func ApplySettings(req *proto_game.Server, settings *proto_game.Settings) *proto_game.Client {
	x := applySettings(req, settings)
	if x != nil {
		x.SuggestId = req.GetSuggest().GetSuggestId()
	}
	return x
}

func applySettings(req *proto_game.Server, settings *proto_game.Settings) *proto_game.Client {
	if req.Suggest == nil {
		return nil
	}
	s := req.GetSuggest()
	switch {
	case settings.GetAutoWin() && s.GetWin():
		return &proto_game.Client{
			OneofClient: &proto_game.Client_Win{true},
		}
	case settings.GetDropTsumo() && s.GetDropMask() != 0,
		settings.GetAutoTempai() && s.GetNoten(),
		settings.GetAutoEndGame() && s.GetContinueGame():
		return &proto_game.Client{
			OneofClient: &proto_game.Client_Cancel{true},
		}
	}
	return nil
}

func (d *SettingsDecorator) Request(req *proto_game.Server) *proto_game.Client {
	if req.Suggest == nil || req.GetSuggest().GetCanceled() {
		return d.Robot.Request(req)
	}
	applyResult := ApplySettings(req, &d.Settings.Settings)
	if applyResult != nil {
		d.Robot.Request(req)
		return applyResult
	}
	cReq := proto.Clone(req).(*proto_game.Server)
	for _, v := range d.Settings.Remove {
		v(cReq.Suggest)
	}
	return d.Robot.Request(cReq)
}
