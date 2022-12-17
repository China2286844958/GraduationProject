package controller

import (
	"GraduationProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@title 管理员-选课管理
@createTime 2022年6月19日00:04:40
@author 薛智敏
*/

func JoinCourseManageController(r *gin.Engine) *gin.Engine {
	var joinCourseManage = "/joinCourseManage"

	r.GET(master+joinCourseManage, func(c *gin.Context) {
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
		c.HTML(http.StatusOK, "joinCourseManage.tmpl", gin.H{
			"id":        admin.A_id,        //id
			"email":     admin.A_email,     //邮箱
			"loginCode": admin.A_loginCode, //登录验证码
			"power":     admin.A_power,     //管理员的权限值
		})
	})
	return r
}
