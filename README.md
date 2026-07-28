# training_golang

对齐 `go-web-roadmap.html` **Part 01**。已完成：**1.1 / 1.2 / 1.3 / 1.4**；进行中：**1.5**。

## 目录

```text
training_golang/           # module: training_golang
├── cmd/demo/
├── internal/
│   ├── calc/
│   └── errorinterfacepractice/
├── pkg/part01/
│   ├── …                  # 起点 / 1.1 / 1.4
│   ├── stdlibpractice/    # 1.4（75–80、82–85）
│   └── concurrencypractice/  # 1.5 测试与并发（86–93，仅题干）
└── test_internal/
```

## 进度对照

| 路线 | 状态 | 落在哪 |
|------|------|--------|
| **1.1** 类型与数据结构 | 完成 | `pkg/part01/structpointerpractice` 等 |
| **1.2** 接口、错误与控制流 | 完成 | `internal/errorinterfacepractice` |
| **1.3** 包 / module / 工程习惯 | 完成 | `cmd/` · `internal/` · `pkg/` · `test_internal/` |
| **1.4** 标准库 SDK | 完成 | `pkg/part01/stdlibpractice` |
| **1.5** 测试与并发基础 | 进行中 | `pkg/part01/concurrencypractice`（86–93） |

## 各包简要说明

| 路径 | 实现了什么 |
|------|------------|
| `pkg/part01/basegrammer` | 基础语法 |
| `pkg/part01/controlpractice` | 流程控制 |
| `pkg/part01/arrslicepractice` | 数组与切片 |
| `pkg/part01/functionmethodpractice` | 函数、闭包、defer |
| `pkg/part01/structpointerpractice` | 结构体/指针等（1.1） |
| `pkg/part01/stdlibpractice` | 标准库 SDK（1.4） |
| `pkg/part01/concurrencypractice` | **1.5**：表驱动测试、基准/Example、goroutine、channel、sync、ConcurrentMap、worker 队列、并发坑 |
| `internal/errorinterfacepractice` | 错误与接口（1.2） |
| `internal/calc` | `Add` / `Price` |
| `cmd/demo` | 调用 `internal` |
| `test_internal` | 验证 `internal` 边界 |

## 1.5 题目与场景对照

| 题 | 知识要点 | 贴近场景 |
|----|----------|----------|
| 86 | 表驱动 / `t.Run` / `t.Helper` | 库存数量 Clamp 单测 |
| 87 | `testing.B` / Example | 同上函数的基准与示例 |
| 88 | goroutine 生命周期 | 并行汇总库存/数值 |
| 89 | channel + select 超时 | 第三方回调 3s，超时降级 |
| 90 | WaitGroup / Mutex / Once；`-race` | 库存并发踩坑再加锁 |
| 91 | 泛型 + RWMutex | `ConcurrentMap` Get/Set/Delete/GetOrSet |
| 92 | worker + cancel | 图片压缩队列，可取消全部 |
| 93 | 闭包循环变量 / 无缓冲阻塞 | 常见并发坑 |

## 常用命令

```bash
go test -v ./pkg/part01/concurrencypractice/
go test -race ./pkg/part01/concurrencypractice/   # 练 90 题 data race
go test -bench=ClampStock ./pkg/part01/concurrencypractice/
go test -run=Example ./pkg/part01/concurrencypractice/

go test -v ./pkg/part01/stdlibpractice/
go run ./cmd/demo
```

实现完 1.5 后进入路线图 **Part 02**（文件 / HTTP / DB 等）。
