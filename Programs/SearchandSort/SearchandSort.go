package main

import (
	"fmt"

	"github.com/SearchandSort"
)

func main() {
	fmt.Printf("Enter 0 for binary search of integer\n")
	fmt.Printf("Enter 1 for binary search of string\n")
	fmt.Printf("Enter 2 for insertion sort of integer\n")
	fmt.Printf("Enter 3 for insertion sort of string\n")
	fmt.Printf("Enter 4 for bubble sort of string\n")
	fmt.Printf("Enter 5 for bubble sort integer\n")

	fmt.Printf("\nEnter your Choice : ")
	var selection int
	fmt.Scanf("%d", &selection)

	switch selection {
	case 0:
		SearchandSort.BinarySearchint()
		break
	case 1:
		SearchandSort.BinarySearchString()
	case 2:
		SearchandSort.InsertionInt()
	case 3:
		SearchandSort.InsertionString()
	case 4:
		SearchandSort.BubblesortString()
	case 5:
		SearchandSort.BubblesortInt()
	}
}
