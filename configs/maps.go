package configs

import "googlemaps.github.io/maps"

func InitMap() (*maps.Client, error){
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyB8Xcw0-bTqcs2vXOQ5SANu65-4IR1rRFc"))

	if err != nil{
		return c, err
	}

	return c, nil
}