package api

import (
	"bytes"
	"encoding/json"
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

func TestGetEventAPI(t *testing.T) {
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

func randomEvent() db.Event {
	return db.Event{
		EventID:          int32(util.RandomInt(1, 100)),
		OrganizerID:      int32(util.RandomInt(1, 100)),
		EventName:        "Test Event",
		EventDescription: "Test Event Description",
		CreatedAt:        time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC),
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
