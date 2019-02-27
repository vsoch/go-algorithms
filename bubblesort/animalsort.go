package main

// go run bubblesort.go

import "fmt"

var sortMe = []string{"dog", "cat", "alligator", "cheetah", "rat", "moose", "cow", "mink", "porcupine", "dung beetle", "camel", "steer", "bat", "hamster", "horse", "colt", "bald eagle", "frog", "rooster"}

// swap two items in an array based on an index (idx)
func swap(input []string, idx int) {
	input[idx-1], input[idx] = input[idx], input[idx-1]
}

// return true if string "a" is greater than b.
func greaterThan(a, b string) bool {
	if a[0] > b[0] {
		return true
	}
	return false
}

func bubbleSort(input []string) {

	var switched bool = true
	var iter int = 0

	// while the list has a switch done
	for switched {

		switched = false

		// loop through the list and bubble sort
		for i := 1; i < len(input); i++ {

			// if the current element is greater than the next, swap and flag
			if greaterThan(input[i-1], input[i]) {

				// This is an in place switch
				swap(input, i)
				switched = true
			}

		}
		fmt.Println("Iteration", iter, input)
		iter += 1
	}
	// Print the answer for the user
	fmt.Println("The   sorted list is: ", input)
}

func main() {
	fmt.Println("Starting bubble sort!")
	fmt.Println("The original list is: ", sortMe)
	bubbleSort(sortMe)
}
