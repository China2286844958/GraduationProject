package controller

import (
	"GraduationProject/dao"
	"GraduationProject/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@Title 超级管理员的处理逻辑区
@Author 薛智敏
@CreateTime 2022年6月16日00:17:56
*/

//超级管理员-重置普通管理员密码的逻辑处理

func SuResetPasswdAdmin(r *gin.Engine) *gin.Engine {
	var SuResetPasswdAdmin = "/SuResetPasswdAdmin"

	r.GET(master+SuResetPasswdAdmin, func(c *gin.Context) {

		Res_id := c.Query("Res_id")             //获取要重置的密码
		Su_id := c.Query("Su_id")               //获取管理员id
		Su_loginCode := c.Query("Su_loginCode") //超级管理员登录验证码

		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(Su_id, Su_loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		//重置的密码设置
		var ResedPasswd = "123456"
		enCode := tools.GetSha256EnCodeByDeCode(ResedPasswd)
		finishedNum := dao.SuResetPasswdAdmin(Res_id, enCode) //数据库执行
		if finishedNum == 0 {

			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "重置失败！刷新后重试", //显示的标题
				"url":     "back",       //按钮按下后，跳转的页面url
				"BtnName": "返回",         //按钮显示的名字
			})

			return
		}

		//	成功的
		c.HTML(http.StatusOK, "Message.tmpl", gin.H{
			"title":   "重置成功！重置的密码为" + ResedPasswd, //显示的标题
			"url":     "back",                      //按钮按下后，跳转的页面url
			"BtnName": "返回",                        //按钮显示的名字
		})

	})

	return r
}

//超级管理员-删除普通管理员的逻辑操作

func SuDeleteAdmin(r *gin.Engine) *gin.Engine {
	var SuDeleteAdmin = "/SuDeleteAdmin"

	r.GET(master+SuDeleteAdmin, func(c *gin.Context) {

		Su_id := c.Query("Su_id")               //获取管理员id
		del_id := c.Query("Del_id")             //获取要删除的普通管理员id
		Su_loginCode := c.Query("Su_loginCode") //超级管理员登录验证码

		//超级管理员-登录验证码校验
		T := dao.AdminLoginCodeVerification(Su_id, Su_loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		num := dao.DeleteAdmin(del_id) //执行数据库
		if num == 0 {

			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "删除失败，请重新刷新后再试！", //显示的标题
				"url":     "back",           //按钮按下后，跳转的页面url
				"BtnName": "返回",             //按钮显示的名字

			})
			return
		}

		c.HTML(http.StatusOK, "Message.tmpl", gin.H{
			"title":   "删除成功，请刷新！", //显示的标题
			"url":     "back",      //按钮按下后，跳转的页面url
			"BtnName": "返回",        //按钮显示的名字
		})
	})

	return r
}
