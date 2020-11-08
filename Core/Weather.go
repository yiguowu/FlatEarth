package Core

import (
	"math/rand"
	"time"
)

func ComputeWeatherProb(weatherMatrix [4][4]float32, vector [4]float32) (result [4]float32) {
	for i := 0; i < 4; i++ {
		var sum float32
		sum = 0.0
		for j := 0; j < 4; j++ {
			sum = sum + weatherMatrix[j][i]*vector[j]
		}
		result[i] = sum
	}
	return result
}

func ConvertProdToWeather(vector [4]float32) Weather {
	sunny := vector[0]
	rainy := sunny + vector[1]
	cloudy := rainy + vector[2]
	rand.Seed(time.Now().UTC().UnixNano())
	value := rand.Float32()
	if value <= sunny {
		return Sunny
	} else if value <= rainy {
		return Rainy
	} else if value <= cloudy {
		return Cloudy
	} else {
		return Stormy
	}
}

func CalculateWeather(weatherMatrix [4][4]float32, currentValue Weather) (futureValue Weather) {
	switch currentValue {
	case Sunny:
		vector := [4]float32{1, 0, 0, 0}
		prob := ComputeWeatherProb(weatherMatrix, vector)
		return ConvertProdToWeather(prob)
	case Rainy:
		vector := [4]float32{0, 1, 0, 0}
		prob := ComputeWeatherProb(weatherMatrix, vector)
		return ConvertProdToWeather(prob)
	case Cloudy:
		vector := [4]float32{0, 0, 1, 0}
		prob := ComputeWeatherProb(weatherMatrix, vector)
		return ConvertProdToWeather(prob)
	case Stormy:
		vector := [4]float32{0, 0, 0, 1}
		prob := ComputeWeatherProb(weatherMatrix, vector)
		return ConvertProdToWeather(prob)
	default:
		return Sunny
	}
}
