// credit высчитывает ежемесяный аннуитетный платеж, переплату и итоговую сумму  по кредиту
package credit

import (
	"math"
)

// Calculate возвращает ежемесяный платеж, переплату и итоговую сумму
func Calculate(amount, percent, months int) (monthly, over, total int) {
	const oneYear = 12
	//middleMonthPercent := calcMiddlePercentYearToMonths(percent) - 1
	middleMonthPercent := (float64(percent) / oneYear) / 100
	//fmt.Printf("Процентная ставка из годoвого значения к месячному: %.5f\n", middleMonthPercent)
	coef := annuityСoefficient(middleMonthPercent, months)
	//fmt.Printf("Коэфициент аннуитета: %.5f\n", coef)
	monthly = int(float64(amount) * coef)
	//fmt.Printf("Ежемесяный платеж %v\n", monthly)
	total = monthly * months
	//fmt.Printf("Итоговая выплата %v\n", total)
	over = total - amount
	//fmt.Printf("Переплата по кредиту составит %d\n", over)

	return monthly, over, total
}

// calcMiddlePercentYearToMonths процентную ставку из годoвого значения к месячному
func calcMiddlePercentYearToMonths(percent int) float64 {
	const oneYear = 12
	middlePercent := 1 + float64(percent)/100.0
	coefPercent := math.Pow(middlePercent, 1.0/oneYear)
	//coefPercent := (float64(percent) / oneYear) / 100
	return coefPercent
}

// annuityСoefficient высчитывает коэфициент аннуитета
func annuityСoefficient(middleMonthPercent float64, months int) float64 {
	k := math.Pow(1.0 + middleMonthPercent, float64(months))
	сoefAnuitet := (middleMonthPercent * k) / (k - 1.0)
	return сoefAnuitet
}
