package main

import (
	"GraduationProject/controller"
	"GraduationProject/dao"
	"GraduationProject/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

/*
*
@Author 薛智敏
@Title 服务器调用
*/
const (
	serviceAddress = "127.0.0.1" //服务器访问地址，默认是127.0.0.1
	servicePort    = "8848"      //服务器端口
)

var mysqlSysFlag = false //调优标记

// StartLog
//
//	@Description: 开启日志记录
func StartLog() {
	// 获取日志文件句柄
	// 以 只写入文件|没有时创建|文件尾部追加 的形式打开这个文件
	currentLogFiles := serviceAddress + "_" + servicePort + ".log"
	logFile, err := os.OpenFile("./Logs/"+currentLogFiles, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//log.Println(err.Error())
	}
	// 设置存储位置
	log.SetOutput(logFile)
	log.Println("【地址" + serviceAddress + "端口" + servicePort + "】===========================日志开始执行=========================")

}

// init 初始化
//
//	@Description:
func init() {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~是否开启日志记录功能和计时业务功能?[输入start开启;其他关闭日志记录]~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	flag := "stop"
	fmt.Scanln(&flag)
	flag = strings.ToLower(flag)
	switch flag {
	case "start":
		log.Println("~_~ ~_~ ~_~ ~_~ ~_~ ~_~[日志记录中...]~_~ ~_~ ~_~ ~_~ ~_~")
		go StartLog() //加入日志记录
		go func() {
			for {
				h := time.Now().Hour()
				m := time.Now().Minute()
				s := time.Now().Second()

				//fmt.Print(h, ":", m, ":", s, "\t")
				//当新的一天的时候，执行的业务逻辑
				if h == 0 && m == h && s == m {
					//重置管理员的登录验证码
					dao.ResetAllAdminLoginCode()
					//重置学生的登录验证码

					dao.ResetAllStudentLoginCode()
					log.Println("重置登录验证码[已执行]")
				}

				time.Sleep(999 * time.Millisecond) //休眠一秒
			}
		}()
	default:
		log.Println("日志功能关闭")
	}

}
func main() {
	JoinGo()
}

var wg sync.WaitGroup
var wg_count = 0

// 加入协程
func JoinGo() {
	startTime := time.Now()
	r := gin.Default()

	//加载静态资源
	r.LoadHTMLGlob("statics/**/*")
	r.Static("/"+dao.Static_Folder_html, "statics/h5")     //映射加载html文件
	r.Static("/"+dao.Static_Folder_css, "statics/css")     //映射加载css文件
	r.Static("/"+dao.Static_Folder_img, "images/")         //映射加载图片资源文件，
	r.Static("/"+dao.Static_Folder_js, "statics/js")       //映射js文件
	r.Static("/"+dao.Static_Folder_upload, "UploadFiles/") //映射用户上传文件

	//服务器设置运行核心数
	cpuNum := runtime.NumCPU() //获取电脑核心数
	log.Println("******[本机CPU核数]:\t", cpuNum)

	// 设置占用cpu个数
	runCpu := tools.ThreeEyesReturnInt(cpuNum > 3, cpuNum/2, cpuNum)

	runtime.GOMAXPROCS(runCpu) //设置参与的核心数

	log.Println("******[设置的核心数]:\t", runCpu)
	//配置Mysql系统设置
	if !mysqlSysFlag {
		dao.MysqlSystemConfig()
		mysqlSysFlag = true
		log.Println("[MYSQL-调优]:[已开启]")
	}

	ServiceRun(r, startTime) //加入协程

	//======本机IP获取==========
	ip := tools.LocalIp()
	if ip == "<nil>" {
		ip = serviceAddress + "[本地]"
	}
	log.Println("\t------------------------------------[局域网访问]" + ip + ":" + servicePort + "-------------------------------------")
	SqlLog := func() string {
		if mysqlSysFlag {
			return "[已开启]"
		}
		return "[未开启]"
	}
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ServerInfo.html", gin.H{
			"CPU_Cores":     cpuNum,                            //本地CPU核心数量
			"Running_Cores": runCpu,                            //使用的核心数量
			"SQL_optimize":  SqlLog(),                          //sql优化
			"Localhost":     ip,                                //本地IP
			"ServerPort":    servicePort,                       //服务器端口
			"wps_doc":       "https://kdocs.cn/l/cm6g3St2xdBH", //网站开发详情的WPS云文档
		})

	})

	//=============加载banner=============
	log.Println("\n", tools.Banner_1)
	fmt.Println(tools.Banner_1)
	log.Println("加载资源耗时:", time.Since(startTime))

	//===============自动跳转==================
	tools.AutoHref("http://" + serviceAddress + ":" + servicePort + controller.RouterUserGroup + controller.UserGroupPathLogin)

	//================运行================
	r.Run(":" + servicePort) //运行

}

