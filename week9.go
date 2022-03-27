package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

/*
总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

LineBasedFrameDecoder
遍历ByteBuf中的可读字节，判断看是否有“\n”或者“\r\n”， 如果有，就以此位置为结束位置

DelimiterBasedFrameDecoder
用特殊符号作为消息的结束符

LengthFieldPrepender（编码）和 LengthFieldBasedFrameDecoder（解码）
将消息分为消息头 + 消息体，在消息头中保存消息体的长度，发送者通过LengthFieldPrepender去编码，接收者通过LengthFieldBasedFrameDecoder去解码。
*/

func Decode(r io.Reader) ([]byte, error) {
	var totalLen int32
	err := binary.Read(r, binary.BigEndian, &totalLen)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, totalLen)
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

type Submit struct {
	PackageLength   []byte
	HeaderLength    []byte
	ProtocolVersion []byte
	Operation       []byte
	Sequence        []byte
	body            []byte
}

func (s *Submit) Decode(pktBody []byte) error {
	s.PackageLength = pktBody[:4]
	s.HeaderLength = pktBody[4:6]
	s.ProtocolVersion = pktBody[6:8]
	s.Operation = pktBody[8:12]
	s.Sequence = pktBody[12:16]
	s.body = pktBody[16:]
	return nil
}

func handlePacket(request []byte) {
	s := Submit{}
	goim := s.Decode(request)
	fmt.Println("get goim", goim)
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		request, err := Decode(c)
		if err != nil {
			fmt.Println("handleConn: request decode error:", err)
			return
		}

		handlePacket(request)
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		go handleConn(c)
	}
}
