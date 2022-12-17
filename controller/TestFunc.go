package controller

import (
	"GraduationProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@Title 试验区
*/

//遍历数据

func Tests(r *gin.Engine) *gin.Engine {

	var test2 = "/test"
	nameList := dao.SelectAllCourseName() //查询所有课程名字 return []string

	r.GET(test2+"/:file", func(c *gin.Context) {

		c.HTML(http.StatusOK, "test.tmpl", gin.H{
			"courseNameList": nameList, //课程名字集合
		})
	})
	return r
}
