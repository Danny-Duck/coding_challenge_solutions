// Given two arrays of strings a1 and a2 return a sorted array r in lexicographical order of the strings of a1 which are substrings of strings of a2.

// #Example 1: a1 = ["arp", "live", "strong"]

// a2 = ["lively", "alive", "harp", "sharp", "armstrong"]

// returns ["arp", "live", "strong"]

// #Example 2: a1 = ["tarp", "mice", "bull"]

// a2 = ["lively", "alive", "harp", "sharp", "armstrong"]

// returns []

package main

import "fmt"

// InArray returns a sorted array in lexicographical order?
func InArray(array1 []string, array2 []string) []string {

	return array1
}

func main() {
	testStringArr1 := []string{"tarp", "mice", "bull"}
	testStringArr2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}
	fmt.Println(InArray(testStringArr1, testStringArr2))
}
