package dao

/**
@Title 分页配置
@Author 薛智敏
@CreateTime 2022年7月26日23:03:08

*/

type PagingNumbers struct {
	showInfoEveryPage int64 //分页查询中每页显示的信息数

}

func (p *PagingNumbers) GetShowInfoEveryPage() int64 {
	return p.showInfoEveryPage
}

func (p *PagingNumbers) SetShowInfoEveryPage(showInfoEveryPage int64) {
	p.showInfoEveryPage = showInfoEveryPage
}

//分页的配置,每页显示的信息数代理

func ShowNumbersProxy() (everyPageShows int) {
	return 5 //分页查询中每页显示的信息数
}
