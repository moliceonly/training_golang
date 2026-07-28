// Command demo：组装并调用本 module 的内部包。
// 运行：在仓库根目录执行  go run ./cmd/demo
package main

import (
	"fmt"

	"training_golang/internal/calc"
	"training_golang/internal/errorinterfacepractice"
)

func main() {

	//同module测试调用internal内pkg
	a, b, c, d := 2, 3, 10, 2
	sum := calc.Add(a, b)
	total := calc.Price(c, d)
	fmt.Printf("Add(%v,%v)=%d  Price(%v,%v)=%d\n", a, b, sum, c, d, total)

	//同module测试不同包可见性
	// errorinterfacepractice.describeValue(1)
	// errorinterfacepractice.Question68() //内部调用了describeValue因此无法调用
	errorinterfacepractice.Question69() //正常运行
}
