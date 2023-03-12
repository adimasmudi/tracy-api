package services

import (
	"context"
	"math"
	"tracy-api/configs"
	"tracy-api/models"
	"tracy-api/repository"

	"googlemaps.github.io/maps"
)

type MapsService interface {
	GetDirection(ctx context.Context, origin string, destination string)([]maps.Route, error)
	GetGeocode(ctx context.Context, location string)([]maps.GeocodingResult, error)
	GetPoliceNearby(ctx context.Context, lat float64, lng float64) (models.PoliceStation, error)
}

type mapsService struct {
	policeRepository repository.PoliceStationRepository
}

func NewMapsService(policeRepository repository.PoliceStationRepository) *mapsService{
	return &mapsService{policeRepository}
}

func (s *mapsService) GetDirection(ctx context.Context, origin string, destination string)([]maps.Route, error){
	c, _ := configs.InitMap()
	// get directions
	r := &maps.DirectionsRequest{
		Origin:      origin,
		Destination: destination,
	}
	route, _, err := c.Directions(ctx, r)
	if err != nil {
		return route, err
	}

	return route, err
}

func (s *mapsService) GetGeocode(ctx context.Context, location string)([]maps.GeocodingResult, error){
	c, _ := configs.InitMap()

	g := &maps.GeocodingRequest{
		Address: location,
	}

	geo, err := c.Geocode(ctx,g)

	if err != nil{
		return geo, err
	}

	return geo, nil
}

func (s *mapsService) GetPoliceNearby(ctx context.Context, lat float64, lng float64)(models.PoliceStation, error){
	c, _ := configs.InitMap()

	policeStation, err := s.policeRepository.GetAllPoliceStation(ctx)

	if err != nil{
		return policeStation[0], err
	}

	var distanceList []float64

	for idxPS := range(policeStation){
		// get geocode
		g := &maps.GeocodingRequest{
			Address: policeStation[idxPS].Alamat,
		}

		geo, err := c.Geocode(ctx,g)

		if err != nil{
			return policeStation[0], err
		}

		distance := (3959 * math.Acos(math.Cos(lat)* math.Cos(geo[0].Geometry.Location.Lat) * math.Cos(geo[0].Geometry.Location.Lng - lng) + math.Sin(lat) * math.Sin(geo[0].Geometry.Location.Lat)))

		distanceList = append(distanceList, distance)
	}

	idxMin := 0

	if len(distanceList) == 1{
		return policeStation[0], nil
	}

	for idxDist := range(distanceList){
		if distanceList[idxDist] < distanceList[idxMin]{
			idxMin = idxDist
		}
	}

	return policeStation[idxMin], nil
}