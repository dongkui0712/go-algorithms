package main

/*
 * Counting sort - https://en.wikipedia.org/wi  ki/Counting_sort
 */

import "fmt"

import "../utils"

func countSort(in []int) (out []int) {
    out = make([]int, len(in))
    if len(in) == 0 {
        return
    }
    min, max := in[0], in[0]
    for _, v := range in {
        if v < min {
            min = v
        }
        if v > max {
            max = v
        }
    }
    k := max - min + 1
    fmt.Printf("min=[%d], max=[%d], k=[%d]\n", min, max, k)

    counts := make([]int, k)
    for _, v := range in {
        counts[v-min]++
    }

    for i := 1; i < len(counts); i++ {
        counts[i] += counts[i-1]
    }

    for _, v := range in {
        out[counts[v-min]-1] = v
    }
    fmt.Printf("results: %v\n", out)
    return
}

func countSort2(in []int) (out []int) {
    out = make([]int, len(in))
    if len(in) == 0 {
        return
    }
    min, max := in[0], in[0]
    for _, v := range in {
        if v < min {
            min = v
        }
        if v > max {
            max = v
        }
    }
    k := max - min + 1
    fmt.Printf("min=[%d], max=[%d], k=[%d]\n", min, max, k)

    counts := make([]int, k)
    for _, v := range in {
        counts[v-min]++
    }

    var j int
    for vaule, count := range counts {
        for ; count > 0; j++ {
            out[j] = vaule + min
            count--
        }
    }

    fmt.Printf("results: %v\n", out)
    return
}

func main(){
    arr := utils.RandArray(10)
    fmt.Printf("Initial array is: %v\n\n", arr)
   
    arr2 := arr[0: len(arr)]
    countSort(arr)
    countSort2(arr2)
}
