package main

import "fmt"

func main() {
	// basic loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic looop with initial/condition/after
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// loop without condition
	for {
		fmt.Println("loop")
		break
	}

	// continue to next iteration of loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
