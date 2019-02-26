package main

// go run bubblesort.go

import "fmt"

var sortMe = []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}

func bubbleSort(input []int) {

	var switched bool = true
        var iter int = 0;

	// while the list has a switch done
	for switched {

		switched = false

		// loop through the list and bubble sort
		for i := 1; i < len(input); i++ {


                        // We can ignore the rest of list if i == iter
                        if iter == i {
                                break
                        }

			// if the current element is greater than the next, swap and flag
			if input[i-1] > input[i] {

                                // This is an in place switch
				input[i-1], input[i] = input[i], input[i-1]
				switched = true
			}

		}

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
