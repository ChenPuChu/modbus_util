// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  朱宗辉  2024/8/15 22:39
// @Update  朱宗辉  2024/8/15 22:39
package rtu

import (
	"fmt"
	"net"
)

func Listen(address string, port int) (net.Listener, error) {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	return listen, err
}

func RtuWrite(conn *net.Conn, data *[]byte) ([]byte, error) {
	_, err := (*conn).Write(*data)
	if err != nil {
		fmt.Errorf("RTU发送数据失败，err:%s", err.Error())
		return nil, err
	}
	var readBuf [128]byte
	n, err := (*conn).Read(readBuf[:])
	if err != nil {
		fmt.Errorf("RTU接收数据失败，err:%s", err.Error())
		return nil, err
	}
	return readBuf[:n], nil
}
