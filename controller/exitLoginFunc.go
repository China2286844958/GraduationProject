package controller

import (
	"GraduationProject/dao"
	"GraduationProject/tools"
	"github.com/gin-gonic/gin"
)

/**
@Title 用户退出登录的逻辑编写
@CreateTime 2022年7月13日23:15:11
*/

const (
	adminExit   = "adminExit"   //管理员退出登录路径
	studentExit = "studentExit" //学生退出登录路径
)

//用户退出登录逻辑，登录码禁用

func UserExitController(r *gin.Engine) *gin.Engine {

	exitMain := r.Group(RouterUserGroup)
	{
		//管理退出登录逻辑
		exitMain.GET(adminExit, func(c *gin.Context) {
			exitId := c.Query("exitId") //获取管理员id

			//执行数据库
			controller := dao.LoginCodeLockedController(dao.AdminLocked, exitId)
			if controller > 0 {
				tools.SetMassage(c, "成功退出登录", "/User/login", "重新登录")

			} else {
				tools.SetMassage(c, "退出登录执行错误,刷新后重试!", "/User/login", "重新登录")

			}
			return

		})

		//学生退出登录逻辑
		exitMain.GET(studentExit, func(c *gin.Context) {
			exitId := c.Query("exitId") //获取学生的学号id

			//执行数据库
			controller := dao.LoginCodeLockedController(dao.StudentLocked, exitId)

			if controller > 0 {
				tools.SetMassage(c, "成功退出登录", "/User/login", "重新登录")

			} else {

				tools.SetMassage(c, "退出登录执行错误,刷新后重试!", "/User/login", "重新登录")
			}
			return

		})

	}

	return r
}
