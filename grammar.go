package main

import "fmt"
/* 全局变量 */
var a int = 10

func main() {

	var a int = 10
	var b int = 11
	fmt.Printf("a = %d , b = %d a/b=%f \n", a, b, float32(a)/float32(11))

	fmt.Println(" loop ---")
	for c:=0; c<10;c++  {
		fmt.Println(c)
	}
	/** 数组 **/
	fmt.Println(" arr ---")
	var arr = [10] float32 {9, 8, 7, 6, 5, 4, 3, 2, 1, 1.111}
	for d,e := range arr {
		fmt.Println(d, e)
	}

	/** map **/
	fmt.Println(" map ---")
	kv := map[int]string{1:"a", 2:"b", 3:"c", 4:"d"}
	for f,g := range kv{
		fmt.Println(f, g)
	}
	fmt.Println(kv[1])
	val,r := kv[1]
	fmt.Println(r)
	fmt.Println(val)

	delete(kv, 2)
	fmt.Println(kv)

	fmt.Println(" 变量指针 ---")
	/** 指针 **/
	fmt.Println(&a, a)

	var h *int = &a
	fmt.Println(h, *h)

	/** 切片 slice 变长 数组 **/
	slice := make([]int, 3)
	slice[0] = 2
	fmt.Println(slice)

	slice1 := []int{1,3, 4}
	fmt.Println(slice1)

	slice2 := arr[:]
	fmt.Println(slice2)

	slice3 := arr[3:7]
	fmt.Println(slice3)
	fmt.Printf("slice3 len=%d cap=%d \n", len(slice3), cap(slice3))

	var slice4 []int
	slice4 = append(slice4, 4, 44, 444)
	fmt.Println(slice4)

	slice5 := make([]int, len(slice4)*2, cap(slice4)*3)
	copy(slice5, slice4)
	fmt.Println(slice5)



}
