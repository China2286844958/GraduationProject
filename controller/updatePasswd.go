package controller

import (
	"GraduationProject/dao"
	"GraduationProject/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@Title 修改密码业务逻辑
*/

func UpdatePasswdController(r *gin.Engine) *gin.Engine {

	//学生端-GET响应修改的请求
	r.GET(RouterUserGroup+"/updatePasswd", func(c *gin.Context) {

		id := c.Query("id")
		loginCode := c.Query("loginCode")
		//登录验证码验证
		loginTrue := dao.TrueLoginByEmailAdLoginCodes(id, loginCode)
		if loginTrue == false {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		c.HTML(http.StatusOK, "updatePasswd.tmpl", gin.H{
			"id":        id,
			"loginCode": loginCode,
		})

	})

	//POST处理
	r.POST(RouterUserGroup+"/updatePasswd", func(c *gin.Context) {
		old_passwd := c.PostForm("old_passwd")
		new1_passwd := c.PostForm("new1_passwd")
		new2_passwd := c.PostForm("new2_passwd")
		id := c.PostForm("id")
		loginCode := c.PostForm("loginCode")

		oldSize := len(old_passwd)
		new1Size := len(new1_passwd)
		new2Size := len(new2_passwd)

		//密码长度判断为空，不符合长度规范的
		if (oldSize < 4 || oldSize > 30) || (new1Size < 4 || new1Size > 30) || (new2Size < 4 || new2Size > 30) {

			//c.Redirect(http.StatusMovedPermanently, "updatePasswd?id="+id+"&loginCode="+loginCode)
			c.HTML(http.StatusOK, "updatePasswd.tmpl", gin.H{
				"id":        id,
				"loginCode": loginCode,
				"errLog":    "请输入4到30位的密码！",
			})
			return
		}

		//两次的密码不一致
		if new1_passwd != new2_passwd {

			//c.Redirect(http.StatusOK, "updatePasswd?id="+id+"&loginCode="+loginCode)
			c.HTML(http.StatusOK, "updatePasswd.tmpl", gin.H{
				"id":        id,
				"loginCode": loginCode,
				"errLog":    "两次的密码不一致",
			})

			return
		}

		//三次密码都一样
		if old_passwd == new1_passwd && old_passwd == new2_passwd {

			//c.Redirect(http.StatusOK, "updatePasswd?id="+id+"&loginCode="+loginCode)
			c.HTML(http.StatusOK, "updatePasswd.tmpl", gin.H{
				"id":        id,
				"loginCode": loginCode,
				"errLog":    "三次密码一致！",
			})

			return
		}

		//旧密码错误
		//对旧密码进行校验，先加密，后检索
		oldPasswdIncode := tools.GetSHA256HashCode(old_passwd)
		old_passwdFlag := dao.SelectTableByColumn(dao.Db_Course_select_sys, dao.Tb_Stud_login, "sl_passwd", oldPasswdIncode)

		if old_passwdFlag == true {

			//数据库执行更新操作
			PasswdEncode := tools.GetSHA256HashCode(new2_passwd)
			logs := dao.UpdatePasswdById(id, PasswdEncode)

			//判断密码是否更新成功！
			if logs != "null" {
				c.HTML(http.StatusOK, "updatePasswd.tmpl", gin.H{
					"id":        id,
					"loginCode": loginCode,
					"errLog":    logs,
				})

			}

			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "更新成功，请重新登录",  //显示的标题
				"url":     "/User/login", //按钮按下后，跳转的页面url
				"BtnName": "返回登录",        //按钮显示的名字
			})
			return
		} else {
			c.HTML(http.StatusOK, "updatePasswd.tmpl", gin.H{
				"id":        id,
				"loginCode": loginCode,
				"errLog":    "原密码不正确！",
			})

			return
		}

	})
	return r

}
