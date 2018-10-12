package network

import (
	"io"
	"fmt"
	"encoding/binary"
	"github.com/gogo/protobuf/proto"
	"slg_game_server/server/game/user"
	proto3 "slg_game_server/proto"
	"net"
)

// 包头4个字节
const (
	CMDSIZE = 4   // 消息ID 2个字节
	BODYSIZE  = 2 // 包体大小 2个字节
	TEMPSIZE = 2  // 预留2个字节
)


// 接收Length-Type-Value格式的封包流程
func RecvLTVPacket(reader net.Conn) (cmd proto3.ProtoCmd, tempId uint16, msg proto.Message, err error) {
	// 8个字节缓冲区
	headBuff := make([]byte, CMDSIZE + BODYSIZE + TEMPSIZE)
	_, err = io.ReadFull(reader, headBuff)
	if err != nil {
		fmt.Println("------------------err ", err)
		return				// 这里并非跳出循环，而是返回值
	}

	// 先读取2字节的消息id
	msgId := int32(binary.LittleEndian.Uint32(headBuff[0:CMDSIZE]))
	// 分配包体缓冲区,再读取2字节的包体长度
	bodyLen := binary.LittleEndian.Uint16(headBuff[CMDSIZE:])
	// 分配包体缓冲区,再读取2字节的预留id
	tempId = binary.LittleEndian.Uint16(headBuff[CMDSIZE + BODYSIZE:])

	// 分配包体大小
	body := make([]byte, bodyLen)
	_, err = io.ReadFull(reader, body)

	if err != nil {
		//reader.Close()		// client disconnect
		fmt.Println("------------------client error ", err)
		return
	}

	// 最终获取消息体内容
	cmd = proto3.ProtoCmd(msgId)
	pbData := user.Handler.GetPbData(cmd)
	msg = reflect.New(reflect.TypeOf(pbData).Elem()).Interface()
	proto.Unmarshal(body, msg.(proto.Message))

	return cmd, tempId, pbData, nil
}

// 发送Length-Type-Value格式的封包流程
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
