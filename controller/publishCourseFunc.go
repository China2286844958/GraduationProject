package controller

import (
	"GraduationProject/dao"
	"GraduationProject/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
@Title 发布课程
@Author 薛智敏
@CreateTime 2022年6月24日00:44:17
*/

//【管理员】-发布课程的业务逻辑

func PublishCourse(r *gin.Engine) *gin.Engine {
	var publishCourse = "/publishCourse"
	var publishCourseSave = "/publishCourseSave"

	//处理请求，返回网页
	r.GET(master+publishCourse, func(c *gin.Context) {
		id := c.Query("id")
		loginCode := c.Query("loginCode")

		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(id, loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		//	响应网页
		c.HTML(http.StatusOK, "publishCourse.tmpl", gin.H{
			"id":        id,        //id
			"loginCode": loginCode, //登录验证码
		})

	})

	//处理请求，保存数据库
	r.POST(master+publishCourseSave, func(c *gin.Context) {
		id := c.PostForm("id")
		loginCode := c.PostForm("loginCode")
		Cname := c.PostForm("Cname")     //课程名字
		Teacher := c.PostForm("Teacher") //授课老师
		Tscore := c.PostForm("Tscore")   //课程总分
		errInfo := ""                    //错误信息

		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(id, loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		Scoref32, err := strconv.ParseFloat(Tscore, 32) //分数转化float32
		dao.CheckErr(err)

		//信息规范
		CNameSize := tools.StrCounts(Cname)     //获取管理员输入的课程长度
		TeacherSize := tools.StrCounts(Teacher) //获取管理员输入的授课老师长度
		TscoreSize := tools.StrCounts(Tscore)   //获取管理员输入的课程总分长度
		MaxSize := 25                           //最大输入

		if CNameSize == 0 || TeacherSize == 0 || TscoreSize == 0 {

			errInfo = "请填写所有信息!"
			//	响应网页
			c.HTML(http.StatusOK, "publishCourse.tmpl", gin.H{
				"id":        id,        //id
				"loginCode": loginCode, //登录验证码
				"err":       errInfo,   //错误
			})
			return
		} else if err != nil || Scoref32 > 150 || Scoref32 < 0 {
			errInfo = "分数不合法!"
			//	响应网页
			c.HTML(http.StatusOK, "publishCourse.tmpl", gin.H{
				"id":        id,        //id
				"loginCode": loginCode, //登录验证码
				"err":       errInfo,   //错误
			})
			return

		} else if CNameSize > MaxSize || TeacherSize >= MaxSize {
			errInfo = fmt.Sprintf("信息的长度不能超过%d位!", MaxSize)
			//	响应网页
			c.HTML(http.StatusOK, "publishCourse.tmpl", gin.H{
				"id":        id,        //id
				"loginCode": loginCode, //登录验证码
				"err":       errInfo,   //错误
			})
			return
		}

		//	执行数据库命令
		finished := dao.AddCourse(Cname, Teacher, Tscore)
		title := fmt.Sprintf("成功添加%d个课程!", finished)
		c.HTML(http.StatusOK, "Message.tmpl", gin.H{
			"title":   title,  //显示的标题
			"url":     "back", //按钮按下后，跳转的页面url
			"BtnName": "返回",   //按钮显示的名字
		})
	})

	return r
}
