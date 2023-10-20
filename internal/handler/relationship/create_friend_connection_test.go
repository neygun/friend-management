package relationship

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
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
			givenRequest:  ``,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "invalid_json_request.json",
		},
		"err - invalid email": {
			givenRequest:  `{"friends":["test1example.com", "test2example.com"]}`,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "invalid_email.json",
		},
		"err - the number of emails must be 2": {
			givenRequest:  `{"friends":["test1@example.com"]}`,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "num_of_emails_must_be_2.json",
		},
		"err - the emails are the same": {
			givenRequest:  `{"friends":["test1@example.com", "test1@example.com"]}`,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "same_emails.json",
		},
		"err - user not found": {
			givenRequest: `{"friends":["test1@example.com", "test2@example.com"]}`,
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
			expResponse:   "user_not_found.json",
		},
		"err - blocking relationship exists": {
			givenRequest: `{"friends":["test1@example.com", "test2@example.com"]}`,
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
			expResponse:   "block_exists.json",
		},
		"service error": {
			givenRequest: `{"friends":["test1@example.com", "test2@example.com"]}`,
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
			expResponse:   "server_error.json",
		},
		"success": {
			givenRequest: `{"friends":["test1@example.com", "test2@example.com"]}`,
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
			expResponse:   "success.json",
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			req := httptest.NewRequest(http.MethodPost, "/friends", strings.NewReader(tc.givenRequest))
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
				expResponse := testdata.LoadTestJSONFile(t, "testdata/"+tc.expResponse)
				require.JSONEq(t, expResponse, res.Body.String())
			}
			mockRelationshipService.AssertExpectations(t)
		})
	}
}
