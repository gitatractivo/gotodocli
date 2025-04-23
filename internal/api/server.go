package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gitatractivo/gotodocli/internal/api/routes"
)

type Server struct {
	router *gin.Engine
	port   string
}

func NewServer(port string) *Server {
	router := gin.Default()

	routes.SetupRoutes(router)

	return &Server{
		router: router,
		port:   port,
	}
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	//start server in a go routine
	go func() {
		log.Printf("Starting server on port %s", s.port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
	//wait for interrupt signal to gracefully shutdown the server
	quit:=make(chan os.Signal,1)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	if err:=srv.Shutdown(ctx);err!=nil{
		log.Fatalf("Error shutting down server: %v",err)
	}
	log.Println("Server exited properly")
	return nil
}
