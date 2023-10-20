package relationship

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/handler/relationship/testdata"
	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/service/relationship"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateFriendConnection(t *testing.T) {
	type mockCreateFriendConnectionService struct {
		expCall bool
		input   relationship.FriendConnectionInput
		output  model.Relationship
		err     error
	}

	type args struct {
		givenRequest                string
		mockCreateFriendConnService mockCreateFriendConnectionService
		expStatusCode               int
		expResponse                 string
	}

	tcs := map[string]args{
		"err - invalid JSON request": {
			givenRequest:  "invalid_json_request.json",
			expStatusCode: http.StatusBadRequest,
			expResponse: handler.ToJsonString(handler.Response{
				Code:        http.StatusBadRequest,
				Description: "Invalid JSON request",
			}),
		},
		"err - invalid email": {
			givenRequest:  "invalid_email.json",
			expStatusCode: http.StatusBadRequest,
			expResponse: handler.ToJsonString(handler.Response{
				Code:        http.StatusBadRequest,
				Description: "Invalid email",
			}),
		},
		"err - the number of emails must be 2": {
			givenRequest:  "num_of_emails_is_not_2.json",
			expStatusCode: http.StatusBadRequest,
			expResponse: handler.ToJsonString(handler.Response{
				Code:        http.StatusBadRequest,
				Description: "The number of emails must be 2",
			}),
		},
		"err - the emails are the same": {
			givenRequest:  "same_emails.json",
			expStatusCode: http.StatusBadRequest,
			expResponse: handler.ToJsonString(handler.Response{
				Code:        http.StatusBadRequest,
				Description: "The emails are the same",
			}),
		},
		"err - user not found": {
			givenRequest: "valid_request.json",
			mockCreateFriendConnService: mockCreateFriendConnectionService{
				expCall: true,
				input: relationship.FriendConnectionInput{
					Friends: []string{
						"test1@example.com",
						"test2@example.com",
					},
				},
				err: relationship.ErrUserNotFound,
			},
			expStatusCode: http.StatusBadRequest,
			expResponse: handler.ToJsonString(handler.Response{
				Code:        http.StatusBadRequest,
				Description: relationship.ErrUserNotFound.Error(),
			}),
		},
		"err - blocking relationship exists": {
			givenRequest: "valid_request.json",
			mockCreateFriendConnService: mockCreateFriendConnectionService{
				expCall: true,
				input: relationship.FriendConnectionInput{
					Friends: []string{
						"test1@example.com",
						"test2@example.com",
					},
				},
				err: relationship.ErrBlockExists,
			},
			expStatusCode: http.StatusBadRequest,
			expResponse: handler.ToJsonString(handler.Response{
				Code:        http.StatusBadRequest,
				Description: relationship.ErrBlockExists.Error(),
			}),
		},
		"service error": {
			givenRequest: "valid_request.json",
			mockCreateFriendConnService: mockCreateFriendConnectionService{
				expCall: true,
				input: relationship.FriendConnectionInput{
					Friends: []string{
						"test1@example.com",
						"test2@example.com",
					},
				},
				err: errors.New("test"),
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse: handler.ToJsonString(handler.Response{
				Code:        http.StatusInternalServerError,
				Description: "Internal Server Error",
			}),
		},
		"success": {
			givenRequest: "valid_request.json",
			mockCreateFriendConnService: mockCreateFriendConnectionService{
				expCall: true,
				input: relationship.FriendConnectionInput{
					Friends: []string{
						"test1@example.com",
						"test2@example.com",
					},
				},
				output: model.Relationship{
					ID:          1,
					RequestorID: 1,
					TargetID:    2,
					Type:        "FRIEND",
				},
			},
			expStatusCode: http.StatusOK,
			expResponse: handler.ToJsonString(handler.SuccessResponse{
				Success: true,
			}),
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			content := testdata.LoadTestJSONFile(t, "testdata/"+tc.givenRequest)
			req := httptest.NewRequest(http.MethodPost, "/friends", strings.NewReader(content))
			routeCtx := chi.NewRouteContext()
			req.Header.Set("Content-Type", "application/json")
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			res := httptest.NewRecorder()

			req = req.WithContext(ctx)

			mockRelationshipService := relationship.NewMockService(t)

			// When
			if tc.mockCreateFriendConnService.expCall {
				mockRelationshipService.ExpectedCalls = []*mock.Call{
					mockRelationshipService.On("CreateFriendConnection", ctx, tc.mockCreateFriendConnService.input).Return(tc.mockCreateFriendConnService.output, tc.mockCreateFriendConnService.err),
				}
			}
			instance := New(mockRelationshipService)
			handler := instance.CreateFriendConnection()
			handler.ServeHTTP(res, req)

			// Then
			require.Equal(t, tc.expStatusCode, res.Code)
			if tc.expResponse != "" {
				require.JSONEq(t, tc.expResponse, res.Body.String())
			}
			mockRelationshipService.AssertExpectations(t)
		})
	}
}
