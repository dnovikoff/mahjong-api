package robots

import (
	"testing"

	"github.com/dnovikoff/tempai-core/compact"
	"github.com/dnovikoff/tempai-core/tile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	proto_api "github.com/dnovikoff/mahjong-api/genproto/api"
	proto_log "github.com/dnovikoff/mahjong-api/genproto/log"
	"github.com/dnovikoff/mahjong-api/pkg/convert"
)

func TestCaller(t *testing.T) {
	d := NewCaller(Left, Pon, Center, Kan, Right)
	d.ClientIndex = 2
	tg := compact.NewTileGenerator()
	tiles, err := tg.CompactFromString("1479m5677789p24s")
	require.NoError(t, err)
	d.Hand = tiles

	req := &proto_api.Server{
		OneofEvents: &proto_api.Server_Drop{
			&proto_log.DropEvent{
				WhoIndex: 1,
				Instance: convert.ProtoInstance(tg.Instance(tile.Pin7)),
			},
		},
	}

	test := func(opts ...SuggestFunc) tile.Instances {
		return callerTest(t, d, req, opts...)
	}
	assert.Equal(t,
		tile.Instances{tile.Pin8.Instance(0), tile.Pin9.Instance(0)},
		test(chi, kan, pon),
	)

	assert.Equal(t,
		tile.Instances{tile.Pin7.Instance(0), tile.Pin7.Instance(1)},
		test(right, center, kan, pon),
	)

	assert.Equal(t,
		tile.Instances{tile.Pin6.Instance(0), tile.Pin8.Instance(0)},
		test(center, right, kan),
	)

	assert.Equal(t,
		tile.Instances{tile.Pin5.Instance(0), tile.Pin6.Instance(0)},
		test(right),
	)
}

func TestCallerDecoration(t *testing.T) {
	caller := NewCaller(Left, Pon, Center, Kan, Right)
	caller.ClientIndex = 2
	dec := NewSettingsDecorator(caller)

	tg := compact.NewTileGenerator()
	tiles, err := tg.CompactFromString("1479m5677789p24s")
	require.NoError(t, err)
	caller.Hand = tiles

	req := &proto_api.Server{
		OneofEvents: &proto_api.Server_Drop{
			&proto_log.DropEvent{
				WhoIndex: 1,
				Instance: convert.ProtoInstance(tg.Instance(tile.Pin7)),
			},
		},
	}
	opts := []SuggestFunc{chi, kan, pon}
	assert.Equal(t,
		tile.Instances{tile.Pin8.Instance(0), tile.Pin9.Instance(0)},
		callerTest(t, dec, req, opts...),
	)
	dec = NewSettingsDecorator(caller, NoChi, NoPon, NoKan)
	assert.Nil(t, callerTest(t, dec, req, opts...))
}

func chi(s *proto_api.Suggest) {
	left(s)
	right(s)
	center(s)
}

func left(s *proto_api.Suggest) {
	s.ChiLeft = true
}

func right(s *proto_api.Suggest) {
	s.ChiRight = true
}

func center(s *proto_api.Suggest) {
	s.ChiCenter = true
}

func pon(s *proto_api.Suggest) {
	s.Pon = true
}

func kan(s *proto_api.Suggest) {
	s.Kan = true
}

func callerTest(t *testing.T, c Robot, req *proto_api.Server, opts ...SuggestFunc) tile.Instances {
	req.Suggest = &proto_api.Suggest{
		SuggestId: 1,
	}
	for _, v := range opts {
		v(req.Suggest)
	}
	x := c.Request(req)
	require.NotNil(t, x)
	h := convert.Hand(x.GetCall())
	if h == nil {
		return nil
	}
	return h.Instances()
}
