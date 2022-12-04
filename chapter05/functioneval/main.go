package main

import (
	"chapter05/nodestack"
	"fmt"
)

func precedence(symbolOne, symbolTwo string) bool {
	// Returns true if symbolOne has a higher precedence than symbolTwo
	if (symbolOne == "+" || symbolOne == "-") && (symbolTwo == "(" || symbolTwo == "/") {
		return false
	} else if (symbolOne == "(" && symbolTwo != ")") || symbolTwo == "(" {
		return false
	} else {
		return true
	}
}

func isPresent(symbol string, operators []string) bool {
	for index := 0; index < len(operators); index++ {
		if symbol == string(operators[index]) {
			return true
		}
	}

	return false
}

func infixpostfix(infix string) (postfix string) {
	operators := []string{"+", "-", "*", "/", "(", ")"}
	postfix = ""
	nodeStack := nodestack.Stack[string]{}
	for index := 0; index < len(infix); index++ {
		newSymbol := string(infix[index])
		if newSymbol == " " || newSymbol == "\n" {
			continue
		}
		if newSymbol >= "a" && newSymbol <= "z" {
			postfix += newSymbol
		}
		if isPresent(newSymbol, operators) {
			if !nodeStack.IsEmpty() {
				topSymbol := nodeStack.Top()
				if precedence(topSymbol, newSymbol) {
					if topSymbol != "(" {
						postfix += topSymbol
					}
					nodeStack.Pop()
				}
			}
			if newSymbol != ")" {
				nodeStack.Push(newSymbol)
			} else { // Pop nodeStack down to first left parenthesis
				for {
					if nodeStack.IsEmpty() {
						break
					}
					ch := nodeStack.Top()
					if ch != "(" {
						postfix += ch
						nodeStack.Pop()
					} else {
						nodeStack.Pop()
						break
					}
				}
			}
		}
	}
	for {
		if nodeStack.IsEmpty() {
			break
		}
		if nodeStack.Top() != "(" {
			postfix += nodeStack.Top()
			nodeStack.Pop()
		}
	}
	return postfix
}

var values map[string]float64

func evaluate(postfix string) float64 {
	operandStack := nodestack.Stack[float64]{}
	for index := 0; index < len(postfix); index++ {
		ch := string(postfix[index])
		if ch >= "a" && ch <= "z" {
			operandStack.Push(values[ch])
		} else { // ch is an operator
			operand1 := operandStack.Pop()
			operand2 := operandStack.Pop()
			if ch == "+" {
				operandStack.Push(operand1 + operand2)
			} else if ch == "-" {
				operandStack.Push(operand2 - operand1)
			} else if ch == "*" {
				operandStack.Push(operand1 * operand2)
			} else if ch == "/" {
				operandStack.Push(operand2 / operand1)
			}
		}
	}
	return operandStack.Top()
}

func main() {
	postfix := infixpostfix("a + (b - c) / (d * e)")
	fmt.Println(postfix)
	values = make(map[string]float64)
	values["a"] = 10
	values["b"] = 5
	values["c"] = 2
	values["d"] = 4
	values["e"] = 3
	result := evaluate(postfix)
	fmt.Println("function evaluates to: ", result)
}

// Output: abc-de*/+
// function evaluates to: 10.25
