package user

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/neygun/friend-management/internal/handler/user/testdata"
	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/service/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateUser(t *testing.T) {
	type mockCreateUserService struct {
		expCall bool
		input   model.User
		output  model.User
		err     error
	}

	type args struct {
		givenRequest          string
		mockCreateUserService mockCreateUserService
		expStatusCode         int
		expResponse           string
	}

	tcs := map[string]args{
		"err - invalid JSON request": {
			givenRequest:  ``,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "invalid_json_request.json",
		},
		"err - missing email field": {
			givenRequest:  `{"test":"test"}`,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "missing_email.json",
		},
		"err - invalid email": {
			givenRequest:  `{"email":"testexample.com"}`,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "invalid_email.json",
		},
		"service error": {
			givenRequest: `{"email":"test@example.com"}`,
			mockCreateUserService: mockCreateUserService{
				expCall: true,
				input: model.User{
					Email: "test@example.com",
				},
				err: errors.New("test"),
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   "server_error.json",
		},
		"success": {
			givenRequest: `{"email":"test@example.com"}`,
			mockCreateUserService: mockCreateUserService{
				expCall: true,
				input: model.User{
					Email: "test@example.com",
				},
				output: model.User{
					ID:    1,
					Email: "test@example.com",
				},
			},
			expStatusCode: http.StatusOK,
			expResponse:   "success.json",
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(tc.givenRequest))
			routeCtx := chi.NewRouteContext()
			req.Header.Set("Content-Type", "application/json")
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			res := httptest.NewRecorder()

			req = req.WithContext(ctx)

			mockUserService := user.NewMockService(t)

			// When
			if tc.mockCreateUserService.expCall {
				mockUserService.ExpectedCalls = []*mock.Call{
					mockUserService.On("CreateUser", ctx, tc.mockCreateUserService.input).Return(tc.mockCreateUserService.output, tc.mockCreateUserService.err),
				}
			}
			instance := New(mockUserService)
			handler := instance.CreateUser()
			handler.ServeHTTP(res, req)

			// Then
			require.Equal(t, tc.expStatusCode, res.Code)
			if tc.expResponse != "" {
				expResponse := testdata.LoadTestJSONFile(t, "testdata/"+tc.expResponse)
				require.JSONEq(t, expResponse, res.Body.String())
			}
			mockUserService.AssertExpectations(t)
		})
	}
}
