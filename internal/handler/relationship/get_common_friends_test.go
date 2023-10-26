package relationship

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/neygun/friend-management/internal/service/relationship"
	"github.com/neygun/friend-management/pkg/util/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetCommonFriends(t *testing.T) {
	type mockGetCommonFriendsService struct {
		expCall       bool
		input         relationship.GetCommonFriendsInput
		commonFriends []string
		count         int
		err           error
	}

	type args struct {
		givenRequest                string
		mockGetCommonFriendsService mockGetCommonFriendsService
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
			mockGetCommonFriendsService: mockGetCommonFriendsService{
				expCall: true,
				input: relationship.GetCommonFriendsInput{
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
		"service error": {
			givenRequest: `{"friends":["test1@example.com", "test2@example.com"]}`,
			mockGetCommonFriendsService: mockGetCommonFriendsService{
				expCall: true,
				input: relationship.GetCommonFriendsInput{
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
			mockGetCommonFriendsService: mockGetCommonFriendsService{
				expCall: true,
				input: relationship.GetCommonFriendsInput{
					Friends: []string{
						"test1@example.com",
						"test2@example.com",
					},
				},
				commonFriends: []string{
					"test3@example.com",
					"test4@example.com",
				},
				count: 2,
			},
			expStatusCode: http.StatusOK,
			expResponse:   "get_common_friends_success.json",
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			req := httptest.NewRequest(http.MethodPost, "/friends/common", strings.NewReader(tc.givenRequest))
			routeCtx := chi.NewRouteContext()
			req.Header.Set("Content-Type", "application/json")
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			res := httptest.NewRecorder()

			req = req.WithContext(ctx)

			mockRelationshipService := relationship.NewMockService(t)

			// When
			if tc.mockGetCommonFriendsService.expCall {
				mockRelationshipService.ExpectedCalls = []*mock.Call{
					mockRelationshipService.On("GetCommonFriends", ctx, tc.mockGetCommonFriendsService.input).
						Return(tc.mockGetCommonFriendsService.commonFriends, tc.mockGetCommonFriendsService.count, tc.mockGetCommonFriendsService.err),
				}
			}
			instance := New(mockRelationshipService)
			handler := instance.GetCommonFriends()
			handler.ServeHTTP(res, req)

			// Then
			require.Equal(t, tc.expStatusCode, res.Code)
			if tc.expResponse != "" {
				expResponse := test.LoadTestJSONFile(t, "testdata/"+tc.expResponse)
				require.JSONEq(t, expResponse, res.Body.String())
			}
		})
	}
}
