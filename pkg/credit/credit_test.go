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
	// 36506 114216 1314216
	// 3716358 33788888 133788888
	// 37163 337868 1337868
}
