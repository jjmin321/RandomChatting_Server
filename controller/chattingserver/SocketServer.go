package chattingserver

import (
	"container/list"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

// socketServerMethod - 구현되어 있는 메서드 모음
type socketServerMethod interface {
	init()
	Socket()
	RecoverServer()
	HandleConnection()
	HandleClient()
	SendJoinMsgToClient()
	GetUserList()
	SendMsgToClient()
	SendMsgToRoomClients()
	SendMsgToAllClients()
	AllocateEmptyRoom()
	RecvMsgFromClient()
	DeleteFromList()
	DupUserCheck()
}

const (
	// LOGIN : Magic Number for Chatting Login
	LOGIN = "1"
	// ROOMCHAT : Magic Number for Room Chatting
	ROOMCHAT = "2"
	// ALLCHAT : Magic Number for All Chatting
	ALLCHAT = "3"
	// MAXUSER : Magic Number for Chatting room max user
	MAXUSER = 2
	// MAXCOUNT : Magic Number for Chatting room max count
	MAXCOUNT = 50
)

// Client - 채팅을 이용하는 사용자의 정보
type Client struct {
	ws           *websocket.Conn
	roomChatting chan string
	allChatting  chan string
	quit         chan int
	name         string
	room         *Room
}

// Room - 채팅방 정보
type Room struct {
	num        int
	clientlist *list.List
}

var (
	// Roomlist - 이중 링크드 리스트
	Roomlist *list.List
	// UserList - 채팅에 참여하고 있는 전체 멤버 목록
	UserList []string
	// Upgrader - http 프로토콜을 ws 프로토콜로 바꿈
	Upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// Init - 50개 채팅방 초기화
func init() {
	Roomlist = list.New()
	for i := 0; i < MAXCOUNT; i++ {
		room := &Room{i + 1, list.New()}
		Roomlist.PushBack(*room)
	}
}

// Socket - 클라이언트가 접속한 HTTP 프로토콜을 WebSocket 프로토콜로 업그레이드 시킨 후, HandleConnection 쓰레드 시작.
func Socket(c echo.Context) error {
	ws, err := Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Print("error occured! : ", err.Error())
		return nil
	}
	go HandleConnection(ws)
	return nil
}

// RecoverServer - 소켓 통신 중 고루틴에서 발생하는 예상치 못한 에러를 무시해준다.
func RecoverServer() {
	socketErr := recover()
	if socketErr != nil {
		log.Print("Recovered", socketErr)
		debug.PrintStack()
	}
}

// HandleConnection - 클라이언트를 객체를 생성한 후, HandleClient 쓰레드 호출
func HandleConnection(ws *websocket.Conn) {
	defer RecoverServer()
	roomChatting := make(chan string)
	allChatting := make(chan string)
	quit := make(chan int)
	client := &Client{ws, roomChatting, allChatting, quit, "익명", &Room{-1, list.New()}}
	go HandleClient(client)
}

// HandleClient - RecvMsgFromClient 쓰레드를 호출하고, 클라이언트의 명령이 오면 메세지 전송, 채팅 종료 구문을 실행시킨다.
func HandleClient(client *Client) {
	defer RecoverServer()
	for {
		select {
		case roomMsg := <-client.roomChatting:
			SendMsgToRoomClients(client.room, client.name, roomMsg)

		case allMsg := <-client.allChatting:
			SendMsgToAllClients(client.name, allMsg)

		case <-client.quit:
			client.ws.Close()
			client.DeleteFromList()
			return

		default:
			go RecvMsgFromClient(client)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

// RecvMsgFromClient - 클라이언트에서 명령이 올 때까지 대기하다가 명령이 오면 채널을 통해 HandleClient에게 값을 전달, 방이 모두 찼거나 한 아이디가 두 번 이상 접속한 경우 연결을 종료시킴.
func RecvMsgFromClient(client *Client) {
	defer RecoverServer()
	_, bytemsg, err := client.ws.ReadMessage()
	if err != nil {
		client.quit <- 0
		return
	}
	msg := string(bytemsg)

	strmsgs := strings.Split(msg, "ᗠ")

	switch strmsgs[0] {
	case LOGIN:
		client.name = strings.TrimSpace(strmsgs[1])

		room := AllocateEmptyRoom()
		if room.num < 1 {
			client.ws.Close()
		}
		client.room = room

		if !client.DupUserCheck() {
			client.quit <- 0
			return
		}
		log.Printf("%s님이 %d번째 방에 입장하셨습니다\n", client.name, client.room.num)

		client.ws.WriteMessage(websocket.TextMessage, []byte("방 번호ᗠ"+strconv.Itoa(client.room.num)))
		client.SendJoinMsgToClient()
		client.GetUserList()
		room.clientlist.PushBack(*client)

	case ROOMCHAT:
		client.roomChatting <- strmsgs[1]

	case ALLCHAT:
		client.allChatting <- strmsgs[1]
	}
}

// SendJoinMsgToClient - 클라이언트가 연결된 후 방이 배정되면 해당 방에 입장하였습니다 메세지를 전송.
func (client *Client) SendJoinMsgToClient() {
	var chatting string
	for e := client.room.clientlist.Front(); e != nil; e = e.Next() {
		c := e.Value.(Client)
		if client.name != c.name {
			chatting = "방 유저ᗠ" + client.name
		}
		c.ws.WriteMessage(websocket.TextMessage, []byte(chatting))
	}
}

// GetUserList - 접속되어 있는 사람들의 정보를 반환함
func (client *Client) GetUserList() {
	var roomUser string
	for e := client.room.clientlist.Front(); e != nil; e = e.Next() {
		c := e.Value.(Client)
		if client.name != c.name {
			roomUser = c.name
			client.ws.WriteMessage(websocket.TextMessage, []byte("방 유저ᗠ"+roomUser))
		}
	}
}

// SendMsgToClient - 클라이언트에게 웹소켓을 통해 메세지를 전송
func SendMsgToClient(client *Client, sender string, msg string, all bool) {
	if all == true {
		chatting := "전체채팅ᗠ" + sender + "ᗠ" + msg
		client.ws.WriteMessage(websocket.TextMessage, []byte(chatting))
	} else {
		chatting := "랜덤채팅ᗠ" + sender + "ᗠ" + msg
		client.ws.WriteMessage(websocket.TextMessage, []byte(chatting))
	}
}

// SendMsgToRoomClients - 이중링크드리스트를 순회하여 클라이언트의 방 인덱스를 찾은 뒤, 방 인덱스와 메세지를 sendMsgToClient에게 전달한다.
func SendMsgToRoomClients(room *Room, sender string, msg string) {
	for e := room.clientlist.Front(); e != nil; e = e.Next() {
		c := e.Value.(Client)
		SendMsgToClient(&c, sender, msg, false)
	}
}

// SendMsgToAllClients - 이중링크드리스트를 순회하여 존재하는 모든 클라이언트의 방 인덱스를 찾은 뒤, 방 인덱스와 메세지를 sendMsgToClient에게 전달한다.
func SendMsgToAllClients(sender string, msg string) {
	for re := Roomlist.Front(); re != nil; re = re.Next() {
		r := re.Value.(Room)
		for e := r.clientlist.Front(); e != nil; e = e.Next() {
			c := e.Value.(Client)
			SendMsgToClient(&c, sender, msg, true)
		}
	}
}

// AllocateEmptyRoom - 이중링크드리스트를 오름차순으로 순회하며 유저가 2명이 아닌 리스트를 배정, 모든 방에 유저가 찼으면 접속하지 못함.
func AllocateEmptyRoom() *Room {
	for e := Roomlist.Front(); e != nil; e = e.Next() {
		r := e.Value.(Room)
		if r.clientlist.Len() < MAXUSER {
			return &r
		}
	}
	return &Room{-1, list.New()}
}

// DupUserCheck - 중복되는 닉네임이 이미 채팅에 접속되어 있을 경우 입장시키지 않음.
func (client *Client) DupUserCheck() bool {
	for re := Roomlist.Front(); re != nil; re = re.Next() {
		r := re.Value.(Room)
		for e := r.clientlist.Front(); e != nil; e = e.Next() {
			c := e.Value.(Client)
			if strings.Compare(client.name, c.name) == 0 {
				return false
			}
		}
	}
	return true
}

// DeleteFromList - 클라이언트의 접속이 끊어지면 링크드리스트에서도 삭제함.
func (client *Client) DeleteFromList() {
	for re := Roomlist.Front(); re != nil; re = re.Next() {
		r := re.Value.(Room)
		for e := r.clientlist.Front(); e != nil; e = e.Next() {
			c := e.Value.(Client)
			if client.ws.RemoteAddr() == c.ws.RemoteAddr() {
				r.clientlist.Remove(e)
			} else if c.name != client.name {
				c.ws.WriteMessage(websocket.TextMessage, []byte("사람 나감ᗠ"+client.name))
			}
		}
	}
	log.Printf("%s님이 %d번째 방에서 퇴장하셨습니다\n", client.name, client.room.num)
}
