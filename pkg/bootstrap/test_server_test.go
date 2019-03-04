package bootstrap_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dnovikoff/mahjong-api/pkg/bootstrap"
)

func TestServer(t *testing.T) {
	ctx := context.Background()
	s := bootstrap.StartTestServer(ctx, func(*bootstrap.InitData) error {
		return nil
	})
	require.NoError(t, s.Stop())
}
