package bootstrap

import (
	"context"
)

type TestServer struct {
	cfg    *ClientConfig
	cancel func()
	err    chan error
}

func (s *TestServer) GetDialer() *ClientConfig {
	return s.cfg
}

func (s *TestServer) Stop() error {
	s.cancel()
	return <-s.err
}

func StartTestServer(ctx context.Context, init func(d *InitData) error) *TestServer {
	cfg := &ConfigStruct{
		Network:       "tcp",
		Address:       ":0",
		LoggerEnabled: false,
	}
	ch := make(chan string, 1)
	ctx, cancel := context.WithCancel(ctx)
	err := make(chan error, 1)
	go func() {
		err <- RunForConfig(ctx, cfg, func(d *InitData) error {
			ch <- d.Listener.Addr().String()
			return init(d)
		})
	}()
	return &TestServer{
		cfg: &ClientConfig{
			Address:  <-ch,
			Insecure: true,
			Secret:   "",
		},
		cancel: cancel,
		err:    err,
	}
}
