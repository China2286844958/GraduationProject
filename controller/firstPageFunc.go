package controller

import (
	"GraduationProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FirstPageController
//
//	@Description: 首页处理的逻辑
//	@param r
//	@return *gin.Engine
func FirstPageController(r *gin.Engine) *gin.Engine {

	var firstPage = "/firstPage"

	r.GET(master+firstPage, func(c *gin.Context) {
		nameList := dao.SelectAllCourseName()               //查询所有课程名字 return []string
		maxScoreList := dao.SelectMaxScoreByCName(nameList) //查询课程对应的最高分 return []float
		var arrMax = maxScoreList[0:len(maxScoreList)]      //最高分数集合
		minScoreList := dao.SelectMinScoreByCName(nameList) //查询课程对应的最低分
		var arrMin = minScoreList[0:len(minScoreList)]      //最低分数集合
		who_href := "login"                                 //跳转的网页,默认是登录
		//身份校验
		who := c.Query("who")             //获取身份
		id := c.Query("id")               //身份id
		loginCode := c.Query("loginCode") //登录验证码
		switch who {
		case "student": //学生
			//学生登录验证码验证
			loginTrue := dao.TrueLoginByEmailAdLoginCodes(id, loginCode)
			if !loginTrue {
				c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
				return
			}
			who_href = "firstPage_student.tmpl"
		case "admin": //管理员
			//管理员登录验证码校验
			loginTrue := dao.AdminLoginCodeVerification(id, loginCode)

			if !loginTrue {
				c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
				return
			}
			who_href = "firstPage_admin.tmpl"
		default:
			c.Redirect(http.StatusMovedPermanently, who_href) //重定向登录界面
			return
		}

		//展示
		c.HTML(http.StatusOK, who_href, gin.H{
			"courseNameList": nameList,            //课程名集合
			"maxScoreList":   arrMax,              //课程最高分集合
			"mixScoreList":   arrMin,              //课程最低分集合
			"id":             id,                  //id
			"loginCode":      loginCode,           //登录验证码
			"size":           len(nameList) * 200, //表格盒子的尺寸
		})

		return
	})
	return r
}
