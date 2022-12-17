package controller

import (
	"GraduationProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
@Title 管理员-课程管理
@Author 薛智敏
@CreateTime 2022年6月18日18:09:31
*/

//课程管理【管理员端】

func CourseManager(r *gin.Engine) *gin.Engine {
	var courseList = "/courseManage"

	group := r.Group(master)
	{

		group.GET(courseList, func(c *gin.Context) {
			id := c.Query("id")
			loginCode := c.Query("loginCode")
			searchWord := c.Query("searchWord")   //搜索词
			search_Page := c.Query("search_Page") //查询页

			//管理员登录验证码校验
			T := dao.AdminLoginCodeVerification(id, loginCode)
			if !T {
				c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
				return
			}

			search_PageInt, errP := strconv.Atoi(search_Page) //分页--类型转化

			searchPageErr := dao.CheckErr(errP) // 边界值判断
			if searchPageErr || search_PageInt <= 0 || search_PageInt > 10000 {
				search_PageInt = 1 //初始是0页
			}

			//根据课程名或者授课老师查询课程
			courseList, CountPage := dao.SelectCourseByCNameAndCTeach(searchWord, search_PageInt-1)

			//==============【课程管理】=======================
			//Top
			c.HTML(http.StatusOK, "courseManageTop.tmpl", gin.H{
				"id":          id,
				"loginCode":   loginCode,
				"currentPage": search_PageInt,         //当前页数
				"PageInfo":    dao.ShowNumbersProxy(), //分页查询中每页显示的信息数
				"CountPage":   CountPage,              //总页数
			})

			//Mid

			for i := 0; i < len(courseList); i++ {
				course := courseList[i] //取课程
				c.HTML(http.StatusOK, "courseManageMid.tmpl", gin.H{
					"id":        id,
					"loginCode": loginCode,
					"c":         course,
				})
			}

			//End
			c.HTML(http.StatusOK, "courseManageEnd.tmpl", gin.H{
				"id":        id,
				"loginCode": loginCode,
			})

		})

	}

	return r
}
