package main

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Buzz)
var upgrader = websocket.Upgrader{}

func main() {
	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static resource
	e.Static("/", "public")

	// Websockets
	e.Any("/ws", handleConn)

	go handleBuzz(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func handleBuzz(e *echo.Echo) {
	for buzz := range broadcast {
		e.Logger.Debug("Buzz occured!")
		for client := range clients {
			err := client.WriteJSON(buzz)
			if err != nil {
				e.Logger.Debug(err)
			}
		}
	}
}

func handleConn(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	c.Logger().Debug("Client connected!")

	clients[ws] = true

	for {
		var buzz Buzz
		err := ws.ReadJSON(&buzz)
		if err != nil {
			c.Logger().Debug("Error happened, deleting client.")
			if e, ok := err.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			delete(clients, ws)
			return err
		}

		broadcast <- buzz
	}
}
