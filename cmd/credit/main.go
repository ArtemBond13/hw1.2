package main

import (
	"fmt"
	"github.com/ArtemBond13/hw1.2.git/pkg/credit"
)

func main() {
	amount := 1_000_000_00
	percent := 20
	months := 12

	monthly, total, over := credit.Calculate(amount, percent, months)
	fmt.Println(monthly, total, over)

}
