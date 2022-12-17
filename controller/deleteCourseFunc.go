package controller

import (
	"GraduationProject/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@Title 管理员-删除课程的业务逻辑
@Author 薛智敏
@CreateTime 2022年6月24日00:38:40
*/

//管理员 -删除课程

func DeleteCourseController(r *gin.Engine) *gin.Engine {

	var deleteCourse = "/deleteCourse"
	r.GET(master+deleteCourse, func(c *gin.Context) {
		id := c.Query("id")
		loginCode := c.Query("loginCode")
		delId := c.Query("delId") //删除的课程id

		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(id, loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}
		//执行数据库语句
		finished := dao.DeleteCourseByCid(delId)

		title := fmt.Sprintf("共删除%d个课程!", finished)
		c.HTML(http.StatusOK, "Message.tmpl", gin.H{
			"title":   title,  //显示的标题
			"url":     "back", //按钮按下后，跳转的页面url
			"BtnName": "返回",   //按钮显示的名字
		})
	})

	return r
}
