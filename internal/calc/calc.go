// Package calc 是本 module 的内部计价工具包
// 路径在 internal/ 下，仅允许本 module（如 cmd/demo）引用；
// 其他 module 无法 import training_golang/internal/calc
package calc

// Add 返回 a + b。
func Add(a, b int) int {
	return a + b
}

// Price 按单价 unitPrice 与数量 qty 计算总价
// qty < 0 时返回 0（或按你自己的规则处理）
func Price(unitPrice, qty int) int {
	Price := 0
	if qty < 0 {
		return Price
	} else {
		Price = unitPrice * qty
	}
	return Price
}
