package main

import "fmt"

func main() {
	printEverythingbaby("Phase 1")
	printEverythingbaby("Phase 2")
	count(10)
}

func printEverythingbaby(Everythingbaby string) {
	fmt.Println(Everythingbaby)
	var input string
	fmt.Scanln(&input)
	fmt.Println(fmt.Sprintf("%s is what you told me", input))
}

func count(num int) {
	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}
}
