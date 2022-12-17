package tools

import (
	"math/rand"
	"time"
)

/**
@Title 随机数提取库
@Author 薛智敏
@CreateTime 2022年8月21日00:12:13
*/

// 小写字母库
var lowerCaseLibrary = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z",
}

// 大写字母库
var upperCaseLibrary = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
}

// 数字库
var numberLibrary = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
}

// 特殊符号库
var specialSymbolLibrary = []string{
	"+", "-", "*", "/", ".",
}

//
//  GetLowerCaseLibrary 获取小写字母库
//  @Description:
//

func GetLowerCaseLibrary() []string {
	return lowerCaseLibrary
}

//  GetUpperCaseLibrary 获取大写字母库
//  @Description:
//

func GetUpperCaseLibrary() []string {
	return upperCaseLibrary
}

//  GetNumberLibrary 获取数字库
//  @Description:
//

func GetNumberLibrary() []string {
	return numberLibrary
}

//  GetSpecialSymbolLibrary 获取特殊符号库
//  @Description:
//

func GetSpecialSymbolLibrary() []string {
	return specialSymbolLibrary
}

//综合字符库

var Arr2dLibrary = [][]string{
	upperCaseLibrary,
	lowerCaseLibrary,
	numberLibrary,
}

func SetRandArr2D() {

}

// 获取随机码库
var getRandomCodesResources = func() []string {

	var valueCount = 0        //值的个数
	m := make(map[int]string) //值的集合

	//将二维数组，转化切片
	//遍历二维数组,得到一维数组
	for _, arr := range Arr2dLibrary {
		//遍历一维数组得到值
		for _, v := range arr {
			//装进集合
			m[valueCount] = v
			//计算值的个数
			valueCount++
		}
	}

	//定义一个和值的集合一样的切片
	randomCodesResourceList := make([]string, valueCount)

	for i := 0; i < valueCount; i++ {
		randomCodesResourceList[i] = m[i] //将集合里面的值装进切片
	}
	return randomCodesResourceList
}

var captchaSize = len(getRandomCodesResources()) //综合随机码库的大小

//
//  GetCaptchaSize 获取综合随机码库的大小
//  @Description:
//  @return int
//

func GetCaptchaSize() int {
	return captchaSize
}

//自定义获取验证码

func CaptchaByTimes(times int) string {

	captcha := ""
	// 给定变化的种子，这是以下所有代码的前提
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < times; i++ {
		captcha += getRandomCodes(rand.Intn(len(getRandomCodesResources())))
	}

	return captcha
}

//随机字母

func getRandomCodes(x int) string {
	return getRandomCodesResources()[x]

}
