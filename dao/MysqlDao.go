package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//处理错误

func CheckErr(err error) bool {

	if err != nil {
		return true
	}

	return false
}

//==================================================增加================================================================

//==================================================查看================================================================

/*查询表中任意存在的字段 出现查询的结果返回真【用于校验注册时填写的信息】*/

func SelectTableByColumn(DbName string, TbName string, column string, columnData string) (TorF bool) {
	db, err := sql.Open(DriverName, DataSourceNameInit(DbName))

	CheckErr(err)
	defer db.Close()

	querySql := fmt.Sprintf("select %s from %s where %s = '%s'", column, TbName, column, columnData)
	rows, err1 := db.Query(querySql)

	CheckErr(err1)

	for rows.Next() {
		return true

	}
	return false

}

//==================================================修改================================================================

//==================================================删除================================================================
