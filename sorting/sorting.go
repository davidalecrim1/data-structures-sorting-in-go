package sorting

// swap exchanges the elements at indices i and j in the array
func swap(arr []int, i, j int) {
    arr[i], arr[j] = arr[j], arr[i]
}

// selectionSort sorts an array using the Selection Sort algorithm

// Selection Sort is a simple comparison-based sorting algorithm. Here's how it works:
// 1. Find the minimum element in the unsorted portion of the array.
// 2. Swap this minimum element with the first element of the unsorted portion.
// 3. Move the boundary of the sorted portion one element to the right.
// 4. Repeat the process until the entire array is sorted.
func SelectionSort(arr []int) {
    n := len(arr) // Get the length of the array

    // Iterate over the array. The outer loop tracks the sorted portion of the array.
    for i := 0; i < n-1; i++ {
        minIndex := i // Assume the current element (arr[i]) is the minimum

        // Iterate over the unsorted portion of the array to find the true minimum element
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j // Update minIndex if a smaller element is found
            }
        }

        // Swap the found minimum element with the first element of the unsorted portion
        swap(arr, i, minIndex)
    }
}

// insertionSort sorts an array using the Insertion Sort algorithm
//
// Insertion Sort is a simple comparison-based sorting algorithm. Here's how it works:
// 1. Start with the second element of the array (the first element is already sorted).
// 2. Compare the current element with the elements in the sorted portion of the array.
// 3. Shift the sorted elements to the right to make space for the current element.
// 4. Insert the current element into its correct position in the sorted portion.
// 5. Move to the next element and repeat the process until the entire array is sorted.
func InsertionSort(arr []int) {
    n := len(arr) // Get the length of the array

    // Iterate over the array starting from the second element
    for i := 1; i < n; i++ {
        key := arr[i] // Current element to be inserted into the sorted portion
        j := i - 1

        // Move elements of arr[0...i-1], that are greater than key, to one position ahead of their current position
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j = j - 1
        }
        arr[j+1] = key // Insert the key into its correct position
    }
}

// bubbleSort sorts an array using the Bubble Sort algorithm
//
// Bubble Sort is a simple comparison-based sorting algorithm. Here's how it works:
// 1. Start from the beginning of the array and compare each pair of adjacent elements.
// 2. Swap the elements if they are in the wrong order (the first element is greater than the second).
// 3. Move to the next pair of adjacent elements and repeat the process until the end of the array.
// 4. After each pass, the largest element in the unsorted portion bubbles up to its correct position.
// 5. Repeat the process for the remaining unsorted portion of the array until the entire array is sorted.
func BubbleSort(arr []int) {
    n := len(arr) // Get the length of the array

    // Iterate over the array n-1 times
    for i := 0; i < n-1; i++ {
        // Iterate over the unsorted portion of the array
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                // Swap the elements if they are in the wrong order
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

// quickSort sorts an array using the Quick Sort algorithm
//
// Quick Sort is a highly efficient sorting algorithm that uses a divide-and-conquer strategy.
// Here's how it works:
// 1. Choose a 'pivot' element from the array.
// 2. Partition the other elements into two sub-arrays: those less than the pivot and those greater than the pivot.
// 3. Recursively apply the same process to the two sub-arrays.
// 4. Combine the sorted sub-arrays and the pivot into a single sorted array.
//
// Quick Sort has an average time complexity of O(n log n), but it can degrade to O(n^2) in the worst case
// if the pivot selections are poor (e.g., always picking the smallest or largest element as the pivot).
func QuickSort(arr []int, low, high int) {
    if low < high {
        pi := partition(arr, low, high)
        QuickSort(arr, low, pi-1)
        QuickSort(arr, pi+1, high)
    }
}

// partition is a helper function for Quick Sort that partitions the array
//
// The partition function selects the last element as the pivot and places
// it in its correct position in the sorted array. All elements less than
// the pivot are moved to the left of the pivot, and all elements greater
// than the pivot are moved to the right.
func partition(arr []int, low, high int) int {
    pivot := arr[high]
    i := low - 1
    for j := low; j < high; j++ {
        if arr[j] < pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1
}

// mergeSort sorts an array using the Merge Sort algorithm
//
// Merge Sort is an efficient, stable, comparison-based sorting algorithm.
// Here's how it works:
// 1. Divide the unsorted list into n sublists, each containing one element.
// 2. Repeatedly merge sublists to produce new sorted sublists until there is only one sublist remaining.
// 3. This will be the sorted list.
//
// Merge Sort has a time complexity of O(n log n) in all cases (worst, average, and best).
func MergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    mid := len(arr) / 2
    left := MergeSort(arr[:mid])
    right := MergeSort(arr[mid:])

    return merge(left, right)
}

// merge is a helper function for Merge Sort that merges two sorted arrays into one sorted array
func merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))
    i, j := 0, 0

    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }

    result = append(result, left[i:]...)
    result = append(result, right[j:]...)

    return result
}

// heapSort sorts an array using the Heap Sort algorithm
//
// Heap Sort is a comparison-based sorting algorithm that uses a binary heap data structure.
// Here's how it works:
// 1. Build a max heap from the input data.
// 2. Repeat the following steps until the heap is empty:
//    a. Swap the root of the max heap with the last element of the heap.
//    b. Reduce the heap size by one.
//    c. Heapify the root of the tree.
//
// Heap Sort has a time complexity of O(n log n) in all cases (worst, average, and best).
func HeapSort(arr []int) {
    n := len(arr)
    // Build a max heap
    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }
    // One by one extract elements from heap
    for i := n - 1; i > 0; i-- {
        arr[0], arr[i] = arr[i], arr[0] // Move current root to end
        heapify(arr, i, 0)              // call max heapify on the reduced heap
    }
}

// heapify is a helper function for heapSort that maintains the max heap property
func heapify(arr []int, n int, i int) {
    largest := i       // Initialize largest as root
    left := 2*i + 1    // left = 2*i + 1
    right := 2*i + 2   // right = 2*i + 2

    // If left child is larger than root
    if left < n && arr[left] > arr[largest] {
        largest = left
    }

    // If right child is larger than largest so far
    if right < n && arr[right] > arr[largest] {
        largest = right
    }

    // If largest is not root
    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i] // swap

        // Recursively heapify the affected sub-tree
        heapify(arr, n, largest)
    }
}