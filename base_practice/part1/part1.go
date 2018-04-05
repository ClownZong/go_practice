package main

import (
	"fmt"
	"time"
	"reflect"
)
// 基本类型
// var a int
// var s string
// var f float64
// var b bool
// var c byte

// 数组初始化和赋值
// var array = [3]int{1, 2, 3}
// var ar1 [3]int
// var ar2 = [2]string{"apple","oragin"}
// var ar3 [2]string

// 切片
// var slice []int


func main() {
	fmt.Println("------------------------ start -----------------")
	fmt.Println("hello", "现在是：", time.Now())
	fmt.Println("------------------------ end -----------------")

	// 如何定义切片和初始化，打印长度、容量、类型和内容
	slice := []int{1, 2, 3, 4}  // 初始化和赋值切片
	sliceLen := len(slice)  // 切片长度
	sliceCap := cap(slice)  // 切片容量
	sliceType := reflect.TypeOf(slice).String()  // 类型
	fmt.Printf("slice len is:%d , cap is %d , type is %s; slice is %v\n", sliceLen, sliceCap, sliceType, slice)
	fmt.Printf("slice=%v\n", slice)
	// 越界panic
	// fmt.Printf("slice[4]=%d\n", slice[4])

	// 初始化和赋值数组，打印长度、容量、类型和内容
	var array = [3]int{1, 2, 3}
	arrayLen := len(array)
	arrayCap := cap(array)
	arrayType := reflect.TypeOf(array)
	fmt.Printf("array len is: %d, Cap is %d, type is %v, array is %v\n", arrayLen, arrayCap, arrayType, array)
	array[1] = 0
	fmt.Printf("array=%v\n", array)

	// 拷贝切片
	copySlice := slice[1:3]
	fmt.Printf("copySlice len is:%d , cap is %d ; sliccopySlicee is %v\n", len(copySlice), cap(copySlice), copySlice)
	// 练习for range 用法，遍历切片
	for index, value := range copySlice {
		fmt.Println(index, value)
	}
	// 切片增加长度, 思考容量变化规律
	copySlice = append(copySlice, 0)
	copySlice = append(copySlice, 9, 9)
	fmt.Printf("copySlice=%v\n", copySlice)
	fmt.Printf("copySlice len is:%d , cap is %d ; sliccopySlicee is %v\n", len(copySlice), cap(copySlice), copySlice)
	copySlice = append(copySlice, slice...)
	fmt.Printf("copySlice len is:%d , cap is %d ; sliccopySlicee is %v\n", len(copySlice), cap(copySlice), copySlice)
}
