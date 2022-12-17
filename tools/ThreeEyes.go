package tools

/**
@Title 三目表达式
@Author 薛智敏
@CreateTime 2022年8月10日07:40:41
*/

//  ThreeEyes  三目表达式
//  @Description:
//  @param expression 表达式
//  @param True 满足的赋值
//  @param False 不满足赋值
//  @return result 返回值
//

func ThreeEyesReturnInt(expression bool, True int, False int) (result int) {

	switch expression {
	case true:
		return True
	}
	return False

}
