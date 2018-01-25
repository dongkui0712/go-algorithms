package main

/*
 * Cocktail sort - https://en.wikipedia.org/wiki/Cocktail_sort
 */

import "fmt"

import "../utils"

func main() {
    arr := utils.RandArray(11)
    fmt.Printf("Initial array is: %v\n", arr)
    
    for i := 0; i < len(arr) / 2; i++ {
        fmt.Printf("\nloop: %d\n", i+1)

        left := 0
        right := len(arr) - 1
        
        for ; left <= right ; {
            
            if arr[left] > arr[left + 1] {
                arr[left], arr[left + 1] = arr[left + 1], arr[left]
            }
			
            left++
            
            if arr[right - 1] > arr[right] {
                arr[right - 1], arr[right] = arr[right], arr[right - 1]
            }

            right--

            // fmt.Printf("left=[%d], right=[%d], arr: %v\n", left, right, arr)
        }
    }
    
    fmt.Println("Sorted array is: ", arr)
}
