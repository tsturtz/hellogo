package main

import "fmt"

func main() {
	fmt.Println("hellooooo go!")
	fmt.Println("testing line 2")

	// format verbs decimal, binary, octal with 0X preface
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	for i := 65; i <= 90; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
	for i := 97; i <= 122; i++ {
		fmt.Printf("%d \t %q \n", i, i)
	}

	// assign
	var e string
	// initialize
	e = "some string"
	// assign and init
	var f = "another string"

	// shorthand variable declaration - infers type
	a := 10
	b := "golang is rad"
	c := 4.17
	d := true

	var g float64

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%v \n", g)
}
