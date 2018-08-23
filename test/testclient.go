package main

import (
	"net"
	"fmt"
	proto3 "slg_game_server/proto"
	"slg_game_server/server/handler"
)

func Error_print(err error) {
	fmt.Println("------------------err ", err)
}

func main() {
	tcpAddress, err := net.ResolveTCPAddr("tcp4", "192.168.60.130:8004")
	if err != nil {
		Error_print(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddress)
	if err != nil {
		Error_print(err)
	}
	writeChan := make(chan []byte, 1024)
	readChan := make(chan []byte, 1024)
	go writeConnection1(conn, writeChan)
	go readConnection1(conn, readChan)
	//go handleReadChannel(readChan)
	for {
		var s string
		fmt.Scan(&s)
		writeChan <- []byte(s)
	}
}
func readConnection1(conn *net.TCPConn, channel chan []byte) {
	defer conn.Close()
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Error_print(err)
			return
		}
		println("Received from:", conn.RemoteAddr(), string(buffer[:n])) //
		channel <- buffer[:n]
	}
}

func writeConnection1(conn *net.TCPConn, channel chan []byte) {
	defer conn.Close()
	for {
		select {
		case data := <-channel:
			switch data {
			case "login":
				
			
			}
			cmd := int32(proto3.CMD_CMD_LoginReq)
			pbData := &proto3.LoginReq{}
			pbData.GameId = 666
			pbData.Token = "hb"
			handler.SendLTVPacket(conn, cmd, pbData)
			_, err := conn.Write(data)
			if err != nil {
				Error_print(err)
			}
			println("Write to:", conn.RemoteAddr(), string(data))        }
	}
}
