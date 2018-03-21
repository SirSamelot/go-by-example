package main

import "fmt"

func main() {
	// make empty slice
	s := make([]string, 3)

	// set and get slices
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// slice length
	fmt.Println("len:", len(s))

	// appending to slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy a slice (deep copy)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	d := c // looks like shallow copy
	fmt.Println("cpy2:", d)

	// modify original and copies
	s[0] = "z"
	c[1] = "y"
	d[2] = "x"
	fmt.Println("original:", s, "first copy:", c, "second copy:", d)

	// slice operator
	l := s[2:5]
	fmt.Println("sl1:", l)
}
