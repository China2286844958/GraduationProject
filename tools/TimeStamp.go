package tools

import (
	"fmt"
	"strconv"
	"time"
)

/*时间戳*/

func GetTimeStamp() (timeStamp int64, exiErr bool) {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	clock, min, sec := time.Now().Clock()
	nanosecond := time.Now().Nanosecond()
	st := fmt.Sprintf("%d%d%d%d%d%d%d", year, month, day, clock, min, sec, nanosecond/10000)
	parseInt, e := strconv.ParseInt(st, 10, 64) //字符串转化十进制int64
	errBool := CheckErr(e)
	return parseInt, errBool

}

//捕捉错误

func CheckErr(err error) bool {

	if err != nil {
		fmt.Println("时间戳错误：\t", err.Error())
		return true
	}

	return false

}
