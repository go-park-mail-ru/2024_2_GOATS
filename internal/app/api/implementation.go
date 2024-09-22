package api

import "context"

type ServiceInterface interface {
	Login(ctx context.Context)
	Register(ctx context.Context)
	GetCollection(ctx context.Context)
}

type Implementation struct {
	ctx     context.Context
	service ServiceInterface
}

func NewImplementation(ctx context.Context, srv ServiceInterface) *Implementation {
	return &Implementation{
		ctx: ctx,
		service: srv,
	}
}
