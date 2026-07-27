package controlpractice

import "fmt"

//判断输入的数字是奇数还是偶数
func Question11(num int) {
	if num % 2 == 0 {
		fmt.Printf("%d为偶数\n", num)
	} else {
		fmt.Printf("%d为奇数\n", num)
	}
}

//输入一个分数判断其等级90+ 优秀，80-89 良好，60-79 及格，<60 不及格
func Question12(score int) {
	fmt.Printf("分数：%d", score)
	switch {
		case score >=90:
			fmt.Println("优秀")
		case score >=80:
			fmt.Println("良好")
		case score >=60:
			fmt.Println("及格")
		case score >=0:
			fmt.Println("不及格")
		default:
			fmt.Println("输入的分数不合法")
	}
}

//输入年份判断是否为闰年，能被4整除且不能被100整除，或者能被400整除的年份
func Question13(year int) {
	if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
		fmt.Printf("%d是闰年\n", year)
	} else {
		fmt.Printf("%d不是闰年\n", year)
	}
}

//输入三个整数输出最大值
func Question14(a, b, c int) {
	max := a
	if b > max {
		max = b
	} else if c > max {
		max = c
	}
	fmt.Printf("最大值为：%d\n", max)
}

//循环遍历从1到100求和
func Question15() {
	sum := 0
	for i := 1; i <=100; i++ {
		sum = sum + i
	}
	fmt.Printf("1到100的和为：%d\n", sum)
}

//打印九九乘法表
func Question16() {
	fmt.Println("九九乘法表：")
	for i :=1; i <=9; i++ {
		for j :=1; j <=9; j++ {
			fmt.Printf("%dX%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

//判断输入的数字是否为素数
func Question17(num int) {
	//此处用最暴力方法
	for i :=2; i < num; i++ {
		if num % i == 0 {
			fmt.Printf("%d不是素数\n", num)
			return
		}
	}
	fmt.Printf("%d是素数\n", num)
}

//打印前20个斐波那契数列
func Question18() {
	a, b := 1, 1
	fmt.Println("前20个斐波那契数列：\t")
	for i :=0; i < 20; i++ {
		fmt.Printf("%d\t", a)
		a, b = b, a+b
	}
	fmt.Println()
}

//找出所有三位水仙花数
func Question19() {
	fmt.Println("三位水仙花数：\t")
	for i :=100;i < 1000; i++ {
		h := int( i / 100)
		d := int( i / 10 % 10)
		g := int(i % 10)
		if h*h*h + d*d*d + g*g*g == i {
			fmt.Printf("%d\t", i)
		}
	}
	fmt.Println()
}

//统计输入的字符串每个字符出现的次数
func Question20(str string) {
	count := make(map[rune]int)
	for _, char := range str {
		count[char]++
	}
	for char, count := range count {
		fmt.Printf("%c: %d\n", char, count)
	}
}