package movie

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/config"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/require"
)

func TestAppIntegration(t *testing.T) {
	ctx := testContext(t)

	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "cannot create new pool")

	pg, err := initPostgres(ctx, pool)
	require.NoError(t, err, "cannot init postgres")

	defer func() {
		require.NoError(t, pool.Purge(pg), "failed to stop postgres container")
	}()

	app, err := New(true)
	require.NoError(t, err, "cannot create app")

	done := make(chan struct{})

	go func() {
		defer close(done)

		app.Run()
	}()

	time.Sleep(1 * time.Second)
	if err := app.database.Ping(); err != nil {
		require.NoError(t, err, "failed to ping postgres from test")
	}

	app.srv.GracefulStop()

	<-done
}

func testContext(t *testing.T) context.Context {
	require.NoError(t, os.Chdir("../.."), "failed to change directory")

	cfg, err := config.New(true)
	require.NoError(t, err, "failed to read config from app_test")

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
