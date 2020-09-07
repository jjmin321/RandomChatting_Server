package main

import (
	"container/list"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	LOGIN          = "1"
	CHAT           = "2"
	ROOM_MAX_USER  = 2
	ROOM_MAX_COUNT = 50
)

// TestClient - 채팅을 이용하는 사용자의 정보
type TestClient struct {
	connection websocket.Conn
	read       chan string
	quit       chan int
	name       string
	room       *Room
}

// TestRoom - 채팅방 정보
type TestRoom struct {
	num        int
	clientlist *list.List
}

var (
	Testroomlist *list.List
	Testupgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func init() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", testsocket)
	e.Logger.Fatal(e.Start(":80"))
}

func testsocket(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	log.Print("채팅 서버를 80번 포트에 열었습니다.")
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

func main() {
}
