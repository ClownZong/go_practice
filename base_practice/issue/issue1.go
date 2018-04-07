// 深入理解slice，函数传参

package main

import "fmt"

func DeleteSliceByIndex (a []int, index int) []int {
	if index >= len(a) {
		return []int{}
	}
	a = append(a[:index], a[index+1:]...)
	a = append(a[:index], a[index+2:]...)
	return a
}

func DeleteSliceByIndex1 (a *[]int, index int) *[]int {
	if index >= len(*a) {
		return &[]int{}
	}
	*a = append((*a)[:index], (*a)[index+1:]...)
	*a = append((*a)[:index], (*a)[index+2:]...)
	return a
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	index := 3
	fmt.Printf("slice is %v\n", slice)
	fmt.Printf("remove slice[%d]; slice is %v\n", index, DeleteSliceByIndex(slice, index))
	fmt.Printf("slice is %v\n", slice)
	fmt.Println("-----------------------------------------------------")
	slice1 := []int{0, 1, 2, 3, 4, 5, 6}
	index1 := 3
	fmt.Printf("slice1 is %v\n", slice1)
	fmt.Printf("remove slice1[%d]; slice1 is %v\n", index1, *DeleteSliceByIndex1(&slice1, index1))
	fmt.Printf("slice1 is %v\n", slice1)
}
