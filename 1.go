package main

import "fmt"

func main() {
	fmt.Println(findAndAddMultiples(3,5,1000));
}

func findAndAddMultiples(multiple1, multiple2, max int) (sum int) {
	sum = 0;
	for i := max-1; i > 0; i-- {
		if(i % multiple1 == 0 || i % multiple2 == 0) {
			sum += i;
		}
	}
	return;
}
