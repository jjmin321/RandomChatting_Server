package chattingserver

import (
	"container/list"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

// socketServerMethod - 구현되어 있는 메서드 모음
type socketServerMethod interface {
	init()
	Socket()
	HandleConnection()
	HandleClient()
	SendJoinMsgToClient()
	sendMsgToClient()
	sendMsgToRoomClients()
	sendMsgToAllClients()
	AllocateEmptyRoom()
	RecvMsgFromClient()
	DeleteFromList()
	DupUserCheck()
}

const (
	// LOGIN : Magic Number for Chatting Login
	LOGIN = "1"
	// CHAT : Magic Number for Chatting
	CHAT = "2"
	// MAXUSER : Magic Number for Chatting room max user
	MAXUSER = 2
	// MAXCOUNT : Magic Number for Chatting room max count
	MAXCOUNT = 50
)

// Client - 채팅을 이용하는 사용자의 정보
type Client struct {
	ws   *websocket.Conn
	read chan string
	quit chan int
	name string
	room *Room
}

// Room - 채팅방 정보
type Room struct {
	num        int
	clientlist *list.List
}

var (
	// Roomlist - 이중 링크드 리스트
	Roomlist *list.List
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

// HandleConnection - 클라이언트를 객체를 생성한 후, HandleClient 쓰레드 호출
func HandleConnection(ws *websocket.Conn) {
	read := make(chan string)
	quit := make(chan int)
	client := &Client{ws, read, quit, "익명", &Room{-1, list.New()}}
	go HandleClient(client)
	log.Printf("%s에서 채팅 서버에 입장하였습니다.\t", ws.RemoteAddr().String())
}

// HandleClient - RecvMsgFromClient 쓰레드를 호출하고, 클라이언트의 명령이 오면 메세지 전송, 채팅 종료 구문을 실행시킨다.
func HandleClient(client *Client) {
	for {
		select {
		case msg := <-client.read:
			if strings.HasPrefix(msg, "[확성기]") {
				sendMsgToAllClients(client.name, msg)
			} else {
				sendMsgToRoomClients(client.room, client.name, msg)
			}

		case <-client.quit:
			log.Printf("%s : %d번째 방의 %s님이 채팅 서버에서 나가셨습니다.", client.ws.RemoteAddr().String(), client.room.num, client.name)
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

		room := AllocateEmptyRoom()
		if room.num < 1 {
			client.ws.Close()
			log.Print("방 인원이 다 찼습니다.")
		}
		client.room = room

		if !client.DupUserCheck() {
			log.Print("닉네임이 중복됨!")
			client.quit <- 0
			return
		}
		log.Printf("안녕하세요 %s님, %d번째 방에 입장하셨습니다.\n", client.name, client.room.num)
		// sendMsgToRoomClients(client.room, client.name, "님이 입장하셨습니다.")
		SendJoinMsgToClient(client.room, client.name)
		room.clientlist.PushBack(*client)

	case CHAT:
		log.Printf("\n"+client.name+" 님의 메시지: %s\n", strmsgs[1])
		client.read <- strmsgs[1]
	}
}

// SendJoinMsgToClient - 클라이언트가 연결된 후 방이 배정되면 해당 방에 입장하였습니다 메세지를 전송.
func SendJoinMsgToClient(room *Room, sender string) {
	chatting := sender + "님이 입장하였습니다"
	for e := room.clientlist.Front(); e != nil; e = e.Next() {
		c := e.Value.(Client)
		err := c.ws.WriteMessage(websocket.TextMessage, []byte(chatting))
		if err != nil {
			log.Print("입장 채팅 전송 중 에러 발생")
		}
		log.Printf("%d방에 전송된 입장 메세지 : %s", room.num, chatting)
	}
}

// sendMsgToClient - 클라이언트에게 웹소켓을 통해 메세지를 전송
func sendMsgToClient(client *Client, sender string, msg string) {
	chatting := sender + " : " + msg
	err := client.ws.WriteMessage(websocket.TextMessage, []byte(chatting))
	if err != nil {
		log.Print("채팅 전송 중 에러 발생")
	}
	log.Printf("%s님에게 전송된 메세지 : %s", client.name, chatting)
}

// sendMsgToRoomClients - 이중링크드리스트를 순회하여 클라이언트의 방 인덱스를 찾은 뒤, 방 인덱스와 메세지를 sendMsgToClient에게 전달한다.
func sendMsgToRoomClients(room *Room, sender string, msg string) {
	for e := room.clientlist.Front(); e != nil; e = e.Next() {
		c := e.Value.(Client)
		sendMsgToClient(&c, sender, msg)
	}
}

// sendMsgToAllClients - 이중링크드리스트를 순회하여 존재하는 모든 클라이언트의 방 인덱스를 찾은 뒤, 방 인덱스와 메세지를 sendMsgToClient에게 전달한다.
func sendMsgToAllClients(sender string, msg string) {
	for re := Roomlist.Front(); re != nil; re = re.Next() {
		r := re.Value.(Room)
		for e := r.clientlist.Front(); e != nil; e = e.Next() {
			c := e.Value.(Client)
			sendMsgToClient(&c, sender, msg)
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
			}
		}
	}
}