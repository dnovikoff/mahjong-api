package bootstrap

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Signals returns context, waiting for for signal
func Signals(ctx context.Context, signals ...os.Signal) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	sigs := make(chan os.Signal, 1)
	if len(signals) == 0 {
		signals = []os.Signal{syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM}
	}
	signal.Notify(sigs, signals...)
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
