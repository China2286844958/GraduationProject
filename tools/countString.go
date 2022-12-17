package tools

/**
@Title 计算字符串个数，包括中文
@Author 薛智敏
@CreteTime 2022年6月24日11:13:41
*/

//统计字符串个数，包括中文，返回数字

func StrCounts(str string) (count int) {
	for _, _ = range str {
		count++
	}
	return count
}
