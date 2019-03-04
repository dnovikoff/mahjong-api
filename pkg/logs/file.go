package logs

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"regexp"

	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	private_log "github.com/dnovikoff/mahjong-api/genproto/private/log"
	public_log "github.com/dnovikoff/mahjong-api/genproto/public/log"
)

var (
	_ LogServer = &FileServer{}
)

// FileServer saves logs on disk
type FileServer struct {
	cfg            *Config
	validateRegexp *regexp.Regexp
	logger         *zap.Logger
}

func NewFileServer(l *zap.Logger, cfg *Config) (*FileServer, error) {
	var r *regexp.Regexp
	if cfg.ValidateIDRegexp != "" {
		var err error
		r, err = regexp.Compile(cfg.ValidateIDRegexp)
		if err != nil {
			return nil, err
		}
	}
	return &FileServer{logger: l, cfg: cfg, validateRegexp: r}, nil
}

func (s *FileServer) SaveLog(ctx context.Context, req *private_log.SaveLogRequest) (*private_log.SaveLogResponse, error) {
	id := req.GetLog().GetInfo().GetId()
	l := s.logger.With(zap.String("log-id", id))
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "No log")
	}
	err := s.saveProto(id, req.GetLog())
	if err != nil {
		l.Error("Error saving log", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error saving log "+id)
	}
	if req.Debug != nil {
		err = s.saveProto(id+".debug", req.GetDebug())
		if err != nil {
			l.Error("Error saving debug log", zap.Error(err))
			return nil, status.Error(codes.Internal, "Error saving debug log "+id)
		}
	}
	l.Info("Log saved")
	if s.cfg.SaveJSON {
		jp := s.jsonPath(id)
		err := os.MkdirAll(jp, 0777)
		if err != nil {
			return nil, status.Error(codes.Internal, "Error saving json log "+id)
		}
		err = SaveJSONLogs(jp+"/", req.Log, req.Debug)
		if err != nil {
			return nil, err
		}
		l.Info("JSON logs saved", zap.String("json-output", jp))
	}
	return &private_log.SaveLogResponse{}, nil
}

func (s *FileServer) GetLog(ctx context.Context, req *public_log.GetLogRequest) (*public_log.GetLogResponse, error) {
	id := req.GetLogId()
	l := s.logger.With(zap.String("log-id", id))
	if err := s.validateID(id); err != nil {
		return nil, err
	}
	resp := &public_log.GetLogResponse{
		Log: &public_log.Log{},
	}
	err := s.readProto(id, resp.Log)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, status.Error(codes.NotFound, "Log not found")
		}
		l.Error("Error loading log", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error loading log")
	}
	return resp, nil
}

func (s *FileServer) GetDebugLog(ctx context.Context, req *private_log.GetDebugLogRequest) (*private_log.GetDebugLogResponse, error) {
	id := req.GetLogId()
	l := s.logger.With(zap.String("log-id", id))
	if err := s.validateID(id); err != nil {
		return nil, err
	}
	resp := &private_log.GetDebugLogResponse{
		Debug: &private_log.DebugLog{},
	}
	err := s.readProto(id+".debug", resp.Debug)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, status.Error(codes.NotFound, "Log not found")
		}
		l.Error("Error loading debug log", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error loading log")
	}
	return resp, nil
}

func (s *FileServer) validateID(id string) error {
	if s.validateRegexp == nil {
		return nil
	}
	if !s.validateRegexp.MatchString(id) {
		return status.Error(codes.Internal, "Invalid ID")
	}
	return nil
}

func (s *FileServer) jsonPath(id string) string {
	return path.Join(s.cfg.Output, "json", id)
}

func (s *FileServer) pathByID(id string) string {
	return path.Join(s.cfg.Output, id)
}

func (s *FileServer) pbPathByID(id string) string {
	return s.pathByID(id) + ".pb"
}

func (s *FileServer) readBytes(id string) ([]byte, error) {
	return ioutil.ReadFile(s.pbPathByID(id))
}

func (s *FileServer) readProto(id string, x proto.Message) error {
	b, err := s.readBytes(id)
	if err != nil {
		return err
	}
	return proto.Unmarshal(b, x)
}

func (s *FileServer) saveBytes(id string, b []byte) error {
	p := s.pbPathByID(id)
	dir := path.Dir(p)
	err := os.MkdirAll(dir, 0777)
	s.logger.Info("Saving log to path", zap.String("path", p), zap.String("dir", dir))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(p, b, 0644)
}

func (s *FileServer) saveProto(id string, x proto.Message) error {
	b, err := proto.Marshal(x)
	if err != nil {
		return err
	}
	return s.saveBytes(id, b)
}
