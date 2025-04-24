package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gitatractivo/gotodocli/configs"
	"github.com/gitatractivo/gotodocli/internal/api/routes"
)

type Server struct {
	router *gin.Engine
	port   string
}

func NewServer(port string) *Server {
	router := gin.Default()

	// need /health endpoint to check if the server is running
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
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

	logFile, err := os.OpenFile(configs.GetConfig().LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		log.Println("Logging to file",configs.GetConfig().LogFile)
		fmt.Println("Logging to file",configs.GetConfig().LogFile)
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	//wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	os.Remove("/tmp/todo-server.pid")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}
	log.Println("Server exited properly")
	return nil
}
