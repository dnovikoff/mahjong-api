package bootstrap

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ClientConfig struct {
	Address  string `yaml:"address"`
	Insecure bool   `yaml:"insecure"`
	Secret   string `yaml:"secret"`
}

func (c *ClientConfig) intercept(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "secret", c.Secret)
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func (cfg *ClientConfig) Dial(opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if cfg.Insecure {
		opts = append(opts, grpc.WithInsecure())
	}
	if cfg.Secret != "" {
		opts = append(opts, grpc.WithUnaryInterceptor(cfg.intercept))
	}
	return grpc.Dial(cfg.Address, opts...)
}
