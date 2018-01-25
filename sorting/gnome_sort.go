package main

/*
 * Gnome sort - https://en.wikipedia.org/wiki/Gnome_sort
 */

import "fmt"

import "../utils"

func main() {
	arr := utils.RandArray(10)
    fmt.Println("Initial array is:", arr)
    fmt.Println("")
	
	for i := 1; i < len(arr); {
		if arr[i] >= arr[i - 1] {
			i++
		} else {
			arr[i], arr[i - 1] = arr[i - 1], arr[i]

			if i > 1 {
				i--
			}
		}
	}

	fmt.Println("Sorted array is: ", arr)
}

