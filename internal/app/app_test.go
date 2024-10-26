package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/require"
)

func TestAppIntegration(t *testing.T) {
	ctx := testContext()
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "cannot create new pool")

	resource, err := initPostgres(ctx, pool)
	require.NoError(t, err, "cannot init postgres")

	defer func() {
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("failed to stop container: %v", err)
		}
	}()

	app, err := New(true)
	require.NoError(t, err, "cannot create app")

	done := make(chan struct{})

	go func() {
		defer close(done)

		app.Run()
	}()

	time.Sleep(100 * time.Millisecond)
	if err := app.Database.Ping(); err != nil {
		require.NoError(t, err, "failed to ping database from test")
	}

	if err := app.GracefulShutdown(); err != nil {
		t.Fatalf("failed to perform gracefulShutdown: %v", err)
	}

	<-done
}

func testContext() context.Context {
	err := os.Chdir("../..")
	if err != nil {
		log.Fatalf("failed to change directory: %v", err)
	}

	cfg, err := config.New(true)
	if err != nil {
		log.Fatalf("failed to read config from Register test: %v", err)
	}

	return config.WrapContext(context.Background(), cfg)
}

func initPostgres(ctx context.Context, pool *dockertest.Pool) (*dockertest.Resource, error) {
	pgCfg := config.FromContext(ctx).Databases.Postgres
	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_PASSWORD=" + pgCfg.Password,
			"POSTGRES_USER=" + pgCfg.User,
			"POSTGRES_DB=" + pgCfg.Name,
		},
		ExposedPorts: []string{"5432/tcp"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432/tcp": {
				{HostIP: "0.0.0.0", HostPort: fmt.Sprintf("%d/tcp", pgCfg.Port)},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		return nil, fmt.Errorf("error while initing postgres: %w", err)
	}

	err = resource.Expire(30)
	if err != nil {
		return nil, fmt.Errorf("auto expiration err: %w", err)
	}

	return resource, nil
}
