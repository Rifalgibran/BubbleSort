package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}
	}
}
func main() {
	bilangan := []int{70, 65, 34, 25, 50, 10, 80}

	fmt.Println("Bilangan sebelum diurutkan", bilangan)
	bubbleSort(bilangan)
	fmt.Println("Bilangan sesudah di urutkan dari besar ke kecil:", bilangan)

	slice := []string{"a", "b", "c", "d", "e"}
	fmt.Println("Huruf sebelum diurutkan :", slice)
	slice = append(slice, "f", "g", "h")
	fmt.Println("Huruf sesudah di urutkan:", slice)

}
