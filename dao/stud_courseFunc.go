package dao

import (
	"GraduationProject/tools"
	"database/sql"
	"fmt"
	"strconv"
)

//==================================================添加================================================================

//根据学生id,课程id产生关联

func AddCourseBySidAndCid(Sid string, Cid string) (NumRows int64) {
	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)
	stamp, _ := tools.GetTimeStamp()
	Cid_int64, err := strconv.ParseInt(Cid, 10, 64) //课程id转int64
	Sid_int64, err := strconv.ParseInt(Sid, 10, 64) //课程id转int64
	CheckErr(err)

	sqlCmd := fmt.Sprintf("insert %s (sc_id,c_id,sd_id)  value (%d,%d,%d); ", Tb_Stud_course, stamp/100, Cid_int64, Sid_int64)
	prepare, err := sql.Prepare(sqlCmd)
	result, err := prepare.Exec()
	rowsAffected, err := result.RowsAffected()

	return rowsAffected
}

//==================================================删除================================================================

//根据学生id以及退课的课程id，删除课程

func StudExitCourseBySid(sdId string, Cid string) (RowsNum int64) {

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)
	defer sql.Close()
	cId_int64, err := strconv.ParseInt(Cid, 10, 64)
	sdId_int64, err := strconv.ParseInt(sdId, 10, 64)
	delSql := fmt.Sprintf("delete from stud_course sc where sc.sd_id = %d and sc.c_id = %d", sdId_int64, cId_int64)
	stmt, err := sql.Prepare(delSql)

	CheckErr(err)
	exec, err := stmt.Exec()

	num, err := exec.RowsAffected()

	CheckErr(err)
	return num
}

//==================================================查询================================================================

//根据学生id查询选课表

func QueryStudCourseBySid(Sid string) StudCourse {

	sc := StudCourse{}

	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)
	defer db.Close()

	SidInt64, _ := strconv.ParseInt(Sid, 10, 64)

	sqlCmd := fmt.Sprintf("select * from %s sc where sc.Sd_id=%d", Tb_Stud_course, SidInt64)
	rows, err := db.Query(sqlCmd)

	for rows.Next() {
		rows.Scan(&sc.SC_id, &sc.C_id, &sc.Sd_id, &sc.Sd_Sscore)
	}

	return sc
}

//根据学生id查询选择的课程

func QueryCourseByStudId(id string, searchData string, searchPage int) (map[int]StudCourses, int) {

	courseList := make(map[int]StudCourses) //返回的集合
	course := StudCourses{}                 //存放课程单个
	count := 0                              //计数器

	db1, err2 := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err2)
	db2, err2 := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err2)

	defer db1.Close()
	defer db2.Close()

	IdInt64, err := strconv.ParseInt(id, 10, 64)
	CheckErr(err)

	//每页显示的信息数,作为N
	everyPageShows := ShowNumbersProxy()
	//当前页面,作为M
	currentPage := searchPage * everyPageShows

	CountSql := fmt.Sprintf("(select c.c_id,c.c_name,c.c_teach,sc.sd_Sscore,c.c_Tscore from stud_del sd , stud_course sc , course c WHERE sc.sd_id = sd.sd_id and sc.c_id = c.c_id and sc.sd_id = %d and (c.c_name like '%%%s%%' or c.c_teach like '%%%s%%') group by c.c_id) order by c.c_name asc", IdInt64, searchData, searchData)
	sqlcmd := CountSql + fmt.Sprintf(" limit %d,%d", currentPage, everyPageShows)
	rows, err2 := db1.Query(sqlcmd)
	CheckErr(err2)
	rows1, err2 := db2.Query(CountSql)
	CheckErr(err2)

	for rows.Next() {

		err2 := rows.Scan(&course.SC_id, &course.SC_name, &course.SC_teach, &course.SC_Sscore, &course.SC_Tscore)
		CheckErr(err2)

		courseList[count] = course //装进map
		course = StudCourses{}     //清除
		count++
	}
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

//跟据学生id ，课程id查询课程

func GetStudCourseBySidAndCid(Sid string, Cid string) (finished int64) {
	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)
	defer db.Close()

	SidInt64, _ := strconv.ParseInt(Sid, 10, 64)
	CidInt64, _ := strconv.ParseInt(Cid, 10, 64)

	sqlCmd := fmt.Sprintf("select * from %s sc where sc.c_id=%d and sc.sd_id=%d", Tb_Stud_course, CidInt64, SidInt64)
	rows, err := db.Query(sqlCmd)

	for rows.Next() {
		return 1
	}

	return 0
}

//==================================================修改================================================================

//【管理员】 给学生的选课评分

func GradesBySidAndCid(sid string, cid string, score float64) (finished int64) {

	Step := GetStudCourseBySidAndCid(sid, cid)

	//当查询不到
	if Step == 0 {
		return 0
	}

	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)
	defer db.Close()

	SidInt64, _ := strconv.ParseInt(sid, 10, 64)
	CidInt64, _ := strconv.ParseInt(cid, 10, 64)

	sqlCmd := fmt.Sprintf("UPDATE stud_course sc set sc.sd_Sscore=%.2f where sc.c_id=%d  and sc.sd_id=%d", score, CidInt64, SidInt64)
	stmt, err := db.Prepare(sqlCmd)
	CheckErr(err)

	exec, err := stmt.Exec()
	CheckErr(err)

	Finished, err := exec.RowsAffected()

	return Finished + Step
}
