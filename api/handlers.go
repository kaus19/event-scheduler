package api

import (
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
