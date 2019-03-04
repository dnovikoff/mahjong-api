package logs

import (
	"context"

	private_log "github.com/dnovikoff/mahjong-api/genproto/private/log"
	public_log "github.com/dnovikoff/mahjong-api/genproto/public/log"
	"github.com/dnovikoff/mahjong-api/pkg/bootstrap"
)

func StartTestServer(ctx context.Context) *bootstrap.TestServer {
	return bootstrap.StartTestServer(ctx, func(d *bootstrap.InitData) error {
		s := NewMemoryServer(&Config{
			SaveDebug: true,
		})
		public_log.RegisterLogServiceServer(d.Server, s)
		private_log.RegisterLogServiceServer(d.Server, s)
		return nil
	})
}
