package main

import "fmt"

func main() {
	fmt.Println("hellooooo world")
	fmt.Println("testing line 2")
	// format verbs decimal, binary, octal with 0X preface
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
	for i := 97; i <= 122; i++ {
		fmt.Printf("%d \t %q \n", i, i)
	}
}
