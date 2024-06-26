package relationship

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/service/relationship"
	"github.com/neygun/friend-management/pkg/util/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateSubscription(t *testing.T) {
	type mockCreateSubscriptionService struct {
		expCall bool
		input   relationship.CreateSubscriptionInput
		output  model.Relationship
		err     error
	}

	type args struct {
		givenRequest                  string
		mockCreateSubscriptionService mockCreateSubscriptionService
		expStatusCode                 int
		expResponse                   string
	}

	tcs := map[string]args{
		"err - invalid JSON request": {
			givenRequest:  ``,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "invalid_json_request.json",
		},
		"err - missing requestor field": {
			givenRequest:  `{"test":"test1@example.com","target":"test2@example.com"}`,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "missing_requestor.json",
		},
		"err - missing target field": {
			givenRequest:  `{"requestor":"test1@example.com","test":"test2@example.com"}`,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "missing_target.json",
		},
		"err - invalid email": {
			givenRequest:  `{"requestor":"test1example.com","target":"test2example.com"}`,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "invalid_email.json",
		},
		"err - requestor and target are the same": {
			givenRequest:  `{"requestor":"test1@example.com","target":"test1@example.com"}`,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "same_requestor_target.json",
		},
		"err - user not found": {
			givenRequest: `{"requestor":"test1@example.com","target":"test2@example.com"}`,
			mockCreateSubscriptionService: mockCreateSubscriptionService{
				expCall: true,
				input: relationship.CreateSubscriptionInput{
					Requestor: "test1@example.com",
					Target:    "test2@example.com",
				},
				err: relationship.ErrUserNotFound,
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   "user_not_found.json",
		},
		"err - subscription exists": {
			givenRequest: `{"requestor":"test1@example.com","target":"test2@example.com"}`,
			mockCreateSubscriptionService: mockCreateSubscriptionService{
				expCall: true,
				input: relationship.CreateSubscriptionInput{
					Requestor: "test1@example.com",
					Target:    "test2@example.com",
				},
				err: relationship.ErrSubscriptionExists,
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   "subscription_exists.json",
		},
		"service error": {
			givenRequest: `{"requestor":"test1@example.com","target":"test2@example.com"}`,
			mockCreateSubscriptionService: mockCreateSubscriptionService{
				expCall: true,
				input: relationship.CreateSubscriptionInput{
					Requestor: "test1@example.com",
					Target:    "test2@example.com",
				},
				err: errors.New("test"),
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   "server_error.json",
		},
		"success": {
			givenRequest: `{"requestor":"test1@example.com","target":"test2@example.com"}`,
			mockCreateSubscriptionService: mockCreateSubscriptionService{
				expCall: true,
				input: relationship.CreateSubscriptionInput{
					Requestor: "test1@example.com",
					Target:    "test2@example.com",
				},
				output: model.Relationship{
					ID:          1,
					RequestorID: 1,
					TargetID:    2,
					Type:        "SUBSCRIBE",
				},
			},
			expStatusCode: http.StatusOK,
			expResponse:   "success.json",
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			req := httptest.NewRequest(http.MethodPost, "/friends/subscription", strings.NewReader(tc.givenRequest))
			routeCtx := chi.NewRouteContext()
			req.Header.Set("Content-Type", "application/json")
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			res := httptest.NewRecorder()

			req = req.WithContext(ctx)

			mockRelationshipService := relationship.NewMockService(t)
			if tc.mockCreateSubscriptionService.expCall {
				mockRelationshipService.ExpectedCalls = []*mock.Call{
					mockRelationshipService.On("CreateSubscription", ctx, tc.mockCreateSubscriptionService.input).Return(tc.mockCreateSubscriptionService.output, tc.mockCreateSubscriptionService.err),
				}
			}

			// When
			instance := New(mockRelationshipService)
			handler := instance.CreateSubscription()
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
