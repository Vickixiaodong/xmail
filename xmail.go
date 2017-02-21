package xmail

import (
    "strings"
    "net/smtp"
)

//描述：发送电子邮件方法
//参数：
//  fromUser：发送者邮箱
//  password：发送者密码
//  host：邮件服务器地址：以qq邮箱为例，需要开启POP3/SMTP协议
//  toUsers：收件人列表
//  title：邮件标题
//  content：邮件内容
//  mailType：邮件类型
func Send(fromUser, password, host, toUsers, title, content, mailType string) error {
    hp := strings.Split(host, ":")
    auth := smtp.PlainAuth("", fromUser, password, hp[0])
    var contentType string
    if mailType == "html" {
        contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
    } else {
        contentType = "Content-Type: text/plain" + "; charset=UTF-8"
    }
    msg := []byte("To: " + toUsers + "\r\nFrom: " + fromUser + "<" + fromUser + ">\r\nSubject: " + title + "\r\n" + contentType + "\r\n\r\n" + content)
    toUsersList := strings.Split(toUsers, ";")
    err := smtp.SendMail(host, auth, fromUser, toUsersList, msg)
    return err
}
