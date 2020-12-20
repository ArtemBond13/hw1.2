// credit высчитывает ежемесяный аннуитетный платеж, переплату и итоговую сумму  по кредиту
package credit

import (
	"math"
)

// Calculate возвращает ежемесяный платеж, переплату и итоговую сумму
func Calculate(amount, percent, months int) (monthly float64, over, total int) {
	middleMonthPercent := calcMiddlePercentYearToMonths(percent) - 1
	// fmt.Printf("Процентная ставка из годoвого значения к месячному: %.5f\n", middleMonthPercent)
	coef := annuityСoefficient(middleMonthPercent, months)
	//fmt.Printf("Коэфициент аннуитета: %.5f\n", coef)
	monthly = float64(amount)* coef
	//fmt.Printf("Ежемесяный платеж %.2f\n", monthly)
	total = int(monthly * float64(months))
	//fmt.Printf("Итоговая выплата %v\n", total)
	over = total - amount
	//fmt.Printf("%.2f, %d, %d", monthly, over, total)

	return monthly, over, total
}

// calcMiddlePercentYearToMonths процентную ставку из годoвого значения к месячному
func calcMiddlePercentYearToMonths(percent int)  float64{
	const oneYear = 12
	middlePercent := 1 + float64(percent) / 100.0
	coefPercent := math.Pow(middlePercent, 1.0 / oneYear)
	return coefPercent
}
// annuityСoefficient высчитывает коэфициент аннуитета
func annuityСoefficient(middleMonthPercent float64, months int) float64 {
	k := math.Pow(1.0 + middleMonthPercent, float64(months))
	сoefAnutet := (middleMonthPercent * k) / (k - 1.0)
	return сoefAnutet
}

