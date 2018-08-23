package main

import (
	"net"
	"fmt"
)

var client_num int = 0

func Error_print1(err error) {
	fmt.Println("------------------err ", err)
}

func main() {
	l, err := net.Listen("tcp4", "127.0.0.1:8006")
	if err != nil {
		Error_print1(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			Error_print1(err)
			continue
		}
		client_num++
		fmt.Printf("A new Connection %d.\n", client_num)
		go handlerConnection(conn)
	}
}
func handlerConnection(conn net.Conn) {
	defer closeConnection(conn)
	readChannel := make(chan []byte, 1024)
	writeChannel := make(chan []byte, 1024)
	go readConnection(conn, readChannel)
	go writeConnection(conn, writeChannel)
	for {
		select {
		case data := <-readChannel:
			if string(data) == "bye" {
				return
			}
			writeChannel <- append([]byte("Back"), data...)        }
	}
}
func writeConnection(conn net.Conn, channel chan []byte) {
	for {
		select {
		case data := <-channel:
			println("Write:", conn.RemoteAddr().String(), string(data))
			_, err := conn.Write(data)
			if err != nil {
				Error_print1(err)
				return
			}        }
	}
}
func readConnection(conn net.Conn, channel chan []byte) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Error_print1(err)
			channel <- []byte("bye") //这里须要进一步改进！
			break
		}
		println("Recei:", conn.RemoteAddr().String(), string(buffer[:n]))
		channel <- buffer[:n]
	}
}
func closeConnection(conn net.Conn) {
	conn.Close()
	client_num--
	fmt.Printf("Now, %d connections is alve.\n", client_num)
}
