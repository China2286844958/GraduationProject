package dao

/*表结构设计*/

//学生登录表

type Stud_login struct {
	Sl_id        int64
	Sl_email     string
	Sl_passwd    string
	Sl_loginCode string
	Sl_logo      string
}

//学生详情表

type Stud_del struct {
	Sd_id      int64
	Sd_relname string
	Sd_gender  string
	Sd_age     int
	Sd_address string
	Sd_sys     string
}

//课程表

type Course struct {
	C_id     int
	C_name   string
	C_teach  string
	C_Tscore float32
}

//学生的选课

type StudCourse struct {
	SC_id     int     //选课的id
	C_id      int64   //课程id
	Sd_id     int64   //学生id
	Sd_Sscore float32 //学生得分

}

//截取字段

type StudCourses struct {
	SC_id     int
	SC_name   string  //课程名字
	SC_teach  string  //授课老师
	SC_Sscore float32 //学生得分
	SC_Tscore float32 //课程总分
}

//管理员表结构

type Admin struct {
	A_id        int64  //管理员id
	A_email     string //管理员邮箱
	A_passwd    string //管理员密码
	A_power     int    //管理员权限
	A_loginCode string //管理员登录码
}

//管理员-选课管理-学生管理截取字段

type StudJoinedCourse struct {
	S_id       string  //学号
	S_name     string  //姓名
	S_Sscore   float64 //学生得分
	S_Tscore   float64 //课程总分
	S_CourseId int64   //课程id
	Sl_logo    string  //学生头像
}
