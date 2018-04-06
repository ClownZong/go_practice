// 函数、传值与传切片
package main

import (
	"fmt"
)

// 比较大小
func MyMax(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// 传值
func SortMaxInt (a, b int) (c, d int) {
	if a < b {
		return b, a
	}
	return a, b
}

// 传切片，修改底层数组
func SortMax(sl []int) []int{
	lenSl := len(sl)
	if lenSl == 0 {
		return sl
	}
	for i := 0; i < lenSl; i++ {
		for j := i+1; j < lenSl; j++ {
			if sl[i] < sl[j] {
				sl[i], sl[j] = sl[j], sl[i]
			}
		}
	}
	return sl
}

// 传切片，用拷贝来避免修改原切片的底层数组
func SortMin(mysll []int) []int{
	mysl := make([]int, len(mysll))
	copy(mysl, mysll)
	lenMySl := len(mysl)
	if lenMySl == 0 {
		return mysl
	}
	for i := 0; i < lenMySl; i++ {
		for j := i+1; j < lenMySl; j++ {
			if mysl[i] > mysl[j] {
				mysl[i], mysl[j] = mysl[j], mysl[i]
			}
		}
	}
	return mysl
}

func main() {
	// 比较大小，返回最大
	a, b := 1, -1
	fmt.Printf("a=%d, b=%d, max=%d\n", a, b, MyMax(int64(a), int64(b)))

	// 将两个数排序，注意有没有修改原来的输入值
	c, d := 1, 3
	e, f := SortMaxInt(c, d)
	fmt.Printf("c=%d\t d=%d\t e=%d\t f=%d\t\n", c, d, e, f)

	// 将切片进行冒泡排序，注意有没有修改原来的输入值
	orgSlice := []int{8, 3, 4, 2, 29, 7, 23, 89, -2, 0, 5}
	copyOrgSlice := make([]int, len(orgSlice))
	// copyOrgSlice = append(copyOrgSlice, orgSlice...)  // 拷贝副本
	copy(copyOrgSlice, orgSlice)
	fmt.Printf("orgSlice=%v\n", orgSlice)
	sortSliceMax := SortMax(orgSlice)
	fmt.Printf("orgSlice=%v\t sortSliceMax=%v\n", orgSlice, sortSliceMax)
	// 将切片进行冒泡排序，注意有没有修改原来的输入值
	fmt.Printf("copyOrgSlice=%v\n", copyOrgSlice)
	sortSliceMin := SortMin(copyOrgSlice)
	fmt.Printf("copyOrgSlice=%v\t sortSliceMin=%v\n", copyOrgSlice, sortSliceMin)

}