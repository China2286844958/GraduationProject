package dao

import (
	"GraduationProject/tools"
	"database/sql"
	"fmt"
	"strconv"
)

//=========================增加=====================================

//【管理员】添加课程

func AddCourse(Cname string, Teacher string, Tscore string) (finished int64) {

	getTimeStamp, _ := tools.GetTimeStamp()

	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)
	defer db.Close()

	TscoreF64, err := strconv.ParseFloat(Tscore, 32)
	CheckErr(err)

	SQLCmd := fmt.Sprintf("insert %s values (%d,'%s','%s',%.2f)", Tb_Course, getTimeStamp, Cname, Teacher, TscoreF64)
	stmt, err := db.Prepare(SQLCmd)
	CheckErr(err)

	exec, err := stmt.Exec()
	CheckErr(err)

	finishedNum, err := exec.RowsAffected()

	return finishedNum
}

//=========================删除=====================================

//【管理员】根据课程id删除课程

func DeleteCourseByCid(Cid string) (finished int64) {
	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)
	defer db.Close()

	CidInt64, err := strconv.ParseInt(Cid, 10, 64)
	CheckErr(err)

	SQLCmd := fmt.Sprintf("delete from %s c where c.c_id = %d", Tb_Course, CidInt64)
	stmt, err := db.Prepare(SQLCmd)
	CheckErr(err)

	exec, err := stmt.Exec()
	CheckErr(err)

	finishedNum, err := exec.RowsAffected()

	return finishedNum

}

//=========================查询=====================================

//查询所有课程

func SelectAllCourse() (courses map[int]Course) {
	course := Course{}
	courseList := make(map[int]Course)

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)

	querySql := fmt.Sprintf("select * from %s c  order by c.c_name", Tb_Course)
	rows, err := sql.Query(querySql)
	CheckErr(err)

	count := 0

	for rows.Next() {
		rows.Scan(&course.C_id, &course.C_name, &course.C_teach, &course.C_Tscore)
		courseList[count] = course
		course = Course{} //清除缓存
		count++
	}

	return courseList
}

//查询课程表所有未选课的课程

func SelectCourses(id string, queryTxt string, searchPage int) (map[int]Course, int) {

	course := Course{}
	courseList := make(map[int]Course)

	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)
	db1, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)

	defer db.Close()
	defer db1.Close()
	CheckErr(err)

	//每页显示的信息数,作为N
	everyPageShows := ShowNumbersProxy()
	//当前页面,作为M
	currentPage := searchPage * everyPageShows

	id_int64, err := strconv.ParseInt(id, 10, 64)
	countSql := fmt.Sprintf("select * from %s c where (c.c_name like '%%%s%%' or c.c_teach like '%%%s%%' ) and (c.c_id  not in (SELECT c_id from stud_course sc WHERE sc.sd_id=%d ))  order by c.c_name", Tb_Course, queryTxt, queryTxt, id_int64)
	querySql := countSql + fmt.Sprintf(" limit %d,%d", currentPage, everyPageShows)

	rows, err := db.Query(querySql)
	CheckErr(err)
	rows1, err := db1.Query(countSql)
	CheckErr(err)

	count := 0

	for rows.Next() {
		rows.Scan(&course.C_id, &course.C_name, &course.C_teach, &course.C_Tscore)
		courseList[count] = course
		course = Course{} //清除缓存
		count++
	}
	//SQL2:查询所有普通管理员的个数
	count = 0
	for rows1.Next() {
		count++
	}
	countPage := 0
	if count%everyPageShows == 0 {
		countPage = count / everyPageShows
	} else {
		countPage = count/everyPageShows + 1
	}
	return courseList, countPage
}

//【管理员】根据课程名或者授课老师查询课程

func SelectCourseByCNameAndCTeach(searchWord string, searchPage int) (map[int]Course, int) {

	course := Course{}
	courseList := make(map[int]Course)

	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)
	db1, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)

	defer db.Close()
	defer db1.Close()

	//每页显示的信息数,作为N
	everyPageShows := ShowNumbersProxy()
	//当前页面,作为M
	currentPage := searchPage * everyPageShows

	countSql := fmt.Sprintf("select * from %s c where c.c_name like '%%%s%%' or c.c_teach like '%%%s%%' order by c.c_name", Tb_Course, searchWord, searchWord)
	querySql := countSql + fmt.Sprintf(" limit %d,%d", currentPage, everyPageShows)

	rows, err := db.Query(querySql)
	CheckErr(err)
	rows1, err := db1.Query(countSql)
	CheckErr(err)

	count := 0
	for rows.Next() {
		rows.Scan(&course.C_id, &course.C_name, &course.C_teach, &course.C_Tscore)
		courseList[count] = course
		course = Course{} //清除缓存
		count++
	}

	//SQL2:查询所有普通管理员的个数
	count = 0
	for rows1.Next() {
		count++
	}
	countPage := 0
	if count%everyPageShows == 0 {
		countPage = count / everyPageShows
	} else {
		countPage = count/everyPageShows + 1
	}

	return courseList, countPage
}

//======================================================================================================================
//根据课程id获取课程名称、课程授课老师-字段截取

type courseNameById struct {
	C_Id     int64   //课程 id
	C_Name   string  //课程名字
	C_Tearch string  //授课老师
	C_Tscore float32 //课程总分
}

//通过课程id查询

func SelectCourseByCourseId(id string) courseNameById {

	cnb := courseNameById{}
	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)
	defer db.Close()

	sqlCmd := fmt.Sprintf("SELECT * from course")

	//当id 有数值传入
	if id != "" {
		idInt64, _ := strconv.ParseInt(id, 10, 64)

		sqlCmd = fmt.Sprintf("SELECT * from course c where c.c_id = %d", idInt64)
	}

	rows, err := db.Query(sqlCmd)

	for rows.Next() {
		rows.Scan(&cnb.C_Id, &cnb.C_Name, &cnb.C_Tearch, &cnb.C_Tscore)
	}

	CheckErr(err)
	return cnb
}

type course struct {
	C_name string
}

// SelectAllCourseName
//
//	@Description: 查询所有课程的名字
//	@return list 返回课程名字的集合
func SelectAllCourseName() (list []string) {

	courseList := []string{}

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)

	querySql := fmt.Sprintf("SELECT c.c_name from %s c", Tb_Course)
	rows, err := sql.Query(querySql)
	CheckErr(err)

	var C_name string
	for rows.Next() {

		rows.Scan(&C_name)
		courseList = append(courseList, C_name)
		C_name = "" //清除缓存
	}

	return courseList
}

//=========================修改=====================================
