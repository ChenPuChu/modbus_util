// @Title  CRC_test
// @Description 测试CEC加密加密相关功能
// @Author  朱宗辉  2024/8/15 21:51
// @Update  朱宗辉  2024/8/15 21:51
package test

import (
	"fmt"
	"modbus_util/crc"
	"testing"
)

func TestCRCModbus(t *testing.T) {
	var buf = []byte{0x01, 0x03, 0x00, 0x00, 0x00, 0x01}
	crc16 := crc.CRC16(buf)
	// 将crc16以16进制方式输出
	fmt.Println(crc16)
}
