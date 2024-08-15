// @Title  RTU_test
// @Description 测试RTU相关功能
// @Author  朱宗辉  2024/8/15 21:47
// @Update  朱宗辉  2024/8/15 21:47
package test

import (
	"modbus_util/rtu"
	"testing"
)

func TestCoilRead(t *testing.T) {
	coilRead := rtu.CoilRead(0x01, 0x1113, 0x0025)
	t.Logf("%X", coilRead)
}

func TestDispersedRead(t *testing.T) {
	dispersedRead := rtu.DispersedRead(0x01, 0x00C4, 0x0016)
	t.Logf("%X", dispersedRead)
}

func TestMaintainRead(t *testing.T) {
	maintainRead := rtu.MaintainRead(0x01, 0x006B, 0x0003)
	t.Logf("%X", maintainRead)
}

func TestInputRead(t *testing.T) {
	inputRead := rtu.InputRead(0x01, 0x0008, 0x0002)
	t.Logf("%X", inputRead)
}

func TestCoilWriteOne(t *testing.T) {
	coilWriteOne := rtu.CoilWriteOne(0x01, 0x00Ac, 0xFF00)
	t.Logf("%X", coilWriteOne)
}

func TestMaintainWriteOne(t *testing.T) {
	maintainWriteOne := rtu.MaintainWriteOne(0x01, 0x0000, 0x0001)
	t.Logf("%X", maintainWriteOne)
}

func TestCoilWriteMore(t *testing.T) {
	coilWriteMore := rtu.CoilWriteMore(0x01, 0x0013, 0x000A, []uint8{0xCD, 0x01})
	t.Logf("%X", coilWriteMore)
}

func TestMaintainWriteMore(t *testing.T) {
	maintainWriteMore := rtu.MaintainWriteMore(0x01, 0x0001, 0x0002, []uint16{0x000A, 0x0102})
	t.Logf("%X", maintainWriteMore)
}

func TestServer(t *testing.T) {
	var data = [][]byte{
		rtu.CoilRead(0x01, 0x1113, 0x0025),
		rtu.MaintainRead(0x01, 0x006B, 0x0003),
	}
	listen, err := rtu.Listen("", 1122)
	if err != nil {
		t.Fatal(err)
		return
	}
	conn, err := listen.Accept()
	if err != nil {
		t.Fatal(err)
		return
	}
	defer conn.Close()
	for i := range data {
		readBuf, err := rtu.RtuWrite(&conn, &data[i])
		if err != nil {
			t.Fatal(err)
			return
		}
		t.Logf("发送%X\n接收%X\n", data[i], readBuf)
	}
}
