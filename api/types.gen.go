// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"time"
)

// User defines model for User.
type User struct {
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
}

// CreateEventJSONBody defines parameters for CreateEvent.
type CreateEventJSONBody struct {
	EventDescription string `json:"event_description"`
	EventName        string `json:"event_name"`
	OrganizerId      int    `json:"organizer_id"`
}

// UpdateEventDescriptionJSONBody defines parameters for UpdateEventDescription.
type UpdateEventDescriptionJSONBody struct {
	EventDescription string `json:"event_description"`
}

// PostUsersJSONBody defines parameters for PostUsers.
type PostUsersJSONBody struct {
	Name *string `json:"name,omitempty"`
}

// CreateEventJSONRequestBody defines body for CreateEvent for application/json ContentType.
type CreateEventJSONRequestBody CreateEventJSONBody

// UpdateEventDescriptionJSONRequestBody defines body for UpdateEventDescription for application/json ContentType.
type UpdateEventDescriptionJSONRequestBody UpdateEventDescriptionJSONBody

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody PostUsersJSONBody
