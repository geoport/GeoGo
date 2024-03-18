package data

import "github.com/geoport/GeoGo/models"

var Pressuremeter = models.Pressuremeter{
	Exps: []models.PressuremeterExp{
		{Depth: 3, LimitPressure: 47.5, NetLimitPressure: 42.5},
		{Depth: 3.5, LimitPressure: 48.2, NetLimitPressure: 43.3},
	},
}
