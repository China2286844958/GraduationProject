package controller

import (
	"GraduationProject/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
@Author 薛智敏
@CreateTime 2022年6月1日16:34:41
@Title 【学生】选课
*/

const (
	joinCourse = "/joinCourse"
)

func JoinCourseController(r *gin.Engine) *gin.Engine {

	r.GET(master+joinCourse, func(c *gin.Context) {
		id := c.Query("id")                     //学生id
		loginCode := c.Query("loginCode")       //登录验证码
		search_words := c.Query("search_words") //搜索词
		search_Page := c.Query("search_Page")   //查询页
		length := 0

		for _, a := range search_words {
			fmt.Sprintf("%d", a)
			length++
		}

		//校验搜索词
		if length > 10 {
			title := fmt.Sprintf("搜索词一共%d个,超过10个字符，请重新填写！", length)
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{

				"title":   title,   //显示的标题
				"url":     "back",  //按钮按下后，跳转的页面url
				"BtnName": "返回上一步", //按钮显示的名字
			})

			return
		}

		//登录验证码验证
		loginTrue := dao.TrueLoginByEmailAdLoginCodes(id, loginCode)
		if loginTrue == false {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		search_PageInt, errP := strconv.Atoi(search_Page) //分页--类型转化

		searchPageErr := dao.CheckErr(errP) // 边界值判断
		if searchPageErr || search_PageInt <= 0 || search_PageInt > 10000 {
			search_PageInt = 1 //初始是0页
		}

		//查询课程表所有未选课的课程
		courses, CountPage := dao.SelectCourses(id, search_words, search_PageInt-1)

		//上
		c.HTML(http.StatusOK, "joinCourseTop.tmpl", gin.H{
			"id":          id,                     //学生id
			"loginCode":   loginCode,              //登录验证码
			"searchWords": search_words,           //搜索词
			"currentPage": search_PageInt,         //当前页数
			"PageInfo":    dao.ShowNumbersProxy(), //分页查询中每页显示的信息数
			"CountPage":   CountPage,              //总页数
		})

		//中
		for i := 0; i < len(courses); i++ {
			v := courses[i]
			c.HTML(http.StatusOK, "joinCourseMid.tmpl", gin.H{
				"c_loops":   v,
				"id":        id,        //学生id
				"loginCode": loginCode, //登录验证码
			})
			v = dao.Course{} //清除缓存
		}
		//下
		c.HTML(http.StatusOK, "joinCourseEnd.tmpl", gin.H{})
	})

	return r
}
