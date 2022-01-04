/*
@Author: nullzz
@Date: 2021/12/30 5:51 下午
@Version: 1.0
@DEC: 区域
*/
package aoi

type Area struct {
	Id   int
	XMin int //左x边界
	XMax int //右x边界
	YMin int //下边界
	YMax int //下边界

	//PreArea  *Area
	//NextArea *Area
	//TopArea  *Area
	//DownArea *Area
}

func NewArea(id, xMin, xMax, yMin, yMax int) *Area {
	area := &Area{
		Id:   id,
		XMin: xMin,
		XMax: xMax,
		YMin: yMin,
		YMax: yMax,
	}
	return area
}
