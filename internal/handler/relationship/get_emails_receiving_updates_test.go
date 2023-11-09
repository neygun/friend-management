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

func TestHandler_GetEmailsReceivingUpdates(t *testing.T) {
	type mockGetEmailsReceivingUpdatesService struct {
		expCall bool
		input   relationship.GetEmailsReceivingUpdatesInput
		output  []string
		err     error
	}

	type args struct {
		givenRequest                         string
		mockGetEmailsReceivingUpdatesService mockGetEmailsReceivingUpdatesService
		expStatusCode                        int
		expResponse                          string
	}

	tcs := map[string]args{
		"err - invalid JSON request": {
			givenRequest:  ``,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "invalid_json_request.json",
		},
		"err - missing sender field": {
			givenRequest:  `{"test":"test1@example.com","text":"test test2@example.com"}`,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "missing_sender.json",
		},
		"err - missing text field": {
			givenRequest:  `{"sender":"test1@example.com","test":"test test2@example.com"}`,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "missing_text.json",
		},
		"err - sender's email is invalid": {
			givenRequest:  `{"sender":"test1example.com","text":"test test2@example.com"}`,
			expStatusCode: http.StatusBadRequest,
			expResponse:   "invalid_sender.json",
		},
		"err - user not found": {
			givenRequest: `{"sender":"test1@example.com","text":"test test2@example.com"}`,
			mockGetEmailsReceivingUpdatesService: mockGetEmailsReceivingUpdatesService{
				expCall: true,
				input: relationship.GetEmailsReceivingUpdatesInput{
					Sender: "test1@example.com",
					Text:   "test test2@example.com",
				},
				err: relationship.ErrUserNotFound,
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   "user_not_found.json",
		},
		"service error": {
			givenRequest: `{"sender":"test1@example.com","text":"test test2@example.com"}`,
			mockGetEmailsReceivingUpdatesService: mockGetEmailsReceivingUpdatesService{
				expCall: true,
				input: relationship.GetEmailsReceivingUpdatesInput{
					Sender: "test1@example.com",
					Text:   "test test2@example.com",
				},
				err: errors.New("test"),
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   "server_error.json",
		},
		"success": {
			givenRequest: `{"sender":"test1@example.com","text":"test test2@example.com"}`,
			mockGetEmailsReceivingUpdatesService: mockGetEmailsReceivingUpdatesService{
				expCall: true,
				input: relationship.GetEmailsReceivingUpdatesInput{
					Sender: "test1@example.com",
					Text:   "test test2@example.com",
				},
				output: []string{
					"test2@example.com",
					"test3@example.com",
				},
			},
			expStatusCode: http.StatusOK,
			expResponse:   "get_emails_receiving_updates_success.json",
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			req := httptest.NewRequest(http.MethodPost, "/friends/recipients", strings.NewReader(tc.givenRequest))
			routeCtx := chi.NewRouteContext()
			req.Header.Set("Content-Type", "application/json")
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			res := httptest.NewRecorder()

			req = req.WithContext(ctx)

			mockRelationshipService := relationship.NewMockService(t)
			if tc.mockGetEmailsReceivingUpdatesService.expCall {
				mockRelationshipService.ExpectedCalls = []*mock.Call{
					mockRelationshipService.On("GetEmailsReceivingUpdates", ctx, tc.mockGetEmailsReceivingUpdatesService.input).Return(tc.mockGetEmailsReceivingUpdatesService.output, tc.mockGetEmailsReceivingUpdatesService.err),
				}
			}

			// When
			instance := New(mockRelationshipService)
			handler := instance.GetEmailsReceivingUpdates()
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
