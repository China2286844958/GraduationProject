package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/***
@Author 薛智敏
@Title 设置信息的提示展示，为了更好的与用户交互
*/

//设置信息提示页面的信息

//c.HTML(http.StatusOK, "Message.tmpl", gin.H{
//	"title":   "登录失效",        //显示的标题
//	"url":     "/User/login", //按钮按下后，跳转的页面url
//	"BtnName": "返回登录",        //按钮显示的名字
//})

func SetMassage(c *gin.Context, title string, url string, BtnName string) {

	c.HTML(http.StatusOK, "Message.tmpl", gin.H{
		"title":   title,   //显示的标题
		"url":     url,     //按钮按下后，跳转的页面url
		"BtnName": BtnName, //按钮显示的名字
	})

	return
}
