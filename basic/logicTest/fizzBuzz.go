package main

import "fmt"


func main() {
	var n int
	fmt.Print("n : ")
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		if i%5== 0 && i%3== 0 {
			fmt.Print("FizzBuzz ")
		}else if i%5== 0 {
			fmt.Print("Buzz ")
		}else if i%3== 0 {
			fmt.Print("Fizz ")
		}else{
			fmt.Print(i, " ")
		}
		
	}
}
