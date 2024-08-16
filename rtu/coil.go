// @Title  coil
// @Description 生成线圈操作相关数据
// @Author  朱宗辉  2024/8/15 21:42
// @Update  朱宗辉  2024/8/15 21:42
package rtu

import "gitee.com/zhu_zonghui/modbus_util/crc"

// CoilRead 读取线圈寄存器
func CoilRead(slaveAddr uint8, startAddr uint16, number uint16) []byte {
	var buf = []byte{slaveAddr, 0x01, byte((startAddr >> 8) & 0xFF), byte(startAddr & 0xFF), byte((number >> 8) & 0xFF), byte(number & 0xFF)}
	crc16 := crc.CRC16(buf)
	buf = append(buf, byte(crc16&0xFF), byte((crc16>>8)&0xFF))
	return buf
}

// CoilWriteOne 写单个线圈
func CoilWriteOne(slaveAddr uint8, startAddr uint16, data uint16) []byte {
	var buf = []byte{slaveAddr, 0x05, byte((startAddr >> 8) & 0xFF), byte(startAddr & 0xFF), byte((data >> 8) & 0xFF), byte(data & 0xFF)}
	crc16 := crc.CRC16(buf)
	buf = append(buf, byte(crc16&0xFF), byte((crc16>>8)&0xFF))
	return buf
}

// CoilWriteMore 写多个线圈
func CoilWriteMore(slaveAddr uint8, startAddr uint16, number uint16, data []uint8) []byte {
	var buf = []byte{slaveAddr, 0x0F, byte((startAddr >> 8) & 0xFF), byte(startAddr & 0xFF), byte((number >> 8) & 0xFF), byte(number & 0xFF), byte(len(data))}
	for i := range data {
		buf = append(buf, data[i])
	}
	crc16 := crc.CRC16(buf)
	buf = append(buf, byte(crc16&0xFF), byte((crc16>>8)&0xFF))
	return buf
}
