package models

import (
	"encoding/json"
	"fmt"
	"micro-maps/db"
	"micro-maps/service"
	"time"
)

var (
	keyMapByAddress = "key-map-get-by-address"
)

type Maps struct {
	Id       int64  `json:"id"`
	Inline   string `json:"address"`
	Street   string `json:"street"`
	District string `json:"district"`
	City     string `json:"city"`
	Country  string `json:"country"`
	State    string `json:"state"`
	Number   string `json:"number"`
	ZipCode  string `json:"zipCode"`
	Lat      string `json:"lat"`
	Lng      string `json:"lng"`
}

func (m *Maps) GenerateInline() {
	m.Inline = fmt.Sprintf("%v, %v, %v, %v, %v, %v", m.Street, m.Number, m.District, m.City, m.State, m.Country)
}

func (m *Maps) GetLocation() error {

	googleServie := service.LatAndLng{}

	if err := googleServie.GetLatAndLngByAddress(m.Inline); err != nil {
		return err
	}

	m.Lat = googleServie.Lat
	m.Lng = googleServie.Lng

	return nil
}

func (m *Maps) CreateMaps() error {
	sql := db.ConnectDatabase()

	query := "insert into maps (inline, street, district, city, zipCode, country, `state`, `number`, lat, lng) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	createDestination, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := createDestination.Exec(m.Inline, m.Street, m.District, m.City, m.ZipCode, m.Country, m.State, m.Number, m.Lat, m.Lng)

	if e != nil {
		return e
	}

	return nil
}

func (m *Maps) setRedisCacheMapGetByInline() error {
	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return err
	}

	marshal, err := json.Marshal(m)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyMapByAddress, m.Inline)

	return redis.Set(key, marshal, 24*time.Hour).Err()
}

func (m *Maps) getRedisCacheMapGetByInline() error {

	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return err
	}

	key := fmt.Sprintf("%v - %v", keyMapByAddress, m.Inline)

	value, err := redis.Get(key).Result()

	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(value), &m); err != nil {
		return err
	}
	return nil
}

func GetMapByAddres(inline string) (Maps, error) {

	m := Maps{}

	if err := m.getRedisCacheMapGetByInline(); err == nil {
		return m, nil
	}

	sql := db.ConnectDatabase()

	query := `select id, inline, street, district, zipCode, city, country, state, number, lat, lng  from maps where inline = ? limit 1;`

	requestConfig, err := sql.Query(query, inline)

	if err != nil {
		return m, err
	}

	for requestConfig.Next() {
		var inline, street, district, zipCode, city, country, state, number, lat, lng string
		var id int64
		_ = requestConfig.Scan(&id, &inline, &street, &district, &zipCode, &city, &country, &state, &number, &lat, &lng)
		if id != 0 {
			m.Id = id
			m.Inline = inline
			m.Street = street
			m.District = district
			m.ZipCode = zipCode
			m.City = city
			m.Country = country
			m.State = state
			m.Number = number
			m.Lat = lat
			m.Lng = lng
		}
	}

	_ = m.setRedisCacheMapGetByInline()

	return m, nil
}
