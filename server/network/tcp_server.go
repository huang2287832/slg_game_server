package network

import (
	"fmt"
	"net"
	"runtime/debug"
	"slg_game_server/server/db"
	"slg_game_server/server/game/rank"
	proto3 "slg_game_server/proto"
	"slg_game_server/server/game/user"
	"slg_game_server/server/global"
)

var Server ServerType

type ServerType struct {
	command     bool
	clientNum 	int
	Done chan struct{}
}

type client struct {
	readClose	bool
	login 		int
	socket		net.Conn
	userPid 	chan interface{}
}

func init()  {
	// create listener
	listener, err := net.Listen("tcp", "127.0.0.1:8005")
	defer listener.Close()
	db.NewIncrease()
	global.NewGlobalPlayers()
	rank.NewRank()

	if err != nil{
		fmt.Println("Error listening", err.Error())
		return
	}

	//listen and receive client's connection
	go func() {
		for {
			conn, err := listener.Accept()
			fmt.Println("------------------ first accept", conn)
			if err != nil {
				fmt.Println("Error accepting", err.Error())
				continue
			}
			Server.clientNum += 1
			client := &client{socket:conn}
			client.handleServer(conn)		// 是否需要开协程
			//conn.SetReadDeadline(time.Now().Add(time.Duration(time.Second * 3)))		// TODO need heartbeat && player timer
		}
	}()

	Server.Done = make(chan struct{})
	if !Server.command {
		Command()
	}

	<- Server.Done
	fmt.Println("close server !!!")

}

func (c *client)handleServer(conn net.Conn) {
	writeChan := make(chan interface{}, 1024)
	fmt.Println("------------------ handleServer")
	go c.readServer(conn, writeChan)
	go c.writeServer(conn, writeChan)		// 是否创建玩家成功后才开这个协程

}

// RecvLTVPacket 阻塞式读，所以不会一直死循环,有数据才循环
// 通过read的标志位来退出阻塞式循环
func (c *client) readServer(conn net.Conn, writeChan chan interface{}) {
	for {
		// TODO try-catch
		defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
			fmt.Println("c")
			if err := recover(); err!=nil {
				fmt.Println("done^^^&&&%%$$$$########### err = ", err) // 这里的err其实就是panic传入的内容，55
				debug.PrintStack()
			}
		}()

		//if c.readClose {	// 这个分支永远进不来，因为RecvLTVPacket是在阻塞式接收
		//	return
		//}


		cmd, _, pbData, err := RecvLTVPacket(conn)
		if err != nil{
			if c.readClose {	// 服务器踢人最后一步就是conn关闭socket后，通知读协程退出
				fmt.Println("--------------- kickplayer close = ")
			}else {				// 客户端主动断开socket
				fmt.Println("--------------- client disconnect = ")
				c.close()
			}
			return
		}

		if pbData == nil {
			c.close()
			return
		}

		if c.login == 0 {				// first login
			fmt.Println("--------c.login")
			player := user.NewPlayer()
			player.SetWriteChan(writeChan)
			c.login = 1
			c.userPid = player.PidChan
		}
		// 异步向玩家协程发送消息
		c.userPid <- &user.Message{Cmd:cmd, PbData:pbData}
	}
}

// writeChan 阻塞式读，所以不会一直死循环,有数据才循环
func (c *client) writeServer(conn net.Conn, writeChan chan interface{}) {
	for {
		msgData, ok := <- writeChan
		if !ok {
			fmt.Println("--------------- writeChan close = ", msgData)
			c.readClose = true
			conn.Close()	// 通知read协程退出
			return
		}
		if msgData != nil {
			switch msg := msgData.(type) {
			case *user.Message:
				SendLTVPacket(conn, msg.Cmd, msg.PbData)
			default:

			}
		}
		fmt.Println("write")
	}
}

// 没有数据过来就会一直打印default
//func (c *client) writeServer(conn net.Conn, writeChan chan interface{}) {
//	for {
//		select {
//		case msgData, ok := <- writeChan:
//			if !ok {
//				fmt.Println("--------------- writeChan close = ", msgData)
//				c.readClose = true
//				return
//			}
//			if msgData != nil {
//				switch msg := msgData.(type) {
//				case *user.Message:
//					SendLTVPacket(conn, int32(msg.Cmd), msg.PbData)
//				default:
//
//				}
//			}
//		default:
//			fmt.Println("wirite", c.readClose)
//
//		}
//	}
//}

func (c *client) close() {
	c.userPid <- &user.Message{Cmd:1003, PbData:&proto3.LogOutReq{}}		// TODO call
	fmt.Println("client call close!!! ")
	<- c.userPid
	c.socket.Close()
	Server.clientNum -= 1
}
