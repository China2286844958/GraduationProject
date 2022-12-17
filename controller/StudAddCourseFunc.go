package controller

import (
	"GraduationProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@Author 薛智敏
@CreateTime 2022年6月2日20:59:11
@Title 学生添加课程

*/

const (
	studAddCourse = "/studAddCourse"
)

func StudAddCourseController(r *gin.Engine) *gin.Engine {

	r.GET(master+studAddCourse, func(c *gin.Context) {
		id := c.Query("id")
		loginCode := c.Query("loginCode")
		C_id := c.Query("C_id")
		//登录验证码验证
		loginTrue := dao.TrueLoginByEmailAdLoginCodes(id, loginCode)
		if loginTrue == false {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}
		//	执行数据库语句
		numRows := dao.AddCourseBySidAndCid(id, C_id)

		//成功返回
		if numRows > 0 {
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "添加课程成功！请刷新", //显示的标题
				"url":     "back",       //按钮按下后，跳转的页面url
				"BtnName": "返回",         //按钮显示的名字
			})
			//c.Redirect(http.StatusMovedPermanently, "joinCourse?id="+id+"&loginCode="+loginCode)
			return
		} else {
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "添加课程失败！请重新操作或者清除缓存后再试", //显示的标题
				"url":     "back",                  //按钮按下后，跳转的页面url
				"BtnName": "返回",                    //按钮显示的名字
			})
		}

	})
	return r
}
