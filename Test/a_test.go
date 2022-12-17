package t_test

import (
	"GraduationProject/dao"
	"fmt"
	"testing"
)

var (
	name = 1
	Name = 1
)

func TestName(t *testing.T) {
	nameList := dao.SelectAllCourseName()               //查询所有课程名字 return []string
	maxScoreList := dao.SelectMaxScoreByCName(nameList) //查询课程对应的最高分 return []float
	var arrMax = maxScoreList[0:len(maxScoreList)]      //最高分数集合
	minScoreList := dao.SelectMinScoreByCName(nameList) //查询课程对应的最低分
	var arrMin = minScoreList[0:len(minScoreList)]      //最低分数集合

	fmt.Println(arrMax)
	fmt.Println(arrMin)

}
