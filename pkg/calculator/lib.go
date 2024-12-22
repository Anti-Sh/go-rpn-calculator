package calculator

import "math"

func opPriority(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3
	}
	return 0
}

func opEval(op rune, a, b float64) (float64, error) {
	switch op {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	case '^':
		return math.Pow(a, b), nil
	default:
		return 0, ErrUnknownOperator
	}
}

func isOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/' || c == '^'
}
