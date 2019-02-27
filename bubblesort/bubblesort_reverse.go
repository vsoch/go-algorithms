package main

// go run bubblesort.go

import "fmt"

var sortMe = []int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}

// swap two items in an array based on an index (idx)
func swap(input []int, idx int) {
        input[idx-1], input[idx] = input[idx], input[idx-1]
}

func bubbleSort(input []int) {

	var switched bool = true
        var iter int = 0;

	// while the list has a switch done
	for switched {

		switched = false

		// loop through the list and bubble sort
		for i := len(input) - 1; i > 0; i-- {

			// if the current element is greater than the next, swap and flag
			if input[i-1] < input[i] {

                                // This is an in place switch
                                swap(input, i)
				switched = true
			}

		}
        	fmt.Println("Iteration", iter, input)
                iter +=1
	}
	// Print the answer for the user
	fmt.Println("The   sorted list is: ", input)
}

func main() {
	fmt.Println("Starting bubble sort!")
	fmt.Println("The original list is: ", sortMe)
	bubbleSort(sortMe)
}
