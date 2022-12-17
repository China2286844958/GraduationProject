package tools

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"strings"
	"time"
)

/**
@Title 发送邮箱
@CreateTime 2022年8月14日23:57:49
@Author 薛智敏
*/

//  SendEmail	    发送邮箱
//  @Description:
/*
	先安装gomail包
	go get -v gopkg.in/gomail.v2
*/

//
//  SendEmailOneByOne 一个代理接一个的代理发送邮箱
//  @Description:
//  @param ToEmail 发送的地址
//  @param catcher 发送的数据验证码
//  @return SendStatus 发送状态 True 为成功，false为失败！
//

var SetEmailCodeExpireUnit = time.Second                       //设置邮箱验证码过期的单位:(m:分钟;s:秒;)
var SetEmailCodeExpireNumber = SetEmailCodeExpireUnit * 5 * 60 //设置邮箱验证码过期时间

// TimeToString 时间转化能看懂的单位
//
//	@Description:
//	@return string
func TimeToString() string {
	t := ""
	str := SetEmailCodeExpireNumber.String()
	t = strings.Replace(str, "d", "天", -1) //替换d
	t = strings.Replace(t, "h", "时", -1)   //替换h
	t = strings.Replace(t, "m", "分", -1)   //替换m
	t = strings.Replace(t, "s", "秒", -1)   //替换m
	return t
}

//
//  GetEmailCodeExpireStr 获取邮箱验证码过期单位
//  @Description:

func GetEmailCodeExpireStr() string {
	return SetEmailCodeExpireNumber.String()

}

const (
	MessageFindPasswd = `
    <p>您好,尊敬的用户%s，您正在申请找回密码的操作!</p>
		验证码为:<h1><strong style='color:red'>%s</strong></h1>有效时间%s,如果不是本人操作，请勿泄露他人使用！
	</p>
	`
	MessageRegister = `
    <p>您好,尊敬的用户%s,您正在注册账号的操作!</p>
		验证码为:<h1><strong style='color:red'>%s</strong></h1>有效时间%s,如果不是本人操作，请勿泄露他人使用！
	</p>
	`
)

//获取发送的验证码信息模版
// SendEmailOneByOne
//  @Description:
//  @param ToEmail 发送的邮箱
//  @param catcher 验证码
//  @param sendMessage 信息模版
//  @return SendStatus
//

func SendEmailOneByOne(ToEmail string, catcher string, sendMessage string) (SendStatus bool) {

	// QQ 邮箱：
	// SMTP 服务器地址：smtp.qq.com（SSL协议端口：465/994 | 非SSL协议端口：25）
	// 163 邮箱：
	// SMTP 服务器地址：smtp.163.com（端口：25）
	host := "smtp.qq.com"
	port := 25

	//标题的内容取决于发送的文本内容
	title := "学生选课系统----"
	if sendMessage == MessageRegister {
		title = "学生选课系统-注册验证码"
	}
	if sendMessage == MessageFindPasswd {
		title = "学生选课系统-找回密码"

	}

	sendEmailManagerList := getSendEmailManagerList()

	for k, v := range sendEmailManagerList {

		m := gomail.NewMessage()

		m.SetHeader("From", v.UserName) // 发件人

		m.SetHeader("To", ToEmail) // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
		m.SetHeader("Cc", ToEmail) // 抄送，可以多个
		//m.SetHeader("Bcc", ToEmail)         // 暗送，可以多个
		m.SetHeader("Subject", title) // 邮件主题

		// text/html 的意思是将文件的 content-type 设置为 text/html 的形式，浏览器在获取到这种文件时会自动调用html的解析器对文件进行相应的处理。
		// 可以通过 text/html 处理文本格式进行特殊处理，如换行、缩进、加粗等等

		m.SetBody("text/html", fmt.Sprintf(sendMessage, ToEmail, catcher, TimeToString()))

		d := gomail.NewDialer(
			host,
			port,
			v.UserName,
			v.PassWord,
		)
		// 关闭SSL协议认证
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		if err := d.DialAndSend(m); err != nil { //有错误

			log.Printf("\n\t【邮箱发送】发件人:\t%s\t发送信息:%s\t收件人:%s\t\t捕捉的错误:%s\n", v.UserName, catcher, ToEmail, err.Error())

			//当最后一个发送管理员也失败，就说明填写信息有误
			if k+1 == len(sendEmailManagerList) {
				log.Println("邮件发送都失败")
				return false
			}

		} else { //没错误

			return true
		}
	}
	return true
}

//发送邮箱的管理员-代理

type SendEmailManager struct {
	UserName string //QQ邮箱
	PassWord string //邮箱校验码
}

//
//  getSendEmailManagerList 获取发送邮箱管理员的集合
//  @Description:
//  @return []SendEmailManager 发送邮箱的管理员们
//

func getSendEmailManagerList() []SendEmailManager {

	var managerList = []SendEmailManager{
		{"261520734@qq.com", "imnbfxwduisucajg"},
		{"2286844958@qq.com", "gvapjrsplowkdibb"},
	}

	return managerList
}
