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

func structToJson(structure interface{}) (response string, err error) {
	resultByte, _ := json.Marshal(structure)
	response = string(resultByte[:])
	return response, nil
}

func (c Api) GetRaceList(dt, city, name string) revel.Result {
	c.modelRace = new(mRace.MRace)
	result, err := c.modelRace.GetRaceList(dt, city, name)
	if err != nil {
		// todo
	}
	response, err := structToJson(result)
	if err != nil {
		// todo
	}
	return c.RenderJson(response)
}

func (c Api) GetRaceInfo(raceUID string) revel.Result {
	c.modelRace = new(mRace.MRace)
	result, err := c.modelRace.GetRaceInfo(raceUID)
	if err != nil {
		// todo
	}
	response, err := structToJson(result)
	if err != nil {
		// todo
	}
	return c.RenderJson(response)
}

func (c Api) GetClassList(raceUID string) revel.Result {
	c.modelRace = new(mRace.MRace)
	result, err := c.modelRace.GetRaceClassesArr(raceUID)
	if err != nil {
		// todo
	}
	response, err := structToJson(result)
	if err != nil {
		// todo
	}
	return c.RenderJson(response)
}

func (c Api) GetMarshalList(raceUID string) revel.Result {
	c.modelRace = new(mRace.MRace)
	result, err := c.modelRace.GetRaceMarshalsArr(raceUID)
	if err != nil {
		// todo
	}
	response, err := structToJson(result)
	if err != nil {
		// todo
	}
	return c.RenderJson(response)
}

func (c Api) GetMarshalInfo(raceUID string, mNumber int) revel.Result {
	c.modelRace = new(mRace.MRace)
	result, err := c.modelRace.GetRaceMarshalInfo(raceUID, mNumber)
	if err != nil {
		// todo
	}
	response, err := structToJson(result)
	if err != nil {
		// todo
	}
	return c.RenderJson(response)
}

func (c Api) GetClassInfo(raceUID string, classUID string) revel.Result {
	c.modelRace = new(mRace.MRace)
	result, err := c.modelRace.GetRaceClassInfo(raceUID, classUID)
	if err != nil {
		// todo
	}
	response, err := structToJson(result)
	if err != nil {
		// todo
	}
	return c.RenderJson(response)
}

func (c Api) GetCheckpointList(raceUID string, classUID string) revel.Result {
	c.modelRace = new(mRace.MRace)
	result, err := c.modelRace.GetCheckpointsArr(raceUID, classUID)
	if err != nil {
		// todo
	}
	response, err := structToJson(result)
	if err != nil {
		// todo
	}
	return c.RenderJson(response)
}

func (c Api) GetCheckpointInfo(raceUID string, classUID string, number int) revel.Result {
	c.modelRace = new(mRace.MRace)
	result, err := c.modelRace.GetCheckpointInfo(raceUID, classUID, number)
	if err != nil {
		// todo
	}
	response, err := structToJson(result)
	if err != nil {
		// todo
	}
	return c.RenderJson(response)
}
