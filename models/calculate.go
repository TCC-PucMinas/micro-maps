package models

import (
	"encoding/json"
	"fmt"
	"micro-maps/db"
	"micro-maps/service"
	"time"
)

type LatAndLng struct {
	Lat string
	Lng string
}

type Calculate struct {
	Id            int64
	Origin        LatAndLng
	Destiny       LatAndLng
	HumanReadable string
	Meters        int
	Duration      time.Duration
}

func (calc *Calculate) CalculateRoute() error {

	origin := fmt.Sprintf("%v, %v", calc.Origin.Lat, calc.Origin.Lng)
	destiny := fmt.Sprintf("%v, %v", calc.Destiny.Lat, calc.Destiny.Lng)

	meters, humanaReadable, duration, err := service.CalculateRoute(origin, destiny)

	if err != nil {
		return err
	}

	calc.Meters = meters
	calc.HumanReadable = humanaReadable
	calc.Duration = duration
	return nil
}

func (l *LatAndLng) GetLatAndLngByAddress(address string) error {

	lat, lng, err := service.GetLatAndLngByAddress(address)

	if err != nil {
		return err
	}

	l.Lat = lat
	l.Lng = lng

	return nil
}

func (calc *Calculate) OriginToString() string {
	mapB, err := json.Marshal(calc.Origin)
	if err != nil {
		return ""
	}
	return string(mapB)
}

func (calc *Calculate) OriginToStruct(val string) {
	latAndLng := LatAndLng{}
	_ = json.Unmarshal([]byte(val), &latAndLng)
	calc.Origin = latAndLng
}

func (calc *Calculate) DestinyToStruct(val string) {
	latAndLng := LatAndLng{}
	_ = json.Unmarshal([]byte(val), &latAndLng)
	calc.Destiny = latAndLng
}

func (calc *Calculate) DestinyToString() string {
	mapB, err := json.Marshal(calc.Destiny)
	if err != nil {
		return ""
	}
	return string(mapB)
}

func (calc *Calculate) InsertCalculate() error {

	sql := db.ConnectDatabase()

	query := "insert into calculates (origin, destiny, humanReadble, meters) values (?, ?, ?, ?)"

	createCalculate, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	origin := calc.OriginToString()
	destiny := calc.DestinyToString()

	_, e := createCalculate.Exec(origin, destiny, calc.HumanReadable, calc.Meters)

	if e != nil {
		return e
	}

	return nil
}

func (calc *Calculate) GetCalculateOriginAndDestiny() error {

	sql := db.ConnectDatabase()

	query := `select id, origin, destiny, humanReadble, meters from calculates 
					where JSON_CONTAINS(origin, ?) and JSON_CONTAINS(destiny, ?) limit 1;`

	requestCalculate, err := sql.Query(query, calc.OriginToString(), calc.DestinyToString())

	if err != nil {
		return err
	}

	for requestCalculate.Next() {
		var origin, destiny, humanReadble string
		var meters int
		var id int64
		_ = requestCalculate.Scan(&id, &origin, &destiny, &humanReadble, &meters)
		if id != 0 {
			calc.Id = id
			calc.OriginToStruct(origin)
			calc.DestinyToStruct(destiny)
			calc.HumanReadable = humanReadble
			calc.Meters = meters
		}
	}

	return nil
}
