package controller

import (
	"GraduationProject/dao"
	"GraduationProject/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@Title 更新密码
@Author 薛智敏
@CreateTime 2022年6月24日12:05:46
*/

//【管理员】-更新密码

func AdminResetPasswd(r *gin.Engine) *gin.Engine {

	var adminResetPasswd = "/adminResetPasswd"
	var adminResetPasswdExe = "/adminResetPasswdExe"

	//处理请求
	r.GET(master+adminResetPasswd, func(c *gin.Context) {
		id := c.Query("id")
		loginCode := c.Query("loginCode")
		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(id, loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		c.HTML(http.StatusOK, "adminResetPasswd.tmpl", gin.H{
			"id":        id,
			"loginCode": loginCode,
		})
	})

	//执行更新
	r.GET(master+adminResetPasswdExe, func(c *gin.Context) {
		id := c.Query("id")
		loginCode := c.Query("loginCode")

		old_p := c.Query("old_p")
		new_p1 := c.Query("new_p1")
		new_p2 := c.Query("new_p2")

		messTitle := "" //违规展示

		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(id, loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		//====数据校验

		//获取参数的长度
		old_size := tools.StrCounts(old_p)
		new1_size := tools.StrCounts(new_p1)
		new2_size := tools.StrCounts(new_p2)
		//位数限制
		if (old_size < 4 || old_size > 30) || (new1_size < 4 || new1_size > 30) || (new2_size < 4 || new2_size > 30) {
			//	返回提示
			messTitle = "请输入4到30位的数据！"
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   messTitle, //显示的标题
				"url":     "back",    //按钮按下后，跳转的页面url
				"BtnName": "返回",      //按钮显示的名字
			})
			return
		}

		//设置的第一次密码与第二次密码不一致时
		if new_p1 != new_p2 {
			//	返回提示
			messTitle = "两次重设的密码不一致!"
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   messTitle, //显示的标题
				"url":     "back",    //按钮按下后，跳转的页面url
				"BtnName": "返回",      //按钮显示的名字
			})
			return
		}

		//三次密码一样的时候
		if (old_p == new_p1) && (old_p == new_p2) {
			//	返回提示
			messTitle = "三次密码一致,取消重置密码!"
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   messTitle, //显示的标题
				"url":     "back",    //按钮按下后，跳转的页面url
				"BtnName": "返回",      //按钮显示的名字
			})
			return
		}

		//执行数据库密码校验
		exist := dao.FindAdminByIdAndEnPasswd(id, tools.GetSHA256HashCode(old_p))
		if !exist {
			//	返回提示
			messTitle = "原密码不正确!"
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   messTitle, //显示的标题
				"url":     "back",    //按钮按下后，跳转的页面url
				"BtnName": "返回",      //按钮显示的名字
			})
			return
		}

		//【管理员】执行密码重置的SQL

		dao.AdminResetPasswd(id, tools.GetSHA256HashCode(new_p2))
		messTitle = fmt.Sprintf("密码更新成功,请重新登陆!")
		c.HTML(http.StatusOK, "Message.tmpl", gin.H{
			"title":   messTitle,     //显示的标题
			"url":     "/User/login", //按钮按下后，跳转的页面url
			"BtnName": "返回登陆",        //按钮显示的名字
		})

	})
	return r
}
