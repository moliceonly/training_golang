// 题号：86 → …
//
// 说明：
//   - 86/87 的「实现」主要在 concurrencypractice_test.go（表驱动 / 基准 / Example）
//   - 90 可用：go test -race ./pkg/part01/concurrencypractice/
package concurrencypractice

import (
	"context"
	"sync"
	"time"
	"fmt"
)

// ---------------------------------------------------------------------------
// 86. 单元测试：表驱动、t.Run、t.Helper、子测试
// 场景引申：库存扣减边界 — 数量夹在 [0, max] 之间
//
// 写函数 ClampStock(n, max int) int：
//   - n < 0 返回 0
//   - n > max 返回 max
//   - 否则返回 n
// 然后在 concurrencypractice_test.go 中完成 TestClampStock：
//   - 使用表驱动（切片里每项含 name / n / max / want）
//   - 用 t.Run(tt.name, ...) 跑子测试
//   - 写辅助函数 assertEq(t, got, want)，内部第一行 t.Helper()
func ClampStock(n, max int) int {
	if n < 0 {
		return 0
	} else if n > max {
		return max
	} else {
		return 0
	}
}

// ---------------------------------------------------------------------------
// 87. 基准测试 testing.B；示例测试 Example
// （实现写在 _test.go）
//
// 在 concurrencypractice_test.go 中：
//   - 写 BenchmarkClampStock(b *testing.B)：循环 b.N 次调用 ClampStock(3, 10)
//   - 写 ExampleClampStock()：调用并打印结果，用 // Output: 注释声明期望输出
// 运行：
//   go test -bench=ClampStock ./pkg/part01/concurrencypractice/
//   go test -run=Example ./pkg/part01/concurrencypractice/

// ---------------------------------------------------------------------------
// 88. goroutine 启动与生命周期直觉
// 场景引申：同时向多个仓查询库存，全部结束后汇总
//
// 写函数 ParallelSum(nums []int) int：
//   - 对每个 num 启动一个 goroutine，把 num 送到结果 channel
//   - 主 goroutine 收齐 len(nums) 个结果后求和返回
//   - 注意：结束后关闭 channel 或用固定次数接收，避免泄漏
// 在 Question88 中对 []int{1,2,3,4} 打印结果（应为 10）
func ParallelSum(nums []int) int {
	var wait sync.WaitGroup
	// var mainwait sync.WaitGroup 
	wait.Add(len(nums))
	ch := make(chan int, len(nums))
	for _, val := range nums {
		go func(v int) {
			defer wait.Done()
			ch <- v 
		}(val)
	}
	wait.Wait()
	close(ch)
	sum := 0
	for val := range ch {
		sum = sum + val
	}
	return sum
}

func Question88() {
	fmt.Println(ParallelSum([]int{1, 2, 3, 4}))
}

// ---------------------------------------------------------------------------
// 89. channel：收发、关闭、select、超时模式
// 场景：第三方回调等待 — 主流程 select 等待结果或 3 秒超时，超时则走降级
//
// 写函数 WaitCallback(result <-chan string, timeout time.Duration) (string, bool)：
//   - select：
//       case v := <-result: 返回 (v, true)
//       case <-time.After(timeout): 返回 ("fallback", false)  // 降级
// 写函数 SimulateCallback(delay time.Duration, value string) <-chan string：
//   - 返回一个 channel；在新 goroutine 里 sleep(delay) 后发送 value 并关闭 channel
// 在 Question89 中：
//   1) delay=100ms, timeout=3s → 应拿到回调值
//   2) delay=2s, timeout=200ms → 应拿到 fallback
func WaitCallback(result <-chan string, timeout time.Duration) (string, bool) {
	select{
	case v :=  <- result:
		return v, true
	case <- time.After(timeout):
		return "fallback", false
	}
}

func SimulateCallback(delay time.Duration, value string) <-chan string {
	ch := make(chan string)
	go func(v string) {
		time.Sleep(delay)
		ch <- v
		close(ch)
	}(value)
	return ch
}

func Question89() {
	fmt.Println(WaitCallback(SimulateCallback(100 * time.Millisecond, "你好"), 3 * time.Second))
	fmt.Println(WaitCallback(SimulateCallback(2 * time.Second, "你好"), 200 * time.Millisecond))
}

