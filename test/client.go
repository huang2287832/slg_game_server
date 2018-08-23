package main

import (
	"net"
	"fmt"
	"github.com/gogo/protobuf/proto"
	proto3 "slg_game_server/proto"
	"encoding/binary"
	"time"
)

// 包头4个字节
const (
	CMDSIZE = 4   // 消息ID 2个字节
	BODYSIZE  = 2 // 包体大小 2个字节
	TEMPSIZE = 2  // 预留2个字节
)


func main() {
	conn, err := net.Dial("tcp", "192.168.60.130:8005")
	fmt.Println("----------- client start !!!")
	if err != nil{
		fmt.Println("Error dialing", err.Error())
		return
	}

	cmd := proto3.ProtoCmd_CMD_LoginReq
	pbData := &proto3.LoginReq{}
	pbData.AcctName = "hb13"

	SendLTVPacket(conn, cmd, pbData)

	cmd = proto3.ProtoCmd_CMD_ShowRankReq
	pbData1 := &proto3.ShowRankReq{Key:proto3.RankEnum_Rank_Power}
	SendLTVPacket(conn, cmd, pbData1)

	time.Sleep(time.Second*1500)
	////
	//SendLTVPacket(conn, cmd, pbData)
	//
	//time.Sleep(time.Second*3)
	//
	//SendLTVPacket(conn, cmd, pbData)
	//
	//time.Sleep(time.Second*3)
	//
	//SendLTVPacket(conn, cmd, pbData)
	//
	//time.Sleep(time.Second*3)
	//
	//SendLTVPacket(conn, cmd, pbData)
	//
	//cmd1 := int32(proto3.ProtoCmd_CMD_LogoutReq)
	//pbData1 := &proto3.LogOutReq{}
	//
	//SendLTVPacket(conn, cmd1, pbData1)



	//for  {
	//	SendLTVPacket(conn, cmd, pbData)
	//}
	//
	//
	//time.Sleep(time.Second*2)
	//
	//SendLTVPacket(conn, cmd, pbData)


}

func SendLTVPacket(writer net.Conn, cmd proto3.ProtoCmd, pbData proto.Message) error {
	msgData, _ := proto.Marshal(pbData)
	pkt := make([]byte, CMDSIZE + BODYSIZE + TEMPSIZE + len(msgData))

	binary.LittleEndian.PutUint32(pkt, uint32(cmd))
	binary.LittleEndian.PutUint16(pkt[CMDSIZE:], uint16(len(msgData)))
	binary.LittleEndian.PutUint16(pkt[CMDSIZE + BODYSIZE:], uint16(255))

	copy(pkt[CMDSIZE + BODYSIZE + TEMPSIZE:], msgData)
	writer.Write(pkt)
	fmt.Println("send success")

	return nil
}