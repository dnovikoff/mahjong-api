package bootstrap

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	private_log "github.com/dnovikoff/mahjong-api/genproto/private/log"
	public_log "github.com/dnovikoff/mahjong-api/genproto/public/log"
)

func TestAuth(t *testing.T) {
	cfg := &ConfigStruct{
		Network:       "tcp",
		Address:       ":0",
		LoggerEnabled: false,
		Secret:        "mysecret",
		SecretWhitelist: []string{
			"/mahjong.log.LogService/GetLog",
		},
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	addrCh := make(chan string, 1)
	done := make(chan interface{})
	go func() {
		defer close(done)
		err := RunForConfig(ctx, cfg, func(data *InitData) error {
			fake := &testServer{}
			private_log.RegisterLogServiceServer(data.Server, fake)
			public_log.RegisterLogServiceServer(data.Server, fake)
			addrCh <- data.Listener.Addr().String()
			return nil
		})
		require.NoError(t, err)
	}()

	gc, err := grpc.Dial(<-addrCh, grpc.WithInsecure())
	require.NoError(t, err)
	defer gc.Close()
	t.Run("private", func(t *testing.T) {
		c := private_log.NewLogServiceClient(gc)
		for _, v := range []struct {
			name string
			md   metadata.MD
			err  error
		}{
			{
				name: "no secret",
				md:   nil,
				err:  status.Error(codes.Unauthenticated, "secret is not supplied"),
			}, {
				name: "empty secret",
				md:   metadata.Pairs("secret", ""),
				err:  status.Error(codes.Unauthenticated, "secret does not match"),
			}, {
				name: "wrong secret",
				md:   metadata.Pairs("secret", "wrong"),
				err:  status.Error(codes.Unauthenticated, "secret does not match"),
			}, {
				name: "correct secret",
				md:   metadata.Pairs("secret", "mysecret"),
				err:  nil,
			},
		} {
			t.Run(v.name, func(t *testing.T) {
				ctx := context.Background()
				if v.md != nil {
					ctx = metadata.NewOutgoingContext(ctx, v.md)
				}
				resp, err := c.GetDebugLog(ctx, &private_log.GetDebugLogRequest{})
				if v.err != nil {
					require.Equal(t, v.err, err)
					require.Nil(t, resp)
				} else {
					require.NoError(t, err)
					require.NotNil(t, resp)
				}
			})
		}
	})
	t.Run("public", func(t *testing.T) {
		c := public_log.NewLogServiceClient(gc)
		for _, v := range []struct {
			name string
			md   metadata.MD
		}{
			{
				name: "no secret",
				md:   nil,
			}, {
				name: "empty secret",
				md:   metadata.Pairs("secret", ""),
			}, {
				name: "wrong secret",
				md:   metadata.Pairs("secret", "wrong"),
			}, {
				name: "correct secret",
				md:   metadata.Pairs("secret", "mysecret"),
			},
		} {
			t.Run(v.name, func(t *testing.T) {
				ctx := context.Background()
				if v.md != nil {
					ctx = metadata.NewOutgoingContext(ctx, v.md)
				}
				resp, err := c.GetLog(ctx, &public_log.GetLogRequest{})
				require.NoError(t, err)
				require.NotNil(t, resp)
			})
		}
	})
	cancel()
	<-done
}

type testServer struct {
}

func (s *testServer) SaveLog(ctx context.Context, req *private_log.SaveLogRequest) (*private_log.SaveLogResponse, error) {
	return &private_log.SaveLogResponse{}, nil
}

func (s *testServer) GetLog(ctx context.Context, req *public_log.GetLogRequest) (*public_log.GetLogResponse, error) {
	return &public_log.GetLogResponse{}, nil
}

func (s *testServer) GetDebugLog(ctx context.Context, req *private_log.GetDebugLogRequest) (*private_log.GetDebugLogResponse, error) {
	return &private_log.GetDebugLogResponse{}, nil
}
