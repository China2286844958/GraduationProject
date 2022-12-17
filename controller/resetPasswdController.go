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
@Title 找回密码-重置密码的业务逻辑
@CreateTime 2022年8月27日21:26:10
@Author 薛智敏
*/

func FindPasswd(r *gin.Engine) *gin.Engine {

	var findPasswdIndex = "/findPasswdIndex" //找回密码的主页面
	var resetPasswd = "/resetPasswd"         //重置密码界面
	//首次找回密码的页面
	r.GET(RouterUserGroup+findPasswdIndex, func(c *gin.Context) {
		c.HTML(http.StatusOK, "findPasswd.tmpl", gin.H{})
	})

	//找回密码带邮箱的页面
	r.GET(RouterUserGroup+findPasswdIndex+"/:emailIn", func(c *gin.Context) {

		//获取JS传来的参数:emailIn
		emailIn := c.Param("emailIn")
		emailIn = strings.ToLower(emailIn) //转小写

		//当emailIn不为空，校验邮箱格式
		if emailIn != "" {
			emailRegExp := tools.EmailRegExp(emailIn) //校验邮箱格式
			//当用户的邮箱格式错误直接返回一个错误
			if !emailRegExp {

				c.HTML(http.StatusOK, "findPasswd.tmpl", gin.H{
					//用户输入的值  错误信息传入前端
					"emailValue": emailIn,
					"ShowLog":    "邮箱格式错误!",
				})
				return

			} else { //符合邮箱格式的业务逻辑
				//查询数据库是否存在该用户,不存在停止提示，用户不存在，返回

				exist := dao.QueryStudLoginByEmail(emailIn) //执行数据库，先查看该用户是否存在

				if !exist {
					c.HTML(http.StatusOK, "findPasswd.tmpl", gin.H{
						//用户输入的值  错误信息传入前端
						"emailValue": emailIn,
						"ShowLog":    emailIn + "该邮箱未注册!",
					})
					return
				}

				//执行发送验证码的函数

				captchaByStr := tools.CaptchaByTimes(4) //生成验证码
				log.Println("60行:验证码:\n", captchaByStr)
				status := tools.SendEmailOneByOne(emailIn, captchaByStr, tools.MessageFindPasswd) //执行STMP协议的发送验证码的功能

				//====测试=====
				//status := true
				//log.Println(tools.SetEmailCodeExpireNumber)

				//当邮箱发送失败
				if !status {
					c.HTML(http.StatusOK, "findPasswd.tmpl", gin.H{
						//用户输入的值  错误信息传入前端
						"emailValue": emailIn,
						"ShowLog":    emailIn + "邮箱发送失败,请检查网络连接或稍后重试!",
					})
					return
				}

				//当用户存在，就生成Token,用户不存在返回

				UserToken := tools.Users{}                                                //创建Token结构体
				UserToken.ExeEmail = emailIn                                              //邮箱注册Token
				UserToken.ExpirationTime = time.Now().Add(tools.SetEmailCodeExpireNumber) //过期时间
				UserToken.Captcha = captchaByStr                                          //验证码

				//生成Token，传入另一个链接解析,当被解析、重置密码后，删除Token
				genToken, _ := tools.GenToken(UserToken, tools.SetEmailCodeExpireNumber) //创建Token

				c.HTML(http.StatusOK, "findPasswd.tmpl", gin.H{
					//用户输入的值  错误信息传入前端
					"emailValue": emailIn,
					"ShowLog":    "验证码已经发送到" + emailIn + "邮箱中，注意查收！",
					"NextStep":   RouterUserGroup + resetPasswd + "/" + emailIn + "/" + genToken, //下一步的链接
					"token":      genToken,                                                       //生成的Token
				})

			}

		}

	})

	//重置密码的业务逻辑
	r.POST(RouterUserGroup+resetPasswd+"/:emailValue/:token", func(c *gin.Context) {
		emailValue := c.Param("emailValue")      //邮箱
		emailValue = strings.ToLower(emailValue) //变成小写
		token := c.Param("token")                //token

		//解析Token
		DeToken, err := tools.ParseToken(token)
		if err != nil {
			c.HTML(http.StatusOK, "findPasswd.tmpl", gin.H{
				//用户输入的值  错误信息传入前端
				"emailValue": emailValue,
				"ShowLog":    "该验证码不存在或者已经失效!",
			})
			return
		}

		//对比校验解析后的Token值:对邮箱校验；过期时间;以及前端验证码;验证码进行校验
		recdCaptcha := c.PostForm("recdCaptcha") //从前端form表单获取的验证码

		//===============================解析的Token进行对比校验=======================
		//当解析的Token用户名不吻合
		if DeToken.ExeEmail != emailValue {
			c.HTML(http.StatusOK, "findPasswd.tmpl", gin.H{
				//用户输入的值  错误信息传入前端
				"emailValue": emailValue,
				"ShowLog":    "该信息有误,邮箱错误!",
			})
			return
		}

		//当解析的验证码Token不吻合
		//都转小写
		if strings.ToLower(DeToken.Captcha) != strings.ToLower(recdCaptcha) {
			c.HTML(http.StatusOK, "findPasswd.tmpl", gin.H{
				//用户输入的值  错误信息传入前端
				"emailValue": emailValue,
				"ShowLog":    "该信息有误,验证码不正确!",
			})
			return
		}

		//当解析的Token超过时间
		if !(time.Now().Before(DeToken.ExpirationTime)) {
			c.HTML(http.StatusOK, "findPasswd.tmpl", gin.H{
				//用户输入的值  错误信息传入前端
				"emailValue": emailValue,
				"ShowLog":    "验证码已经过期!",
			})
			return
		}
		//========================================================END==================================

		//执行重置密码的业务逻辑
		//拿到前端的密码
		NewPasswdStr := c.PostForm("NewPasswd2")

		//对密码进行加密
		sha256HashCode := tools.GetSHA256HashCode(NewPasswdStr)

		//数据库更新
		logs := dao.UpdatePasswdByEmail(emailValue, sha256HashCode)
		if !logs {
			c.HTML(http.StatusOK, "findPasswd.tmpl", gin.H{
				//用户输入的值  错误信息传入前端
				"emailValue": emailValue,
				"ShowLog":    "更新失败，请尝试更换密码或者稍后再试!",
			})
			return
		}

		//	提示修改成功
		c.HTML(http.StatusOK, "Message.tmpl", gin.H{
			"title":   "修改成功",        //显示的标题
			"url":     "/User/login", //按钮按下后，跳转的页面url
			"BtnName": "赶紧登录",        //按钮显示的名字
		})
		return

	})

	r.POST(RouterUserGroup+resetPasswd, func(c *gin.Context) {

		tools.SetMassage(c, "页面找不到哦! ", "back", "返回")
	})

	return r
}
