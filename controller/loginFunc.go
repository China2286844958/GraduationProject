package controller

import (
	"GraduationProject/dao"
	"GraduationProject/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/**
@Author 薛智敏
@Title 登录界面的功能以及逻辑代码
@Date 2022年5月12日14:42:30
*/

var loginWarming = ""                           //登录信息警告
var UserLoginGroupMaster = RouterUserGroup      //用户登录路由组主干
var UserGroupPathLogin = "/login"               //登录
var UserGroupPathLoginErr = "/loginErr"         //登录失败标记
var UserGroupPathLoginCaptcha = "/loginCaptcha" //登录验证码查询生成

const (
	uLogin_Student = "Student"
	uLogin_Admin   = "Admin"
)

func UserLoginController(r *gin.Engine) *gin.Engine {
	//学生身份登录get转发
	r.POST("/User/Index", func(c *gin.Context) {

		c.HTML(http.StatusOK, "userIndex.tmpl", nil)
	})

	//================路由组==============
	UserGroup := r.Group(UserLoginGroupMaster)
	{

		//----------登录的Get请求处理，响应一个网页
		UserGroup.GET(UserGroupPathLogin, func(c *gin.Context) {

			captcha := tools.CaptchaByTimes(4)      //生成验证码
			LoginNoValue := c.Query("LoginNoValue") //判断登录是否失效

			if LoginNoValue == "true" {
				tools.SetMassage(c, "登录失效!", "/User/login", "返回登录")

				return
			}
			c.HTML(http.StatusOK, "login.tmpl", gin.H{

				"captcha": captcha,
			})
		})

		//---------对登录Post请求进行处理，获取登录的数据
		UserGroup.POST(UserGroupPathLogin, func(c *gin.Context) {

			email := c.PostForm("email")         //获取登录界面邮箱的填写
			passwd := c.PostForm("passwd")       //获取登录界面密码的填写
			class := c.PostForm("class_login")   //判断用户登录的类型：学生登录还是管理员登录
			captcha := c.PostForm("captcha")     //获取验证码
			captchaIn := c.PostForm("captchaIn") //获取用户输入的验证码

			//判断用户是否输入
			if len(email) == 0 || len(passwd) == 0 {

				loginWarming = "请填写信息！"
				c.Request.URL.Path = UserLoginGroupMaster + UserGroupPathLoginErr
				r.HandleContext(c) //继续后续的处理
				return

			}

			//对管理员的身份进行校验
			if class == uLogin_Admin {
				//验证码判断，不区分大小写
				if strings.ToUpper(captchaIn) != strings.ToUpper(captcha) || strings.ToLower(captchaIn) != strings.ToLower(captcha) {

					loginWarming = "验证码错误！"
					c.Request.URL.Path = UserLoginGroupMaster + UserGroupPathLoginErr
					r.HandleContext(c) //继续后续的处理
					return
				}
				emailFlag := dao.SelectTableByColumn(dao.Db_Course_select_sys, dao.Tb_Admin, "a_email", email) //对管理员表email进行检索

				if emailFlag == true {
					//对密码进行检索，先进行加密
					sha256HashCode := tools.GetSHA256HashCode(passwd)
					passwdFlag := dao.SelectTableByColumn(dao.Db_Course_select_sys, dao.Tb_Admin, "a_passwd", sha256HashCode) //密码加密后对管理员表passwd进行检索

					//密码正确的
					if passwdFlag == true {

						//根据邮箱获取管理员登录表的数据库信息
						admin := dao.GetAdminLoginByEmail(email)

						//设置登录状态码
						c1 := tools.CaptchaByTimes(2) //随机码1
						c2 := tools.CaptchaByTimes(2) //随机码2

						newLoginCode := c1 + captchaIn + c2
						Row := dao.SetAdminLoginCodeById(admin.A_id, newLoginCode)

						if Row == 0 {
							loginWarming = "登录状态码错误！"
							c.Request.URL.Path = UserLoginGroupMaster + UserGroupPathLoginErr
							r.HandleContext(c) //继续后续的处理
							return
						}

						str := fmt.Sprintf("id=%d&loginCode=%s", admin.A_id, newLoginCode)
						c.Redirect(http.StatusMovedPermanently, "firstPage?who=admin&"+str)

					} else {
						loginWarming = "密码错误"
						c.Request.URL.Path = UserLoginGroupMaster + UserGroupPathLoginErr
						r.HandleContext(c) //继续后续的处理
						return
					}

				} else {

					loginWarming = "用户不存在！"
					c.Request.URL.Path = UserLoginGroupMaster + UserGroupPathLoginErr
					r.HandleContext(c) //继续后续的处理
					return

				}

				//	对学生身份进行校验
			} else if class == uLogin_Student {

				//验证码判断，不区分大小写
				if strings.ToUpper(captchaIn) != strings.ToUpper(captcha) || strings.ToLower(captchaIn) != strings.ToLower(captcha) {

					loginWarming = "验证码错误！"
					c.Request.URL.Path = UserLoginGroupMaster + UserGroupPathLoginErr
					r.HandleContext(c) //继续后续的处理
					return
				}

				//对学生表email进行查询
				emailFlag := dao.SelectTableByColumn(dao.Db_Course_select_sys, dao.Tb_Stud_login, "sl_email", email)
				if emailFlag == true {
					//后对密码进行查找：因为需要进行加密，比较耗费时间，所以排在后

					passwdEncode := tools.GetSHA256HashCode(passwd) //加密密码
					//查询数据库
					passwdFlag := dao.QueryByEmailAndPasswd(email, passwdEncode)

					if passwdFlag == true {

						//根据邮箱获取学生登录表的数据库信息
						studLogin := dao.GetStudentLoginByEmail(email)

						//设置学生的登录状态码
						dao.SetLoginCode(studLogin.Sl_id, captcha)
						//登录成功查询所有的数据，数据进行传输id和登录状态验证码

						//重定向为get请求，结果转发到个人登录界面
						parmStr := fmt.Sprintf("id=%d&loginCode=%s", studLogin.Sl_id, captcha)
						c.Redirect(http.StatusMovedPermanently, "firstPage?who=student&"+parmStr)

					} else {
						loginWarming = "密码错误！！"
						c.Request.URL.Path = UserLoginGroupMaster + UserGroupPathLoginErr
						r.HandleContext(c) //继续后续的处理
					}
				} else {

					loginWarming = "邮箱错误或者不存在！"
					c.Request.URL.Path = UserLoginGroupMaster + UserGroupPathLoginErr
					r.HandleContext(c) //继续后续的处理
				}

			}

		})

		//---------登录失败的Post处理---------------
		UserGroup.POST(UserGroupPathLoginErr, func(c *gin.Context) {
			captcha := tools.CaptchaByTimes(4) //生成验证码

			c.HTML(http.StatusOK, "login.tmpl", gin.H{
				"loginWarming": loginWarming,
				"captcha":      captcha,
			})
		})

		//	-----------------登录验证码重新生成------------------
		UserGroup.POST(UserGroupPathLoginCaptcha, func(c *gin.Context) {
			captcha := tools.CaptchaByTimes(5) //生成验证码

			c.HTML(http.StatusOK, "login.tmpl", gin.H{
				"loginWarming": loginWarming,
				"captcha":      captcha,
			})
		})

	}

	//

	return r
}
