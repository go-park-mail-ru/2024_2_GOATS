package app

import (
	"context"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestAppIntegration(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		Env:          map[string]string{"POSTGRES_USER": "test_user", "POSTGRES_PASSWORD": "test_password", "POSTGRES_DB": "test_db"},
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Networks:     []string{"cassette-world"},
	}

	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	port, err := postgresC.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatalf("failed to get mapped port: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to start container: %v", err)
	}

	defer postgresC.Terminate(ctx)

	err = os.Chdir("../..")
	if err != nil {
		t.Fatalf("failed to change directory: %v", err)
	}

	app, err := New(true, &port)
	if err != nil {
		t.Fatalf("failed to create app: %v", err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)

		app.Run()
	}()

	time.Sleep(100 * time.Millisecond)
	if err := app.Database.Ping(); err != nil {
		t.Fatalf("failed to ping database: %v", err)
	}

	app.GracefulShutdown()
	<-done
}
