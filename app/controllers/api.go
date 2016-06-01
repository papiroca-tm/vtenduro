package controllers

import (
	"github.com/revel/revel"
	mRace "vtEnduro/app/models/races"
)

type Api struct {
	*revel.Controller
	modelRace *mRace.MRace
}

func (c Api) GetRaceList(dt, city, name string) revel.Result {
	revel.INFO.Println("dt", dt)
	revel.INFO.Println("city", city)
	revel.INFO.Println("name", name)

	c.modelRace = new(mRace.MRace)
	res := c.modelRace.GetRaceList(dt, city, name)
	return c.RenderJson(res)
}

func (c Api) GetRaceInfo(raceUID string) revel.Result {
	revel.INFO.Println("raceUID", raceUID)
	c.modelRace = new(mRace.MRace)
	res := c.modelRace.GetRaceInfo(raceUID)
	return c.RenderJson(res)
}
