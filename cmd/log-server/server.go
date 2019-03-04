package main

import (
	"context"
	"flag"
	"log"
	"path/filepath"

	"go.uber.org/zap"

	private_log "github.com/dnovikoff/mahjong-api/genproto/private/log"
	public_log "github.com/dnovikoff/mahjong-api/genproto/public/log"
	"github.com/dnovikoff/mahjong-api/pkg/bootstrap"
	"github.com/dnovikoff/mahjong-api/pkg/logs"
)

type config struct {
	bootstrap.ConfigStruct `yaml:",squash"`
	logs.Config            `yaml:",squash"`
}

func main() {
	configPath := flag.String("config", "", "path to yaml config file")
	flag.Parse()
	var cfg config
	ctx := bootstrap.Signals(context.Background())
	err := bootstrap.RunForPath(ctx, *configPath, &cfg, func(d *bootstrap.InitData) error {
		var s logs.LogServer
		if cfg.Output == "" {
			d.Logger.Info("Staring server in memory mode")
			s = logs.NewMemoryServer(&cfg.Config)
		} else {
			p, err := filepath.Abs(cfg.Output)
			if err != nil {
				return err
			}
			d.Logger.Info("Staring server in file mode", zap.String("output", p))
			s, err = logs.NewFileServer(d.Logger, &cfg.Config)
			if err != nil {
				return err
			}
		}
		public_log.RegisterLogServiceServer(d.Server, s)
		private_log.RegisterLogServiceServer(d.Server, s)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
