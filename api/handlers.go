package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/kaus19/event-scheduler/db/sqlc"
)

type Server struct {
	store db.Store
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	return server
}

// (POST /users/)
func (server Server) PostUsers(ctx *gin.Context) {

	var req *PostUsersJSONBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := server.store.CreateUser(ctx, *req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// (GET /users/)
func (server Server) GetUsers(ctx *gin.Context) {

	users, err := server.store.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// (GET /users/{id})
func (server Server) GetUsersId(ctx *gin.Context, id int) {

	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := server.store.GetUserByID(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// (POST /events/)
func (server Server) CreateEvent(ctx *gin.Context) {

	var req *CreateEventJSONRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	arg := db.CreateEventParams{
		OrganizerID:      int32(req.OrganizerId),
		EventName:        req.EventName,
		EventDescription: req.EventDescription,
		Duration:         int32(req.Duration),
	}

	resp, err := server.store.CreateEvent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(req.StartTime) != len(req.EndTime) {
		ctx.JSON(http.StatusBadRequest, "Start and end times should have the same length")
		return
	}
	for i := range req.StartTime {
		// add check for start time < end time
		arg := db.CreateTimeSlotEventParams{
			EventID:   int32(resp.EventID),
			StartTime: req.StartTime[i],
			EndTime:   req.EndTime[i],
		}

		_, err := server.store.CreateTimeSlotEvent(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		arg1 := db.CreateTimeSlotUserParams{
			UserID:    int32(req.OrganizerId),
			StartTime: req.StartTime[i],
			EndTime:   req.EndTime[i],
		}
		_, err = server.store.CreateTimeSlotUser(ctx, arg1)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

// (GET /events/list)
func (server Server) ListEvents(ctx *gin.Context) {

	events, err := server.store.ListEvents(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, events)
}

// (GET /events/{event_id})
func (server Server) GetEventByID(ctx *gin.Context, eventId int) {

	if err := ctx.ShouldBindUri(&eventId); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	event, err := server.store.GetEventByID(ctx, int32(eventId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, event)
}

// (DELETE /events/{event_id})
func (server Server) DeleteEvent(ctx *gin.Context, eventId int) {

	if err := ctx.ShouldBindUri(&eventId); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err := server.store.DeleteEvent(ctx, int32(eventId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// (GET /events/organiser/{organiser_id}/)
func (server Server) ListEventsByOrganizer(ctx *gin.Context, organizerId int) {

	if err := ctx.ShouldBindUri(&organizerId); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	events, err := server.store.ListEventsByOrganizer(ctx, int32(organizerId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, events)
}

// (PUT /events/{event_id})
func (server Server) UpdateEvent(ctx *gin.Context, eventId int) {

	if err := ctx.ShouldBindUri(&eventId); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	var req *UpdateEventJSONRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	arg := db.UpdateEventParams{
		EventID:          int32(eventId),
		EventName:        req.EventName,
		EventDescription: req.EventDescription,
		Duration:         int32(req.Duration),
		UpdatedAt:        time.Now(),
	}

	err := server.store.UpdateEvent(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// (GET /time-slots/user)
func (server Server) GetTimeSlotsByUser(ctx *gin.Context, userId int) {

	if err := ctx.ShouldBindUri(&userId); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	Slots, err := server.store.GetTimeSlotsByUser(ctx, int32(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, Slots)
}

// (POST /time-slots/user)
func (server Server) CreateTimeSlotUser(ctx *gin.Context) {
	var req *CreateTimeSlotUserJSONRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if len(req.StartTime) != len(req.EndTime) {
		ctx.JSON(http.StatusBadRequest, "Start and end times should have the same length")
		return
	}
	for i := range req.StartTime {
		// add check for start time < end time
		arg := db.CreateTimeSlotUserParams{
			UserID:    int32(req.UserId),
			StartTime: req.StartTime[i],
			EndTime:   req.EndTime[i],
		}

		_, err := server.store.CreateTimeSlotUser(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	ctx.Status(http.StatusOK)
}

// (PUT /time-slots/user)
func (server Server) UpdateTimeSlotUser(ctx *gin.Context) {
	var req *UpdateTimeSlotUserJSONRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	arg := db.UpdateTimeSlotUserParams{
		ID:        int32(req.Id),
		UserID:    int32(req.UserId),
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	err := server.store.UpdateTimeSlotUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// (DELETE /time-slots/user)
func (server Server) DeleteTimeSlotUser(ctx *gin.Context, params DeleteTimeSlotUserParams) {
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	arg := db.DeleteTimeSlotUserParams{
		ID:     int32(params.Id),
		UserID: int32(params.UserId),
	}

	err := server.store.DeleteTimeSlotUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// (GET /time-slots/event)
func (server Server) GetTimeSlotsByEvent(ctx *gin.Context, eventId int) {

	if err := ctx.ShouldBindUri(&eventId); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	Slots, err := server.store.GetTimeSlotsByEvent(ctx, int32(eventId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, Slots)
}

// (PUT /time-slots/event)
func (server Server) UpdateTimeSlotEvent(ctx *gin.Context) {
	var req *UpdateTimeSlotEventJSONRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	arg := db.UpdateTimeSlotEventParams{
		ID:        int32(req.Id),
		EventID:   int32(req.EventId),
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	err := server.store.UpdateTimeSlotEvent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// (DELETE /time-slots/event)
func (server Server) DeleteTimeSlotEvent(ctx *gin.Context, params DeleteTimeSlotEventParams) {
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	arg := db.DeleteTimeSlotEventParams{
		ID:      int32(params.Id),
		EventID: int32(params.EventId),
	}

	err := server.store.DeleteTimeSlotEvent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// (GET /matching-slots/event)
func (server Server) GetMatchingTimeSlotsForEvent(ctx *gin.Context, eventId int) {
	if err := ctx.ShouldBindUri(&eventId); err != nil {
		ctx.JSON(http.StatusBadRequest, "Error 1")
		return
	}

	rows, err := server.store.GetTimeSlotsByEvent(ctx, int32(eventId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error 2")
		return
	}

	eventTimeSlots := []TimeSlot{}

	for _, row := range rows {
		var slot TimeSlot
		slot.Start = row.StartTime
		slot.End = row.EndTime
		eventTimeSlots = append(eventTimeSlots, slot)
	}

	// get time slots for all users

	rows1, err := server.store.GetTimeSlotsForAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error 3")
		return
	}

	userTimeSlots := make(map[int][]TimeSlot)

	for _, row1 := range rows1 {
		var userID int
		var slot TimeSlot
		userID = int(row1.UserID)
		slot.Start = row1.StartTime
		slot.End = row1.EndTime
		userTimeSlots[userID] = append(userTimeSlots[userID], slot)
	}

	// get duration
	event, err := server.store.GetEventByID(ctx, int32(eventId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error 4")
		return
	}
	timeSlotDuration := time.Duration(event.Duration) * time.Hour

	// Divide Event Time Slots into Equal Durations
	dividedSlots := []TimeSlot{}
	for _, slot := range eventTimeSlots {
		for t := slot.Start; t.Before(slot.End); t = t.Add(timeSlotDuration) {
			end := t.Add(timeSlotDuration)
			if end.After(slot.End) {
				end = slot.End
			}
			dividedSlots = append(dividedSlots, TimeSlot{Start: t, End: end})
		}
	}

	// Generate 2D Matrix
	users := make([]int, 0, len(userTimeSlots))
	for userID := range userTimeSlots {
		users = append(users, userID)
	}

	matrix := make([][]int, len(users))
	for i, userID := range users {
		matrix[i] = make([]int, len(dividedSlots))
		for j, slot := range dividedSlots {
			matrix[i][j] = 0
			for _, userSlot := range userTimeSlots[userID] {
				if (slot.Start.After(userSlot.Start) || slot.Start.Equal(userSlot.Start)) &&
					(slot.End.Before(userSlot.End) || slot.End.Equal(userSlot.End)) {
					matrix[i][j] = 1
					break
				}
			}
		}
	}

	// Find Best Time Slots
	res := FindBestTimeSlots(matrix, dividedSlots, users)

	ctx.JSON(http.StatusOK, res)
}

func FindBestTimeSlots(matrix [][]int, dividedSlots []TimeSlot, users []int) AvailabilityResult {
	allAvailableSlots := []TimeSlot{}
	partialSlots := []PartialSlotDetail{}
	maxAvailableUsers := len(matrix)

	// Track unavailable users per slot
	for j, slot := range dividedSlots {
		countAvailable := 0
		unavailableUsers := []int{}

		for i, row := range matrix {
			if row[j] == 1 {
				countAvailable++
			} else {
				unavailableUsers = append(unavailableUsers, users[i])
			}
		}

		if countAvailable == maxAvailableUsers {
			allAvailableSlots = append(allAvailableSlots, slot)
		} else {
			partialSlots = append(partialSlots, PartialSlotDetail{
				Slot:             slot,
				UnavailableUsers: unavailableUsers,
			})
		}
	}

	return AvailabilityResult{
		AllAvailableSlots: allAvailableSlots,
		PartialSlots:      partialSlots,
	}
}

type TimeSlot struct {
	Start time.Time
	End   time.Time
}

type AvailabilityResult struct {
	AllAvailableSlots []TimeSlot          `json:"all_available_slots"`
	PartialSlots      []PartialSlotDetail `json:"partial_slots"`
}

type PartialSlotDetail struct {
	Slot             TimeSlot `json:"slot"`
	UnavailableUsers []int    `json:"unavailable_users"`
}
