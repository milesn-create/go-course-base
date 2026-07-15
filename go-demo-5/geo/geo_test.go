package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {

	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	got, err := geo.GetMyLocation(city)

	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v , получилось %v", expected, got)

	}

}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Londonshdocsc"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v , получилось %v", geo.ErrNoCity, err)
	}

}
