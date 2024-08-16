// @Title  dispersed
// @Description 生成离散操作相关数据
// @Author  朱宗辉  2024/8/15 22:14
// @Update  朱宗辉  2024/8/15 22:14
package rtu

import "gitee.com/zhu_zonghui/modbus_util/crc"

// DispersedRead 读离散输入寄存器
func DispersedRead(slaveAddr uint8, startAddr uint16, number uint16) []byte {
	var buf = []byte{slaveAddr, 0x02, byte((startAddr >> 8) & 0xFF), byte(startAddr & 0xFF), byte((number >> 8) & 0xFF), byte(number & 0xFF)}
	crc16 := crc.CRC16(buf)
	buf = append(buf, byte(crc16&0xFF), byte((crc16>>8)&0xFF))
	return buf
}