func ServiceRun(r *gin.Engine, startTime time.Time) {
	//=========================================

	wg.Add(29) //添加N个协程
	//========注册功能=======1

	go func() {
		r = controller.UserRegisterController(r) //注册
		wg.Done()
	}()

	//========登录功能=======2
	go func() {
		r = controller.UserLoginController(r) //登录
		wg.Done()
	}()

	//=========用户界面==============3
	go func() {
		r = controller.UserIndexController(r)
		wg.Done()
	}()
	//=========用户退出登录==============4
	go func() {
		controller.UserExitController(r)
		wg.Done()
	}()

	//===================================================学生端====================================

	//=========个人信息界面-信息展示===============5

	go func() {
		r = controller.PrivateInfoShowController(r)
		wg.Done()
	}()

	//===========个人信息界面-修改个人信息===========6

	go func() {
		r = controller.PrivateInfoUpdate(r)
		wg.Done()
	}()

	//=========	我的课程【学生端】===============7

	go func() {
		r = controller.MyCourse(r)
		wg.Done()
	}()

	//=========学生退课=======================8

	go func() {
		r = controller.StudExitCourseController(r)
		wg.Done()
	}()

	//=============加入课程===========9
	go func() {
		r = controller.JoinCourseController(r)
		wg.Done()
	}()

	//=============学生选课===========10
	go func() {
		r = controller.StudAddCourseController(r)
		wg.Done()
	}()

	//===========修改密码==============11
	go func() {
		r = controller.UpdatePasswdController(r)
		wg.Done()
	}()

	//===================================================管理员====================================
	//===============管理员个人界面================12
	go func() {
		r = controller.PersonalManageController(r)
		wg.Done()
	}()

	//==============管理员-选课管理=========13
	go func() {
		r = controller.JoinCourseManageController(r)
		wg.Done()
	}()

	//==============管理员-课程管理=========14
	go func() {
		r = controller.CourseManager(r)
		wg.Done()
	}()

	//==============管理员-选课管理--学生管理=========15
	go func() {
		r = controller.StudentMangeController(r)
		wg.Done()
	}()

	//==============管理员-选课管理--学生管理：密码重置=========16
	go func() {
		r = controller.StudResetPasswdController(r)
		wg.Done()
	}()

	//==============管理员-选课管理--学生管理：批改分数=========17
	go func() {
		r = controller.GradesController(r)
		wg.Done()
	}()
	//==============管理员-选课管理--学生管理：注销此号=========18
	go func() {
		r = controller.DeleteStudController(r)
		wg.Done()
	}()
	//==============管理员-选课管理--课程管理：删除课程=========19
	go func() {
		r = controller.DeleteCourseController(r)
		wg.Done()
	}()
	//==============管理员-选课管理--课程管理：发布课程=========20
	go func() {
		r = controller.PublishCourse(r)
		wg.Done()
	}()

	//==============管理员-选课管理--修改密码=========21
	go func() {
		r = controller.AdminResetPasswd(r)
		wg.Done()
	}()

	//==============超级管理员添加管理员===================22
	go func() {
		r = controller.AddAdminController(r)
		wg.Done()
	}()

	//==============超级管理员-删除普通管理员的逻辑操作=========23
	go func() {
		r = controller.SuDeleteAdmin(r)
		wg.Done()
	}()

	//==============超级管理员-重置普通管理员密码的逻辑操作=========24
	go func() {
		r = controller.SuResetPasswdAdmin(r)
		wg.Done()
	}()

	//===========================学生-忘记密码=============25
	go func() {
		r = controller.FindPasswd(r)
		wg.Done()
	}()
	//===================测试======================26
	//测试功能:发送验证码
	go func() {
		r = controller.Tests(r) //
		wg.Done()
	}()
	//======================发送验证码=====================27
	go func() {
		r = controller.SendEmailByNetwork(r) //
		wg.Done()
	}()

	//======================首页=====================28
	go func() {
		r = controller.FirstPageController(r) //
		wg.Done()
	}()
	//======================更改头像=====================29

	go func() {
		r = controller.UpLoadFile(r) //
		wg.Done()
	}()
	//UpLoadFile

	wg.Wait()

}
