package utils

import (
    "math/rand"
    "time"
)

func RandArray(n int) []int {
    r := rand.New(rand.NewSource(time.Now().Unix()))
    arr := make([]int, n)
    for i := 0; i <= n - 1; i++ {
        arr[i] = r.Intn(100)
    }
    return arr
}
