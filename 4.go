package main

import "fmt"
import "strconv"

func main() {
	fmt.Println(largestThreeDigitPalindrome())
}

func largestThreeDigitPalindrome() (int) {
	var largestPalindrome int
	for f1 := 999; f1 > 0; f1-- {
		for f2 := 999; f2 > 0; f2-- {
			candidate := f1 * f2
			if candidate > largestPalindrome && isPalindrome(candidate) {
				largestPalindrome = candidate
			}
		}
	}
	return largestPalindrome;
}

func isPalindrome(candidate int) (bool) {
	candidateStr := strconv.Itoa(candidate)
	numberOfDigits := len(candidateStr)
	var numberOfComparisons int
	if numberOfComparisons = numberOfDigits / 2; numberOfDigits % 2 != 0  {
		numberOfComparisons = (numberOfDigits - 1) / 2
	}
	for i, j := 0, numberOfDigits-1; i <= numberOfComparisons; i, j = i+1, j-1 {
		if(candidateStr[i] != candidateStr[j]) {
			return false
		}
	}
	return true;
}

