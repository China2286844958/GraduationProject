package dao

import (
	"database/sql"
	"fmt"
	"strconv"
)

/**
Stud_login

*/

//=========================增加=====================================

//学生登录表，添加学生

func AddStudent(stu Stud_login) {

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))

	defer sql.Close()
	CheckErr(err)

	addSql := fmt.Sprintf("insert %s (sl_id,sl_email,sl_passwd,sl_loginCode) values (%d,'%s','%s','%s')", Tb_Stud_login, stu.Sl_id, stu.Sl_email, stu.Sl_passwd, stu.Sl_loginCode)
	fmt.Println(addSql)
	prepare, err := sql.Prepare(addSql)
	CheckErr(err)

	exec, err := prepare.Exec()
	CheckErr(err)
	_, err = exec.LastInsertId()
	CheckErr(err)
}

//=========================删除=====================================
//=========================查询=====================================

//获取学生登录表

func GetStudentLoginByEmail(email string) Stud_login {
	studLogin := Stud_login{}

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)

	sqlCmd := fmt.Sprintf("select sl.sl_id,sl.sl_email,sl.sl_passwd,sl.sl_loginCode from stud_login sl where sl.sl_email='%s'", email)

	rows, err := sql.Query(sqlCmd)
	CheckErr(err)

	for rows.Next() {
		rows.Scan(&studLogin.Sl_id, &studLogin.Sl_email, &studLogin.Sl_passwd, &studLogin.Sl_loginCode)
	}
	return studLogin
}

//学生端-根据id和网页登录状态码查询学生登录状态码，判断是真登录还是伪登录

func TrueLoginByEmailAdLoginCodes(id string, loginCodes string) (loginTrue bool) {
	if loginCodes == LockedWord { //当等于锁登词，直接返回false
		return false
	}

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)

	//字符串转化int64,无损
	id_int, err := strconv.ParseInt(id, 10, 64)
	CheckErr(err)
	cmd := fmt.Sprintf("select * from %s where sl_id = %d and sl_loginCode = '%s'", Tb_Stud_login, id_int, loginCodes)
	rows, err := sql.Query(cmd)
	if err != nil {
		return false
	}
	for rows.Next() {
		return true
	}

	return false
}

//登录校验,根据email和加密后的数据查找

func QueryByEmailAndPasswd(email string, passwdEncode string) bool {
	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))

	CheckErr(err)
	defer db.Close()

	querySql := fmt.Sprintf("select * from %s where sl_email ='%s' and sl_passwd = '%s' ", Tb_Stud_login, email, passwdEncode)
	rows, err1 := db.Query(querySql)

	CheckErr(err1)

	for rows.Next() {
		return true

	}
	return false
}

//
//  QueryStudLoginByEmail 通过邮箱和验证码查询学生表
//  @Description:
//  @param Email 学生的邮箱
//  @param captchaCode 验证码
//  @return exist 执行数据库后，存在返回true，不存在返回false
//

func QueryStudLoginByEmail(Email string) (exist bool) {
	//if captchaCode == LockedWord { //当等于锁登词，直接返回false
	//	return false
	//}

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)

	cmd := fmt.Sprintf("select sl.sl_email,sl.sl_loginCode from %s sl where sl.sl_email = '%s' ", Tb_Stud_login, Email)
	rows, err := sql.Query(cmd)
	if err != nil {
		return false
	}
	for rows.Next() {
		return true
	}

	return false
}

func QueryStuLoginByIdAndLoginCode(id string, loginCode string) Stud_login {
	studLogin := Stud_login{}

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

	}

	sqlCmd := fmt.Sprintf("select sl.sl_id,sl.sl_email,sl.sl_passwd,sl.sl_loginCode,sl.sl_logo from stud_login sl where sl.sl_id='%d' and sl.sl_loginCode='%s'", idInt, loginCode)

	rows, err := sql.Query(sqlCmd)
	CheckErr(err)

	for rows.Next() {
		rows.Scan(&studLogin.Sl_id, &studLogin.Sl_email, &studLogin.Sl_passwd, &studLogin.Sl_loginCode, &studLogin.Sl_logo)
	}
	return studLogin
}

//=========================修改=====================================