// ---------------------------------------------------------------------------
// 90. sync.WaitGroup / Mutex / RWMutex / Once；go test -race
// 场景：库存计数踩坑 — 多 goroutine 同时改共享 map
//
// 1) 写 UnsafeIncrStock(stock map[string]int, sku string, n, goroutines int)
//    - 启动 goroutines 个协程，每个对 stock[sku] += n（无锁）
//    - 用 WaitGroup 等待全部结束
//    - 预期：go test -race 会报 data race
//
// 2) 写 SafeIncrStock(...)：用 Mutex 保护 map 写入，逻辑同上
//    - -race 应干净；最终 stock[sku] == n*goroutines
//
// 3) 写 LoadConfigOnce(loader func() string) string：
//    - 用 sync.Once 保证 loader 只执行一次；多次调用返回同一结果
//
// 在 Question90 中演示 SafeIncrStock 与 LoadConfigOnce；Unsafe 可注释掉以免 race 干扰日常测试
func UnsafeIncrStock(stock map[string]int, sku string, n, goroutines int) {
	var wait sync.WaitGroup
	wait.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			stock[sku] += n
			wait.Done()
		}()
	}
	wait.Wait()
}

func SafeIncrStock(stock map[string]int, sku string, n, goroutines int) {
	var wait sync.WaitGroup
	var lock sync.Mutex
	wait.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			lock.Lock()
			stock[sku] += n
			lock.Unlock()
			wait.Done()
		}()
	}
	wait.Wait()
}

var (
	configOnce sync.Once
	configCached string
)

func LoadConfigOnce(loader func() string) string {
	configOnce.Do(func() {
		configCached = loader()
	})
	return configCached
}

// ConfigLoader 用 Once 缓存配置加载结果
type ConfigLoader struct {
	once   sync.Once
	value  string
	loader func() string
}

func NewConfigLoader(loader func() string) *ConfigLoader {
	return &ConfigLoader{loader: loader}
}

func (c *ConfigLoader) Get() string {
	c.once.Do(func() {
		c.value = c.loader()
	})
	return c.value
}

func Question90() {
	mp := map[string]int{"a":3, "b":4, "c":1}
	UnsafeIncrStock(mp, "a", 1, 4)
	fmt.Println(mp) //go test -v -race
	mp = map[string]int{"a":3, "b":4, "c":1}
	SafeIncrStock(mp, "a", 1, 4)
	fmt.Println(mp)
	loader := NewConfigLoader(func() string {
		fmt.Println("你好")
		return ""
	})
	fmt.Println(loader.Get())
}

// ---------------------------------------------------------------------------
// 91. 泛型 + 锁：手写并发安全容器（读多写少用 RWMutex）
// 场景：并发安全 Map — 不使用 sync.Map
//
// 实现 ConcurrentMap[K comparable, V any]：
//   - 内部：sync.RWMutex + map[K]V
//   - 方法：Get / Set / Delete / GetOrSet
//     GetOrSet(key, val)：若 key 已存在返回 (old, false)；否则写入返回 (val, true)
// 在 Question91 中：多 goroutine Set/Get，打印 GetOrSet 两次的结果差异
type ConcurrentMap[K comparable, V any] struct {
	lock sync.RWMutex
	mp map[K]V
}

func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		mp: make(map[K]V),
	}
}

func (m *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	m.lock.RLock()
	v, ok := m.mp[key]
	m.lock.RUnlock()
	return v, ok
}

func (m *ConcurrentMap[K, V]) Set(key K, val V) {
	m.lock.Lock()
	m.mp[key] = val
	m.lock.Unlock()
}

func (m *ConcurrentMap[K, V]) Delete(key K) {
	m.lock.Lock()
	delete(m.mp, key)
	m.lock.Unlock()
}

func (m *ConcurrentMap[K, V]) GetOrSet(key K, val V) (V, bool) {
	m.lock.Lock()
	if v, ok := m.mp[key]; ok {
		return v, ok
	}
	m.mp[key] = val
	m.lock.Unlock()
	return val, true
}

