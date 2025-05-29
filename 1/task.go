package main

import "fmt"

func bubble(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    arr := []int{1, 4, 2, 12, 9, 3, 6}
    
    fmt.Println("Исходный массив:", arr)
    
    bubble(arr)
    
    fmt.Println("Отсортированный массив:", arr)
}