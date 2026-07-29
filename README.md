# training_golang

对齐 `go-web-roadmap.html`。

- **Part 01**：1.1–1.5 练习在 `pkg/part01/`
- **Part 02**：2.1 文件 IO；进行中 **2.2** HTTP

## 目录

```text
training_golang/
├── cmd/demo/
├── internal/
├── pkg/part01/                    # Part 01
├── pkg/part02/
│   ├── fileiopractice/            # 2.1（94–101）
│   └── httppractice/              # 2.2（102–108，仅题干）
└── test_internal/
```

## 进度对照

| 路线 | 状态 | 落在哪 |
|------|------|--------|
| 1.1–1.5 | 练习库已齐 | `pkg/part01/...` |
| **2.1** 文件与本地 IO | 练习中 | `pkg/part02/fileiopractice` |
| **2.2** HTTP 客户端与服务端 | 进行中 | `pkg/part02/httppractice`（102–108） |

## 2.2 题目与场景对照

| 题 | 知识要点 | 贴近场景 |
|----|----------|----------|
| 102 | 方法 / 状态码 / Header / Body / Cookie | 协议要点上手 |
| 103 | ServeMux、Handler、JSON 读写 | 备忘录 Todo CRUD |
| 104 | 中间件：日志、耗时、recover | 请求进出访问日志 |
| 105 | `http.Client` + Timeout + 文件缓存 | 天气聚合（2s 超时） |
| 106 | `context` 取消；GET 重试 | 重试与幂等意识 |
| 107 | CORS / Content-Type / 4xx·5xx | 了解即可 |
| 108 | 手写 `/users/:id` | 非法 id → 400 |

## 常用命令

```bash
# 2.1
go test -v ./pkg/part02/fileiopractice/
go doc fileiopractice Question98

# 1.5
go test -v ./pkg/part01/concurrencypractice/
go test -race ./pkg/part01/concurrencypractice/

go run ./cmd/demo
```

实现完 2.2 后进入路线图 **2.3** Protobuf。
