package logs

import (
	private_log "github.com/dnovikoff/mahjong-api/genproto/private/log"
	public_log "github.com/dnovikoff/mahjong-api/genproto/public/log"
)

type LogServer interface {
	private_log.LogServiceServer
	public_log.LogServiceServer
}
