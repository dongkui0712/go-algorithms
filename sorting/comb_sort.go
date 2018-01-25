package main

/*
 * Comb sort - https://en.wikipedia.org/wiki/Combsort
 */

import "fmt"

import "../utils"

func main() {
	arr := utils.RandArray(10)
	fmt.Printf("Initial array is: %v\n", arr)

	gap := len(arr)

	for gap > 1 {
		gap = gap * 100 / 124

		for i := 0; i+gap < len(arr); i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
			}
		}
		fmt.Printf("gap=[%d], arr: %v\n", gap, arr)
	}

	fmt.Println("Sorted array is: ", arr)
}
