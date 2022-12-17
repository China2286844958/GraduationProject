package tools

import (
	"regexp"
)

var (
	reEmail = `[\w\.]+@\w+\.[a-z]{2,3}(\.[a-z]{2,3})?` //邮箱
	/**
	  [\w\.]+  表示字母字符或者.其中一个出现 +1到多次
	  [a-z]{2,3} 表示a-z任意其中一个字符 出现2到3次
	  (\.[a-z]{2,3})?  ()表示分组   a到z中任意一个出现2到3次  ? 该分组出现0到1次
	*/
)

/*正则表达式*/

//校验邮箱的正则表达式

func EmailRegExp(email string) bool {
	compile := regexp.MustCompile(reEmail)
	allString := compile.FindAllString(email, -1)
	for _, v := range allString {
		if email == v {
			return true
		}
	}
	return false
}

//根据一串字符串，校验数据是否有空白字符

func ExprBlank(str string) bool {

	compile := regexp.MustCompile(`[a-zA-Z\S]`)
	allString := compile.FindAllString(str, -1) //符合表达式的字符串

	if len(str) == len(allString) { //当符合表达式的字符串与传入的字符串长度一致，表示源字符串所有数据符合
		return true
	}

	return false
}
