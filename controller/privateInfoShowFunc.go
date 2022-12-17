package controller

import (
	"GraduationProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@Title 用户信息展示
@author 薛智敏
@CreateTime 2022年6月8日21:08:00

*/

const (
	privateInfoShow = "/privateInfoShow"
)

//用户信息展示

func PrivateInfoShowController(r *gin.Engine) *gin.Engine {

	group := r.Group(master)
	{

		//学生端展示
		group.GET(privateInfoShow, func(c *gin.Context) {
			id := c.Query("id")
			loginCode := c.Query("loginCode")
			//登录验证码验证
			loginTrue := dao.TrueLoginByEmailAdLoginCodes(id, loginCode)
			if loginTrue == false {
				c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
				return
			}

			_, stud_del := dao.QueryUnionStudentById(id)

			c.HTML(http.StatusOK, "privateInfoShow.tmpl", gin.H{
				"id":        stud_del.Sd_id,      //学生ID
				"loginCode": loginCode,           //学生登录状态码
				"name":      stud_del.Sd_relname, //学生姓名
				"age":       stud_del.Sd_age,     //学生年龄
				"gender":    stud_del.Sd_gender,  //学生性别
				"address":   stud_del.Sd_address, //学生地址
				"sys":       stud_del.Sd_sys,     //学生系别
			})
		})

	}

	return r
}
