package main

import (
	"fmt"
	"github.com/ArtemBond13/hw1.2.git/pkg/credit"
)

func main() {
	amount := 12000
	percent := 6
	months := 36

	monthly, total, over := credit.Calculate(amount, percent, months)
	fmt.Println(monthly, total, over)

}