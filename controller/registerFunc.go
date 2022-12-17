package controller

import (
	"GraduationProject/dao"
	"GraduationProject/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/**
@Title 注册界面功能以及逻辑
@Date 2022年5月11日20:20:33
*/

var registerWarming = ""                          //错误警告
var PasswdSizeMax = 30                            //密码最多位
var PasswdSizeMin = 4                             //密码最小位
var UserRegisterGroupPathMaster = RouterUserGroup //用户主路由名字
var UserRegisterGroupPathRegister = "/register"   //注册
var UserGroupPathRegisterErr = "/registerErr"     //注册失败标记

func UserRegisterController(r *gin.Engine) *gin.Engine {

	//===================================路由组名=======================
	//----------------用户组----------------
	//====================================用户组========================

	UserGroup := r.Group(UserRegisterGroupPathMaster)
	{

		//------------响应用户获取的请求
		UserGroup.GET(UserRegisterGroupPathRegister, func(c *gin.Context) {
			captcha := tools.CaptchaByTimes(5)

			c.HTML(http.StatusOK, "register.tmpl", gin.H{
				"captcha": captcha, //验证码
			})
		})

		//------------处理用户提交表单
		UserGroup.POST(UserRegisterGroupPathRegister, func(c *gin.Context) {

			//通过name获取表单数据
			email := c.PostForm("email")               //邮箱
			token := c.PostForm("token")               //邮箱token
			passwd1 := c.PostForm("passwd1")           //第一次密码
			passwd2 := c.PostForm("passwd2")           //第二次密码
			captcha := c.PostForm("captcha")           //验证码
			captchaIn := c.PostForm("captchaIn")       //用户输入的验证码
			emailChapter := c.PostForm("emailChapter") //用户输入的邮箱验证码

			if passwd1 != passwd2 {
				c.Request.URL.Path = UserRegisterGroupPathRegister

				registerWarming = "两次密码不一致!请重新输入"
				c.HTML(http.StatusOK, "Message.tmpl", gin.H{
					"title":   registerWarming, //显示的标题
					"url":     "back",          //按钮按下后，跳转的页面url
					"BtnName": "返回",            //按钮显示的名字
				})
				return
			} else { //密码一致
				//邮箱正则表达式
				exp := tools.EmailRegExp(email)

				if !exp {
					registerWarming = "邮箱格式错误！"
					c.HTML(http.StatusOK, "Message.tmpl", gin.H{
						"title":   registerWarming, //显示的标题
						"url":     "back",          //按钮按下后，跳转的页面url
						"BtnName": "返回",            //按钮显示的名字
					})
				} else {

					//校验密码是否为指定尺寸
					size := len(passwd2) //获取密码的位数

					if size < PasswdSizeMin || size > PasswdSizeMax {

						registerWarming = fmt.Sprintf("密码不能低于%d位,且不能大于%d位！", PasswdSizeMin, PasswdSizeMax)
						c.HTML(http.StatusOK, "Message.tmpl", gin.H{
							"title":   registerWarming, //显示的标题
							"url":     "back",          //按钮按下后，跳转的页面url
							"BtnName": "返回",            //按钮显示的名字
						})
						return
					} else if !tools.ExprBlank(passwd2) {
						registerWarming = "密码不能为空字符！"
						c.HTML(http.StatusOK, "Message.tmpl", gin.H{
							"title":   registerWarming, //显示的标题
							"url":     "back",          //按钮按下后，跳转的页面url
							"BtnName": "返回",            //按钮显示的名字
						})
					} else {

						//校验验证码是否正确,不区分大小写
						if strings.ToUpper(captchaIn) != strings.ToUpper(captcha) || strings.ToLower(captchaIn) != strings.ToLower(captcha) {
							registerWarming = "图形验证码错误！"
							c.HTML(http.StatusOK, "Message.tmpl", gin.H{
								"title":   registerWarming, //显示的标题
								"url":     "back",          //按钮按下后，跳转的页面url
								"BtnName": "返回",            //按钮显示的名字
							})
							return
						}

						//解析token，并进行对比数据校验
						user, err := tools.ParseToken(token)
						//fmt.Println("解析的token[112]+", user)
						if err != nil {
							registerWarming = "验证码已过期或者验证码不存在，请尝试重新操作！"
							c.HTML(http.StatusOK, "Message.tmpl", gin.H{
								"title":   registerWarming, //显示的标题
								"url":     "back",          //按钮按下后，跳转的页面url
								"BtnName": "返回",            //按钮显示的名字
							})
							return
						}
						//当解析的token数据与输入的数据不一致:
						//当token的邮箱与输入的邮箱不一致
						if user.ExeEmail != email {

							registerWarming = "token的邮箱与输入的邮箱不一致!"
							c.HTML(http.StatusOK, "Message.tmpl", gin.H{
								"title":   registerWarming, //显示的标题
								"url":     "back",          //按钮按下后，跳转的页面url
								"BtnName": "返回",            //按钮显示的名字
							})
							return
						}
						// token的验证码与输入的验证码不一致
						if (strings.ToLower(user.Captcha)) != (strings.ToLower(emailChapter)) {
							registerWarming = "token的验证码与输入的验证码不一致!"
							c.HTML(http.StatusOK, "Message.tmpl", gin.H{
								"title":   registerWarming, //显示的标题
								"url":     "back",          //按钮按下后，跳转的页面url
								"BtnName": "返回",            //按钮显示的名字
							})
							return
						}

						//保存数据库
						timeStamp, timeInErr := tools.GetTimeStamp() //获取时间戳作为id
						if timeInErr == true {
							registerWarming = "注册失败,生成时间戳错误"
							c.HTML(http.StatusOK, "Message.tmpl", gin.H{
								"title":   registerWarming, //显示的标题
								"url":     "back",          //按钮按下后，跳转的页面url
								"BtnName": "返回",            //按钮显示的名字
							})
							return
						}
						passwdInCode := tools.GetSHA256HashCode(passwd2) //加密密码
						studLogin := dao.Stud_login{}
						studLogin.Sl_id = timeStamp
						studLogin.Sl_email = email
						studLogin.Sl_passwd = passwdInCode
						studLogin.Sl_loginCode = "YYDS"
						dao.AddStudent(studLogin)
						c.HTML(http.StatusOK, "Message.tmpl", gin.H{
							"title":   "注册成功",        //显示的标题
							"url":     "/User/login", //按钮按下后，跳转的页面url
							"BtnName": "返回登录",        //按钮显示的名字
						})
						return
					}

				}

			}

		})

		//--------对注册失败的Post处理
		UserGroup.POST(UserGroupPathRegisterErr, func(c *gin.Context) {
			captcha := tools.CaptchaByTimes(5)
			//重新提交表单，并且打印信息
			c.HTML(http.StatusBadRequest, "register.tmpl", gin.H{

				"massageErr": registerWarming, //错误提示
				"captcha":    captcha,         //验证码
			})
		})

	}

	return r
}
