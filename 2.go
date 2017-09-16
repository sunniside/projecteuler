package main

import "fmt"

func getActiveFibonaccis(x,y int) (int, int) {
	return y, x+y
}

func main() {
	sum := 0;
	for prev,current := getActiveFibonaccis(0,1); current  < 4000000; prev, current = getActiveFibonaccis(prev,current) {
		if(current % 2 == 0) {
			sum += current;
		}
	}
	fmt.Println(sum);
}

