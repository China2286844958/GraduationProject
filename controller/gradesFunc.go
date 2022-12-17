package controller

import (
	"GraduationProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
@Title 【管理员】为学生批改分数
@Author 薛智敏
@CreatTime 2022年6月21日00:48:03
*/

//管理员-对指定的学生课程评分

func GradesController(r *gin.Engine) *gin.Engine {

	var grades = "/grades"

	//处理get请求
	r.GET(master+grades, func(c *gin.Context) {

		id := c.Query("id")
		loginCode := c.Query("loginCode")
		sid := c.Query("sid") //学生学号
		cid := c.Query("cid") //课程号

		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(id, loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		//根据课程id获取课程信息
		courseById := dao.SelectCourseByCourseId(cid) //查询所有课程，返回课程结构体

		studDel := dao.QueryStudCourseBySid(sid)

		//	执行数据库命令

		c.HTML(http.StatusOK, "adminGrades.tmpl", gin.H{
			"cbi":       courseById,        //课程信息
			"id":        id,                //id
			"loginCode": loginCode,         //登录验证码
			"sid":       sid,               //学号
			"c":         courseById,        //课程
			"Sscore":    studDel.Sd_Sscore, //学生得分

		})

	})

	//处理评分post请求
	r.GET(master+grades+"Post", func(c *gin.Context) {

		id := c.Query("id")
		loginCode := c.Query("loginCode")
		sid := c.Query("sid")           //学生学号
		cid := c.Query("cid")           //课程号
		score := c.Query("score")       //获取用户输入的评分
		maxScore := c.Query("maxScore") //最大分数限制，课程总分

		//当没有数据传输，设置为0
		if score == "" {
			score = "0"
		}

		//输入分数极值限制

		//转化
		InScoreF32, err := strconv.ParseFloat(score, 32)
		dao.CheckErr(err)

		MaxScoreF32, err := strconv.ParseFloat(maxScore, 32)
		dao.CheckErr(err)

		//最大值限制
		if InScoreF32 > MaxScoreF32 {
			InScoreF32 = MaxScoreF32
		}

		//最小值限制
		if InScoreF32 < 0.00 {
			InScoreF32 = 0.00
		}

		//管理员登录验证码校验
		T := dao.AdminLoginCodeVerification(id, loginCode)
		if !T {
			c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
			return
		}

		//当评分不为空，说明评分了
		if score != "" {
			//	执行数据库
			finished := dao.GradesBySidAndCid(sid, cid, InScoreF32)

			//有标记，代表成功
			if finished > 0 {
				//	返回提示
				c.HTML(http.StatusOK, "Message.tmpl", gin.H{
					"title":   "保存成功！", //显示的标题
					"url":     "back",  //按钮按下后，跳转的页面url
					"BtnName": "返回",    //按钮显示的名字
				})
			} else {
				//	返回提示
				c.HTML(http.StatusOK, "Message.tmpl", gin.H{
					"title":   "保存失败！稍后再试", //显示的标题
					"url":     "back",      //按钮按下后，跳转的页面url
					"BtnName": "返回",        //按钮显示的名字
				})
			}
		}

	})

	return r
}
