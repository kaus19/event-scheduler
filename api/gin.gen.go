// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new event
	// (POST /events)
	CreateEvent(c *gin.Context)
	// List all events
	// (GET /events/list)
	ListEvents(c *gin.Context)
	// List events by organizer
	// (GET /events/organizer/{organizer_id})
	ListEventsByOrganizer(c *gin.Context, organizerId int)
	// Delete an event
	// (DELETE /events/{event_id})
	DeleteEvent(c *gin.Context, eventId int)
	// Get an event by ID
	// (GET /events/{event_id})
	GetEventByID(c *gin.Context, eventId int)
	// Update an event's name, description and duration
	// (PUT /events/{event_id})
	UpdateEvent(c *gin.Context, eventId int)
	// Get matching time slots for event
	// (GET /matching-slots/event/{event_id})
	GetMatchingTimeSlotsForEvent(c *gin.Context, eventId int)
	// Delete an event time slot
	// (DELETE /time-slots/event)
	DeleteTimeSlotEvent(c *gin.Context, params DeleteTimeSlotEventParams)
	// Update an event time slot
	// (PUT /time-slots/event)
	UpdateTimeSlotEvent(c *gin.Context)
	// Get time slots by event
	// (GET /time-slots/event/{event_id})
	GetTimeSlotsByEvent(c *gin.Context, eventId int)
	// Delete a user time slot
	// (DELETE /time-slots/user)
	DeleteTimeSlotUser(c *gin.Context, params DeleteTimeSlotUserParams)
	// Create time slots for a user
	// (POST /time-slots/user)
	CreateTimeSlotUser(c *gin.Context)
	// Update a user time slot
	// (PUT /time-slots/user)
	UpdateTimeSlotUser(c *gin.Context)
	// Get time slots by user
	// (GET /time-slots/user/{user_id})
	GetTimeSlotsByUser(c *gin.Context, userId int)
	// List all users
	// (GET /users)
	GetUsers(c *gin.Context)
	// Create a new user
	// (POST /users)
	PostUsers(c *gin.Context)
	// Get a user by ID
	// (GET /users/{id})
	GetUsersId(c *gin.Context, id int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// CreateEvent operation middleware
func (siw *ServerInterfaceWrapper) CreateEvent(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateEvent(c)
}

// ListEvents operation middleware
func (siw *ServerInterfaceWrapper) ListEvents(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListEvents(c)
}

// ListEventsByOrganizer operation middleware
func (siw *ServerInterfaceWrapper) ListEventsByOrganizer(c *gin.Context) {

	var err error

	// ------------- Path parameter "organizer_id" -------------
	var organizerId int

	err = runtime.BindStyledParameterWithOptions("simple", "organizer_id", c.Param("organizer_id"), &organizerId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter organizer_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListEventsByOrganizer(c, organizerId)
}

// DeleteEvent operation middleware
func (siw *ServerInterfaceWrapper) DeleteEvent(c *gin.Context) {

	var err error

	// ------------- Path parameter "event_id" -------------
	var eventId int

	err = runtime.BindStyledParameterWithOptions("simple", "event_id", c.Param("event_id"), &eventId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter event_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteEvent(c, eventId)
}

// GetEventByID operation middleware
func (siw *ServerInterfaceWrapper) GetEventByID(c *gin.Context) {

	var err error

	// ------------- Path parameter "event_id" -------------
	var eventId int

	err = runtime.BindStyledParameterWithOptions("simple", "event_id", c.Param("event_id"), &eventId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter event_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetEventByID(c, eventId)
}

// UpdateEvent operation middleware
func (siw *ServerInterfaceWrapper) UpdateEvent(c *gin.Context) {

	var err error

	// ------------- Path parameter "event_id" -------------
	var eventId int

	err = runtime.BindStyledParameterWithOptions("simple", "event_id", c.Param("event_id"), &eventId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter event_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateEvent(c, eventId)
}

// GetMatchingTimeSlotsForEvent operation middleware
func (siw *ServerInterfaceWrapper) GetMatchingTimeSlotsForEvent(c *gin.Context) {

	var err error

	// ------------- Path parameter "event_id" -------------
	var eventId int

	err = runtime.BindStyledParameterWithOptions("simple", "event_id", c.Param("event_id"), &eventId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter event_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetMatchingTimeSlotsForEvent(c, eventId)
}

// DeleteTimeSlotEvent operation middleware
func (siw *ServerInterfaceWrapper) DeleteTimeSlotEvent(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteTimeSlotEventParams

	// ------------- Required query parameter "id" -------------

	if paramValue := c.Query("id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "id", c.Request.URL.Query(), &params.Id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "event_id" -------------

	if paramValue := c.Query("event_id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument event_id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "event_id", c.Request.URL.Query(), &params.EventId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter event_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteTimeSlotEvent(c, params)
}

// UpdateTimeSlotEvent operation middleware
func (siw *ServerInterfaceWrapper) UpdateTimeSlotEvent(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateTimeSlotEvent(c)
}

// GetTimeSlotsByEvent operation middleware
func (siw *ServerInterfaceWrapper) GetTimeSlotsByEvent(c *gin.Context) {

	var err error

	// ------------- Path parameter "event_id" -------------
	var eventId int

	err = runtime.BindStyledParameterWithOptions("simple", "event_id", c.Param("event_id"), &eventId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter event_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTimeSlotsByEvent(c, eventId)
}

// DeleteTimeSlotUser operation middleware
func (siw *ServerInterfaceWrapper) DeleteTimeSlotUser(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteTimeSlotUserParams

	// ------------- Required query parameter "id" -------------

	if paramValue := c.Query("id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "id", c.Request.URL.Query(), &params.Id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "user_id" -------------

	if paramValue := c.Query("user_id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument user_id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "user_id", c.Request.URL.Query(), &params.UserId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteTimeSlotUser(c, params)
}

// CreateTimeSlotUser operation middleware
func (siw *ServerInterfaceWrapper) CreateTimeSlotUser(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateTimeSlotUser(c)
}

// UpdateTimeSlotUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateTimeSlotUser(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateTimeSlotUser(c)
}

// GetTimeSlotsByUser operation middleware
func (siw *ServerInterfaceWrapper) GetTimeSlotsByUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId int

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", c.Param("user_id"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTimeSlotsByUser(c, userId)
}

// GetUsers operation middleware
func (siw *ServerInterfaceWrapper) GetUsers(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsers(c)
}

// PostUsers operation middleware
func (siw *ServerInterfaceWrapper) PostUsers(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUsers(c)
}

// GetUsersId operation middleware
func (siw *ServerInterfaceWrapper) GetUsersId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsersId(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/events", wrapper.CreateEvent)
	router.GET(options.BaseURL+"/events/list", wrapper.ListEvents)
	router.GET(options.BaseURL+"/events/organizer/:organizer_id", wrapper.ListEventsByOrganizer)
	router.DELETE(options.BaseURL+"/events/:event_id", wrapper.DeleteEvent)
	router.GET(options.BaseURL+"/events/:event_id", wrapper.GetEventByID)
	router.PUT(options.BaseURL+"/events/:event_id", wrapper.UpdateEvent)
	router.GET(options.BaseURL+"/matching-slots/event/:event_id", wrapper.GetMatchingTimeSlotsForEvent)
	router.DELETE(options.BaseURL+"/time-slots/event", wrapper.DeleteTimeSlotEvent)
	router.PUT(options.BaseURL+"/time-slots/event", wrapper.UpdateTimeSlotEvent)
	router.GET(options.BaseURL+"/time-slots/event/:event_id", wrapper.GetTimeSlotsByEvent)
	router.DELETE(options.BaseURL+"/time-slots/user", wrapper.DeleteTimeSlotUser)
	router.POST(options.BaseURL+"/time-slots/user", wrapper.CreateTimeSlotUser)
	router.PUT(options.BaseURL+"/time-slots/user", wrapper.UpdateTimeSlotUser)
	router.GET(options.BaseURL+"/time-slots/user/:user_id", wrapper.GetTimeSlotsByUser)
	router.GET(options.BaseURL+"/users", wrapper.GetUsers)
	router.POST(options.BaseURL+"/users", wrapper.PostUsers)
	router.GET(options.BaseURL+"/users/:id", wrapper.GetUsersId)
}
