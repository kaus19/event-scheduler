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
	"github.com/stretchr/testify/require"
)

func TestGetTimeSlotsByEvent(t *testing.T) {
	eventTimeSlots := randomEventTimeSlots()

	testCases := []struct {
		name          string
		eventID       int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			eventID: eventTimeSlots[0].EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTimeSlotsByEvent(gomock.Any(), gomock.Eq(eventTimeSlots[0].EventID)).
					Times(1).
					Return(eventTimeSlots, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchEventTimeSlots(t, recorder.Body, eventTimeSlots)
			},
		},
		{
			name:    "NOTFOUND",
			eventID: eventTimeSlots[0].EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTimeSlotsByEvent(gomock.Any(), gomock.Eq(eventTimeSlots[0].EventID)).
					Times(1).
					Return([]db.TimeSlotsEvent{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "InternalServerError",
			eventID: eventTimeSlots[0].EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTimeSlotsByEvent(gomock.Any(), gomock.Eq(eventTimeSlots[0].EventID)).
					Times(1).
					Return([]db.TimeSlotsEvent{}, sql.ErrConnDone)
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

			url := fmt.Sprintf("/time-slots/event/%d", tc.eventID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			r := gin.Default()

			RegisterHandlers(r, server)

			r.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestUpdateTimeSlotEvent(t *testing.T) {
	eventTimeSlot := randomEventTimeSlots()[0]

	testCases := []struct {
		name          string
		eventID       int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			eventID: eventTimeSlot.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateTimeSlotEvent(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, params db.UpdateTimeSlotEventParams) error {
						// Validate other fields explicitly if needed
						require.Equal(t, eventTimeSlot.ID, params.ID)
						require.Equal(t, eventTimeSlot.EventID, params.EventID)
						require.Equal(t, eventTimeSlot.StartTime, params.StartTime)
						require.Equal(t, eventTimeSlot.EndTime, params.EndTime)
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
			eventID: eventTimeSlot.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateTimeSlotEvent(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, params db.UpdateTimeSlotEventParams) error {
						// Validate other fields explicitly if needed
						require.Equal(t, eventTimeSlot.ID, params.ID)
						require.Equal(t, eventTimeSlot.EventID, params.EventID)
						require.Equal(t, eventTimeSlot.StartTime, params.StartTime)
						require.Equal(t, eventTimeSlot.EndTime, params.EndTime)
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
			eventID: eventTimeSlot.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateTimeSlotEvent(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, params db.UpdateTimeSlotEventParams) error {
						// Validate other fields explicitly if needed
						require.Equal(t, eventTimeSlot.ID, params.ID)
						require.Equal(t, eventTimeSlot.EventID, params.EventID)
						require.Equal(t, eventTimeSlot.StartTime, params.StartTime)
						require.Equal(t, eventTimeSlot.EndTime, params.EndTime)
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

			url := "/time-slots/event"
			jsonBody1 := UpdateTimeSlotEventJSONRequestBody{
				Id:        int(eventTimeSlot.ID),
				EventId:   int(eventTimeSlot.EventID),
				StartTime: eventTimeSlot.StartTime,
				EndTime:   eventTimeSlot.EndTime,
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

func TestDeleteTimeSlotEvent(t *testing.T) {
	eventTimeSlot := randomEventTimeSlots()[0]

	testCases := []struct {
		name          string
		eventID       int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			eventID: eventTimeSlot.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTimeSlotEvent(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, recorder.Code)
			},
		},
		{
			name:    "NOT_FOUND",
			eventID: eventTimeSlot.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTimeSlotEvent(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "InternalServerError",
			eventID: eventTimeSlot.EventID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTimeSlotEvent(gomock.Any(), gomock.Any()).
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

			url := fmt.Sprintf("/time-slots/event?id=%d&event_id=%d", eventTimeSlot.ID, eventTimeSlot.EventID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)
			r := gin.Default()

			RegisterHandlers(r, server)

			r.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomEventTimeSlots() []db.TimeSlotsEvent {
	return []db.TimeSlotsEvent{
		{
			ID:        1,
			EventID:   1,
			StartTime: time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			ID:        2,
			EventID:   1,
			StartTime: time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
		},
	}
}

func requireBodyMatchEventTimeSlots(t *testing.T, body *bytes.Buffer, events []db.TimeSlotsEvent) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotTimeSlotsEvents []db.TimeSlotsEvent
	err = json.Unmarshal(data, &gotTimeSlotsEvents)
	require.NoError(t, err)
	require.Equal(t, events, gotTimeSlotsEvents)
}
