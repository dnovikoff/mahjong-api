package logs

import (
	"context"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	private_log "github.com/dnovikoff/mahjong-api/genproto/private/log"
	public_log "github.com/dnovikoff/mahjong-api/genproto/public/log"
)

var (
	_ LogServer = &MemoryServer{}
)

// MemoryServer just saves logs in memory
type MemoryServer struct {
	cfg   *Config
	logs  map[string]*public_log.Log
	debug map[string]*private_log.DebugLog
	mu    sync.RWMutex
}

func NewMemoryServer(cfg *Config) *MemoryServer {
	return &MemoryServer{
		cfg:   cfg,
		logs:  map[string]*public_log.Log{},
		debug: map[string]*private_log.DebugLog{},
	}
}

func (s *MemoryServer) SaveLog(ctx context.Context, req *private_log.SaveLogRequest) (*private_log.SaveLogResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := req.GetLog().GetInfo().GetId()
	s.logs[id] = req.GetLog()
	if s.cfg.SaveDebug {
		s.debug[id] = req.GetDebug()
	}
	return &private_log.SaveLogResponse{}, nil
}

func (s *MemoryServer) GetLog(ctx context.Context, req *public_log.GetLogRequest) (*public_log.GetLogResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	id := req.GetLogId()
	log := s.logs[id]
	if log == nil {
		return nil, status.Error(codes.NotFound, "Log not found")
	}
	return &public_log.GetLogResponse{
		Log: log,
	}, nil
}

func (s *MemoryServer) GetDebugLog(ctx context.Context, req *private_log.GetDebugLogRequest) (*private_log.GetDebugLogResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	id := req.GetLogId()
	log := s.debug[id]
	if log == nil {
		return nil, status.Error(codes.NotFound, "DebugLog not found")
	}
	return &private_log.GetDebugLogResponse{
		Debug: log,
	}, nil
}
