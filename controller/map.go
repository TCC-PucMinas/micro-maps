package controller

import (
	"context"
	"micro-maps/communicate"
	"micro-maps/models"
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

	origin := models.LatAndLng{
		Lat: request.Origin.Lat,
		Lng: request.Origin.Lng,
	}

	destiny := models.LatAndLng{
		Lat: request.Destiny.Lat,
		Lng: request.Destiny.Lng,
	}
	serviceCalculate := models.Calculate{
		Origin:  origin,
		Destiny: destiny,
	}

	// get cache calculate or database
	if err := serviceCalculate.GetCalculateOriginAndDestiny(); err == nil || serviceCalculate.Id == 0 {
		originResponse := &communicate.LatAndLng{
			Lat: origin.Lat,
			Lng: origin.Lng,
		}

		destinyResponse := &communicate.LatAndLng{
			Lat: destiny.Lat,
			Lng: destiny.Lng,
		}

		res.HumanReadable = serviceCalculate.HumanReadable
		res.Meters = int64(serviceCalculate.Meters)
		res.HumanReadable = serviceCalculate.HumanReadable

		res.Origin = originResponse
		res.Destiny = destinyResponse
		return res, nil
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

	res.HumanReadable = serviceCalculate.HumanReadable
	res.Meters = int64(serviceCalculate.Meters)
	res.HumanReadable = serviceCalculate.HumanReadable

	res.Origin = originResponse
	res.Destiny = destinyResponse

	return res, nil
}
