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

func TestHandler_GetFriendsList(t *testing.T) {
	type mockGetFriendsListService struct {
		expCall     bool
		input       relationship.GetFriendsInput
		friendsList []string
		count       int
		err         error
	}

	type args struct {
		givenRequest              string
		mockGetFriendsListService mockGetFriendsListService
		expStatusCode             int
		expResponse               string
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
		"err - user not found": {
			givenRequest: `{"email":"test@example.com"}`,
			mockGetFriendsListService: mockGetFriendsListService{
				expCall: true,
				input: relationship.GetFriendsInput{
					Email: "test@example.com",
				},
				err: relationship.ErrUserNotFound,
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   "user_not_found.json",
		},
		"service error": {
			givenRequest: `{"email":"test@example.com"}`,
			mockGetFriendsListService: mockGetFriendsListService{
				expCall: true,
				input: relationship.GetFriendsInput{
					Email: "test@example.com",
				},
				err: errors.New("test"),
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   "server_error.json",
		},
		"success": {
			givenRequest: `{"email":"test@example.com"}`,
			mockGetFriendsListService: mockGetFriendsListService{
				expCall: true,
				input: relationship.GetFriendsInput{
					Email: "test@example.com",
				},
				friendsList: []string{
					"test2@example.com",
					"test3@example.com",
				},
				count: 2,
			},
			expStatusCode: http.StatusOK,
			expResponse:   "get_friends_success.json",
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			req := httptest.NewRequest(http.MethodPost, "/friends/list", strings.NewReader(tc.givenRequest))
			routeCtx := chi.NewRouteContext()
			req.Header.Set("Content-Type", "application/json")
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			res := httptest.NewRecorder()

			req = req.WithContext(ctx)

			mockRelationshipService := relationship.NewMockService(t)
			if tc.mockGetFriendsListService.expCall {
				mockRelationshipService.ExpectedCalls = []*mock.Call{
					mockRelationshipService.On("GetFriendsList", ctx, tc.mockGetFriendsListService.input).Return(tc.mockGetFriendsListService.friendsList, tc.mockGetFriendsListService.count, tc.mockGetFriendsListService.err),
				}
			}

			// When
			instance := New(mockRelationshipService)
			handler := instance.GetFriendsList()
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
