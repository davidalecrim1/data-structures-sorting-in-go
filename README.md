# Data Structures in Go

Just trying and learning about it with Go.

## Getting Started

1. Open the Terminal
2. go run main.go

## Sorting Theory

### Selection Sort:

Time Complexity: **O(n^2)** in all cases (worst, average, and best). It has nested loops and always iterates through the entire array.

**How it works:**
1. Iterate over the array to find the smallest element.
2. Swap the smallest element with the first element.
3. Move to the next element and repeat the process, treating the already sorted portion separately.
4. Continue until the entire array is sorted.

### Bubble Sort:

Time Complexity: **O(n^2)** in all cases (worst, average, and best). Similar to Selection Sort, it has nested loops and always iterates through the entire array.

**How it works:**
1. Iterate over the array and compare adjacent elements.
2. Swap adjacent elements if they are in the wrong order.
3. Repeat this process until no more swaps are needed, indicating the array is sorted.

### Insertion Sort:

Time Complexity: **O(n^2)** in the worst and average cases. However, in the best case (when the array is already sorted), it has a time complexity of **O(n)**. It performs well on small arrays or nearly sorted arrays.

**How it works:**
1. Iterate over the array, starting from the second element.
2. Insert each element into its correct position among the elements to its left that are already sorted.
3. Continue until the entire array is sorted.

### Merge Sort:

Time Complexity: **O(n log n)** in all cases (worst, average, and best). It divides the array into halves recursively and then merges the sorted halves, resulting in a time complexity of **O(n log n)**.

**How it works:**
1. Divide the array into halves recursively until each sub-array has one element.
2. Merge the sorted sub-arrays pairwise until there is only one sorted array remaining.

### Quick Sort:

Time Complexity: **O(n log n)** in the average and best cases. However, in the worst case, it can degrade to **O(n^2)** if the pivot selection is poor and it always picks the smallest or largest element as pivot. Despite this, Quick Sort is often faster in practice than other **O(n^2)** algorithms like Selection Sort and Bubble Sort.

**How it works:**
1. Select a pivot element from the array.
2. Partition the array into two sub-arrays: elements less than the pivot and elements greater than the pivot.
3. Recursively apply the same process to the sub-arrays.
4. Combine the sorted sub-arrays with the pivot to form the final sorted array.

### Heap Sort:

Time Complexity: **O(n log n)** in all cases (worst, average, and best). It builds a heap data structure and repeatedly extracts the maximum element to sort the array.

**How it works:**
1. Build a max-heap from the input array.
2. Extract the maximum element (root) from the heap and place it at the end of the array.
3. Heapify the remaining elements.
4. Repeat steps 2 and 3 until the heap is empty.

### Conclusion

These time complexities provide an understanding of how the sorting algorithms perform as the size of the input array increases. In practice, the choice of sorting algorithm depends on various factors such as the size of the data, the nature of the data (nearly sorted, mostly sorted, or completely random), memory constraints, and the need for stability (maintaining the relative order of equal elements).

### Online Examples
- ![ComparisonSort Website](https://www.cs.usfca.edu/~galles/visualization/ComparisonSort.html)