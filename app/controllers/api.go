package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	mRace "vtEnduro/app/models/races"
)

type Api struct {
	*revel.Controller
	modelRace *mRace.MRace
}

func (c Api) GetRaceList(dt, city, name string) revel.Result {
	c.modelRace = new(mRace.MRace)
	result, err := c.modelRace.GetRaceList(dt, city, name)
	if err != nil {
		// todo
	}
	resultByte, _ := json.Marshal(result)
	response := string(resultByte[:])
	return c.RenderJson(response)
}

func (c Api) GetRaceInfo(raceUID string) revel.Result {
	c.modelRace = new(mRace.MRace)
	result, err := c.modelRace.GetRaceInfo(raceUID)
	if err != nil {
		// todo
	}
	resultByte, _ := json.Marshal(result)
	response := string(resultByte[:])
	return c.RenderJson(response)
}

func (c Api) GetClassList(raceUID string) revel.Result {
	c.modelRace = new(mRace.MRace)
	result, err := c.modelRace.GetRaceClassesArr(raceUID)
	if err != nil {
		// todo
	}
	resultByte, _ := json.Marshal(result)
	response := string(resultByte[:])
	return c.RenderJson(response)
}
