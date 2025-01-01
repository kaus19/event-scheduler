package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	api "github.com/kaus19/event-scheduler/api"
	db "github.com/kaus19/event-scheduler/db/sqlc"
	"github.com/kaus19/event-scheduler/util"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Cannot load config from path: %v, error: %v", ".", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	r := gin.Default()

	r.Use(cors.Default())

	api.RegisterHandlers(r, server)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%s", config.ServerAddress),
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")

	if err := s.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}
}
