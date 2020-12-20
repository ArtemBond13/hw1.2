package main

import (
	"fmt"
	"math"
)

func main() {
	amount := 12000
	percent := 6
	months := 36
	//min, max := credit.Calculate(amount)
	//fmt.Println(min, max)
	//
	//montly, total, over := credit.CalculateCredit(amount, percent, months)
	//fmt.Println(montly, total, over)
	const oneYear = 12
	middleMonthPercent := calcMiddlePercentYearToMonths(percent) - 1
	fmt.Printf("Процентная ставка из годoвого значения к месячному %.5f\n", middleMonthPercent)
	coef := annuitСoefficient(middleMonthPercent, months)
	fmt.Printf("Коэфициент ануетета %.5f\n", coef)
	monthly := float64(amount)* coef
	fmt.Printf("Ежемесяный платеж %.2f\n", monthly)
	var total = int64(monthly * float64(months))
	fmt.Printf("Итоговая выплата %v\n", total)
	var over = total - int64(amount)
	fmt.Printf("%.2f, %d, %d", monthly, over, total)
}

func calcMiddlePercentYearToMonths(percent int)  float64{
	middlePercent := 1 + float64(percent) / 100.0
	coefPercent := math.Pow(middlePercent, 1.0 / 12.0)
	return coefPercent
}

func annuitСoefficient(middleMonthPercent float64, months int) float64 {
	k := math.Pow(1.0 + middleMonthPercent, float64(months))
	koefAnutet := (middleMonthPercent * k) / (k - 1.0)
	return koefAnutet
}