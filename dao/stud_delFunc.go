package dao

import (
	"database/sql"
	"fmt"
	"strconv"
)

//=========================增加=====================================
//=========================删除=====================================

//【管理员】注销学生账号

func DeleteStudBySid(Sid string) (finished int64) {

	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)
	defer db.Close()

	idInt64, err := strconv.ParseInt(Sid, 10, 64)
	sqlCmd := fmt.Sprintf("delete from stud_del sd where sd.sd_id=%d", idInt64)

	stmt, err := db.Prepare(sqlCmd)
	CheckErr(err)

	exec, err := stmt.Exec()
	CheckErr(err)

	num, err := exec.RowsAffected()
	CheckErr(err)

	return num
}

//=========================查询=====================================

//根据学号id查询学生

func SelectStudentById(id string) Stud_del {

	student_del := Stud_del{}

	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	defer sql.Close()
	CheckErr(err)

	idInt64, err := strconv.ParseInt(id, 10, 64)
	CheckErr(err)

	querySql := fmt.Sprintf("select * from %s sd where sd.sd_id=%d", Tb_Stud_del, idInt64)
	rows, err := sql.Query(querySql)
	CheckErr(err)

	for rows.Next() {

		err := rows.Scan(&student_del.Sd_id, &student_del.Sd_relname, &student_del.Sd_gender, &student_del.Sd_age, &student_del.Sd_address, &student_del.Sd_sys)
		CheckErr(err)
	}

	return student_del
}

//=========================修改=====================================

func StudDelUpdateById(id string, name string, age string, gender string, address string, sys string) (finishedNum int64) {
	sql, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)

	//年龄转int
	age_int, err := strconv.ParseInt(age, 10, 32)
	CheckErr(err)

	//str转int64
	id_int64, err := strconv.ParseInt(id, 10, 64)
	CheckErr(err)
	updateSql := fmt.Sprintf("update stud_del set sd_relname = '%s' , sd_age =%d ,sd_gender = '%s' ,sd_address = '%s' ,sd_sys = '%s' where sd_id = %d ", name, age_int, gender, address, sys, id_int64)
	defer sql.Close()

	result, err := sql.Exec(updateSql)
	CheckErr(err)
	num, err := result.RowsAffected()
	CheckErr(err)
	return num

}
