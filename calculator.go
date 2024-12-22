package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrDivisionByZero    = errors.New("division by zero")
)

func Calc(expression string) (float64, error) {
	tokens := tokenize(expression)
	return parseExpression(tokens)
}

func tokenize(expression string) []string {
	var tokens []string
	var current strings.Builder

	for _, ch := range expression {
		if unicode.IsSpace(ch) {
			continue
		}
		if isOperator(ch) || ch == '(' || ch == ')' {
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			tokens = append(tokens, string(ch))
		} else {
			current.WriteRune(ch)
		}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}

func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func parseExpression(tokens []string) (float64, error) {
	var stack []float64
	var ops []string

	for _, token := range tokens {
		if isNumber(token) {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		} else if token == "(" {
			ops = append(ops, token)
		} else if token == ")" {
			for len(ops) > 0 && ops[len(ops)-1] != "(" {
				result, err := applyOperator(&stack, &ops)
				if err != nil {
					return 0, err
				}
				stack = append(stack, result)
			}
			if len(ops) == 0 {
				return 0, errors.New("mismatched parentheses")
			}
			ops = ops[:len(ops)-1] // Убираем '('
		} else if isOperator(rune(token[0])) {
			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(token) {
				result, err := applyOperator(&stack, &ops)
				if err != nil {
					return 0, err
				}
				stack = append(stack, result)
			}
			ops = append(ops, token)
		} else {
			return 0, ErrInvalidExpression
		}
	}

	for len(ops) > 0 {
		result, err := applyOperator(&stack, &ops)
		if err != nil {
			return 0, err
		}
		stack = append(stack, result)
	}

	if len(stack) != 1 {
		return 0, ErrInvalidExpression
	}
	return stack[0], nil
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

func applyOperator(stack *[]float64, ops *[]string) (float64, error) {
	if len(*stack) < 2 {
		return 0, errors.New("not enough values in stack")
	}
	rhs := (*stack)[len(*stack)-1]
	lhs := (*stack)[len(*stack)-2]
	*stack = (*stack)[:len(*stack)-2]

	op := (*ops)[len(*ops)-1]
	*ops = (*ops)[:len(*ops)-1]

	switch op {
	case "+":
		return lhs + rhs, nil
	case "-":
		return lhs - rhs, nil
	case "*":
		return lhs * rhs, nil
	case "/":
		if rhs == 0 {
			return 0, ErrDivisionByZero
		}
		return lhs / rhs, nil
	}
	return 0, errors.New("invalid operator")
}

func IsValidExpression(expression string) bool {
	for _, r := range expression {
		if !(unicode.IsDigit(r) || r == '+' || r == '-' || r == '*' || r == '/' || r == '(' || r == ')') {
			return false
		}
	}
	return true
}
