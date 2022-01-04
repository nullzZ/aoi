/*
@Author: nullzz
@Date: 2021/12/30 9:00 下午
@Version: 1.0
@DEC:
*/
package test

import (
	"encoding/json"
	"server/pkg/aoi"
	"testing"
)

func Test(t *testing.T) {
	mgr := aoi.NewAoiManager(0, 100, 0, 100, 5, 5, true)
	mgr.Load()
}

func TestGetSurroundAreasByAreaId(t *testing.T) {
	mgr := aoi.NewAoiManager(0, 100, 0, 100, 5, 5, true)
	mgr.Load()
	areas, _ := mgr.GetSurroundAreasByAreaId(9)
	str, _ := json.Marshal(areas)
	t.Logf("areas=%s", string(str))
}
