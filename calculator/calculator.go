package calculator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Calculator struct {
	infixExpr   string
	postfixExpr string
}

func NewCalculator(expr string) *Calculator {
	c := &Calculator{infixExpr: expr}
	c.updatePostfix()

	return c
}

func (c *Calculator) Execute() (float64, error) {
	var stack []float64

	tokens := strings.Fields(c.postfixExpr)

	for _, token := range tokens {
		if isOperator(rune(token[0])) {
			// Если токен - оператор, извлекаем два верхних элемента из стека
			if len(stack) < 2 {
				return 0, fmt.Errorf("недостаточно операндов для операции %s", token)
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// Выполняем операцию и помещаем результат обратно в стек
			result := opEval(rune(token[0]), a, b)
			stack = append(stack, result)
		} else {
			// Если токен - операнд, преобразуем его в число и помещаем в стек
			value, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("ошибка преобразования токена %s: %v", token, err)
			}
			stack = append(stack, value)
		}
	}

	// В конце в стеке должен остаться один элемент - результат
	if len(stack) != 1 {
		return 0, fmt.Errorf("неверное постфиксное выражение")
	}

	return stack[0], nil
}

func (c *Calculator) updatePostfix() {
	var stack []rune
	var postfix []string

	// Регулярное выражение для поиска операндов и операторов
	re := regexp.MustCompile(`(\d+|[+\-*/^()])`)
	tokens := re.FindAllString(c.infixExpr, -1)

	for _, token := range tokens {
		if isOperator(rune(token[0])) {
			// Если токен - оператор
			for len(stack) > 0 && opPriority(stack[len(stack)-1]) >= opPriority(rune(token[0])) {
				if token[0] == '^' && stack[len(stack)-1] == '^' {
					break // Правосторонний оператор
				}
				postfix = append(postfix, string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, rune(token[0]))
		} else if token == "(" {
			// Если токен - открывающая скобка
			stack = append(stack, '(')
		} else if token == ")" {
			// Если токен - закрывающая скобка
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				postfix = append(postfix, string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] // Удаляем открывающую скобку
		} else {
			// Если токен - операнд
			postfix = append(postfix, token)
		}
	}

	for len(stack) > 0 {
		postfix = append(postfix, string(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}

	c.postfixExpr = strings.Join(postfix, " ")
}
