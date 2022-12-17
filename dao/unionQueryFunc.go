package dao

import (
	"database/sql"
	"fmt"
	"strconv"
)

/**
联合查询
*/

//=======================【Stud_login 和 Stud_del】======================
/*多表查询学生所有信息【查询学生所有信息】*/

func QueryUnionStudentById(id string) (Stud_login, Stud_del) {

	//=======================查询Stud_login表=======================

	studLogin := Stud_login{}

	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))

	CheckErr(err)

	id_int, err := strconv.Atoi(id)
	querySql := fmt.Sprintf("select sl.sl_id,sl.sl_email,sl.sl_passwd,sl.sl_loginCode from %s sl where sl.sl_id = %d ", Tb_Stud_login, id_int)
	rows, err1 := db.Query(querySql)
	CheckErr(err1)

	for rows.Next() {
		rows.Scan(&studLogin.Sl_id, &studLogin.Sl_email, &studLogin.Sl_passwd, &studLogin.Sl_loginCode)

	}
	db.Close()

	//========================查询Stud_del表==========================
	studDel := Stud_del{}

	db2, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer db2.Close()
	querySql2 := fmt.Sprintf("select * from %s where sd_id = %d ", Tb_Stud_del, studLogin.Sl_id)

	rows2, err1 := db2.Query(querySql2)

	CheckErr(err1)

	for rows2.Next() {
		rows2.Scan(&studDel.Sd_id, &studDel.Sd_relname, &studDel.Sd_gender, &studDel.Sd_age, &studDel.Sd_address, &studDel.Sd_sys)
	}

	return studLogin, studDel
}

//管理员- 选课管理 [学号 姓名  课程id ]>学生学号>学生姓名>得分

func QueryUnionStudByJoinedCourse(courseId string, searchWord string, searchPage int) (list map[int]StudJoinedCourse, countPages int) {

	SjoinedCourse := StudJoinedCourse{}
	SjoinedCourseList := make(map[int]StudJoinedCourse)

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

	//定义SQL语句
	querySql := ""   //执行Sql语句
	SqlNoLimit := "" //不加Limit的SQL语句
	// := "" //查询数据Sql语句

	//当课程id搜索所有，就查看所有学生
	if courseId == "ALL" {

		SqlNoLimit = fmt.Sprintf(" SELECT sd.sd_id,sd.sd_relname,sl.sl_logo from stud_del sd INNER JOIN stud_login sl ON sd.sd_id=sl.sl_id where sd.sd_id like '%%%s%%' or sd.sd_relname like '%%%s%%'  ", searchWord, searchWord)
		querySql = SqlNoLimit + fmt.Sprintf(" limit %d,%d", currentPage, everyPageShows)
	} else {
		parseInt, _ := strconv.ParseInt(courseId, 10, 64)
		//id 有的时候:指定课程查询
		SqlNoLimit = fmt.Sprintf(" SELECT sc.sd_id,sd.sd_relname,sc.sd_Sscore,c.c_Tscore,c.c_id,sl.sl_logo from stud_course sc,course c,stud_del sd,stud_login sl  where  sc.c_id=%d and sc.c_id=c.c_id and sc.sd_id =sd.sd_id and sl.sl_id=sd.sd_id and (sd.sd_relname LIKE '%%%s%%' or sc.sd_id like '%%%s%%') order by sc.sd_Sscore desc ", parseInt, searchWord, searchWord)
		querySql = SqlNoLimit + fmt.Sprintf(" limit %d,%d", currentPage, everyPageShows)

	}

	rows, err := db.Query(querySql)
	CheckErr(err)
	rows1, err := db1.Query(SqlNoLimit)
	CheckErr(err)

	count := 0

	for rows.Next() {

		//当查询了所有选课
		if courseId == "ALL" {
			//课程得分  课程总分  课程id 不传入数据，因为技术有限
			rows.Scan(&SjoinedCourse.S_id, &SjoinedCourse.S_name, &SjoinedCourse.Sl_logo)

		} else {
			rows.Scan(&SjoinedCourse.S_id, &SjoinedCourse.S_name, &SjoinedCourse.S_Sscore, &SjoinedCourse.S_Tscore, &SjoinedCourse.S_CourseId, &SjoinedCourse.Sl_logo)
		}

		SjoinedCourseList[count] = SjoinedCourse
		SjoinedCourse = StudJoinedCourse{} //清除缓存
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

	return SjoinedCourseList, countPage
}

// SelectMaxScoreByCourseName
//
//	@Description: 根据课程名查询对应的最高分
//	@param courseNameList 课程名字集合
//	@return maxScore

func SelectMaxScoreByCName(courseNameList []string) (maxList []float32) {
	list := []float32{}

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)

	//=======
	var maxScore float32 = 0
	for i := 0; i < len(courseNameList); i++ {
		querySql := fmt.Sprintf("SELECT MAX(sd_Sscore) as maxScore from %s sc, %s  c WHERE sc.c_id=c.c_id and c.c_name='%s' ", Tb_Stud_course, Tb_Course, courseNameList[i])
		rows, err := sql.Query(querySql)
		CheckErr(err)

		for rows.Next() {
			rows.Scan(&maxScore)
			list = append(list, maxScore)
			maxScore = 0 //清除
		}
		rows.Close()

	}

	return list
}

// SelectMinScoreByCName
//	@Description:根据课程名查询对应的最低分
//	@param courseNameList 课程名字集合
//	@return minList

func SelectMinScoreByCName(courseNameList []string) (minList []float32) {

	list := []float32{}

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)

	//=======
	var minScore float32 = 0
	for i := 0; i < len(courseNameList); i++ {
		querySql := fmt.Sprintf("SELECT MIN(sd_Sscore) as maxScore from %s sc, %s  c WHERE sc.c_id=c.c_id and c.c_name='%s' ", Tb_Stud_course, Tb_Course, courseNameList[i])
		rows, err := sql.Query(querySql)
		CheckErr(err)

		for rows.Next() {
			rows.Scan(&minScore)
			list = append(list, minScore)
			minScore = 0 //清除
		}
		rows.Close()

	}

	return list
}
