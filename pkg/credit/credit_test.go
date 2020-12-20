package credit_test

import (
	"fmt"
	"github.com/ArtemBond13/hw1.2.git/pkg/credit"
)

func ExampleCalculate() {
	fmt.Println(credit.Calculate( 12_000_00, 6, 36))
	fmt.Println(credit.Calculate(1_000_000_00, 20, 36))
	fmt.Println(credit.Calculate(10_000_00, 20, 36))
	// Output:
	// 36419 111084 1311084
	// 3718400 33862300 133862300
	// 37184 338623 1338623
}
