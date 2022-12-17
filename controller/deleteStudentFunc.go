package controller

import (
	"GraduationProject/dao"
	"GraduationProject/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@Title 【管理员】删除学生的账号
@Author 薛智敏
@CreatTime 2022年6月21日00:49:32
*/

//【管理员】-删除学生账号的业务逻辑

func DeleteStudController(r *gin.Engine) *gin.Engine {

	var deleteStud = "/deleteStud"

	r.GET(master+deleteStud, func(c *gin.Context) {

		id := c.Query("id")
		loginCode := c.Query("loginCode")
		sid := c.Query("sid")            //要注销的学生学号
		delLogoUrl := c.Query("delLogo") //要删除的学生头像
		if sid == "1" {

		}
		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(id, loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}
		//执行删除静态资源命令
		err := tools.RemoveLogo(delLogoUrl)
		if err != nil {
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "删除失败!请稍后重试", //显示的标题
				"url":     "back",       //按钮按下后，跳转的页面url
				"BtnName": "返回",         //按钮显示的名字
			})
			return
		}
		//	执行数据库命令
		finished := dao.DeleteStudBySid(sid)

		//	返回提示
		if finished == 0 {
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "删除失败!请刷新后重试", //显示的标题
				"url":     "back",        //按钮按下后，跳转的页面url
				"BtnName": "返回",          //按钮显示的名字
			})
		} else {
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "删除成功！", //显示的标题
				"url":     "back",  //按钮按下后，跳转的页面url
				"BtnName": "返回",    //按钮显示的名字
			})
		}
	})

	return r
}
