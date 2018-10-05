package weather

import "testing"

func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{APIKey: ""}

	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}

	if weather.ID != 1851632 {
		t.Errorf("Chine is 3117735, not %d\n", weather.ID)
	}
}

func TestOpenWeatherMap(t *testing.T) {

}
