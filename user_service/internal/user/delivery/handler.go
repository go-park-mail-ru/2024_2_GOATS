package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/errs"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/delivery/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/validation"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
	"github.com/rs/zerolog/log"
)

type UserHandler struct {
	user.UnimplementedUserRPCServer
	userService UserServiceInterface
	lclStrg     *config.LocalStorage
}

func NewUserHandler(ctx context.Context, usrSrv UserServiceInterface) user.UserRPCServer {
	lclStrg := config.FromContext(ctx).Databases.LocalStorage

	return &UserHandler{
		userService: usrSrv,
		lclStrg:     &lclStrg,
	}
}

func (uh *UserHandler) Create(ctx context.Context, createUsrReq *user.CreateUserRequest) (*user.ID, error) {
	logger := log.Ctx(ctx)
	err := validation.ValidateCreateUserRequest(createUsrReq)
	if err != nil {
		logger.Error().Msgf("validation_error: %v", err)
		return nil, err
	}

	srvData := converter.ConvertToSrvCreateUser(createUsrReq)
	if srvData == nil {
		logger.Error().Msgf("convert error")
		return nil, errs.ErrBadRequest
	}

	usrID, err := uh.userService.Create(ctx, srvData)
	if err != nil {
		logger.Error().Interface("createUserError", err).Msg("failed_to_create_user")
		return nil, err
	}

	logger.Info().Uint64("createUserSuccess", usrID).Msg("successfully_create_user")
	return &user.ID{ID: usrID}, nil
}

func (uh *UserHandler) UpdateProfile(ctx context.Context, updateProfileReq *user.UserData) (*user.Nothing, error) {
	logger := log.Ctx(ctx)
	ctx = config.WrapLocalStorageContext(ctx, uh.lclStrg)

	srvData := converter.ConvertToSrvUpdateProfile(updateProfileReq)
	err := uh.userService.UpdateProfile(ctx, srvData)
	if err != nil {
		logger.Error().Err(err).Msg("failed_to_update_profile")
		return nil, err
	}

	return &user.Nothing{Dummy: true}, nil
}

func (uh *UserHandler) UpdatePassword(ctx context.Context, updatePasswdReq *user.UpdatePasswordRequest) (*user.Nothing, error) {
	logger := log.Ctx(ctx)
	err := validation.ValidateUpdatePasswordRequesr(updatePasswdReq)
	if err != nil {
		logger.Error().Msgf("validation_error: %v", err)
		return nil, err
	}

	srvData := converter.ConvertToSrvUpdatePassword(updatePasswdReq)
	if srvData == nil {
		return nil, errs.ErrBadRequest
	}

	err = uh.userService.UpdatePassword(ctx, srvData)
	if err != nil {
		logger.Error().Interface("updatePasswordError", err).Msg("failed_to_update_password")
		return nil, err
	}

	logger.Info().Msg("successfully_update_password")
	return &user.Nothing{Dummy: true}, nil
}

func (uh *UserHandler) GetFavorites(ctx context.Context, getFavReq *user.ID) (*user.GetFavoritesResponse, error) {
	logger := log.Ctx(ctx)
	if getFavReq.ID == 0 {
		return nil, errs.ErrBadRequest
	}

	mvIDs, err := uh.userService.GetFavorites(ctx, getFavReq.ID)
	if err != nil {
		logger.Error().Interface("getFavoritesError", err).Msg("failed_to_get_user_favorites")
		return nil, err
	}

	logger.Error().Interface("getFavoritesSuccess", mvIDs).Msg("successfully_get_user_favorites")
	return &user.GetFavoritesResponse{MovieIDs: mvIDs}, nil
}

func (uh *UserHandler) SetFavorite(ctx context.Context, setFavReq *user.HandleFavorite) (*user.Nothing, error) {
	return uh.toggleFavorite(ctx, setFavReq, "set")
}

func (uh *UserHandler) ResetFavorite(ctx context.Context, resetFavReq *user.HandleFavorite) (*user.Nothing, error) {
	return uh.toggleFavorite(ctx, resetFavReq, "reset")
}

func (uh *UserHandler) toggleFavorite(ctx context.Context, req *user.HandleFavorite, op string) (*user.Nothing, error) {
	logger := log.Ctx(ctx)
	err := validation.ValidateFavoriteRequest(req)
	if err != nil {
		logger.Error().Msgf("validation_error: %v", err)
		return nil, err
	}

	srvData := converter.ConvertToSrvFavorite(req)
	if srvData == nil {
		return nil, errs.ErrBadRequest
	}

	if op == "set" {
		err = uh.userService.SetFavorite(ctx, srvData)
	} else {
		err = uh.userService.ResetFavorite(ctx, srvData)
	}

	if err != nil {
		logger.Error().Str("operation", op).Interface("toggleFavoriteError", err).Msg("failed_to_toggle_favorite")
		return nil, err
	}

	logger.Error().Str("toggleFavoriteSuccess", op).Msg("successfully_toggle_favorite")
	return &user.Nothing{Dummy: true}, nil
}

func (uh *UserHandler) CheckFavorite(ctx context.Context, checkFavReq *user.HandleFavorite) (*user.Nothing, error) {
	logger := log.Ctx(ctx)
	err := validation.ValidateFavoriteRequest(checkFavReq)
	if err != nil {
		logger.Error().Msgf("validation_error: %v", err)
		return nil, err
	}

	srvData := converter.ConvertToSrvFavorite(checkFavReq)
	if srvData == nil {
		return nil, errs.ErrBadRequest
	}

	present, err := uh.userService.CheckFavorite(ctx, srvData)
	if err != nil {
		logger.Error().Interface("checkFavoriteError", err).Msg("failed_to_check_favorite")
		return nil, err
	}

	logger.Error().Bool("checkFavoriteSuccess", present).Msg("successfully_check_favorite")
	return &user.Nothing{Dummy: present}, nil
}

func (uh *UserHandler) FindByID(ctx context.Context, usrID *user.ID) (*user.UserData, error) {
	logger := log.Ctx(ctx)
	if usrID.ID == 0 {
		return nil, errs.ErrBadRequest
	}

	usrData, err := uh.userService.FindByID(ctx, usrID.ID)
	if err != nil {
		logger.Error().Interface("findUserByIDError", err).Msg("failed_to_find_user_by_id")
		return nil, err
	}

	logger.Error().Interface("findUserByIDSuccess", usrData).Msg("successfully_find_user_by_id")
	return converter.ConvertToGRPCUser(usrData), nil
}

func (uh *UserHandler) FindByEmail(ctx context.Context, usrEmail *user.Email) (*user.UserData, error) {
	logger := log.Ctx(ctx)
	if usrEmail.Email == "" {
		return nil, errs.ErrBadRequest
	}

	usrData, err := uh.userService.FindByEmail(ctx, usrEmail.Email)
	if err != nil {
		logger.Error().Interface("findUserByEmailError", err).Msg("failed_to_find_user_by_email")
		return nil, err
	}

	logger.Error().Interface("findUserByEmailSuccess", usrData).Msg("successfully_find_user_by_email")
	return converter.ConvertToGRPCUser(usrData), nil
}