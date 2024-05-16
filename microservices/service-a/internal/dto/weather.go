package dto

import "fmt"

type WeatherResponse struct {
	TempC Number `json:"temp_C"`
	TempF Number `json:"temp_F"`
	TempK Number `json:"temp_K"`
}

type Number float64

func (n Number) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.1f", n)), nil
}
