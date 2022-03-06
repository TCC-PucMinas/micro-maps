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
