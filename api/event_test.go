package api

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/kaus19/event-scheduler/db/mock"
	db "github.com/kaus19/event-scheduler/db/sqlc"
	"github.com/kaus19/event-scheduler/util"
	"github.com/stretchr/testify/require"
)

func TestGetEvent(t *testing.T) {
	event := randomEvent()

	testCases := []struct {
		name          string
		eventID       int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			eventID: event.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEventByID(gomock.Any(), gomock.Eq(event.EventID)).
					Times(1).
					Return(event, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchEvent(t, recorder.Body, event)
			},
		},
		{
			name:    "NOTFOUND",
			eventID: event.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEventByID(gomock.Any(), gomock.Eq(event.EventID)).
					Times(1).
					Return(db.Event{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "InternalServerError",
			eventID: event.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEventByID(gomock.Any(), gomock.Eq(event.EventID)).
					Times(1).
					Return(db.Event{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/events/%d", tc.eventID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			r := gin.Default()

			RegisterHandlers(r, server)

			r.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListEvents(t *testing.T) {
	events := []db.Event{
		randomEvent(),
		randomEvent(),
		randomEvent(),
	}
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListEvents(gomock.Any()).
					Times(1).
					Return(events, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchEvents(t, recorder.Body, events)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, "/events/list", nil)
			require.NoError(t, err)
			r := gin.Default()

			RegisterHandlers(r, server)

			r.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestUpdateEvent(t *testing.T) {
	event := randomEvent()

	testCases := []struct {
		name          string
		eventID       int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			eventID: event.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateEvent(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, params db.UpdateEventParams) error {
						// Validate other fields explicitly if needed
						require.Equal(t, event.EventID, params.EventID)
						require.Equal(t, "Update event name", params.EventName)
						require.Equal(t, "Update event description", params.EventDescription)
						require.Equal(t, int32(5), params.Duration)
						// Ignore validation of the Timestamp
						return nil
					}).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:    "NOT_FOUND",
			eventID: event.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateEvent(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, params db.UpdateEventParams) error {
						// Validate other fields explicitly if needed
						require.Equal(t, event.EventID, params.EventID)
						require.Equal(t, "Update event name", params.EventName)
						require.Equal(t, "Update event description", params.EventDescription)
						require.Equal(t, int32(5), params.Duration)
						// Ignore validation of the Timestamp
						return sql.ErrNoRows
					}).
					Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "InternalServerError",
			eventID: event.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateEvent(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, params db.UpdateEventParams) error {
						// Validate other fields explicitly if needed
						require.Equal(t, event.EventID, params.EventID)
						require.Equal(t, "Update event name", params.EventName)
						require.Equal(t, "Update event description", params.EventDescription)
						require.Equal(t, int32(5), params.Duration)
						// Ignore validation of the Timestamp
						return errors.New("internal server error")
					}).
					Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/events/%d", tc.eventID)
			jsonBody1 := UpdateEventJSONRequestBody{
				Duration:         5,
				EventDescription: "Update event description",
				EventName:        "Update event name",
			}
			jsonBody, err := json.Marshal(jsonBody1)
			if err != nil {
				t.Fatalf("failed to marshal json: %v", err)
			}

			request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBody))
			require.NoError(t, err)
			r := gin.Default()

			RegisterHandlers(r, server)

			r.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListEventsByOrganizer(t *testing.T) {
	events := []db.Event{
		randomEvent(),
		randomEvent(),
		randomEvent(),
	}
	testCases := []struct {
		name          string
		organizerID   int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			organizerID: events[0].OrganizerID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListEventsByOrganizer(gomock.Any(), gomock.Eq(events[0].OrganizerID)).
					Times(1).
					Return(events, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchEvents(t, recorder.Body, events)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/events/organizer/%d", tc.organizerID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			r := gin.Default()

			RegisterHandlers(r, server)

			r.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestDeleteEvent(t *testing.T) {
	event := randomEvent()

	testCases := []struct {
		name          string
		eventID       int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			eventID: event.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteEvent(gomock.Any(), gomock.Eq(event.EventID)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, recorder.Code)
			},
		},
		{
			name:    "NOT_FOUND",
			eventID: event.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteEvent(gomock.Any(), gomock.Eq(event.EventID)).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "InternalServerError",
			eventID: event.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteEvent(gomock.Any(), gomock.Eq(event.EventID)).
					Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/events/%d", tc.eventID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)
			r := gin.Default()

			RegisterHandlers(r, server)

			r.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomEvent() db.Event {
	return db.Event{
		EventID:          int32(util.RandomInt(1, 100)),
		OrganizerID:      int32(util.RandomInt(1, 100)),
		EventName:        "Test Event",
		EventDescription: "Test Event Description",
		Duration:         1,
		CreatedAt:        time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
	}
}

func requireBodyMatchEvent(t *testing.T, body *bytes.Buffer, event db.Event) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotEvent db.Event
	err = json.Unmarshal(data, &gotEvent)
	require.NoError(t, err)
	require.Equal(t, event, gotEvent)
}

func requireBodyMatchEvents(t *testing.T, body *bytes.Buffer, events []db.Event) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotEvents []db.Event
	err = json.Unmarshal(data, &gotEvents)
	require.NoError(t, err)
	require.Equal(t, events, gotEvents)
}
