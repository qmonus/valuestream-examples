package server

import (
	"context"
	"demo-backend-app/pkg/database"
	"demo-backend-app/pkg/env"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	srv    *http.Server
	router *gin.Engine
	db     *gorm.DB
}

func (s *Server) Run() {
	s.initDB()
	s.initSrv()
	s.initRouter()
	s.run()
}

func (s *Server) run() {
	go func() {
		// service connections
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

func (s *Server) initSrv() {
	s.router = gin.Default()
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", env.Host, env.Port),
		Handler: s.router,
	}
	s.srv = srv
}

func (s *Server) initDB() {
	db, err := database.ConnectMySQL()
	if err != nil {
		panic(err)
	}
	database.Initialize(db)
	s.db = db
}
