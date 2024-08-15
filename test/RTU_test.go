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
	read := rtu.CoilRead(0x01, 0x1113, 0x0025)
	t.Log(read)
}

func TestCoilWriteMore(t *testing.T) {
	coilWriteMore := rtu.CoilWriteMore(0x01, 0x0013, 0x000A, []uint8{0xCD, 0x01})
	t.Log(coilWriteMore)

}