//设置学生登录状态

func SetLoginCode(id int64, updateLoginCode string) {

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)
	sqlCmd := fmt.Sprintf("update %s  set sl_loginCode ='%s' where Sl_id =%d", Tb_Stud_login, updateLoginCode, id)
	result, err := sql.Exec(sqlCmd)
	CheckErr(err)
	_, err = result.RowsAffected()
	CheckErr(err)
}

//根据id，修改密码，错误信息返回到日志

func UpdatePasswdById(id string, newPasswd string) (Logs string) {

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)
	id_int64, err := strconv.ParseInt(id, 10, 64)
	cmd := fmt.Sprintf("update %s set sl_passwd = '%s' where sl_id =%d", Tb_Stud_login, newPasswd, id_int64)
	result, err := sql.Exec(cmd)

	CheckErr(err)

	//根据执行sql语句判断影响的结果，结果大于0，则成功执行
	num, err := result.RowsAffected()
	if num == 0 {
		return "密码更新失败"
	}

	return "null"
}

//  UpdatePasswdByEmail 根据邮箱修改学生的密码
//  @Description:
//  @param email 邮箱
//  @param newPasswd 新密码
//  @return Logs 返回的执行状态信息:true,表示执行成功;false,执行失败
//

func UpdatePasswdByEmail(email string, newPasswd string) (Logs bool) {

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)
	cmd := fmt.Sprintf("update %s set sl_passwd = '%s' where sl_email ='%s' ", Tb_Stud_login, newPasswd, email)
	result, err := sql.Exec(cmd)
	CheckErr(err)

	//根据执行sql语句判断影响的结果，结果大于0，则成功执行
	num, err := result.RowsAffected()
	if num == 0 {
		return false
	}

	return true
}

//【管理员】-重置学生密码

func ResetStudPasswd(Sid string, ResetPasswd string) int {

	Step := 0

	num := QueryStudBySid(Sid) //先查询
	if num == 0 {              //如果没有直接为返回0
		return 0
	} else {
		Step++
	}

	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer db.Close()
	CheckErr(err)

	SidInt64, err := strconv.ParseInt(Sid, 10, 64)

	SqlCmd := fmt.Sprintf("update %s sl set sl.sl_passwd ='%s' where sl.sl_id= %d", Tb_Stud_login, ResetPasswd, SidInt64)
	stmt, err := db.Prepare(SqlCmd)
	CheckErr(err)

	exec, err := stmt.Exec()
	CheckErr(err)
	FinishedNum, err := exec.RowsAffected()
	if FinishedNum != 0 {
		Step++
	}

	return Step
}

//【管理员】 根据学生id查询学生登录表,有则返回结果1;0

func QueryStudBySid(Sid string) (num int64) {

	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)

	SidInt64, err := strconv.ParseInt(Sid, 10, 64)

	SqlCmd := fmt.Sprintf("select * from  %s sl where sl.sl_id=%d ", Tb_Stud_login, SidInt64)

	rows, err := db.Query(SqlCmd)
	CheckErr(err)

	for rows.Next() {
		return 1
	}

	return 0
}

// ResetAllStudentLoginCode
//
//	@Description: 重置所有学生的登录验证码
//	@return int64
func ResetAllStudentLoginCode() int64 {
	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)

	sqlCmd := fmt.Sprintf("UPDATE %s set sl_loginCode ='%s'", Tb_Stud_login, LoginLocked)
	result, err := sql.Exec(sqlCmd)
	CheckErr(err)
	num, err := result.RowsAffected()
	CheckErr(err)
	return num
}

// UpdateStuLogoByIdAndLoginCode 更新头像路径
//
//	@Description:
func UpdateStuLogoByIdAndLoginCode(id string, loginCode string, logoUrl string) (resultNum int64) {
	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)

	IdInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
	}

	sqlCmd := fmt.Sprintf(" UPDATE %s sl set sl_logo='%s' where sl.sl_id=%d and sl.sl_loginCode='%s' ", Tb_Stud_login, logoUrl, IdInt, loginCode)
	result, err := sql.Exec(sqlCmd)
	CheckErr(err)
	num, err := result.RowsAffected()
	CheckErr(err)
	return num
}
