package models

import (
	"micro-maps/communicate"
	"sort"
)

func recursiveCalculateRoute(arrayRoutes []*communicate.Routes) (Calculate, error) {

	var arrayCalc []Calculate
	for _, v := range arrayRoutes {
		calc := Calculate{}
		calc.Id = v.IdCourier
		calc.Origin = LatAndLng{Lat: v.Origin.Lat, Lng: v.Origin.Lng}
		calc.Destiny = LatAndLng{Lat: v.Destiny.Lat, Lng: v.Destiny.Lng}

		if err := calc.CalculateRoute(); err != nil {
			return arrayCalc[0], err
		}

		arrayCalc = append(arrayCalc, calc)
	}

	sort.SliceStable(arrayCalc, func(i, j int) bool {
		return arrayCalc[i].Meters < arrayCalc[j].Meters
	})

	return arrayCalc[0], nil
}

func removeIndex(s []*communicate.Routes, index int) []*communicate.Routes {
	return append(s[:index], s[index+1:]...)
}

func setNewOrigin(c Calculate, array []*communicate.Routes) []*communicate.Routes {

	for _, v := range array {
		v.Origin.Lat = c.Destiny.Lat
		v.Origin.Lng = c.Destiny.Lng
	}

	return array
}

func setOriginCalc(arrayRoutes, newArrayOrder []*communicate.Routes) ([]*communicate.Routes, error) {

	if len(arrayRoutes) > 0 {
		calc, err := recursiveCalculateRoute(arrayRoutes)
		if err != nil {
			return arrayRoutes, err
		}
		for index, v := range arrayRoutes {
			if v.IdCourier == calc.Id {
				v.Order = int64(len(newArrayOrder) + 1)
				newArrayOrder = append(newArrayOrder, arrayRoutes[index])
				arrayRoutes = removeIndex(arrayRoutes, index)
				return setOriginCalc(setNewOrigin(calc, arrayRoutes), newArrayOrder)
			}
		}
	}

	return newArrayOrder, nil
}

func CalculateRoutes(arrayRoutesOrder []*communicate.Routes) ([]*communicate.Routes, error) {

	var arrayRoutes []*communicate.Routes

	return setOriginCalc(arrayRoutes, arrayRoutesOrder)
}
