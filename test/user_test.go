package test

import (
	"cmdb/dao"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) { //在Test开始之前先运行，调用run方法才会开始测试
	fmt.Println("测试userdao")
	m.Run()
}
func TestUser(t *testing.T) {
	fmt.Println("测试userdao中的参数")
	t.Run("保存用户", testSave)
}

// 测试能否保存用户密码
func testSave(t *testing.T) {
	dao.SaveUser("admin", "123", "123@qq.com")
}
