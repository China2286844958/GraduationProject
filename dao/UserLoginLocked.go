package dao

import (
	"database/sql"
	"fmt"
	"strconv"
)

/**
@Title 用户登录上锁，必须成功登录，才能解开（刷新登录码）
@CreateTime 2022年7月14日00:27:22
@Author 薛智敏
*/

//用户上锁身份

const (
	LockedWord    = "--Locked" //锁登词
	AdminLocked   = "admin"    //管理员登录上锁
	StudentLocked = "student"  //学生登录上锁
)

//用户上锁 逻辑

func LoginCodeLockedController(WhoLocked string, LockedId string) int64 {

	db, err2 := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err2)
	defer db.Close()

	SqlCmd_ := "" //SQL语句命令
	lockedId, err := strconv.ParseInt(LockedId, 10, 64)
	CheckErr(err)

	switch WhoLocked {
	case AdminLocked:
		SqlCmd_ = fmt.Sprintf("update %s a set a.a_loginCode='%s' where a.a_id =%d", Tb_Admin, LockedWord, lockedId)

	case StudentLocked:
		SqlCmd_ = fmt.Sprintf("update %s sl set sl.sl_loginCode='%s' where sl.sl_id =%d", Tb_Stud_login, LockedWord, lockedId)
	}
	stmt, _ := db.Prepare(SqlCmd_)
	exec, _ := stmt.Exec()
	affected, _ := exec.RowsAffected()

	return affected
}
