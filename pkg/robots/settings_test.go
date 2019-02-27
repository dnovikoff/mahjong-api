package robots

import (
	"testing"

	"github.com/dnovikoff/tempai-core/compact"
	"github.com/dnovikoff/tempai-core/tile"
	"github.com/stretchr/testify/assert"

	proto_api "github.com/dnovikoff/mahjong-api/genproto/api"
	proto_log "github.com/dnovikoff/mahjong-api/genproto/log"
	"github.com/dnovikoff/mahjong-api/pkg/convert"
)

func TestSettingsAutoWin(t *testing.T) {
	d := NewSettingsDecorator(&Tsumo{})
	req := &proto_api.Server{
		OneofEvents: &proto_api.Server_Drop{
			&proto_log.DropEvent{
				WhoIndex: 1,
				Instance: convert.ProtoInstance(tile.Sou4.Instance(0)),
			},
		},
		Suggest: &proto_api.Suggest{
			SuggestId: 1,
			Win:       true,
		},
	}
	d.Settings.AutoWin = true
	assert.True(t, d.Request(req).GetWin())

	d.Settings.AutoWin = false
	assert.True(t, d.Request(req).GetCancel())
}

func TestSettingsZaichikDrop(t *testing.T) {
	d := NewSettingsDecorator(&Zaichik{&Tsumo{}})

	req := &proto_api.Server{
		OneofEvents: &proto_api.Server_Drop{
			&proto_log.DropEvent{
				WhoIndex: 1,
				Instance: convert.ProtoInstance(tile.Sou4.Instance(0)),
			},
		},
		Suggest: &proto_api.Suggest{
			SuggestId: 1,
			Win:       true,
			Kan:       true,
		},
	}

	d.Settings.AutoWin = false
	assert.Equal(t, &proto_api.Client{
		SuggestId: 1,
		OneofClient: &proto_api.Client_Call{
			convert.ProtoInstances(tile.Instances{
				tile.Sou4.Instance(1),
				tile.Sou4.Instance(2),
				tile.Sou4.Instance(3),
			}),
		},
	}, d.Request(req))

	d.Settings.AutoWin = false
	d.Settings.Remove = []SuggestFunc{
		NoKan, NoChi, NoPon,
	}
	assert.Equal(t, &proto_api.Client{
		SuggestId: 1,
		OneofClient: &proto_api.Client_Cancel{
			true,
		},
	}, d.Request(req))
}

func TestSettingsZaichikTake(t *testing.T) {
	d := NewSettingsDecorator(&Zaichik{&Tsumo{}})

	req := &proto_api.Server{
		OneofEvents: &proto_api.Server_Take{
			&proto_log.TakeEvent{
				WhoIndex: 1,
				Instance: convert.ProtoInstance(tile.Sou4.Instance(0)),
			},
		},
		Suggest: &proto_api.Suggest{
			SuggestId:     1,
			DropMask:      convert.ProtoMask(compact.AllTiles),
			ClosedKanMask: convert.ProtoMask(compact.Tiles(0).Set(tile.Man2).Set(tile.Red)),
			Win:           true,
			Kan:           true,
		},
	}

	d.Settings.AutoWin = false
	d.Settings.DropTsumo = false
	kt := tile.Man2
	assert.Equal(t,
		convert.ProtoInstances(tile.Instances{
			kt.Instance(0),
			kt.Instance(1),
			kt.Instance(2),
			kt.Instance(3),
		}), d.Request(req).GetCall())

	d.Settings.AutoWin = false
	d.Settings.DropTsumo = true
	assert.True(t, d.Request(req).GetCancel())
}
