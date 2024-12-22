package calculator_test

import (
	"errors"
	"github.com/Anti-Sh/go-rpn-calculator/pkg/calculator"
	"testing"
)

func TestCalculator_Execute(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   float64
		isError    bool
		err        error
	}{
		{"simple level - addition", "1+5", 6, false, nil},
		{"simple level - subtraction", "15-4", 11, false, nil},
		{"simple level - multiplication", "4*6", 24, false, nil},
		{"simple level - division", "16/2", 8, false, nil},
		{"simple level - degree", "4^3", 64, false, nil},

		{"simple+ level - complex op priority", "2 + 3 * 2", 8, false, nil},
		{"simple+ level - complex parentheses", "(2 + 3) * 2", 10, false, nil},
		{"simple+ level - complex nested parentheses", "(21+(17*2))-10", 45, false, nil},

		{"middle level - complex with all operations", "152+1400-10*1512/10-5^6", -15585, false, nil},
		{"middle level - complex nested parentheses", "(152+1400-10)*1512/((10-5)^6)", 149.216256, false, nil},

		{"error - division by zero", "100/0", 0, true, calculator.ErrDivisionByZero},
		{"error - invalid expression", "(133+23", 0, true, calculator.ErrUnknownToken},
		{"error - ErrUnknownToken", "133&23", 0, true, calculator.ErrInvalidExpression},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := calculator.NewCalculator(tt.expression)
			executionResult, err := calc.Execute()

			if tt.isError {
				if err == nil {
					t.Fatalf("Expression %s should return err", tt.expression)
				} else if !errors.Is(err, tt.err) {
					t.Fatalf("Expression %s should return error %v, got %v", tt.expression, tt.err.Error(), err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("Expression %s should not return err", tt.expression)
				}

				if executionResult != tt.expected {
					t.Fatalf("Expression %s should return %v, got %v", tt.expression, tt.expected, executionResult)
				}
			}
		})
	}
}
