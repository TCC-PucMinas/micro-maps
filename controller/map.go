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

	if m, err := models.GetMapByAddres(maps.Inline); err == nil && m.Id != 0 {
		res = &communicate.GelocationResponse{
			Lat: m.Lat,
			Lng: m.Lng,
		}
		return res, nil
	}

	if err := maps.GetLocation(); err != nil {
		return res, err
	}

	_ = maps.CreateMaps()

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

	if err := serviceCalculate.GetCalculateOriginAndDestiny(); err == nil && serviceCalculate.Id != 0 {
		originResponse := &communicate.LatAndLng{
			Lat: origin.Lat,
			Lng: origin.Lng,
		}

		destinyResponse := &communicate.LatAndLng{
			Lat: destiny.Lat,
			Lng: destiny.Lng,
		}

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

	if err := serviceCalculate.InsertCalculate(); err != nil {
		return res, nil
	}

	res.HumanReadable = serviceCalculate.HumanReadable
	res.Meters = int64(serviceCalculate.Meters)

	res.Origin = originResponse
	res.Destiny = destinyResponse

	return res, nil
}

func (s *MapServer) OrderRoutes(ctx context.Context, request *communicate.OrderRoutesRequest) (*communicate.OrderRoutesResponse, error) {
	res := &communicate.OrderRoutesResponse{
		Routes: request.Routes,
	}
	// TODO: calcular as rotas mais próximas com base no array.
	return res, nil
}
