package main

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"net"
	"strings"
	"time"
)

const (
	LOGIN          = "1"
	CHAT           = "2"
	ROOM_MAX_USER  = 100
	ROOM_MAX_COUNT = 2
)

// Client - 채팅을 이용하는 사용자의 정보
type Client struct {
	connection net.Conn
	read       chan string
	quit       chan int
	name       string
	room       *Room
}

// Room - 채팅방
type Room struct {
	num        int
	clientlist *list.List
}

var roomlist *list.List

func main() {
	roomlist = list.New()
	for i := 0; i < ROOM_MAX_COUNT; i++ {
		room := &Room{i + 1, list.New()}
		roomlist.PushBack(*room)
	}

	chatting, err := net.Listen("tcp", ":5000")
	if err != nil {
		handleErrorServer(nil, err, "채팅 서버를 여는 데 실패하였습니다.")
	}
	defer chatting.Close()

	for {
		// waiting connection
		connection, err := chatting.Accept()
		if err != nil {
			handleErrorServer(connection, err, "사용자가 채팅 서버에 들어오는 데 실패하였습니다.")
		}

		go handleConnection(connection)
	}
}

func handleErrorServer(conn net.Conn, err error, errmsg string) {
	if conn != nil {
		conn.Close()
	}
	fmt.Println(err)
	fmt.Println(errmsg)
}

func handleConnection(connection net.Conn) {
	read := make(chan string)
	quit := make(chan int)
	client := &Client{connection, read, quit, "unknown", &Room{-1, list.New()}}

	go handleClient(client)

	// fmt.Printf("remote Addr = %s\n", connection.RemoteAddr().String())
}

func handleClient(client *Client) {
	for {
		select {
		case msg := <-client.read:
			if strings.HasPrefix(msg, "[G]") {
				sendToAllClients(client.name, msg)
			} else if strings.HasPrefix(msg, "[R]") {
				sendToRoomClients(client.room, client.name, msg)
			} else if strings.HasPrefix(msg, "[W]") {
				sendToClientToClient(client, msg)
			} else {
				sendToClient(client, client.name, msg)
			}

		case <-client.quit:
			fmt.Println("disconnect client")
			client.conn.Close()
			client.deleteFromList()
			return

		default:
			go recvFromClient(client)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func recvFromClient(client *Client) {
	recvmsg, err := bufio.NewReader(client.conn).ReadString('\n')
	if err != nil {
		handleErrorServer(client.conn, err, "string read error..")
		client.quit <- 0
		return
	}

	strmsgs := strings.Split(recvmsg, "|")

	switch strmsgs[0] {
	case LOGIN:
		client.name = strings.TrimSpace(strmsgs[1])

		room := allocateEmptyRoom()
		if room.num < 1 {
			handleErrorServer(client.conn, nil, "max user limit!")
		}
		client.room = room

		if !client.dupUserCheck() {
			handleErrorServer(client.conn, nil, "duplicate user!!"+client.name)
			client.quit <- 0
			return
		}
		fmt.Printf("\nhello = %s, your room number is = %d\n", client.name, client.room.num)
		room.clientlist.PushBack(*client)

	case CHAT:
		fmt.Printf("\nrecv message = %s\n", strmsgs[1])
		client.read <- strmsgs[1]
	}
}

func sendToClient(client *Client, sender string, msg string) {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString(sender)
	buffer.WriteString("] ")
	buffer.WriteString(msg)

	fmt.Printf("client = %s ==> %s", client.name, buffer.String())

	fmt.Fprintf(client.conn, "%s", buffer.String())
}

func sendToAllClients(sender string, msg string) {
	fmt.Printf("global broad cast message = %s", msg)
	for re := roomlist.Front(); re != nil; re = re.Next() {
		r := re.Value.(Room)
		for e := r.clientlist.Front(); e != nil; e = e.Next() {
			c := e.Value.(Client)
			sendToClient(&c, sender, msg)
		}
	}
}

func (client *Client) deleteFromList() {
	for re := roomlist.Front(); re != nil; re = re.Next() {
		r := re.Value.(Room)
		for e := r.clientlist.Front(); e != nil; e = e.Next() {
			c := e.Value.(Client)
			if client.conn == c.conn {
				r.clientlist.Remove(e)
			}
		}
	}
}

func (client *Client) dupUserCheck() bool {
	for re := roomlist.Front(); re != nil; re = re.Next() {
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

func allocateEmptyRoom() *Room {
	for e := roomlist.Front(); e != nil; e = e.Next() {
		r := e.Value.(Room)

		fmt.Printf("clientlist len = %d", r.clientlist.Len())
		if r.clientlist.Len() < ROOM_MAX_USER {
			return &r
		}
	}

	// full room
	return &Room{-1, list.New()}
}

func sendToRoomClients(room *Room, sender string, msg string) {
	fmt.Printf("room broad cast message = %s", msg)
	for e := room.clientlist.Front(); e != nil; e = e.Next() {
		c := e.Value.(Client)
		sendToClient(&c, sender, msg)
	}
}

func findClientByName(name string) *Client {
	for re := roomlist.Front(); re != nil; re = re.Next() {
		r := re.Value.(Room)
		for e := r.clientlist.Front(); e != nil; e = e.Next() {
			c := e.Value.(Client)
			if strings.Compare(c.name, name) == 0 {
				return &c
			}
		}
	}

	return &Client{nil, nil, nil, "unknown", nil}
}

func sendToClientToClient(client *Client, msg string) {
	strmsgs := strings.Split(msg, " ")

	target := findClientByName(strmsgs[1])
	if target.conn == nil {
		fmt.Println("Can't find target User")
		return
	}

	sendToClient(target, client.name, strmsgs[2])
}
