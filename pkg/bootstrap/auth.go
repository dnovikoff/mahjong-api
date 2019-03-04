package bootstrap

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type secretChecker struct {
	cfg Config
}

func newSecretChecker(cfg Config) *secretChecker {
	return &secretChecker{cfg: cfg}
}

func (c *secretChecker) check(ctx context.Context, full string) error {
	if c.cfg.GetSecret() == "" {
		return nil
	}
	for _, v := range c.cfg.GetSecretWhitelist() {
		if v == full {
			return nil
		}
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}
	return c.checkMD(md)
}

func (c *secretChecker) checkMD(md metadata.MD) error {
	authHeader, ok := md["secret"]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "secret is not supplied")
	}

	token := authHeader[0]
	if token != c.cfg.GetSecret() {
		return status.Errorf(codes.Unauthenticated, "secret does not match")
	}
	return nil
}

func (c *secretChecker) Intercept(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	if err := c.check(ctx, info.FullMethod); err != nil {
		return nil, err
	}
	return handler(ctx, req)
}
