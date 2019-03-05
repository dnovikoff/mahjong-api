package logs_test

import (
	"context"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	private_game "github.com/dnovikoff/mahjong-api/genproto/private/game"
	private_log "github.com/dnovikoff/mahjong-api/genproto/private/log"
	public_log "github.com/dnovikoff/mahjong-api/genproto/public/log"
	"github.com/dnovikoff/mahjong-api/pkg/logs"
)

func TestServer(t *testing.T) {
	ctx := context.Background()
	s := logs.StartTestServer(ctx)
	defer func() {
		require.NoError(t, s.Stop())
	}()
	c, err := s.GetDialer().Dial()
	require.NoError(t, err)
	defer c.Close()
	pr := private_log.NewLogServiceClient(c)
	pu := public_log.NewLogServiceClient(c)
	req := &private_log.SaveLogRequest{
		Log: &public_log.Log{
			Info: &public_log.GameInfo{
				Id: "myid",
			},
		},
		Debug: &private_log.DebugLog{
			Create: &private_game.CreateRequest{
				Seed: &wrappers.Int64Value{
					Value: 1234,
				},
			},
		},
	}
	_, err = pr.SaveLog(ctx, req)
	require.NoError(t, err)
	t.Run("get debug", func(t *testing.T) {
		resp, err := pr.GetDebugLog(ctx, &private_log.GetDebugLogRequest{
			LogId: "myid",
		})
		require.NoError(t, err)
		assert.True(t, proto.Equal(req.Debug, resp.Debug))
	})
	t.Run("get log", func(t *testing.T) {
		resp, err := pu.GetLog(ctx, &public_log.GetLogRequest{
			LogId: "myid",
		})
		require.NoError(t, err)
		assert.True(t, proto.Equal(req.Log, resp.Log))
	})
	t.Run("get wrong debug", func(t *testing.T) {
		_, err := pr.GetDebugLog(ctx, &private_log.GetDebugLogRequest{
			LogId: "wrong",
		})
		require.Error(t, err)
	})
	t.Run("get wrong log", func(t *testing.T) {
		_, err := pu.GetLog(ctx, &public_log.GetLogRequest{
			LogId: "wrong",
		})
		require.Error(t, err)
	})
}
