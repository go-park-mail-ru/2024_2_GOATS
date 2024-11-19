package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/config"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
)

type UserHandler struct {
	user.UnimplementedUserRPCServer
	userService UserServiceInterface
	lclStrg     *config.LocalStorage
}

func NewUserHandler(ctx context.Context, usrSrv UserServiceInterface) user.UserRPCServer {
	lclStrg := config.FromLocalStorageContext(ctx)

	return &UserHandler{
		userService: usrSrv,
		lclStrg:     lclStrg,
	}
}

func (uh *UserHandler) Create(ctx context.Context, createUsrReq *user.CreateUserRequest) (*user.ID, error) {

	return nil, nil
}

func (uh *UserHandler) UpdateProfile(ctx context.Context, updateProfileReq *user.Nothing) (*user.Nothing, error) {

	return nil, nil
}

func (uh *UserHandler) UpdatePassword(ctx context.Context, updatePasswdReq *user.UpdatePasswordRequest) (*user.Nothing, error) {

	return nil, nil
}

func (uh *UserHandler) GetFavorites(ctx context.Context, getFavReq *user.Nothing) (*user.Nothing, error) {

	return nil, nil
}

func (uh *UserHandler) SetFavorite(ctx context.Context, setFavReq *user.Nothing) (*user.Nothing, error) {

	return nil, nil
}

func (uh *UserHandler) ResetFavorite(ctx context.Context, resetFavReq *user.Nothing) (*user.Nothing, error) {

	return nil, nil
}

func (uh *UserHandler) CheckFavorite(ctx context.Context, checkFavReq *user.Nothing) (*user.Nothing, error) {

	return nil, nil
}

func (ua *UserHandler) FindByID(ctx context.Context, usrID *user.ID) (*user.UserData, error) {

	return nil, nil
}

func (ua *UserHandler) FindByEmail(ctx context.Context, usrEmail *user.Email) (*user.UserData, error) {

	return nil, nil
}
