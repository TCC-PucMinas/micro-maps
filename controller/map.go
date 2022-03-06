package controller

import (
	"context"
	"fmt"
	"micro-maps/communicate"
	"micro-maps/models"
	"micro-maps/service"
)

type MapServer struct{}

func (s *MapServer) GetLocation(ctx context.Context, request *communicate.GelocationRequest) (*communicate.GelocationResponse, error) {

	res := &communicate.GelocationResponse{}

	maps := models.Map{
		Street:   request.Street,
		District: request.District,
		City:     request.City,
		Country:  request.Country,
		State:    request.State,
		Number:   request.Number,
		ZipCode:  request.ZipCode,
	}

	googleServie := service.LatAndLng{}

	address := fmt.Sprintf("%v, %v, %v, %v, %v, %v", maps.Street, maps.Number, maps.District, maps.City, maps.State, maps.Country)

	if err := googleServie.GetLatAndLngByAddress(address); err != nil {
		return res, err
	}

	res = &communicate.GelocationResponse{
		Lat: googleServie.Lat,
		Lng: googleServie.Lng,
	}

	return res, nil
}
