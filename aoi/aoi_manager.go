/*
@Author: nullzz
@Date: 2021/12/30 5:48 下午
@Version: 1.0
@DEC:
*/
package aoi

import (
	"bitbucket.org/funplus/golib/zaplog"
	"bytes"
	"encoding/json"
	"fmt"
)

type AOIManager struct {
	MinX   int
	areas  map[int]*Area
	MaxX   int
	MinY   int
	MaxY   int
	CountX int //从X方向的格子数
	CountY int //从Y方向的格子数
	isLog  bool
}

func NewAoiManager(minX, maxX, minY, maxY, countX, countY int, isLog bool) *AOIManager {
	mgr := &AOIManager{
		MinX:   minX,
		MaxX:   maxX,
		MinY:   minY,
		MaxY:   maxY,
		CountX: countX,
		CountY: countY,
		areas:  make(map[int]*Area),
		isLog:  isLog,
	}
	return mgr
}

func (a *AOIManager) Load() {
	var log bytes.Buffer
	log.WriteString("\n-----------------------------------\n")
	for y := 0; y < a.CountY; y++ {
		var areaIdLog bytes.Buffer
		for x := 0; x < a.CountY; x++ {
			id := a.GetAreaId(x, y, a.CountX)
			xMin := a.MinX + x*a.areaWidth()
			xMax := a.MinX + (x+1)*a.areaWidth()
			yMin := a.MinY + y*a.areaLength()
			yMax := a.MinY + (y+1)*a.areaLength()
			area := NewArea(id, xMin, xMax, yMin, yMax)
			a.logArea(area)
			a.areas[id] = area
			ids := fmt.Sprintf("%5d", id)
			areaIdLog.WriteString(ids + ",")
		}
		log.WriteString(areaIdLog.String() + "\n")
		log.WriteString("-----------------------------------\n")
	}
	log.WriteString("-----------------------------------\n")
	if a.isLog {
		zaplog.LoggerSugar.Debugf("show log=%s", log)
	}
	a.logAoiManager()
}

func (a *AOIManager) GetAreaId(x, y, countX int) int {
	return y*countX + x
}

// 得到每个格子在x轴方向的宽度
func (a *AOIManager) areaWidth() int {
	return (a.MaxX - a.MinX) / a.CountX
}

// 得到每个格子在x轴方向的长度
func (a *AOIManager) areaLength() int {
	return (a.MaxY - a.MinY) / a.CountY
}

// 通过区域id获取9宫格
func (a *AOIManager) GetSurroundAreasByAreaId(areaId int) ([]*Area, bool) {
	area, h := a.areas[areaId]
	if !h {
		return nil, false
	}
	var areas []*Area
	var areaIdXs []int
	areas = append(areas, area)
	//根据areaId求X
	x := areaId % a.CountX
	//先判断左边是否存在格子
	if x > 0 {
		areas = append(areas, a.areas[areaId-1])
	}
	//判断右边是否有格子
	if x < a.CountX-1 {
		areas = append(areas, a.areas[areaId+1])
	}

	for _, a := range areas {
		areaIdXs = append(areaIdXs, a.Id)
	}
	for _, areaIdX := range areaIdXs {
		idY := areaIdX / a.CountY //取出第几列
		if idY > 0 {              //判断上方是否有区域
			areas = append(areas, a.areas[areaIdX-a.CountX])
		}

		//判断下方是否有区域
		if idY < a.CountY-1 {
			areas = append(areas, a.areas[areaIdX+a.CountX])
		}
	}

	return areas, true
}

// 通过坐标查找区域id
func (a *AOIManager) GetAreaIdByPos(x, y float32) int {
	gx := (int(x) - a.MinX) / a.areaWidth()
	gy := (int(y) - a.MinY) / a.areaLength()
	return gy*a.CountX + gx
}

func (a *AOIManager) logArea(area *Area) {
	if !a.isLog {
		return
	}
	log, _ := json.Marshal(area)
	zaplog.LoggerSugar.Infof("AOIManager gen area=%s", string(log))
}

func (a *AOIManager) logAoiManager() {
	if !a.isLog {
		return
	}
	log, _ := json.Marshal(a)
	zaplog.LoggerSugar.Infof("AOIManager gen aoiManager=%s", string(log))
}
