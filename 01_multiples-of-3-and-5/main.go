package main

import "fmt"

func main() {
	var summul int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			summul += i
		}
	}
	fmt.Println(summul)
}
