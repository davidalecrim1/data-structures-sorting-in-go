package main

import (
	"data_structures/helpers"
	"fmt"
	"os"
)


func chooseAlgorithm() int{
    var algorithm int
    fmt.Print(
              "1. Insertion Sort\n" +
              "2. Bubble Sort\n" +
              "3. Selection Sort\n" +
              "4. Quick Sort\n" +
              "5. Merge Sort\n" +
              "6. Heap Sort\n" +
              "7. All of them\n" +
              "Choose the Sorting Algorithm: ")

    _, err := fmt.Scanln(&algorithm)
    
    if err != nil {
        fmt.Println("Invalid input:", err)
        os.Exit(1)
    }

    return algorithm
}

func chooseSize() int{
    var size int
    fmt.Print("Enter the size of the array: ")

    _, err := fmt.Scanln(&size)

    if err != nil {
        fmt.Println("Invalid input:", err)
        os.Exit(1)
    }

    return size
}

func main() {
    algorithm := chooseAlgorithm()
    size := chooseSize()
    arr := helpers.GenerateRandomArray(size)

    switch algorithm {
    case 1:
        helpers.InsertionSortHelper(arr)
    case 2:
        helpers.BubbleSortHelper(arr)
    case 3:
        helpers.SelectionSortHelper(arr)
    case 4:
        helpers.QuickSortHelper(arr)
    case 5:
        helpers.MergeSortHelper(arr)
    case 6:
        helpers.HeapSortHelper(arr)
    case 7:
        helpers.InsertionSortHelper(arr)
        helpers.BubbleSortHelper(arr)
        helpers.SelectionSortHelper(arr)
        helpers.QuickSortHelper(arr)
        helpers.MergeSortHelper(arr)
        helpers.HeapSortHelper(arr)
    default:
        fmt.Println("Algorithm not implemented yet or invalid selection.")
    }
}