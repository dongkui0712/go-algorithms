package main

/*
 * Bubble sort - http://en.wikipedia.org/wiki/Bubble_sort
 */

import "fmt"

import "../utils"

func main() {
    arr := utils.RandArray(10)
    fmt.Printf("Initial array is: %v\n", arr)
    
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr) - 1; j++ {
            if arr[j] > arr[j + 1] {
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
            }
        }
    }
    
    fmt.Printf("Sorted array is:  %v", arr)
}
