// challenges/generics/begin/main.go
package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring

// non-generic print functions

// func printString(s string) { fmt.Println(s) }

// func printInt(i int) { fmt.Println(i) }

// func printBool(b bool) { fmt.Println(b) }

func PrintAny [T string | int | bool] (val T) {
    fmt.Println(val)
}

// Part 2 sum function refactoring
type numeric interface {
//	int | float32 | float64
    constraints.Integer | constraints.Float
}

// sum sums a slice of any type
func sum[T numeric](numbers ...T) T {
	var result T
	for _, n := range numbers {
			result += n
		}
	return result
}

func sumOrig(numbers []interface{}) interface{} {
	var result float64
	for _, n := range numbers {
		switch n.(type) {
		case int:
			result += float64(n.(int))
		case float32, float64:
			result += n.(float64)
		default:
			continue
		}
	}
	return result
}


func main() {
	// call non-generic print functions
	// printString("Hello")
	// printInt(42)
	// printBool(true)

	// call generic printAny function for each value above
	PrintAny("Hello")
	PrintAny(42)
	PrintAny(true)



	fmt.Println()
	fmt.Println("Integer-ish")
	fmt.Println("result", sumOrig([]interface{}{1, 2, 3}))
	fmt.Println("result", sum(1, 2, 3))


	fmt.Println()
	fmt.Println("Float-ish")
	fmt.Println("result", sumOrig([]interface{}{4.1, 5.2, 6.3}))
	fmt.Println("result", sum(4.1, 5.2, 6.3))
}
