package main

import "fmt"

func romanToInt(s string) int {
	// Here we use a map for O(1) access time to a match
	romanMap := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	//Holds the sum
	sum := 0
	//Here we use a for loop with an index because we want to always check the next value to see if it is valid.
	for idx := 0; idx < len(s)-1; idx++ {
		// check to see if current numeral is less than next one if so then it is a subtraction case
		if romanMap[rune(s[idx])] < romanMap[rune(s[idx+1])] {
			sum -= romanMap[rune(s[idx])]
		} else { //else we add this case
			sum += romanMap[rune(s[idx])]
		}
	}
	//Add the last string that was not accounted for since it will not be a subtraction case
	return sum + romanMap[rune(s[len(s)-1])]
}

func main() {
	fmt.Print(romanToInt("IVI"))
}
