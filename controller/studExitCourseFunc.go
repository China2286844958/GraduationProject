package controller

import (
	"GraduationProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@createTime 2022年5月31日21:57:41
@title 学生退课处理
*/

const (
	exitCourse = "/exitCourse" //路由
)

//学生退课处理

func StudExitCourseController(r *gin.Engine) *gin.Engine {

	//执行删除数据库命令，返回
	r.GET(master+exitCourse, func(c *gin.Context) {

		id := c.Query("id")               //学生id
		loginCode := c.Query("loginCode") //登录验证码
		SC_id := c.Query("SC_id")         //课程id

		//登录验证码验证
		loginTrue := dao.TrueLoginByEmailAdLoginCodes(id, loginCode)
		if loginTrue == false {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		rowsNum := dao.StudExitCourseBySid(id, SC_id)
		if rowsNum > 0 {
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "退课成功！请刷新", //显示的标题
				"url":     "back",     //按钮按下后，跳转的页面url
				"BtnName": "返回",       //按钮显示的名字
			})
		} else {
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "退课失败！请重新操作或者清除缓存后再试", //显示的标题
				"url":     "back",                //按钮按下后，跳转的页面url
				"BtnName": "返回",                  //按钮显示的名字
			})
		}

		//c.Redirect(http.StatusMovedPermanently, "/User/myCourse?id="+id+"&loginCode="+loginCode)
		return
	})

	return r
}
