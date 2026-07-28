# training_golang

对齐 `go-web-roadmap.html` **Part 01**。已完成：**1.1 / 1.2 / 1.3**；进行中：**1.4**。

## 目录

```text
training_golang/           # module: training_golang
├── cmd/demo/              # 本 module 入口，调用 internal
├── internal/
│   ├── calc/              # 内部计价 Add / Price
│   └── errorinterfacepractice/  # 1.2 错误与接口练习
├── pkg/part01/            # 练习包
│   ├── …                  # 起点 / 1.1
│   └── stdlibpractice/    # 1.4 标准库 SDK（题 75–85，仅题干）
└── test_internal/         # 独立 module，验证无法 import internal
```

## 进度对照

| 路线 | 状态 | 落在哪 |
|------|------|--------|
| **1.1** 类型与数据结构 | 完成 | `pkg/part01/structpointerpractice` 等 |
| **1.2** 接口、错误与控制流 | 完成 | `internal/errorinterfacepractice` |
| **1.3** 包 / module / 工程习惯 | 完成 | `cmd/` · `internal/` · `pkg/` · `test_internal/` |
| **1.4** 标准库 SDK | 进行中 | `pkg/part01/stdlibpractice`（75–85） |

## 各包简要说明

| 路径 | 实现了什么 |
|------|------------|
| `pkg/part01/basegrammer` | 变量、常量、类型转换、作用域等基础语法 |
| `pkg/part01/controlpractice` | if / switch / for 等流程控制 |
| `pkg/part01/arrslicepractice` | 数组与切片（增删、遍历、扩容） |
| `pkg/part01/functionmethodpractice` | 函数、可变参数、闭包、defer |
| `pkg/part01/structpointerpractice` | 结构体/指针/接收者、深浅拷贝、JSON、标签 map、iota、匿名嵌入等 |
| `pkg/part01/stdlibpractice` | **1.4**：文本库、time、JSON、Base64、SHA256、AES-GCM、CSV、regexp/slices、flag、slog、context |
| `internal/errorinterfacepractice` | error、接口、Is/As、断言、recover 等 |
| `internal/calc` | `Add`、`Price` |
| `cmd/demo` | 同 module 调用 `internal` |
| `test_internal` | 外 module 验证 `internal` 边界 |

## 1.4 题目与场景对照

| 题 | 知识要点 | 贴近场景 |
|----|----------|----------|
| 75 | fmt / strings / strconv / bytes / unicode | API 入参清洗 |
| 76 | time | 活动截止：时区 + Duration |
| 77 | encoding/json | API 契约：往返、忽略密码、自定义 Role |
| 78 | encoding/base64 | 令牌进 Header / URL |
| 79 | crypto/sha256 | 盐 + 密码摘要落库 |
| 80 | crypto/aes + GCM | 配置密钥字段加解密 |
| 81 | encoding/csv（+xml 选做） | 用户列表导出 |
| 82 | regexp / sort / slices / maps | 手机号校验、标签规范化 |
| 83 | flag | 运维工具启动参数 |
| 84 | log/slog | 关键操作结构化日志 |
| 85 | context | 超时取消后续处理 |

## 常用命令

```bash
go run ./cmd/demo

cd test_internal && go run ./cmd/demo   # 应失败：internal 不可用

go test -v ./pkg/part01/stdlibpractice/
go test -v ./pkg/part01/...
go test -v ./internal/...
```

实现完 1.4 后进入路线图 **1.5** 测试与并发基础。
