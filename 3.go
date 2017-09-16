package main

import "fmt"

func isPrime(primeCandidate int) (bool) {
	y:= primeCandidate-1;
	for y > 1 {
		if(primeCandidate % y == 0) {
			return false;
		}
	}
	return true;
}

func getLargestPrime(x int) (int) {
	y := x-1;
	for y > 1 {
		if(x % y == 0 && isPrime(y)) {
			fmt.Println(x,y);
			return y;
		}
		y--;
	}
	return -1;
}

func main() {
	fmt.Println(getLargestPrime(13195));
}
