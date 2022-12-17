package controller

import (
	"GraduationProject/dao"
	"GraduationProject/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@Title 【管理员】对学生重置密码
@Author 薛智敏
@CreatTime 2022年6月21日00:47:41
*/

//【管理员】重置学生账号的密码

func StudResetPasswdController(r *gin.Engine) *gin.Engine {
	var studResetPasswd = "/StudResetPasswd"
	//http://localhost:8848/User/StudResetPasswd?id=20020228&loginCode=Ra5mvhRa&sid=202262017626499610
	r.GET(master+studResetPasswd, func(c *gin.Context) {
		id := c.Query("id")
		loginCode := c.Query("loginCode")
		sid := c.Query("sid") //要重置密码的学生的学号

		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(id, loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		resetPasswd := "1234"
		enCode := tools.GetSha256EnCodeByDeCode(resetPasswd)
		//	执行数据库命令
		num := dao.ResetStudPasswd(sid, enCode) //更改密码的SQL，返回一个标记，该标记的执行数据库后，影响的行数，只有更改成功才能邮箱

		if num == 0 {
			//	返回提示
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "重置密码失败,请刷新后重试！", //显示的标题
				"url":     "back",           //按钮按下后，跳转的页面url
				"BtnName": "返回",             //按钮显示的名字
			})
			return
		} else {
			//	返回提示
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "重置密码成功,重置的密码为" + resetPasswd, //显示的标题
				"url":     "back",                        //按钮按下后，跳转的页面url
				"BtnName": "返回",                          //按钮显示的名字
			})

		}

	})

	return r
}
