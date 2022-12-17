package controller

import (
	"GraduationProject/dao"
	"GraduationProject/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userIndexMaster = RouterUserGroup //主路由
var adminRouter = "/adminIndex"       //管理身份登录的主界面
var studentRouter = "/studentIndex"   //学生身份登录的主界面

//传输的数据

var StudEmail = ""  //学生表邮箱
var AdminEmail = "" //管理员邮箱
//个人界面

func UserIndexController(r *gin.Engine) *gin.Engine {

	userIndexGroup := r.Group(userIndexMaster)
	{

		//管理员身份登录
		userIndexGroup.GET(adminRouter, func(c *gin.Context) {
			timeStatus := tools.TimeStatus()

			id := c.Query("id")
			loginCode := c.Query("loginCode")

			//管理员登录验证码校验
			T := dao.AdminLoginCodeVerification(id, loginCode)
			if !T {
				c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
				return
			}

			//根据id获取管理员
			admin := dao.GetAdminByIdAndLoginCode(id, loginCode)

			c.HTML(http.StatusOK, "adminIndex.tmpl", gin.H{
				"timeStatus": timeStatus + "好!", //时间状态
				"id":         admin.A_id,        //id
				"email":      admin.A_email,     //邮箱
				"loginCode":  admin.A_loginCode, //登录验证码
				"power":      admin.A_power,     //管理员的权限值
			})

		})

		//学生身份登录，访问个人界面
		userIndexGroup.GET(studentRouter, func(c *gin.Context) {
			timeStatus := tools.TimeStatus()

			id := c.Query("id")
			StudEmail = c.Query("email")      //获取url传输的数据
			loginCode := c.Query("loginCode") //获取登录状态码

			//登录验证码验证
			loginTrue := dao.TrueLoginByEmailAdLoginCodes(id, loginCode)
			if !loginTrue {
				c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
				return
			}
			stud_login, stud_del := dao.QueryUnionStudentById(id) //联合查询

			//当学生未实名
			times := 0
			//姓名判断
			if stud_del.Sd_relname == "" {
				times++
			}
			//年龄判断
			if stud_del.Sd_age == 0 {
				times++
			}
			//性别判断
			if stud_del.Sd_gender == "" {
				times++
			}
			//地址
			if stud_del.Sd_address == "" {
				times++
			}
			//性别
			if stud_del.Sd_sys == "" {
				times++
			}

			//当有一个没填时,判断为信息不完善
			if times >= 1 {
				stud_del.Sd_relname = "[信息未完善]"
			}
			//获取用户头像
			studLogin := dao.QueryStuLoginByIdAndLoginCode(id, loginCode)
			c.HTML(http.StatusOK, "userIndex.tmpl", gin.H{

				"timeStatus": timeStatus + "好!",   //时间状态
				"id":         id,                  //id
				"loginCode":  loginCode,           //登录校验码
				"email":      stud_login.Sl_email, //邮箱
				"name":       stud_del.Sd_relname, //真实姓名
				"age":        stud_del.Sd_age,     //年龄
				"gender":     stud_del.Sd_gender,  //性别
				"address":    stud_del.Sd_address, //地址
				"sys":        stud_del.Sd_sys,     //系别
				"logoUrl":    studLogin.Sl_logo,   //用户头像

				//"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
			})

		})

	}
	//学生身份登录
	userIndexGroup.POST(studentRouter, func(c *gin.Context) {

		c.HTML(http.StatusOK, "userIndex.tmpl", nil)

		c.JSON(http.StatusOK, gin.H{
			"WHO": "Student",
		})
	})

	return r
}
