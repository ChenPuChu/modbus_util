// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  朱宗辉  2024/8/15 22:23
// @Update  朱宗辉  2024/8/15 22:23
package test

import (
	"modbus_util/rtu"
	"testing"
)

func TestMaintainWriteMore(t *testing.T) {
	maintainWriteMore := rtu.MaintainWriteMore(0x01, 0x0001, 0x0002, []uint16{0x000a, 0x0102})
	t.Log(maintainWriteMore)
}
