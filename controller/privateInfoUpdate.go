package controller

import (
	"GraduationProject/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	privateInfoUpdate = "/privateInfoUpdate"
)

//修改个人信息

func PrivateInfoUpdate(r *gin.Engine) *gin.Engine {

	r.GET(master+privateInfoUpdate, func(c *gin.Context) {

		id := c.Query("id")
		loginCode := c.Query("loginCode")
		//登录验证码验证
		loginTrue := dao.TrueLoginByEmailAdLoginCodes(id, loginCode)
		if loginTrue == false {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		_, stud_del := dao.QueryUnionStudentById(id)
		c.HTML(http.StatusOK, "privateInfoUpdate.tmpl", gin.H{
			"id":        stud_del.Sd_id,      //学生ID
			"loginCode": loginCode,           //学生登录状态码
			"name":      stud_del.Sd_relname, //学生姓名
			"age":       stud_del.Sd_age,     //学生年龄
			"gender":    stud_del.Sd_gender,  //学生性别
			"address":   stud_del.Sd_address, //学生地址
			"sys":       stud_del.Sd_sys,     //学生系别
		})

	})

	//提交表单后的数据库更改
	r.POST(master+privateInfoUpdate, func(c *gin.Context) {
		rel_name := c.PostForm("rel_name") //获取真实姓名
		age := c.PostForm("age")           //获取学生的年龄
		gender := c.PostForm("gender")     //获取学生的性别
		address := c.PostForm("address")   //获取学生的住址
		sys := c.PostForm("sys")           //获取学生的院系
		id := c.Query("id")                //获取学生的学号id

		//根据学号查询修改信息
		old := dao.SelectStudentById(id) //以前的信息

		//当修改的信息无变化
		if old.Sd_relname == rel_name && fmt.Sprintf("%d", old.Sd_age) == age && old.Sd_gender == gender && old.Sd_address == address && old.Sd_sys == sys {
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "无修改数据", //显示的标题
				"url":     "back",  //按钮按下后，跳转的页面url
				"BtnName": "返回",    //按钮显示的名字
			})
			return
		}

		//执行修改的SQL命令
		finishedNum := dao.StudDelUpdateById(id, rel_name, age, gender, address, sys)
		if finishedNum > 0 {
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "更新成功", //显示的标题
				"url":     "back", //按钮按下后，跳转的页面url
				"BtnName": "返回",   //按钮显示的名字
			})
			return
		} else {
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "更新失败", //显示的标题
				"url":     "back", //按钮按下后，跳转的页面url
				"BtnName": "返回",   //按钮显示的名字
			})
			return

		}

	})

	return r
}
