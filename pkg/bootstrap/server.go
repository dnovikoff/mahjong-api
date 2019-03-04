package bootstrap

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type InitData struct {
	Logger   *zap.Logger
	Server   *grpc.Server
	Listener net.Listener
}

// RunForPath starts the server, loading config from given path
func RunForPath(ctx context.Context, path string, cfg Config, init func(*InitData) error) error {
	if path == "" {
		return fmt.Errorf("Config path is empty")
	}
	if err := DecodeConfigFile(path, cfg); err != nil {
		return err
	}
	return RunForConfig(ctx, cfg, init)
}

// RunForConfig starts the server for already set config
func RunForConfig(ctx context.Context, cfg Config, init func(*InitData) error) error {
	logger, err := newLogger(cfg)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	logger.Info("Starting server",
		zap.String("network", cfg.GetNetwork()),
		zap.String("address", cfg.GetAddress()),
	)
	listener, err := net.Listen(cfg.GetNetwork(), cfg.GetAddress())
	if err != nil {
		return err
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(newSecretChecker(cfg).Intercept))
	if err = init(&InitData{
		Logger:   logger,
		Server:   s,
		Listener: listener,
	}); err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		s.GracefulStop()
	}()
	err = s.Serve(listener)
	logger.Info("Server stopped")
	return err
}

// Signals returns context, waiting for for signal
func Signals(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	go func() {
		defer signal.Reset()
		select {
		case sig := <-sigs:
			log.Printf("Signal received: %s\n", sig)
			cancel()
		case <-ctx.Done():
			log.Println("Context done")
		}
	}()
	return ctx
}

func newLogger(cfg Config) (*zap.Logger, error) {
	if cfg.IsLoggerEnabled() {
		return zap.NewDevelopment()
	}
	return zap.NewNop(), nil
}
