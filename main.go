package main

import (
	"fmt"
	"go-rpn-calculator/calculator"
	"time"
)

func main() {
	start := time.Now()
	input := "9*(10+2)/6+8+19^12+14+20+15000+124+142+241421+412321*125/124+123^123"

	for i := 0; i < 1000; i++ {
		calc := calculator.NewCalculator(input)

		_, _ = calc.Execute()
	}
	//calc := calculator.NewCalculator(input)
	//
	//res, _ := calc.Execute()

	elapsed := time.Since(start)
	fmt.Printf("______________________\ncalc: %v \nresult: %f\n______________________\nspent time: %d mks", input, 0, elapsed.Microseconds())
}
