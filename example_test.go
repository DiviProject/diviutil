package diviutil_test

import (
	"fmt"
	"math"

	"github.com/DiviProject/diviutil"
)

func ExampleAmount() {

	a := diviutil.Amount(0)
	fmt.Println("Zero Satoshi:", a)

	a = diviutil.Amount(1e8)
	fmt.Println("100,000,000 Satoshis:", a)

	a = diviutil.Amount(1e5)
	fmt.Println("100,000 Satoshis:", a)
	// Output:
	// Zero Satoshi: 0 DIVI
	// 100,000,000 Satoshis: 1 DIVI
	// 100,000 Satoshis: 0.001 DIVI
}

func ExampleNewAmount() {
	amountOne, err := diviutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := diviutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := diviutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := diviutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 DIVI
	// 0.01234567 DIVI
	// 0 DIVI
	// invalid bitcoin amount
}

func ExampleAmount_unitConversions() {
	amount := diviutil.Amount(44433322211100)

	fmt.Println("Satoshi to kDIVI:", amount.Format(diviutil.AmountKiloDIVI))
	fmt.Println("Satoshi to DIVI:", amount)
	fmt.Println("Satoshi to MilliDIVI:", amount.Format(diviutil.AmountMilliDIVI))
	fmt.Println("Satoshi to MicroDIVI:", amount.Format(diviutil.AmountMicroDIVI))
	fmt.Println("Satoshi to Satoshi:", amount.Format(diviutil.AmountSatoshi))

	// Output:
	// Satoshi to kDIVI: 444.333222111 kDIVI
	// Satoshi to DIVI: 444333.222111 DIVI
	// Satoshi to MilliDIVI: 444333222.111 mDIVI
	// Satoshi to MicroDIVI: 444333222111 ÎDIVIC
	// Satoshi to Satoshi: 44433322211100 Satoshi
}
