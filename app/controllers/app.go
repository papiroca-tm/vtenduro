package controllers

import (
	"github.com/revel/revel"
	mRace "vtEnduro/app/models/races"
)

type App struct {
	*revel.Controller
	modelRace *mRace.MRace
}

func (c App) Index() revel.Result {
	return c.Render()
}
