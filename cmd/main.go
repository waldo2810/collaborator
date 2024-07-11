package main

import (
	"github.com/labstack/echo/v4"
	"github.com/waldo2810/collaborator/internal/handler"
)

func main() {
    e := echo.New()

    e.Logger.Fatal(e.Start(":8080"))
}
