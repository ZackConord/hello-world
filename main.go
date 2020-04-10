package main

import "fmt"

func printEverythingbaby(Everythingbaby string) {
	fmt.Println(Everythingbaby)
	var input string
	fmt.Scanln(&input)
	fmt.Println(fmt.Sprintf("%s is what you told me", "zack"))
}

func count(num int) {
	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}
}
func main() {
	printEverythingbaby("Phase 1")
	count(15)
	printEverythingbaby("Phase 2")

}
