package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang.org/x/net/context"
	"io/ioutil"
	_ "ms/spatial/cmd/api/docs"
	"ms/spatial/internal/api/handler"
	"ms/spatial/pkg/contract"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	_ = godotenv.Load()
	server := echo.New()
	server.Use(middleware.Secure())
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.Gzip())

	payload, err := ioutil.ReadFile(os.Getenv("API_POINTS_DATA"))
	if err != nil {
		server.Logger.Fatalf("ioutil.ReadFile failed with %s\n", err)
	}

	var points []contract.Point
	if err = json.Unmarshal(payload, &points); err != nil {
		server.Logger.Fatalf("json.Unmarshal failed with %s\n", err)
	}

	pointsHandler := handler.NewPoint(handler.PointOpts{
		Points: points,
	})

	server.GET("/docs/*", echoSwagger.WrapHandler)
	server.GET(handler.PointGet, pointsHandler.Get)

	go func() {
		binding := os.Getenv("API_PORT")
		if err := server.Start(binding); err != nil && err != http.ErrServerClosed {
			server.Logger.Fatalf("server.Start failed with %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatalf("shutdown failed with %s\n", err)
	}
}
