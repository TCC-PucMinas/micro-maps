package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"googlemaps.github.io/maps"
)

var keyGoogle = os.Getenv("GOOGLE_MAP_KEY")

func CalculateRoute(origin, destiny string) (int, string, time.Duration, error) {

	var meters int
	var humanReadable string
	var duration time.Duration

	c, err := maps.NewClient(maps.WithAPIKey(keyGoogle))

	if err != nil {
		return meters, humanReadable, duration, err
	}

	r := &maps.DirectionsRequest{
		Origin:      origin,
		Destination: destiny,
	}

	rout, _, err := c.Directions(context.Background(), r)

	if err != nil {
		return meters, humanReadable, duration, err
	}

	for _, v := range rout {
		for _, a := range v.Legs {
			meters = a.Meters
			humanReadable = a.HumanReadable
			duration = a.Duration
		}
	}
	return meters, humanReadable, duration, nil
}

func GetLatAndLngByAddress(address string) (string, string, error) {

	var lat, lng string
	c, err := maps.NewClient(maps.WithAPIKey(keyGoogle))

	if err != nil {
		return lat, lng, err
	}

	r := &maps.GeocodingRequest{
		Address: address,
	}

	rout, err := c.Geocode(context.Background(), r)

	if err != nil {
		return lat, lng, err
	}

	for _, v := range rout {
		lat = fmt.Sprintf("%f", v.Geometry.Location.Lat)
		lng = fmt.Sprintf("%f", v.Geometry.Location.Lng)
	}

	return lat, lng, nil
}
