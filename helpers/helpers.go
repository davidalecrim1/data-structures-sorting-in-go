package helpers

import (
	"data_structures/sorting"
	"fmt"
	"math/rand"
	"time"
)

func InsertionSortHelper(arr []int){
    start := time.Now()
    sorting.InsertionSort(arr)
    duration := time.Since(start)
    fmt.Printf("Time taken to sort the array using Insertion Sort: %v\n", duration)
}

func BubbleSortHelper (arr []int){
    start := time.Now()
    sorting.BubbleSort(arr)
    duration := time.Since(start)
    fmt.Printf("Time taken to sort the array using Bubble Sort: %v\n", duration)
}

func SelectionSortHelper(arr []int){
    start := time.Now()
    sorting.SelectionSort(arr)
    duration := time.Since(start)
    fmt.Printf("Time taken to sort the array using Selection Sort: %v\n", duration)
}

func QuickSortHelper(arr []int){
    start := time.Now()
    sorting.QuickSort(arr, 0, len(arr)-1)
    duration := time.Since(start)
    fmt.Printf("Time taken to sort the array using Quick Sort: %v\n", duration)
}

func MergeSortHelper(arr []int){
    start := time.Now()
    sorting.MergeSort(arr)
    duration := time.Since(start)
    fmt.Printf("Time taken to sort the array using Merge Sort: %v\n", duration)
}

func HeapSortHelper(arr []int){
    start := time.Now()
    sorting.HeapSort(arr)
    duration := time.Since(start)
    fmt.Printf("Time taken to sort the array using Heap Sort: %v\n", duration)
}

// generateRandomArray generates an array of a given size filled with random integers
func GenerateRandomArray(size int) []int {
    arr := make([]int, size)
    for i := range arr {
        arr[i] = rand.Intn(size * 10) // Random numbers between 0 and size*10
    }
    return arr
}