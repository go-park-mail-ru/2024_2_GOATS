package app

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestAppIntegration(t *testing.T) {
	ctx := testContext(t)

	app, err := New(true)
	require.NoError(t, err, "cannot create app")

	done := make(chan struct{})

	go func() {
		defer close(done)

		app.Run()
	}()

	time.Sleep(1 * time.Second)

	require.NoError(t, app.Server.Shutdown(ctx), "failed to shut down server")

	<-done
}

func testContext(t *testing.T) context.Context {
	require.NoError(t, os.Chdir("../.."), "failed to change directory")

	cfg, err := config.New(true)
	require.NoError(t, err, "failed to read config from app_test")

	return config.WrapContext(context.Background(), cfg)
}
