package main

import "fmt"

func quickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    
    left, right := 0, len(arr)-1
    pivot := len(arr) / 2
    
    arr[pivot], arr[right] = arr[right], arr[pivot]
    
    for i := range arr {
        if arr[i] < arr[right] {
            arr[left], arr[i] = arr[i], arr[left]
            left++
        }
    }
    
    arr[left], arr[right] = arr[right], arr[left]
    
    quickSort(arr[:left])
    quickSort(arr[left+1:])
    
    return arr
}

func main() {
    arr := []int{10, 5, 2, 3, 8, 7, 1, 4, 6, 9}
    fmt.Println("До сортировки:", arr)
    quickSort(arr)
    fmt.Println("После сортировки:", arr)
}