func Question91() {
	m := NewConcurrentMap[string, int]()
	m.Set("a", 1)
	fmt.Println(m.Get("a"))
	m.Set("b", 2)
	fmt.Println(m)
	m.Delete("b")
	fmt.Println(m)
	fmt.Println(m.GetOrSet("c", 3))
	fmt.Println(m)
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 92. channel + WaitGroup 综合
// 场景：图片压缩队列 — 固定 N 个 worker 从任务 channel 取活，支持「取消全部」
//
// 写函数 RunCompressWorkers(n int, tasks <-chan string, ctx context.Context) (done []string)：
//   - 启动 n 个 worker；每个循环：select
//       case <-ctx.Done(): return
//       case job, ok := <-tasks: 若 !ok 则 return；否则「压缩」后把 job 送入结果收集
//   - 用 WaitGroup 等所有 worker 结束后，返回已完成的 job 列表（顺序不限）
//   - 「压缩」可用 time.Sleep(10*time.Millisecond) 模拟
// 在 Question92 中：
//   1) 正常跑完若干任务
//   2) 用 cancel 中途取消，观察完成数变少
func RunCompressWorkers(n int, tasks <-chan string, ctx context.Context) (done []string) {
	var wait sync.WaitGroup
	var lock sync.Mutex
	wait.Add(n)
	for i:=0; i < n; i++ {
		go func() {
			defer wait.Done()
			for {
				select {
				case <-ctx.Done():
					return 
				case job, ok := <- tasks:
					if !ok {
						return 
					}
					time.Sleep( 10 * time.Millisecond )
					lock.Lock()
					done = append(done, job)
					lock.Unlock()
				}
			}
		}()
	}
	wait.Wait()
	return done
}

func Question92() {
	task1 := make(chan string, 5)
	for _, name := range []string{"pict1", "pict2", "pict3", "pict4"} {
		task1 <- name
	}
	close(task1)
	done1 := RunCompressWorkers(2, task1, context.Background())
	fmt.Println(done1)

	tasks2 := make(chan string, 20)
	for i := 0; i < 20; i++ {
		tasks2 <- "a"
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()
	done2 := RunCompressWorkers(2, tasks2, ctx)
	fmt.Println(done2)

	fmt.Println()
}

// ---------------------------------------------------------------------------
// 93. 常见并发坑：闭包循环变量、泄漏、无缓冲死锁
//
// 写函数 FixedLoopPrint(n int) []int：
//   - 错误写法直觉：for i := 0; i < n; i++ { go func(){ ch <- i }() } 可能全是同一个 i
//   - 正确：go func(v int){ ch <- v }(i) 或 Go 1.22+ 每轮新变量；本题按「显式传参」写
//   - 收齐 n 个数返回（可用 WaitGroup + 带缓冲 channel）
//
// 写函数 TryUnbufferedDeadlock() (panicked bool)：
//   - 演示：无缓冲 channel 上，同一 goroutine 先发后收会死锁
//   - 用 defer + recover 捕获 runtime 的 deadlock 不可靠（死锁是 Fatal）；
//     因此改为：在子 goroutine 里对无缓冲 ch 发送且无人接收，主流程用 select+timeout
//     判断「发送未在 50ms 内完成」则视为演示成功，返回 true；并注释说明真死锁会卡死进程
// 在 Question93 中打印 FixedLoopPrint(5)；调用 TryUnbufferedDeadlock 并说明
func FixedLoopPrint(n int) []int {
	var slice []int
	var wait sync.WaitGroup
	ch := make(chan int, n)
	wait.Add(n)
	for i := 0; i < n; i++ {
		go func(v int) {
			ch <- v
			wait.Done()
		}(i)
	}
	wait.Wait()
	close(ch)
	for i := range ch {
		slice = append(slice, i)
	}
	return slice
}

func TryUnbufferedDeadlock() (blocked bool) {
	ch := make(chan int)
	for i:=0; i <2; i++ {
		go func(v int) {
			ch <- v
		}(i)
	}
	select {
	case <-time.After(50 * time.Millisecond):
		return false
	}
	return true
}

func Question93() {
	fmt.Println(FixedLoopPrint(6))
	fmt.Println(TryUnbufferedDeadlock())
}
