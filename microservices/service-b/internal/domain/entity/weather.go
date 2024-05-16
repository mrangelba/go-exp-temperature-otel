package entity

import (
	"fmt"
	"strconv"
)

type Weather struct {
	TempC float32
}

func (weather *Weather) TempF() float32 {
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", (weather.TempC*1.8)+32), 32)

	return float32(v)
}

func (weather *Weather) TempK() float32 {
	return weather.TempC + 273
}
