package controller

import (
	"GraduationProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
@Title 管理员-学生管理
@Author 薛智敏
@CreateTime 2022年6月19日09:56:17
*/

//学生管理，对学生进行重置密码，账号的注销，选课查询、批改分数

func StudentMangeController(r *gin.Engine) *gin.Engine {

	var studentManage = "/studentManage"

	r.GET(master+studentManage, func(c *gin.Context) {
		id := c.Query("id")
		loginCode := c.Query("loginCode")
		searchWord := c.Query("searchWord")   //搜索词
		courseId := c.Query("courseId")       //课程id
		search_Page := c.Query("search_Page") //查询页数

		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(id, loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		//当id 有数值传入
		if courseId != "" {

		}

		//根据课程id获取课程信息
		courseById := dao.SelectCourseByCourseId(courseId) //查询所有课程，返回课程结构体

		searchFlag := searchWord

		//当搜索词为空，搜索的就是全部
		if searchFlag == "" {
			searchFlag = "[全部]"
		}
		//Top开始
		c.HTML(http.StatusOK, "studentManageTop.tmpl", gin.H{
			"cbi":        courseById, //课程信息
			"id":         id,         //id
			"loginCode":  loginCode,  //登录验证码
			"searchFlag": searchFlag, //搜索词复制

		})

		//Mid1课程
		courses := dao.SelectAllCourse() //查询所有课程，返回课程结构体
		for i := 0; i < len(courses); i++ {

			v := courses[i]
			c.HTML(http.StatusOK, "studentManageMid1.tmpl", gin.H{
				"c": v, //传输数据
			})
		}

		search_PageInt, errP := strconv.Atoi(search_Page) //分页--类型转化

		searchPageErr := dao.CheckErr(errP)
		if searchPageErr || search_PageInt <= 0 || search_PageInt > 10000 { // 分页查询的边界值判断
			search_PageInt = 1 //初始是0页
		}
		list, CountPage := dao.QueryUnionStudByJoinedCourse(courseId, searchWord, search_PageInt-1) //筛选的学生

		//Mid2搜索按钮
		c.HTML(http.StatusOK, "studentManageMid2.tmpl", gin.H{
			"id":          id,                     //id
			"loginCode":   loginCode,              //登录验证码
			"courseId":    courseById.C_Id,        //课程信息
			"searchWord":  searchWord,             //搜索词
			"currentPage": search_PageInt,         //当前页数
			"PageInfo":    dao.ShowNumbersProxy(), //分页查询中每页显示的信息数
			"CountPage":   CountPage,              //总页数
		})

		//Mid3筛选的学生信息

		if search_Page == "" { //当搜索页面为空
			search_Page = "1"
		}

		for i := 0; i < len(list); i++ {
			v := list[i]
			c.HTML(http.StatusOK, "studentManageMid3.tmpl", gin.H{
				"s":         v,          //管理员-选课管理-学生管理截取字段
				"id":        id,         //id
				"loginCode": loginCode,  //登录验证码
				"c":         courseById, //课程信息

			})
		}

		//End结束
		c.HTML(http.StatusOK, "studentManageEnd.tmpl", nil)

	})

	return r
}
