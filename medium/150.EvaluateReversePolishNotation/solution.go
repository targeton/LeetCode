package solution

import (
	"container/list"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := list.New()
	for _, str := range tokens {
		if str == "+" || str == "-" || str == "*" || str == "/" {
			oper1 := stack.Back().Value.(int)
			stack.Remove(stack.Back())
			oper2 := stack.Back().Value.(int)
			stack.Remove(stack.Back())
			switch str {
			case "+":
				stack.PushBack(oper2 + oper1)
			case "-":
				stack.PushBack(oper2 - oper1)
			case "*":
				stack.PushBack(oper2 * oper1)
			case "/":
				stack.PushBack(oper2 / oper1)
			}
		} else {
			oper, _ := strconv.Atoi(str)
			stack.PushBack(oper)
		}
	}
	return stack.Back().Value.(int)
}
