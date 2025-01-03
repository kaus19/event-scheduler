package api

import (
	"fmt"
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

func TestGetMatchingTimeSlotsForEvent(t *testing.T) {
	eventTimeSlots := randomEventTimeSlotsForMatching()
	usersTimeSlots := randomUserTimeSlotsForMatching()
	event := randomEventForMatching()

	testCases := []struct {
		name          string
		eventID       int
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			eventID: int(event.EventID),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTimeSlotsByEvent(gomock.Any(), gomock.Eq(event.EventID)).
					Times(1).
					Return(eventTimeSlots, nil)

				store.EXPECT().
					GetTimeSlotsForAllUsers(gomock.Any()).
					Times(1).
					Return(usersTimeSlots, nil)

				store.EXPECT().
					GetEventByID(gomock.Any(), gomock.Eq(event.EventID)).
					Times(1).
					Return(event, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/matching-slots/event/%d", tc.eventID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			r := gin.Default()
			RegisterHandlers(r, server)

			r.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomEventTimeSlotsForMatching() []db.TimeSlotsEvent {
	return []db.TimeSlotsEvent{
		{
			ID:        1,
			EventID:   1,
			StartTime: time.Now(),
			EndTime:   time.Now().Add(2 * time.Hour),
		},
		{
			ID:        2,
			EventID:   2,
			StartTime: time.Now(),
			EndTime:   time.Now().Add(2 * time.Hour),
		},
	}
}

func randomUserTimeSlotsForMatching() []db.GetTimeSlotsForAllUsersRow {
	return []db.GetTimeSlotsForAllUsersRow{
		{
			UserID:    1,
			StartTime: time.Now(),
			EndTime:   time.Now().Add(3 * time.Hour),
		},
		{
			UserID:    2,
			StartTime: time.Now().Add(time.Hour),
			EndTime:   time.Now().Add(4 * time.Hour),
		},
	}
}

func randomEventForMatching() db.Event {
	return db.Event{
		EventID:          1,
		OrganizerID:      1,
		EventName:        "Test Event",
		EventDescription: "Test Event Description",
		Duration:         1,
		CreatedAt:        time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
	}
}
