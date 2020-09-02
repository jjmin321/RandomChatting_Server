package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	LOGIN = "1"
	CHAT  = "2"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		handleErrorClient(connection, "채팅 서버에 연결하는 데 실패하였습니다.")
	}

	message := make(chan string)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input your name :")
	name, err := reader.ReadString('\n')
	if err != nil {
		handleErrorClient(connection, "read input failed..")
	}

	// login
	fmt.Fprintf(connection, "%s|%s", LOGIN, name)

	// chat
	go handleRecvMsg(connection, message)
	handleSendMsg(connection)
}

func handleErrorClient(conn net.Conn, errmsg string) {
	if conn != nil {
		conn.Close()
	}
	fmt.Println(errmsg)
}

func handleSendMsg(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send : ")
		text, err := reader.ReadString('\n')
		if err != nil {
			handleErrorClient(conn, "read input failed..")
		}

		fmt.Fprintf(conn, "%s|%s", CHAT, text)
	}
}

func handleRecvMsg(conn net.Conn, msgch chan string) {
	for {
		select {
		case msg := <-msgch:
			fmt.Printf("\nMessage from server : %s\n", msg)
		default:
			go recvFromServer(conn, msgch)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func recvFromServer(conn net.Conn, msgch chan string) {
	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		handleErrorClient(conn, "read msg failed..")
		os.Exit(2)
		return
	}
	msgch <- msg
}
