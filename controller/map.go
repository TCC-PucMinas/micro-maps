package controller

import (
	"context"
	"micro-maps/communicate"
	"micro-maps/models"
	"micro-maps/service"
	"time"
)

type MapServer struct{}

func (s *MapServer) GetLocation(ctx context.Context, request *communicate.GelocationRequest) (*communicate.GelocationResponse, error) {

	res := &communicate.GelocationResponse{}

	maps := models.Maps{
		Street:   request.Street,
		District: request.District,
		City:     request.City,
		Country:  request.Country,
		State:    request.State,
		Number:   request.Number,
		ZipCode:  request.ZipCode,
	}

	maps.GenerateInline()

	if m, _ := models.GetMapByAddres(maps.Inline); m.Id != 0 {
		res = &communicate.GelocationResponse{
			Lat: maps.Lat,
			Lng: maps.Lng,
		}
		return res, nil
	}

	if err := maps.GetLocation(); err != nil {
		return res, err
	}

	if err := maps.CreateMaps(); err != nil {
		return res, err
	}

	res = &communicate.GelocationResponse{
		Lat: maps.Lat,
		Lng: maps.Lng,
	}

	return res, nil
}

func (s *MapServer) DirectionLocation(ctx context.Context, request *communicate.DirectionLocationRequest) (*communicate.DirectionLocationResponse, error) {
	res := &communicate.DirectionLocationResponse{}

	origin := service.LatAndLng{
		Lat: request.Origin.Lat,
		Lng: request.Origin.Lng,
	}

	destiny := service.LatAndLng{
		Lat: request.Destiny.Lat,
		Lng: request.Destiny.Lng,
	}
	serviceCalculate := service.Calculate{
		Origin:  origin,
		Destiny: destiny,
	}

	if err := serviceCalculate.CalculateRoute(); err != nil {
		return res, err
	}

	originResponse := &communicate.LatAndLng{
		Lat: origin.Lat,
		Lng: origin.Lng,
	}

	destinyResponse := &communicate.LatAndLng{
		Lat: destiny.Lat,
		Lng: destiny.Lng,
	}

	res.Duration = int64(serviceCalculate.Duration)

	h, _ := time.ParseDuration(serviceCalculate.Duration.String())
	res.Duration = int64(h.Minutes())
	res.HumanReadable = serviceCalculate.HumanReadable
	res.Meters = int64(serviceCalculate.Meters)
	res.HumanReadable = serviceCalculate.HumanReadable

	res.Origin = originResponse
	res.Destiny = destinyResponse

	return res, nil
}
