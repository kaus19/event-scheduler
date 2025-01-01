package api

import (
	"database/sql"
	"net/http"

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

	var req *CreateEventJSONBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	arg := db.CreateEventParams{
		OrganizerID:      int32(req.OrganizerId),
		EventName:        req.EventName,
		EventDescription: sql.NullString{String: req.EventDescription, Valid: true},
	}

	resp, err := server.store.CreateEvent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
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
func (server Server) UpdateEventDescription(ctx *gin.Context, eventId int) {

	if err := ctx.ShouldBindUri(&eventId); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	var req *UpdateEventDescriptionJSONBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	arg := db.UpdateEventDescriptionParams{
		EventID:          int32(eventId),
		EventDescription: sql.NullString{String: req.EventDescription, Valid: true},
	}

	err := server.store.UpdateEventDescription(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
