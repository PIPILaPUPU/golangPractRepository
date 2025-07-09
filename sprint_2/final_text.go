package sprint2

import (
	"fmt"
)

const (
	CmdAdd   = iota // сложить два числа в стеке (top-1) = (top-1) + top
	CmdSub          // вычесть (top-1) = (top-1) - top
	CmdMul          // умножить (top-1) = (top-1) * top
	CmdDiv          // разделить (top-1) = (top-1) / top
	CmdPush         // поместить следующее число в стек
	CmdPop          // убрать число из стека
	CmdPrint        // печать верхнего элемента стека
	CmdSave         // сохранить число из стека в ячейку
	CmdLoad         // переместить число из ячейки в стек
)

func sprint2() {
	program := []int{CmdPush, 33, CmdPush, 44, CmdAdd, CmdPush, 567, CmdSub, CmdPush,
		-13, CmdMul, CmdPush, 5, CmdDiv, CmdPush, 45, CmdPush, 21, CmdAdd, CmdMul,
		CmdPrint, CmdSave, 'А', CmdPop, CmdPush, 3, CmdPush, 9, CmdPush, 7,
		CmdSub, CmdMul, CmdLoad, 'А', CmdMul, CmdPrint, CmdSave, 'Б',
		CmdLoad, 'А', CmdPush, 10230, CmdLoad, 'Б', CmdSub, CmdSub,
		CmdPush, 1000, CmdDiv, CmdPrint}

	// напишите код калькулятора
	stack := make([]int, 0, 100)
	registry := make(map[rune]int)
	var top int

	for opp := 0; opp < len(program); opp++ {

		switch program[opp] {
		case CmdAdd:
			top = len(stack) - 1
			stack[top-1] += stack[top]
			stack = stack[:top]
			break
		case CmdSub:
			top = len(stack) - 1
			stack[top-1] -= stack[top]
			stack = stack[:top]
			break
		case CmdMul:
			top = len(stack) - 1
			stack[top-1] *= stack[top]
			stack = stack[:top]
			break
		case CmdDiv:
			top = len(stack) - 1
			stack[top-1] /= stack[top]
			stack = stack[:top]
			break
		case CmdPush:
			opp++
			stack = append(stack, program[opp])
			break
		case CmdPop:
			top = len(stack) - 1
			stack = stack[:top]
			break
		case CmdPrint:
			top = len(stack) - 1
			fmt.Println(stack[top])
			break
		case CmdSave:
			opp++
			top = len(stack) - 1
			cell := rune(program[opp])
			registry[cell] = stack[top]
			break
		case CmdLoad:
			opp++
			top = len(stack) - 1
			cell := rune(program[opp])
			stack = append(stack, registry[cell])
			break
		default:
			continue
		}

	}
}
