package dao

import (
	"database/sql"
	"fmt"
	"log"
)

//MySql配置相关信息

var LoginLocked = "--Locked"

const (
	DriverName  = "mysql"     //数据库驱动
	DbAddress   = "localhost" //数据库IP地址默认是127.0.0.1
	DbPort      = "3306"      //数据库的端口
	MysqlUser   = "root"      //Mysql 用户名
	MysqlPasswd = "123456"    //Mysql 密码
)

//数据源初始化

func DataSourceNameInit(DbName string) string {
	sprintf := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", MysqlUser, MysqlPasswd, DbAddress, DbPort, DbName)
	return sprintf

}

//修改mysql系统配置

func MysqlSystemConfig() {

	db, err := sql.Open(DriverName, DataSourceNameInit(Db_Course_select_sys))
	CheckErr(err)
	defer db.Close()
	stmt, err := db.Prepare("set global  interactive_timeout=1") //设置交互超时
	CheckErr(err)
	stm2, err := db.Prepare("set global wait_timeout=1") //设置连接超时
	CheckErr(err)
	stm3, err := db.Prepare("set global max_connections=15000") //设置最大连接数
	CheckErr(err)
	exec, err := stmt.Exec()
	CheckErr(err)
	exec2, err := stm2.Exec()
	CheckErr(err)
	exec3, err := stm3.Exec()
	exe1, _ := exec.RowsAffected()
	exe2, _ := exec2.RowsAffected()
	exe3, err := exec3.RowsAffected()
	if exe1 >= 0 && exe2 >= 0 && exe3 >= 0 {
		log.Println("修改mysql系统配置[成功]")
	} else {
		log.Println("修改mysql系统配置[失败]")
	}

}
