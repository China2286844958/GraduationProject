package controller

import (
	"GraduationProject/dao"
	"GraduationProject/tools"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

/**
@Title 发送验证码
@author 薛智敏
@CreateTime 2022年10月8日17:45:42
*/

//
// SendEmailByNetwork 通过网络发送邮箱验证码
//  @Description:
//  @param r
//  @return *gin.Engine
//

var sendEmailIndex = "/SendEmailIndex"

func SendEmailByNetwork(r *gin.Engine) *gin.Engine {

	r.GET(RouterUserGroup+sendEmailIndex+"/:v/:token", func(c *gin.Context) {

		v := c.Param("v")
		token := c.Param("token") //获取token

		captchaByStr := tools.CaptchaByTimes(5) //生成验证码
		c.HTML(http.StatusOK, "register.tmpl", gin.H{
			"captchaCodes": captchaByStr,
			"email":        v,
			"massageErr":   "验证码已发送至" + v + "邮箱中",
			"token":        token, //将token传入
		})
	})

	r.POST(RouterUserGroup+sendEmailIndex+"/:email", func(c *gin.Context) {
		email_ := c.Param("email")
		email := strings.ToLower(email_) //转化小写后的邮箱
		//当邮箱不符合邮箱正则表达式
		if !tools.EmailRegExp(email) {
			c.JSON(500, gin.H{
				"消息": "邮箱错误！",
				"状态": tools.EmailRegExp(email),
			})
			return
		}
		//查找该邮箱是否注册过
		exist := dao.QueryStudLoginByEmail(email)
		if exist {
			c.HTML(http.StatusOK, "register.tmpl", gin.H{
				"massageErr": "邮箱已注册！",
			})
			return
		}

		//执行发送验证码的函数
		captchaByStr := tools.CaptchaByTimes(5) //生成邮箱验证码
		//fmt.Println("60+行:验证码:\t", captchaByStr)
		log.Println("邮箱验证码:\t", captchaByStr)
		status := tools.SendEmailOneByOne(email, captchaByStr, tools.MessageRegister) //执行STMP协议的发送验证码的功能
		//status := true:测试专用

		//当邮箱发送失败
		if !status {
			log.Println("发送失败,邮箱不存在或者网络连接有问题")
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "邮箱发送失败,请检查网络连接、邮箱是否存在或稍后重试", //显示的标题
				"url":     "back",                       //按钮按下后，跳转的页面url,back代表返回上一历史
				"BtnName": "返回",                         //按钮显示的名字
			})
			return
		}

		//邮箱发送成功后生成Token
		users := tools.Users{}
		users.ExeEmail = email
		users.LoginStatus = false
		users.SendEmailStatus = true
		users.Captcha = captchaByStr

		token, err := tools.GenToken(users, time.Minute*5)
		if err != nil {
			c.HTML(http.StatusOK, "register.tmpl", gin.H{
				"massageErr": "发送邮箱验证码生成Token错误!",
			})
			return
		}

		//重定向get请求
		c.Redirect(http.StatusMovedPermanently, RouterUserGroup+sendEmailIndex+"/"+email+"/"+token)

	})

	return r
}
