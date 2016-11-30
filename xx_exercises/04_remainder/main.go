package main

import "fmt"

func main() {
	var smallnum int
	var largenum int

	fmt.Print("Enter a large number: ")
	fmt.Scan(&largenum)
	fmt.Print("Enter a small number: ")
	fmt.Scan(&smallnum)
	fmt.Println(largenum, "%", smallnum, " = ", largenum%smallnum)

}
