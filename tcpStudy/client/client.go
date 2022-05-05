package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	dial, err := net.Dial("tcp", "127.0.0.1:6789")
	if err != nil {
		fmt.Println("client connect error,", err)
	}
	channelForRoutine := make(chan bool, 0)
	var channelForRoutineInput chan<- bool = channelForRoutine
	go inputFromConsole(dial, &channelForRoutineInput)
	needClose := <-channelForRoutine
	if needClose {
		fmt.Println(needClose)
		return
	}
}

func inputFromConsole(conn net.Conn, channelForRoutine *chan<- bool) {
	defer func() {
		*channelForRoutine <- true
		err := conn.Close()
		if err != nil {
			fmt.Println("client connection close error,", err)
		}
	}()
	for true {
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("console read error,", err)
			return
		}
		if string(line) == "quit" {
			return
		}
		_, err = conn.Write(line)
		if err != nil {
			fmt.Println("connection write error,", err)
			return
		}
	}
}
