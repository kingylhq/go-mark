package example

import "fmt"

func For() {

	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Another way of accomplishing the basic "do this
	// N times" iteration is `range` over an integer.
	arr := []uint{7, 10, 5, 2, 3}
	for i := range arr {
		fmt.Println("range", i)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	//for {
	//	fmt.Println("loop")
	//	break
	//}

	// You can also `continue` to the next iteration of
	// the loop.
	//arr := []uint{7, 9 ,5, 2, 3}
	for n, value := range arr {
		if value%2 == 0 {
			continue
		}
		fmt.Printf("n = %v, value = %v \n", n, value)
	}
}
