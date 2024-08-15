package test

import (
	"net"
	"testing"
)

// 测试函数
func TestTcpServer(t *testing.T) {
	listen, err := net.Listen("tcp", ":1122")
	if err != nil {
		t.Error("监听失败", err.Error())
		return
	}
	defer listen.Close()
	conn, err := listen.Accept()
	if err != nil {
		t.Error("链接客户端失败", err.Error())
		return
	}
	defer conn.Close()
	var buf = []byte{0x01, 0x03, 0x00, 0x00, 0x00, 0x0A, 0xC5, 0xCD}
	_, err = conn.Write(buf)
	if err != nil {
		t.Error("发送消息失败", err.Error())
		return
	}
	t.Log("发送消息成功")
	var readBuf [128]byte
	n, err := conn.Read(readBuf[:])
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("收到消息", readBuf[:n])
}

// 压力测试
func BenchmarkTcpServer(b *testing.B) {
	listen, err := net.Listen("tcp", ":1122")
	if err != nil {
		b.Error("监听失败", err.Error())
		return
	}
	defer listen.Close()
	conn, err := listen.Accept()
	if err != nil {
		b.Error("链接客户端失败", err.Error())
		return
	}
	defer conn.Close()
	var buf = []byte{0x01, 0x03, 0x00, 0x00, 0x00, 0x0A, 0xC5, 0xCD}
	_, err = conn.Write(buf)
	if err != nil {
		b.Error("发送消息失败", err.Error())
		return
	}
	b.Log("发送消息成功")
}
