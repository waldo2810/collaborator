package main

import (
	"github.com/labstack/echo/v4"
	"github.com/waldo2810/collaborator/internal/handler"
	myWS "github.com/waldo2810/collaborator/internal/websocket"
	"golang.org/x/net/websocket"
)

func main() {
	e := echo.New()

	// Websocket server
	myWSServer := myWS.NewServer()
	e.GET("/ws", func(c echo.Context) error {
		websocket.Handler(myWSServer.HandleWS).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	})

	// Static files
	e.Static("/static", "static")

	homeHandler := handler.HomeHandler{}

	e.GET("/", homeHandler.HandleHome)

	e.Logger.Fatal(e.Start(":8080"))
}
