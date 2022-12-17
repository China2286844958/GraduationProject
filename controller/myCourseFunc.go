package controller

import (
	"GraduationProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**学生选课【我的课程】

 */

var MyCourseLis = "/myCourse"

//我的课程【学生端】

func MyCourse(r *gin.Engine) *gin.Engine {

	group := r.Group(master)
	{

		group.GET(MyCourseLis, func(c *gin.Context) {
			id := c.Query("id")
			loginCode := c.Query("loginCode")     //登录验证码
			search_data := c.Query("search_data") //查询字段
			search_Page := c.Query("search_Page") //查询页

			//登录验证码验证
			loginTrue := dao.TrueLoginByEmailAdLoginCodes(id, loginCode)
			if loginTrue == false {
				c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
				return
			}

			//==============【我的课程】=======================

			search_PageInt, errP := strconv.Atoi(search_Page) //分页--类型转化

			searchPageErr := dao.CheckErr(errP) // 边界值判断
			if searchPageErr || search_PageInt <= 0 || search_PageInt > 10000 {
				search_PageInt = 1 //初始是0页
			}

			//查询course表
			list, CountPage := dao.QueryCourseByStudId(id, search_data, search_PageInt-1)
			//上
			c.HTML(http.StatusOK, "myCourseTop.tmpl", gin.H{
				"id":          id,
				"loginCode":   loginCode,
				"search_data": search_data,            //搜索词
				"currentPage": search_PageInt,         //当前页数
				"PageInfo":    dao.ShowNumbersProxy(), //分页查询中每页显示的信息数
				"CountPage":   CountPage,              //总页数
			})

			//中
			for i := 0; i < len(list); i++ {

				course := list[i]
				c.HTML(http.StatusOK, "myCourseMid.tmpl", gin.H{
					"c_loops":   course,
					"id":        id,
					"loginCode": loginCode,
				})
			}
			//下
			c.HTML(http.StatusOK, "myCourseEnd.tmpl", gin.H{})

		})

	}

	return r
}
