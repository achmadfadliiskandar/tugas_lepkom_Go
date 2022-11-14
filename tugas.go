package main

import "fmt"

func main() {
	// variabel
	var a, b, c, x, y, z int
	a = 12
	b = 8
	c = a + b
	x = a + b - c
	y = a * b
	z = a / b
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	// operator asigment

	// looping golang
	for i := 1; i <= 10; i++ {
		fmt.Println("i ke ", i)
	}

	// conditional golang
	variabel := 10
	if variabel > 10 {
		fmt.Println("you big")
	} else {
		fmt.Println("you small")
	}

	// switch case golang
	var nilai int
	nilai = 90

	switch nilai {
	case 100:
		fmt.Println("Excellent")
		break
	case 90:
		fmt.Println("Good")
		break
	case 80:
		fmt.Println("Better")
		break
	case 70:
		fmt.Println("Sorry You Fail")
		break
	}
}
