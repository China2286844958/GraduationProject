package controller

import (
	"GraduationProject/dao"
	"GraduationProject/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@Title 【超级管理员】-添加管理员界面的逻辑处理
@Author 薛智敏
@CreateTime 2022年6月14日00:03:18
*/

//【超级管理员】-添加管理员

func AddAdminController(r *gin.Engine) *gin.Engine {
	var addAdmin = "/addAdmin"

	//处理get请求
	r.GET(master+addAdmin, func(c *gin.Context) {
		id := c.Query("id")
		loginCode := c.Query("loginCode")

		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(id, loginCode)

		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		c.HTML(http.StatusOK, "addAdmin.tmpl", gin.H{
			"id":        id,
			"loginCode": loginCode,
		})

	})

	//post执行添加管理员的操作
	r.POST(master+addAdmin, func(c *gin.Context) {

		id := c.PostForm("id")
		loginCode := c.PostForm("loginCode")

		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(id, loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		MinLength := 4                         //管理员邮箱最短
		MaxLength := 30                        //管理员邮箱最长
		email_in := c.PostForm("email_in")     //获取验证码
		passwd_in1 := c.PostForm("passwd_in1") //获取第一次的密码
		passwd_in2 := c.PostForm("passwd_in2") //获取第二次的密码

		errTitle := "" //错误提示
		errTimes := 0  //错误的标记次数

		//输入的邮箱,字符限制根据数据库设计在4到30位;第一次密码必须和第二次密码一样，而且字符限制在4到30位
		if len(email_in) < MinLength {
			errTimes++
			errTitle = fmt.Sprintf("输入的邮箱长度低于%d位!", MinLength)
		}
		if len(email_in) > MaxLength {
			errTimes++
			errTitle += fmt.Sprintf("输入的邮箱长度高于%d位!", MaxLength)
		}
		if passwd_in1 != passwd_in2 {
			errTimes++
			errTitle += "两次密码不一致!"
		}
		if len(passwd_in2) < MinLength {
			errTimes++
			errTitle += fmt.Sprintf("输入的密码长度低于%d位!", MinLength)
		}
		if len(passwd_in2) > MaxLength {
			errTimes++
			errTitle += fmt.Sprintf("输入的密码长度高于%d位!", MaxLength)
		}

		//校验验证码
		if !tools.EmailRegExp(email_in) {
			errTimes++
			errTitle += "不符合邮箱格式！"
		}

		if errTimes > 0 { //errTimes>0表示不符合规则的
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   errTitle, //显示的标题
				"url":     "back",   //按钮按下后，跳转的页面url
				"BtnName": "返回",     //按钮显示的名字
			})
			return
		}

		//	对合法的数据进行保存数据库【登录成功，输入的数据合法】

		//先进行数据库搜索，注册过的邮箱是否和准备注册的是否冲突！
		admin := dao.GetAdminLoginByEmail(email_in)
		if admin.A_email == email_in {
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "邮箱已被注册！", //显示的标题
				"url":     "back",    //按钮按下后，跳转的页面url
				"BtnName": "返回",      //按钮显示的名字
			})
			return
		}

		timeStamp, _ := tools.GetTimeStamp()                //根据时间戳生成id
		EncodePasswd := tools.GetSHA256HashCode(passwd_in2) //密码加密

		dao.AddAdminByRoot(timeStamp, email_in, EncodePasswd) //执行数据库添加
		c.HTML(http.StatusOK, "Message.tmpl", gin.H{
			"title":   "添加成功！请刷新", //显示的标题
			"url":     "back",     //按钮按下后，跳转的页面url
			"BtnName": "返回",       //按钮显示的名字
		})
	})

	return r
}
