package concurrencypractice

import (
	"testing"
	"fmt"
)
// ---------------------------------------------------------------------------
// 86. 表驱动 + t.Run + t.Helper
// TODO: 实现 assertEq 与 TestClampStock
func assertEq(t *testing.T, got, want int) {
	t.Helper() // TODO: 保留；比较失败时 t.Fatalf(...)
	_ = got
	_ = want
}

func TestClampStock(t *testing.T) {
	// TODO: 表驱动示例结构：
	tests := []struct {
		name string
		n, max, want int
	}{
		{"below zero", -1, 10, 0},
		{"normal", 3, 10, 3},
		{"above max", 20, 10, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertEq(t, ClampStock(tt.n, tt.max), tt.want)
		})
	}
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 87. Benchmark + Example
// TODO: 取消注释并补全

func BenchmarkClampStock(b *testing.B) {
	for i := 0; i < b.N; i++ {
			ClampStock(3, 10)
	}
}

func ExampleClampStock() {
	fmt.Println(ClampStock(15, 10))
	// Output:
	// 10
}

// ---------------------------------------------------------------------------
// 演示题入口（实现 Question* 后会有输出）
func Test_all_question_concurrencypractice(t *testing.T) {
	Question88()
	Question89()
	Question90()
	Question91()
	Question92()
	Question93()
}
