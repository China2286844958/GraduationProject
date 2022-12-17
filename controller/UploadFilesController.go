package controller

import (
	"GraduationProject/dao"
	"GraduationProject/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

/*
*
用户上传文件控制器
*/
var upload = "/upload"
var student = "/student"
var defaultLogo = "/UploadFiles/logo.gif" //默认的头像

func UpLoadFile(r *gin.Engine) *gin.Engine {
	uploadGroup := r.Group(upload)
	{
		uploadGroup.POST(student, func(c *gin.Context) {
			id := c.PostForm("id")
			loginCode := c.PostForm("loginCode")
			oldFile := c.PostForm("oldFile")
			//登录验证码验证
			loginTrue := dao.TrueLoginByEmailAdLoginCodes(id, loginCode)
			if !loginTrue {
				c.Redirect(http.StatusMovedPermanently, "login?LoginNoValue=true")
				return
			}

			file, er := c.FormFile("newFile")
			filename := file.Filename
			//检测文件是否有后缀,没有后缀.返回-1，后缀最后一次出现的索引
			index := strings.LastIndex(filename, ".")
			uri := "" //头像格式
			if index == -1 {
				c.HTML(http.StatusOK, "Message.tmpl", gin.H{
					"title":   "文件非法!", //显示的标题
					"url":     "back",  //按钮按下后，跳转的页面url
					"BtnName": "返回",    //按钮显示的名字
				})
				return
			} else if index != -1 { //当有后缀,截取后缀

				uri = filename[index:] //格式
				//转小写
				uri = strings.ToLower(uri)
				if uri == ".png" || uri == ".jpeg" || uri == ".jpg" || uri == ".ico" || uri == ".gif" {

				} else {
					c.HTML(http.StatusOK, "Message.tmpl", gin.H{
						"title":   "文件格式不符合【png,jpeg,jpg,ico,gif】!", //显示的标题
						"url":     "back",                           //按钮按下后，跳转的页面url
						"BtnName": "返回",                             //按钮显示的名字
					})
					return
				}

			}
			//文件名校验以及大小校验
			imgSize := file.Size / 1024
			if imgSize > 1024 {
				c.HTML(http.StatusOK, "Message.tmpl", gin.H{
					"title":   "文件过大，文件不能超过 1 MB", //显示的标题
					"url":     "back",             //按钮按下后，跳转的页面url
					"BtnName": "返回",               //按钮显示的名字
				})
				return
			}

			if er != nil {
				log.Println(er.Error())
			}
			//保存图片到本地
			//dst := "./UploadFiles/" + id + ".png"
			//获取时间戳
			stamp, err := tools.GetTimeStamp()
			if err {
				return
			}
			timeStr := fmt.Sprintf("%d", stamp)

			//当提交的头像不是旧的,也不是之前的，执行删除旧
			//if filename != oldFile && filename != defaultLogo {
			if oldFile != defaultLogo && filename != oldFile {
				er := tools.RemoveLogo(oldFile) //执行删除本地资源的I/O操作
				if er != nil {
					dao.UpdateStuLogoByIdAndLoginCode(id, loginCode, defaultLogo)
					log.Println("旧头像" + file.Filename + "删除失败！")
					c.HTML(http.StatusOK, "Message.tmpl", gin.H{
						"title":   "旧头像" + file.Filename + "删除失败！", //显示的标题
						"url":     "back",                          //按钮按下后，跳转的页面url
						"BtnName": "返回",                            //按钮显示的名字
					})
					return
				}
			}

			dst := "./" + dao.Static_Folder_upload + "/" + id + "_" + timeStr + uri
			FileFlag := c.SaveUploadedFile(file, dst)
			dst = dst[1:]
			//数据库更新
			num := dao.UpdateStuLogoByIdAndLoginCode(id, loginCode, dst)
			//判断
			if FileFlag != nil || num < 0 {

				c.HTML(http.StatusOK, "Message.tmpl", gin.H{
					"title":   "上传失败,请稍后再试！", //显示的标题
					"url":     "back",        //按钮按下后，跳转的页面url
					"BtnName": "返回",          //按钮显示的名字
				})
				return
			}
			c.HTML(http.StatusOK, "Message.tmpl", gin.H{
				"title":   "上传成功", //显示的标题
				"url":     "back", //按钮按下后，跳转的页面url
				"BtnName": "返回",   //按钮显示的名字
			})

		})

	}
	return r
}
