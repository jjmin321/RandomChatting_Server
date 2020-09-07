package main

import (
	"bytes"
	"container/list"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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
	ws   websocket.Conn
	read chan string
	quit chan int
	name string
	room *TestRoom
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
	log.Print("누군가가 입장하였습니다")
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}
		go testhandleConnection(ws)
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "접속 끊김 ㅃㅇ",
	})
}

func testhandleConnection(ws *websocket.Conn) {
	read := make(chan string)
	quit := make(chan int)
	client := &TestClient{*ws, read, quit, "익명", &TestRoom{-1, list.New()}}
	go testhandleClient(client)
}

func testhandleClient(client *TestClient) {
	for {
		select {
		case msg := <-client.read:
			if strings.HasPrefix(msg, "[확성기]") {
				testsendToAllClients(client.name, msg)
			} else {
				testsendToRoomClients(client.room, client.name, msg)
			}

		case <-client.quit:
			log.Print(client.name + " 님이 나갔습니다.")
			client.ws.Close()
			client.deleteFromList()
			return

		default:
			go recvFromClient(client)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func testrecvFromClient(client *TestClient) {
	// Read
	_, bytemsg, err := client.ws.ReadMessage()
	if err != nil {
		client.quit <- 0
		return
	}
	msg := string(bytemsg)
	log.Print("1 : 로그인, 2 : 채팅 ", msg)

	strmsgs := strings.Split(msg, "|")

	switch strmsgs[0] {
	case LOGIN:
		client.name = strings.TrimSpace(strmsgs[1])

		room := testallocateEmptyRoom()
		if room.num < 1 {
			client.ws.Close()
			log.Print("방 인원이 다 찼습니다.")
		}
		client.room = room

		if !client.testdupUserCheck() {
			client.ws.Close()
			log.Print("닉네임 중복!")
			return
		}
		log.Printf("안녕하세요 %s님, %d번째 방에 입장하셨습니다.\n", client.name, client.room.num)
		testsendToRoomClients(client.room, client.name, "님이 입장하셨습니다.")
		room.clientlist.PushBack(*client)

	case CHAT:
		log.Printf("\n"+client.name+" 님의 메시지: %s\n", strmsgs[1])
		client.read <- strmsgs[1]
	}
}

func testsendToRoomClients(room *TestRoom, sender string, msg string) {
	for e := room.clientlist.Front(); e != nil; e = e.Next() {
		c := e.Value.(TestClient)
		testsendToClient(&c, sender, msg)
	}
}

func testsendToAllClients(sender string, msg string) {
	for re := roomlist.Front(); re != nil; re = re.Next() {
		r := re.Value.(Room)
		for e := r.clientlist.Front(); e != nil; e = e.Next() {
			c := e.Value.(TestClient)
			testsendToClient(&c, sender, msg)
		}
	}
}

// 원래는 버퍼로 해서 작성했다면 이제는 웹소켓으로 작성하게끔 짜야됨.
func testsendToClient(client *TestClient, sender string, msg string) {
	// err := client.ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
	// if err != nil {
	// 	c.Logger().Error(err)
	// }
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString(sender)
	buffer.WriteString("] ")
	buffer.WriteString(msg)

	log.Printf("%s님에게 전송된 메세지 : %s", client.name, buffer.String())
	fmt.Fprintf(client.connection, "%s", buffer.String())
}

func testallocateEmptyRoom() *TestRoom {
	for e := roomlist.Front(); e != nil; e = e.Next() {
		r := e.Value.(TestRoom)
		if r.clientlist.Len() < ROOM_MAX_USER {
			return &r
		}
	}
	// 방 다참
	return &TestRoom{-1, list.New()}
}

func (client *TestClient) testdupUserCheck() bool {
	for re := roomlist.Front(); re != nil; re = re.Next() {
		r := re.Value.(TestRoom)
		for e := r.clientlist.Front(); e != nil; e = e.Next() {
			c := e.Value.(Client)
			if strings.Compare(client.name, c.name) == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
}