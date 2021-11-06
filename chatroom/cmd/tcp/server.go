package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

var (
	// 存储加入的用户
	enteringChannel = make(chan *User)
	leavingChannel  = make(chan *User)
	messageChannel  = make(chan *Message, 8)
)

type User struct {
	ID             int
	Addr           string
	EnterAt        time.Time
	MessageChannel chan string
}

func (u *User) String() string {
	return fmt.Sprintf("%d@%s", u.ID, u.Addr)
}

type Message struct {
	OwnerID int
	Content string
}

// broadcaster 用于记录聊天室用户，并进行消息广播和用户处理
func broadcaster() {
	users := make(map[*User]struct{})

	for {
		select {
		case user := <-enteringChannel:
			users[user] = struct{}{}
		case user := <-leavingChannel:
			delete(users, user)
			close(user.MessageChannel)
		case msg := <-messageChannel:
			for user := range users {
				if user.ID == msg.OwnerID {
					continue
				}
				user.MessageChannel <- msg.Content
			}
		}
	}
}

func GenUserID() int {
	return time.Now().Nanosecond()
}

// handleConn 处理链接
func handleConn(conn net.Conn) {
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	// 构建用户实例
	user := &User{
		ID:             GenUserID(),
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now(),
		MessageChannel: make(chan string, 8),
	}

	go sendMessage(conn, user.MessageChannel)

	user.MessageChannel <- "Welcome, " + user.String()

	msg := &Message{
		OwnerID: user.ID,
	}
	msg.Content = "User: " + strconv.Itoa(user.ID) + " has enter."
	messageChannel <- msg
	enteringChannel <- user

	var userActive = make(chan struct{})
	go func() {
		d := 1 * time.Minute
		timer := time.NewTimer(d)
		for {
			select {
			case <-timer.C:
				_ = conn.Close()
			case <-userActive:
				timer.Reset(d)
			}
		}
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg.Content = strconv.Itoa(user.ID) + ": " + input.Text()
		messageChannel <- msg
		userActive <- struct{}{}
	}

	if err := input.Err(); err != nil {
		log.Println("读取错误: ", err)
	}
	leavingChannel <- user
	msg.Content = "user: " + strconv.Itoa(user.ID) + " has left"
	messageChannel <- msg
}

func sendMessage(conn net.Conn, messageChannel <-chan string) {
	for msg := range messageChannel {
		_, _ = fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		panic(err)
	}

	// 启动广播
	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}
