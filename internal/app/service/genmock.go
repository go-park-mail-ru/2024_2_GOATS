package service

//go:generate mockery --name=(.+)Mock --case=underscore --with-expecter=true --unroll-variadic=false

type RepositoryInterfaceMock interface {
	RepositoryInterface
}
