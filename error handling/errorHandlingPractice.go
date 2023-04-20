package main

import (
	"errors"
	"fmt"
)

func main() {
	age := 23
	funcRes, errPart := firstFunction(age)
	fmt.Println(funcRes, errPart) // 0 Age is greater than 15 which is not allowed

	a, b := 12, 0
	divRes, divErrPart := division(a, b)
	_ = divRes

	fmt.Println()
	if divErrPart != nil {
		// checking error type
		switch divErrPart.(type) {
		case *DivErr:
			fmt.Println("its an div error which we defined", divErrPart)
			break
		default:
			fmt.Println("its not an div error which we defined", divErrPart)

		}
		// divErr := DivErr{a: a, b: b}
		// if errors.Is(divErrPart, &DivErr{a: a, b: b}) {
		// 	fmt.Println("its an div error which we defined", divErrPart)
		// } else {
		// 	fmt.Println("its not an div error which we defined", divErrPart)

		// }
	}
}

func firstFunction(age int) (int, error) {
	if age > 15 {
		return 0, errors.New("Age is greater than 15 which is not allowed")
	} else {
		return age + 12, nil
	}
}

type DivErr struct {
	a int
	b int
}

func (d *DivErr) Error() string {
	return fmt.Sprintf("Cannot Divide %v by %v", d.a, d.b)
}

func division(a int, b int) (int, error) {
	if b == 0 {
		return 0, &DivErr{a, b}
	} else {
		return a / b, nil
	}
}
