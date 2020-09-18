package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	login = "1"
	chat  = "2"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		handleErrorClient(connection, "채팅 서버에 연결하는 데 실패하였습니다.")
		os.Exit(2)
		return
	}

	message := make(chan string)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("닉네임을 입력해주세요 :")
	name, err := reader.ReadString('\n')
	if err != nil {
		handleErrorClient(connection, "닉네임을 읽는 데 실패하였습니다.")
	}

	fmt.Fprintf(connection, "%s|%s", LOGIN, name)

	go handleRecvMsg(connection, message)
	handleSendMsg(connection)
}

func handleErrorClient(conn net.Conn, errmsg string) {
	if conn != nil {
		conn.Close()
	}
	log.Println(errmsg)
}

func handleSendMsg(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("채팅을 입력하세요 : ")
		text, err := reader.ReadString('\n')
		if err != nil {
			handleErrorClient(conn, "메세지를 읽는 데 실패하였습니다.")
		}

		fmt.Fprintf(conn, "%s|%s", CHAT, text)
	}
}

func handleRecvMsg(conn net.Conn, msgch chan string) {
	for {
		select {
		case msg := <-msgch:
			fmt.Printf("\n%s\n", msg)
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
