package main

import (
	"fmt"
	"math"
	"strconv"
)

// This takes integers in the bounds 0<= n <= 3999
func intToRoman(n int) string {
	specCaseMap := map[int]string{4: "IV", 9: "IX", 40: "XL", 90: "XC", 400: "CD", 900: "CM"}
	romanMap := [4]map[int]byte{
		{
			1: 'I',
			5: 'V',
		},
		{
			1: 'X',
			5: 'L',
		},
		{
			1: 'C',
			5: 'D',
		},
		{
			1: 'M',
		},
	}
	builtStr := "" //Empty String
	//For each digit in the number
	for idx, digit := range strconv.Itoa(n) {
		// We need to store the sigFigs
		// imagine idx is a pointer to current sig in list then we can do len - sig to get
		sigfig := len(fmt.Sprint(n)) - idx - 1
		digitInt := int((digit - '0'))

		// if digit is 4 or 9 it is a special case
		if digitInt == 4 || digitInt == 9 {
			// We need to find what the number is to access hmap
			// distance from the end
			builtStr += specCaseMap[digitInt*int(math.Pow10(sigfig))]
		} else {
			// Here we know that the number is not a special case

			//In the case the digit is greater than 5 or equal to 5 we know that we can use a whole number
			if digitInt >= 5 {
				builtStr += string(romanMap[sigfig][5])
				digitInt -= 5
			}
			// If it is zero dont do anything
			// if it greater 0 and less than 5 we know we have to iterate
			if digitInt != 0 && digitInt <= 5 {
				for x := 0; x < digitInt; x++ {
					builtStr += string(romanMap[sigfig][1])
				}
			}
		}

	}
	return builtStr
}
