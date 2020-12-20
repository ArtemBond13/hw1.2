package anuetet

func Calculate(amount  int) (min, max int) {
	const minPercent = 6
	minResult := (amount * (100 +minPercent)) / 100
	const maxPercent = 12
	maxResult := (amount * (100 +maxPercent)) / 100
	return minResult, maxResult
}

func CalculateFloat(amount  float32) (min, max float32) {
	const minPercent = 6
	min = (amount * (minPercent + 100)) / 100
	const maxPercent = 12
	max= (amount * (100 + maxPercent)) / 100
	return
}