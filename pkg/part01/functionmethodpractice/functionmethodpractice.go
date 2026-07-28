package functionmethodpractice

import "fmt"

// 写一个函数add(a，b int)int计算两数之和
func add(a int, b int) int {
	return a + b
}
func Question31(a int, b int) {
	fmt.Printf("a%d和b%d两数之和为%d\n", a, b, add(a, b))
}

// 写一个函数，返回两个整数的和与差
func add_minus(a int, b int) (int, int) {
	return a + b, a - b
}
func Question32(a int, b int) {
	addnum, minusnum := add_minus(a, b)
	fmt.Printf("a%d和b%d两数之和为%d，两数之差为%d\n", a, b, addnum, minusnum)
}

// 用命名返回值重写上一题
func add_minus_named(a int, b int) (an int, mn int) {
	an, mn = a+b, a-b
	return
}
func Question33(a int, b int) {
	addnum, minusnum := add_minus_named(a, b)
	fmt.Printf("a%d和b%d两数之和为%d，两数之差为%d\n", a, b, addnum, minusnum)
}

// 写一个函数利用可变参数特性计算所有参数的和
func sumAll(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum = sum + val
	}
	return sum
}
func Question34(nums ...int) {
	fmt.Printf("传入%d个参数和为%d\n", len(nums), sumAll(nums...))
}

// 写一个函数，返回一个闭包，每次调用计数器加1
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func Question35() {
	a := counter()
	b := counter()

	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
}

// 用递归计算n!(阶乘)
func n(a int) int {
	if a <= 1 {
		return 1
	} else {
		return a * n(a-1)
	}
}
func Question36(num int) {
	fmt.Printf("%d的阶乘为%d\n", num, n(num))
}

// 用递归实现斐波那契数列
func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func Question37(numsize int) {
	fmt.Printf("前%d个斐波那契数列元素为:\t", numsize)
	for i := 0; i < numsize; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()
}

// 写一个函数，用defer在函数退出时打印"Done!"
func Question38() {
	defer fmt.Println("Done!")
	fmt.Println(fibonacci(3))
}

// 定义一个Rectangle结构体，实现计算面积和周长的方法
type Rectangle struct {
	w float64
	h float64
}

func (r Rectangle) area() float64 {
	return r.w * r.h
}
func (r Rectangle) perimeter() float64 {
	return (r.w + r.h) * 2
}
func Question39(weight float64, height float64) {
	r := Rectangle{
		w: weight,
		h: height,
	}
	fmt.Printf("长为%.3f, 宽为%.3f的长方形周长为%.3f, 面积为%.3f\n", r.h, r.w, r.perimeter(), r.area())
}

// 定义一个Shape接口(包含 Area()方法)，让 Rectangle 和 circle实现它
type Shape interface {
	Area() float64
}
type Circle struct {
	r float64
}

func (R Rectangle) Area() float64 {
	return R.w * R.h
}
func (C Circle) Area() float64 {
	return 3.1415 * C.r * C.r
}
func Question40(weight float64, height float64, radius float64) {
	rect := Rectangle{
		h: height,
		w: weight,
	}
	cir := Circle{
		r: radius,
	}
	var sr Shape = rect
	var sc Shape = cir
	fmt.Printf("半径为%.3f的圆面积为%.3f\n", cir.r, sc.Area())
	fmt.Printf("长为%.3f, 宽为%.3f的长方形面积为%.3f\n", rect.h, rect.w, sr.Area())
}
