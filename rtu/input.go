// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  朱宗辉  2024/8/15 22:27
// @Update  朱宗辉  2024/8/15 22:27
package rtu

import "modbus_util/crc"

func InputRead(slaveAddr uint8, startAddr uint16, number uint16) []byte {
	var buf = []byte{slaveAddr, 0x04, byte((startAddr >> 8) & 0xFF), byte(startAddr & 0xFF), byte((number >> 8) & 0xFF), byte(number & 0xFF)}
	crc16 := crc.CRC16(buf)
	buf = append(buf, byte(crc16&0xFF), byte((crc16>>8)&0xFF))
	return buf
}
