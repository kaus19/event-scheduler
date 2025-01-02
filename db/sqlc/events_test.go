package db

import (
	"context"
	"testing"

	"github.com/kaus19/event-scheduler/util"
	"github.com/stretchr/testify/require"
)

func createRandomEvent(t *testing.T) Event {
	_, err := testQueries.CreateUser(context.Background(), "Aditya")
	arg := CreateEventParams{
		OrganizerID:      int32(1),
		EventName:        util.RandomString(5),
		EventDescription: util.RandomString(10),
	}

	event, err := testQueries.CreateEvent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, event)

	require.Equal(t, arg.OrganizerID, event.OrganizerID)
	require.Equal(t, arg.EventName, event.EventName)
	require.Equal(t, arg.EventDescription, event.EventDescription)

	require.NotZero(t, event.EventID)
	require.NotZero(t, event.CreatedAt)

	return event
}

func TestCreateEvent(t *testing.T) {
	createRandomEvent(t)
}
