package dao

import (
	"database/sql"
	"fmt"
	"strconv"
)

//=========================添加=====================================

//根据超级管理员权限，添加普通管理员

func AddAdminByRoot(idAutoGene int64, inEmail string, passwd string) {
	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer db.Close()
	CheckErr(err)

	stmt, err := db.Prepare("insert admin (a_id,a_email,a_passwd) VALUES (?,?,?) ")

	stmt.Exec(idAutoGene, inEmail, passwd)
}

//=========================删除=====================================

//超级管理员-删除普通管理员

func DeleteAdmin(id string) (finished int64) {

	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)
	defer db.Close()
	id_int64, _ := strconv.ParseInt(id, 10, 64)

	sqlCmd := fmt.Sprintf("delete from %s a where a.a_id = %d ", Tb_Admin, id_int64)
	result, err := db.Exec(sqlCmd)

	CheckErr(err)
	num, err := result.RowsAffected()

	return num
}

//=========================查询=====================================

//查询所有的普通管理员

func GetAllAdminPower1(searchWord string, searchPage int) (map[int]Admin, int) {

	admin := Admin{}

	adminList := make(map[int]Admin)
	db, err2 := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err2)

	db1, err2 := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err2)

	defer db.Close()
	defer db1.Close()

	//每页显示的信息数
	everyPageShows := ShowNumbersProxy()

	//当前页面
	currentPage := searchPage * everyPageShows
	//limit 页数*每页显示的数=当前页currentPage,每页显示的数据

	sqlCmd := fmt.Sprintf("select * from %s a where a.a_power !=10 and a.a_email like '%%%s%%' limit %d,%d", Tb_Admin, searchWord, currentPage, everyPageShows)
	rows, _ := db.Query(sqlCmd)

	countSql := fmt.Sprintf("select * from Admin a where a.a_power !=10 and a.a_email like '%%%s%%'", searchWord)
	rows1, _ := db1.Query(countSql)
	count := 0
	for rows.Next() {

		rows.Scan(&admin.A_id, &admin.A_email, &admin.A_passwd, &admin.A_power, &admin.A_loginCode)
		adminList[count] = admin
		admin = Admin{} //清除
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

	return adminList, countPage

}

//根据id和登录验证码查询admin

func GetAdminByIdAndLoginCode(id string, loginCodeIn string) Admin {
	if loginCodeIn == LockedWord { //当等于锁登词，直接返回false
		return Admin{} //无信息
	}
	admin := Admin{}
	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer db.Close()

	CheckErr(err)

	id_int64, err := strconv.ParseInt(id, 10, 64)
	CheckErr(err)
	sqlCmd := fmt.Sprintf("select * from %s a where a.a_id = %d  and a.a_logincode ='%s'", Tb_Admin, id_int64, loginCodeIn)
	rows, err := db.Query(sqlCmd)

	for rows.Next() {
		rows.Scan(&admin.A_id, &admin.A_email, &admin.A_passwd, &admin.A_power, &admin.A_loginCode)
	}
	return admin

}

//通过email获取管理员的信息-

func GetAdminLoginByEmail(email string) Admin {

	admin := Admin{}
	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)
	defer sql.Close()
	sqlCmd := fmt.Sprintf("select * from  %s a where a.A_email = '%s'", Tb_Admin, email)
	rows, err := sql.Query(sqlCmd)
	CheckErr(err)
	for rows.Next() {
		rows.Scan(&admin.A_id, &admin.A_email, &admin.A_passwd, &admin.A_power, &admin.A_loginCode)
	}
	return admin
}

//管理员登录验证码校验

func AdminLoginCodeVerification(id string, loginCodeIn string) (T bool) {

	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer db.Close()
	CheckErr(err)

	id_int64, err := strconv.ParseInt(id, 10, 64)
	CheckErr(err)
	sqlCmd := fmt.Sprintf("select a.a_logincode from %s a where a.a_id = %d  and a.a_logincode ='%s'", Tb_Admin, id_int64, loginCodeIn)
	rows, err := db.Query(sqlCmd)
	for rows.Next() {
		return true
	}
	return false

}

//根据id查询管理员

func QueryAdminById(id string) (finished int64) {
	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer db.Close()
	CheckErr(err)

	Id_Int64, err := strconv.ParseInt(id, 10, 64)
	sqlCmd := fmt.Sprintf("select * from admin a where a.a_id =%d", Id_Int64)
	prepare, err := db.Query(sqlCmd)

	CheckErr(err)
	for prepare.Next() {
		return 1
	}
	return 0

}

//根据id查询管理员,是否存在

func FindAdminByIdAndEnPasswd(id string, EnPasswd string) (Exist bool) {
	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer db.Close()
	CheckErr(err)

	Id_Int64, err := strconv.ParseInt(id, 10, 64)
	sqlCmd := fmt.Sprintf("select * from admin a where a.a_id =%d and a.a_passwd='%s'", Id_Int64, EnPasswd)
	prepare, err := db.Query(sqlCmd)

	CheckErr(err)
	for prepare.Next() {
		return true
	}
	return false

}

//=========================修改=====================================

//登录成功后刷新管理员的登录状态码

func SetAdminLoginCodeById(id int64, loginCode string) (res int64) {

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()

	CheckErr(err)
	sqlCmd := fmt.Sprintf("update  %s a set a.A_loginCode ='%s' where a.a_id = %d", Tb_Admin, loginCode, id)
	result, err := sql.Exec(sqlCmd)
	CheckErr(err)
	num, err := result.RowsAffected()
	CheckErr(err)
	return num
}

//超级管理员-重置普通管理员密码

func SuResetPasswdAdmin(ResetId string, ResetPasswd string) (finishedNum int64) {

	var num int64 = 0
	finished := QueryAdminById(ResetId)
	if finished == 0 {
		return 0
	} else {
		num = 1 //能查到设置1，先假设密码已经被重置了，再次被重置，它的影响行数为0，设置1，表示有这个行，即使被重置，也不会提示错误
	}

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()

	CheckErr(err)
	ResetId_Int64, _ := strconv.ParseInt(ResetId, 10, 64)

	sqlCmd := fmt.Sprintf("update  %s a set a.a_passwd ='%s' where a.a_id = %d", Tb_Admin, ResetPasswd, ResetId_Int64)
	result, err := sql.Exec(sqlCmd)
	CheckErr(err)
	affected, err := result.RowsAffected()

	num += affected //影响的行数，加上查找的行数，精确确保密码是否被重置过，或者该号根本不存在
	CheckErr(err)
	return num

}

//【管理员】根据id,修改密码

func AdminResetPasswd(id string, resetPasswd string) (finishedNum int64) {
	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()

	CheckErr(err)

	ResetId_Int64, _ := strconv.ParseInt(id, 10, 64)
	sqlCmd := fmt.Sprintf("update  %s a set a.A_passwd ='%s' where a.a_id = %d", Tb_Admin, resetPasswd, ResetId_Int64)
	result, err := sql.Exec(sqlCmd)
	CheckErr(err)
	num, err := result.RowsAffected()
	CheckErr(err)
	return num
}

// ResetAllAdminLoginCode
//
//	@Description: 重置所有管理员登录验证码
//	@return int64
func ResetAllAdminLoginCode() int64 {
	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)

	sqlCmd := fmt.Sprintf("UPDATE %s set a_loginCode ='%s'", Tb_Admin, LoginLocked)
	result, err := sql.Exec(sqlCmd)
	CheckErr(err)
	num, err := result.RowsAffected()
	CheckErr(err)
	return num
}
