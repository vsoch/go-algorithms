# Bubblesort

The first effort, [bubblesort.go](bubblesort.go), iterates through the list
and compares each pair, switching the two if the first is greater than the second.
We keep going until there are no switches for an entire run, indicating
that the list is sorted.

Then I realized that we don't need to go to the end every time, because the biggest
gets pushed to the end. Each time we only need to go up until the length minus
whatever index we are looking at.
