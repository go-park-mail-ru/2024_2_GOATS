package app

import (
	"context"
	auth "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/auth"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository/cache"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/repository/repository"
	authServicePkg "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/service/auth"
	desc "github.com/go-park-mail-ru/2024_2_GOATS/internal/pb/auth"
	pg "github.com/go-park-mail-ru/2024_2_GOATS/internal/pkg/database/postgres"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	grpcAddress = "localhost:50051"
	httpAddress = "localhost:8080"
)

type IApp interface {
	Run(ctx context.Context) error
}

type App struct {
	db          *pg.PGDatabase
	authService auth.IAuth
}

// New returns application initialized with services.
func New(ctx context.Context) (*App, *auth.Implementation, error) {

	db, err := pg.New(ctx)
	if err != nil {
		log.Fatalf("error in database init: %v", err)
	}

	//addr, err := redisInit.Init()
	//if err != nil {
	//	log.Fatalf("error on connection to redis: %v", err)
	//}

	cacheRepo := redis.NewClient(&redis.Options{
		//Addr: addr,
		Addr: "localhost:6379",
	})

	repo := repository.NewRepo(db)

	cache := cache.NewRedisCache(cacheRepo)

	authService := authServicePkg.NewAuthService(repo, cache)

	authSvcImpl := auth.NewAuth(
		authService,
	)

	a := App{
		db:          db,
		authService: authService,
	}

	return &a, authSvcImpl, nil
}

// Run runs application with its services.
func (a *App) Run(ctx context.Context, authImpl *auth.Implementation) error {
	ctx = context.Background()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		if err := startGrpcServer(authImpl); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()

		if err := startHttpServer(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	defer a.authService.Close()
	wg.Wait()
	return nil
}

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	m, err := handler(ctx, req)

	// Если есть токен, добавьте его в контекст
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		token := md.Get("auth-token")
		if len(token) > 0 {
			// Добавляем токен в новый контекст для последующей передачи в HTTP ответ
			newCtx := metadata.AppendToOutgoingContext(ctx, "auth-token", token[0])
			return handler(newCtx, req) // Передаем новый контекст дальше
		}
	}

	return m, err
}

func startGrpcServer(authImpl *auth.Implementation) error {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(authInterceptor),
	)

	reflection.Register(grpcServer)

	desc.RegisterAuthServer(grpcServer, authImpl)

	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	log.Printf("gRPC server listening at %v\n", grpcAddress)

	return grpcServer.Serve(list)
}

func startHttpServer(ctx context.Context) error {
	mux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(setAuthCookie),
	)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := desc.RegisterAuthHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	log.Printf("http server listening at %v\n", httpAddress)

	return http.ListenAndServe(httpAddress, mux)
}

func setAuthCookie(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	if md, ok := runtime.ServerMetadataFromContext(ctx); ok {
		log.Printf("Metadata received in setAuthCookie: %v", md)
		if token := md.HeaderMD.Get("auth-token"); len(token) > 0 {
			log.Printf("Setting cookie with token: %s", token[0])
			http.SetCookie(w, &http.Cookie{
				Name:     "auth_token",
				Value:    token[0],
				HttpOnly: true,
				Path:     "/",
				Expires:  time.Now().Add(24 * time.Hour),
			})
		}
	} else {
		log.Println("No metadata found in setAuthCookie")
	}
	return nil
}
