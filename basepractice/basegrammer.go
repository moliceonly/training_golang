package basepractice

import "fmt"

//1.Hello, world
func Question1() {
	fmt.Println("Hello, world!")
}

//2.声明一个整数，字符串，布尔变量并打印出来
func Question2() {
	var (
		a int = 23
		b string = "Hello"
		c bool = true
	)
	fmt.Println(a,b,c)
}

//3.声明一个常量Pi并计算半径为5的圆的面积
func Question3() {
	const Pi float64 = 3.1415926
	fmt.Println(5*5 *Pi)
}

//4.输出【姓名: 张三, 年龄: 25, 成绩: 92.5】
func Question4() {
	var (
		name string = "张三"
		age int = 25
		score float64 = 92.5
	)
	fmt.Printf("姓名：%s，年龄：%d，成绩：%f", name, age, score)
}

//5.从命令行读取两个整数，并输出和
func Question5() {
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Println(a+b)
} //go test不给输入，此处可go run一下

//6.将float64的3.14转换为整型
func Question6() {
	var pi float64 = 3.64
	fmt.Println(int(pi))
}

//7.交换a和b的值
func Question7() {
	var (
		a int = 7
		b int = 8
	)
	a, b = b, a
	fmt.Println(a, b)
}

//8.在代码内外各声明一个同名变量，观察输出
func Question8() {
	var a float64 = 3.14
	{
		var a float64 = 3.15
		fmt.Println(a)
	}
	fmt.Println(a)
}

//9.声明整型，字符串，布尔值，切片变量但不赋值，观察输出
func Question9() {
	var (
		a int
		s string
		b bool
	)
	fmt.Println(a, s,b)
}

//10.在if控制语句中声明变量，验证作用域
func Question10() {
	if a := 1; a<0{
		fmt.Println("a小于0: %d", a)
	} else {
		fmt.Println("a大于等于0: %d", a)
	}
	// fmt.Printf("a的值为: %d", a) 此处用于测试验证域，若取消注释则会导致编译失败，因为a的作用域在if语句中
}
