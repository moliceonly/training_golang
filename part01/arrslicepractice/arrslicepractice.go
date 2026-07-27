package arrslicepractice

import "fmt"

//定义一个[10]int数组，计算所有元素和
func Question21() {
	sum := 0
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, val := range arr {
		sum = sum + val
	}
	fmt.Printf("%v的和为%d\n", arr, sum)
}

//不使用额外数组反转一个数组
func Question22() {
	arr := [10]int{23, 15, 24, 431, 432, 56, 1234, 123, 90, 87}
	fmt.Printf("原数组为%v: \n",arr)
	arr_length := len(arr)
	for i:=0; i < arr_length - i - 1; i++ {
		arr[i], arr[arr_length- i - 1] = arr[arr_length - i -1], arr[i]
	}
	fmt.Printf("反转后数组为%v\n", arr)
}

//创建一个切片 [1,2,3,4,5]，追加元素 6、7，删除索引 2 的元素
func Question23() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("原切片为%v\n", slice)
	slice = append(slice, []int{6, 7}...)
	fmt.Printf("追加元素后切片为%v\n", slice)
	slice = append(slice[:2], slice[3:]...)
	fmt.Printf("删除索引为2的元素后为%v\n", slice)
}

//用 copy 复制一个切片，验证修改是否相互影响。
func Question24() {
	slice1 := []int{23, 51, 21, 87, 65}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)
	slice2[2]=30
	fmt.Printf("切片1为%v\n", slice1)
	fmt.Printf("切片2为%v，修改后不影响\n", slice2)
}

//观察 append 多次后容量的变化（打印 len 和 cap）
func Question25() {
	// slice := make([]int,0 ,0)
	// slice = append(slice, 5)
	// fmt.Printf("长度：%d, 容量：%d", len(slice), cap(slice))
	// slice = append(slice, 8, 10)
	// fmt.Printf("长度：%d, 容量：%d", len(slice), cap(slice))
	slice2 := make([]int, 0, 4)
	slice2 = append(slice2, 9, 1, 2)
	fmt.Printf("长度：%d, 容量：%d\n", len(slice2), cap(slice2))
	slice2 = append(slice2, 4, 5)
	fmt.Printf("长度：%d, 容量：%d\n", len(slice2), cap(slice2))
	slice2 = append(slice2, 12, 321, 23, 45)
	fmt.Printf("长度：%d, 容量：%d\n", len(slice2), cap(slice2))
}

//找出切片中的最大值和最小值
func Question26(slice []int) {
	max, min := slice[0], slice[0]
	for _, val := range slice {
		if val > max {
			max = val
		} else if val < min{ 
			min = val		
		}
	}
	fmt.Printf("切片%v中最大值为%d，最小值为%d\n", slice, max, min)
}

//删除切片中的重复元素（如 [1,2,2,3,3,3] → [1,2,3]）
func Question27() {
	slice := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	fmt.Printf("原切片为%v\n", slice)
	unival := make(map[int]bool)
	var newslice []int
	for _, val := range slice {
		if !unival[val] {
			unival[val] = true
			newslice = append(newslice, val)
		}
	}
	fmt.Printf("去重切片为%v\n", newslice)
}

//定义一个 3x3 的二维数组，计算主对角线元素之和
func Question28() {
	slice2d := make([][]int, 3)
	for i := 0; i < len(slice2d); i++{
		slice2d[i] = []int{2, 4, 6}
	}
	sum := 0
	for i := 0; i < len(slice2d); i++{
		sum = sum + slice2d[i][i]
	}
	fmt.Printf("二维切片%v的主对角线元素之和为%d\n",slice2d , sum)
}

//对一个整数切片进行冒泡排序（升序）
func Question29(slice []int) {
	fmt.Printf("原数组为%v", slice)
	for i := 0; i < len(slice) - 1; i++ {
		for j := i+1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	fmt.Printf("数组排序后为%v\n", slice)
}

//过滤出切片中所有偶数，返回新切片
func Question30(slice []int) {
	fmt.Printf("原数组为%v", slice)
	newslice := make([]int, 0, 0)
	for _, val := range slice {
		if val % 2 == 0 {
			newslice = append(newslice, val)
		}
	}
	fmt.Printf("切片中所有偶数有%v\n", newslice)
}