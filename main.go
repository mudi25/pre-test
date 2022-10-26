package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mudi25/pretest/controller"
)

func shutdown(ctx context.Context, e *echo.Echo) {
	if err := e.Shutdown(ctx); err != nil {
		log.Println("failed shutdown echo")
		return
	}
	log.Println("shutdown server")
}
func main() {
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8080")
	}
	e := controller.NewController()
	go func() {
		if err := e.Start(":" + os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	shutdown(ctx, e)
}
