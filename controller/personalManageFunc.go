package controller

import (
	"GraduationProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var personalManage = "/personalManage"

//人员管理

func PersonalManageController(r *gin.Engine) *gin.Engine {

	//请求
	r.GET(master+personalManage, func(c *gin.Context) {

		id := c.Query("id")
		loginCode := c.Query("loginCode")
		search_words := c.Query("search_words")
		search_Page := c.Query("search_Page") //搜索的页数

		search_PageInt, errP := strconv.Atoi(search_Page)
		searchPageErr := dao.CheckErr(errP)
		if searchPageErr || search_PageInt <= 0 || search_PageInt >= 10000 {
			search_PageInt = 1 //初始是0页
		}

		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(id, loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}
		admin := dao.GetAdminByIdAndLoginCode(id, loginCode)
		if admin.A_power < 10 {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		//查询所有管理员
		adminPower1List, CountPage := dao.GetAllAdminPower1(search_words, search_PageInt-1)
		//上
		c.HTML(http.StatusOK, "personalManageTop.tmpl", gin.H{
			"id":          admin.A_id,             //id
			"loginCode":   admin.A_loginCode,      //登录验证码
			"searchWords": search_words,           //搜索词
			"currentPage": search_PageInt,         //当前页数
			"PageInfo":    dao.ShowNumbersProxy(), //分页查询中每页显示的信息数
			"CountPage":   CountPage,              //总页数

		})

		//中
		adminClass := "超级管理员"
		for i := 0; i < len(adminPower1List); i++ {
			admin := adminPower1List[i]
			if admin.A_power == 1 {
				adminClass = "普通管理员"

			}
			c.HTML(http.StatusOK, "personalManageMid.tmpl", gin.H{
				"admin":       admin,          //管理员
				"adminClass":  adminClass,     //管理员类型 1 为普通管理员;2 为超级管理员
				"Su_id":       id,             //超级管理员的id
				"loginCode":   loginCode,      //登录验证码
				"currentPage": search_PageInt, //当前页数

			})
		}

		//下
		c.HTML(http.StatusOK, "personalManageEnd.tmpl", gin.H{})
	})

	return r
}
