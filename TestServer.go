package main

import (
	"container/list"
	"fmt"
	"
	"github.com/gorilla/websocket"e"
	"log"
	"net/http"
	"strings"
	"tim
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
	room       *TestRoom
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
	Testupgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := Testupgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	log.Print("채팅 서버를 80번 포트에 열었습니다.")
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(ws.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		go testhandleConnection(ws)

		
		fmt.Printf("%s\n", msg)
	}
}

func testhandleConnection(connection *websocket.Conn) {
	read := make(chan string)
	quit := make(chan int)
	client := &Client{*connection, read, quit, "익명", &Room{-1, list.New()}}
	go testhandleClient(client)
}

func testhandleClient(client *Client) {
	for {
		select {
		case msg := <-client.read:
			if strings.HasPrefix(msg, "[확성기]") {
				sendToAllClients(client.name, msg)
			} else if strings.HasPrefix(msg, "[귓속말]") {
				sendToClientToClient(client, msg)
			} else {
				sendToRoomClients(client.room, client.name, msg)
			}

		case <-client.quit:
			log.Print(client.name + " 님이 나갔습니다.")
			client.connection.Close()
			client.deleteFromList()
			return

		default:
			go recvFromClient(client)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func testrecvFromClient(client *Client) {
	// Read
	_, msg, err := ws.ReadMessage()
	if err != nil {
		c.Logger().Error(err)
	}
}
func main() {
}
