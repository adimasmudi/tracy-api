package configs

import (
	"os"

	"googlemaps.github.io/maps"
)

func InitMap() (*maps.Client, error){
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_CREDENTIAL")))

	if err != nil{
		return c, err
	}

	return c, nil
}