package xmail

import "testing"

// 测试Send()方法发送邮件
func TestSend(t *testing.T)  {
    if err := Send("770****12@qq.com", "yrg********ddf", "smtp.qq.com:25", "1****48@qq.com;1*****43@qq.com", "邮件标题", "邮件内容", "html"); err != nil {
        t.Error("发送邮件失败：", err.Error())
    } else {
        t.Log("发送邮件成功！")
    }
}
