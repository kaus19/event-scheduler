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

func TestGetTimeSlotsByUser(t *testing.T) {
	userTimeSlots := randomUserTimeSlots()

	testCases := []struct {
		name          string
		userID        int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			userID: userTimeSlots[0].UserID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTimeSlotsByUser(gomock.Any(), gomock.Eq(userTimeSlots[0].UserID)).
					Times(1).
					Return(userTimeSlots, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUserTimeSlots(t, recorder.Body, userTimeSlots)
			},
		},
		{
			name:   "NOTFOUND",
			userID: userTimeSlots[0].UserID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTimeSlotsByUser(gomock.Any(), gomock.Eq(userTimeSlots[0].UserID)).
					Times(1).
					Return([]db.TimeSlotsUser{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "InternalServerError",
			userID: userTimeSlots[0].UserID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTimeSlotsByUser(gomock.Any(), gomock.Eq(userTimeSlots[0].UserID)).
					Times(1).
					Return([]db.TimeSlotsUser{}, sql.ErrConnDone)
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

			url := fmt.Sprintf("/time-slots/user/%d", tc.userID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			r := gin.Default()

			RegisterHandlers(r, server)

			r.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestCreateTimeSlotUser(t *testing.T) {
	testCases := []struct {
		name          string
		requestBody   gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			requestBody: gin.H{
				"user_id":    1,
				"start_time": []time.Time{time.Now()},
				"end_time":   []time.Time{time.Now().Add(time.Hour)},
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateTimeSlotUser(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ context.Context, arg db.CreateTimeSlotUserParams) (db.TimeSlotsUser, error) {
						require.Equal(t, int32(1), arg.UserID)
						require.WithinDuration(t, time.Now(), arg.StartTime, time.Second)
						require.WithinDuration(t, time.Now().Add(time.Hour), arg.EndTime, time.Second)
						return db.TimeSlotsUser{}, nil
					}).
					Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "InvalidInput",
			requestBody: gin.H{
				"user_id":    1,
				"start_time": []time.Time{time.Now()},
				// Missing end_time
			},
			buildStubs: func(store *mockdb.MockStore) {
				// No calls to the mock store expected
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "MismatchedStartEndTimes",
			requestBody: gin.H{
				"user_id":    1,
				"start_time": []time.Time{time.Now()},
				"end_time":   []time.Time{time.Now().Add(2 * time.Hour), time.Now().Add(3 * time.Hour)},
			},
			buildStubs: func(store *mockdb.MockStore) {
				// No calls to the mock store expected
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalError",
			requestBody: gin.H{
				"user_id":    1,
				"start_time": []time.Time{time.Now()},
				"end_time":   []time.Time{time.Now().Add(time.Hour)},
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateTimeSlotUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.TimeSlotsUser{}, errors.New("internal server error"))
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

			// Marshal body data to JSON
			data, err := json.Marshal(tc.requestBody)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, "/time-slots/user", bytes.NewBuffer(data))
			require.NoError(t, err)
			r := gin.Default()

			RegisterHandlers(r, server)

			r.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestUpdateTimeSlotUser(t *testing.T) {
	userTimeSlot := randomUserTimeSlots()[0]

	testCases := []struct {
		name          string
		userID        int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			userID: userTimeSlot.UserID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateTimeSlotUser(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, params db.UpdateTimeSlotUserParams) error {
						// Validate other fields explicitly if needed
						require.Equal(t, userTimeSlot.ID, params.ID)
						require.Equal(t, userTimeSlot.UserID, params.UserID)
						require.Equal(t, userTimeSlot.StartTime, params.StartTime)
						require.Equal(t, userTimeSlot.EndTime, params.EndTime)
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
			name:   "NOT_FOUND",
			userID: userTimeSlot.UserID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateTimeSlotUser(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, params db.UpdateTimeSlotUserParams) error {
						// Validate other fields explicitly if needed
						require.Equal(t, userTimeSlot.ID, params.ID)
						require.Equal(t, userTimeSlot.UserID, params.UserID)
						require.Equal(t, userTimeSlot.StartTime, params.StartTime)
						require.Equal(t, userTimeSlot.EndTime, params.EndTime)
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
			name:   "InternalServerError",
			userID: userTimeSlot.UserID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateTimeSlotUser(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, params db.UpdateTimeSlotUserParams) error {
						// Validate other fields explicitly if needed
						require.Equal(t, userTimeSlot.ID, params.ID)
						require.Equal(t, userTimeSlot.UserID, params.UserID)
						require.Equal(t, userTimeSlot.StartTime, params.StartTime)
						require.Equal(t, userTimeSlot.EndTime, params.EndTime)
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

			url := "/time-slots/user"
			jsonBody1 := UpdateTimeSlotUserJSONRequestBody{
				Id:        int(userTimeSlot.ID),
				UserId:    int(userTimeSlot.UserID),
				StartTime: userTimeSlot.StartTime,
				EndTime:   userTimeSlot.EndTime,
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

func TestDeleteTimeSlotuser(t *testing.T) {
	userTimeSlot := randomUserTimeSlots()[0]

	testCases := []struct {
		name          string
		userID        int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			userID: userTimeSlot.UserID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTimeSlotUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, recorder.Code)
			},
		},
		{
			name:   "NOT_FOUND",
			userID: userTimeSlot.UserID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTimeSlotUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "InternalServerError",
			userID: userTimeSlot.UserID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTimeSlotUser(gomock.Any(), gomock.Any()).
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

			url := fmt.Sprintf("/time-slots/user?id=%d&user_id=%d", userTimeSlot.ID, userTimeSlot.UserID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)
			r := gin.Default()

			RegisterHandlers(r, server)

			r.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomUserTimeSlots() []db.TimeSlotsUser {
	return []db.TimeSlotsUser{
		{
			ID:        1,
			UserID:    1,
			StartTime: time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			ID:        2,
			UserID:    1,
			StartTime: time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
		},
	}
}

func requireBodyMatchUserTimeSlots(t *testing.T, body *bytes.Buffer, users []db.TimeSlotsUser) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotTimeSlotsUsers []db.TimeSlotsUser
	err = json.Unmarshal(data, &gotTimeSlotsUsers)
	require.NoError(t, err)
	require.Equal(t, users, gotTimeSlotsUsers)
}
