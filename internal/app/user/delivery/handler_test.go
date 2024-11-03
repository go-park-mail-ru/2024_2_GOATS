package delivery

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/logger"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	mockSrv "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery/mocks"
)

func TestUserHandler_UpdatePassword(t *testing.T) {
	tests := []struct {
		name       string
		reqBody    string
		mockReturn *models.UpdateUserRespData
		mockErr    *models.ErrorRespData
		statusCode int
		resp       string
	}{
		{
			name:    "Success",
			reqBody: `{"password": "newpass123", "passwordConfirmation": "newpass123"}`,
			mockReturn: &models.UpdateUserRespData{
				StatusCode: http.StatusOK,
			},
			resp:       `{"success":true}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "Parse_error",
			reqBody:    `{"password": "newpass123"}`,
			resp:       `{"success":false,"errors":[{"Code":"user_request_parse_error","Error":"request error: updateProfile action: Path params err - strconv.Atoi: parsing \"\": invalid syntax"}]}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Validation error",
			reqBody:    `{"password": "newpass123", "passwordConfirmation": "wrongpass"}`,
			resp:       `{"success":false,"errors":[{"Code":"user_validation_error","Error":"request error: updatePassword action: Password err - password doesnt match with passwordConfirmation"}]}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:    "Service error",
			reqBody: `{"password": "newpass123", "passwordConfirmation": "newpass123"}`,
			mockErr: &models.ErrorRespData{
				StatusCode: http.StatusInternalServerError,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: errors.New("Some database error")})},
			},
			resp:       `{"success":false,"errors":[{"Code":"something_went_wrong","Error":"Some database error"}]}`,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			req := httptest.NewRequest(http.MethodPost, "/users/1/password", bytes.NewBufferString(test.reqBody))
			if test.name != "Parse_error" {
				req = mux.SetURLVars(req, map[string]string{"id": "1"})
			}
			req.Header.Set("Content-Type", "application/json")

			srv := mockSrv.NewMockUserServiceInterface(ctrl)
			handler := NewUserHandler(testContext(), srv)

			if test.mockReturn != nil || test.mockErr != nil {
				srv.EXPECT().UpdatePassword(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)
			}

			w := httptest.NewRecorder()

			handler.UpdatePassword(w, req)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
			assert.JSONEq(t, test.resp, w.Body.String())
		})
	}
}

func TestUserHandler_UpdateProfile(t *testing.T) {
	tests := []struct {
		name       string
		formData   map[string]string
		fileData   []byte
		mockReturn *models.UpdateUserRespData
		mockErr    *models.ErrorRespData
		statusCode int
		resp       string
	}{
		{
			name: "Success",
			formData: map[string]string{
				"email":    "test@mail.ru",
				"username": "testuser",
			},
			fileData: []byte("fake image data"),
			mockReturn: &models.UpdateUserRespData{
				StatusCode: http.StatusOK,
			},
			resp:       `{"success":true}`,
			statusCode: http.StatusOK,
		},
		{
			name: "Parse_error",
			formData: map[string]string{
				"email": "test@mail.ru",
			},
			resp:       `{"success":false,"errors":[{"Code":"user_request_parse_error","Error":"request error: updateProfile action: Path params err - strconv.Atoi: parsing \"\": invalid syntax"}]}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Validation error",
			formData: map[string]string{
				"email":    "invalid-email",
				"username": "testuser",
			},
			resp:       `{"success":false,"errors":[{"Code":"user_validation_error","Error":"updateProfile action: Email err - email is incorrect"}]}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Service error",
			formData: map[string]string{
				"email":    "test@example.com",
				"username": "testuser",
			},
			fileData: []byte("fake image data"),
			mockErr: &models.ErrorRespData{
				StatusCode: http.StatusInternalServerError,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: errors.New("Some database error")})},
			},
			resp:       `{"success":false,"errors":[{"Code":"something_went_wrong","Error":"Some database error"}]}`,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			for key, val := range test.formData {
				_ = writer.WriteField(key, val)
			}

			if test.fileData != nil {
				part, _ := writer.CreateFormFile("avatar", "avatar.png")
				part.Write(test.fileData)
			}

			writer.Close()

			req := httptest.NewRequest(http.MethodPost, "/users/1/update_profile", body)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			if test.name != "Parse_error" {
				req = mux.SetURLVars(req, map[string]string{"id": "1"})
			}

			srv := mockSrv.NewMockUserServiceInterface(ctrl)
			handler := NewUserHandler(testContext(), srv)

			if test.mockReturn != nil || test.mockErr != nil {
				srv.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)
			}

			w := httptest.NewRecorder()
			handler.UpdateProfile(w, req)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, test.statusCode, res.StatusCode)
			assert.JSONEq(t, test.resp, w.Body.String())
		})
	}
}

func testContext() context.Context {
	err := os.Chdir("../../../..")
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("failed to change directory: %v", err))
	}

	cfg, err := config.New(logger.NewLogger(), false)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("failed to read config: %v", err))
	}

	ctx := config.WrapContext(context.Background(), cfg)

	return ctx
}
