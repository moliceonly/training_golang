# training_golang

对齐 `go-web-roadmap.html`。

- **Part 01**：1.1–1.5 在 `pkg/part01/`
- **Part 02**：2.1–2.2 已有练习；进行中 **2.3** Protobuf

## 目录

```text
training_golang/
├── cmd/demo/
├── internal/
├── pkg/part01/
├── pkg/part02/
│   ├── fileiopractice/            # 2.1（94–101）
│   ├── httppractice/              # 2.2（102–108）
│   └── protobufpractice/          # 2.3（109–113）
│       ├── proto/user.proto
│       └── userpb/                # protoc 生成
└── test_internal/
```

## 进度对照

| 路线 | 状态 | 落在哪 |
|------|------|--------|
| 1.1–1.5 | 练习库已齐 | `pkg/part01/...` |
| **2.1** 文件 IO | 练习中 | `pkg/part02/fileiopractice` |
| **2.2** HTTP | 练习中 | `pkg/part02/httppractice` |
| **2.3** Protobuf | 进行中 | `pkg/part02/protobufpractice`（109–113） |

## 2.3 题目与场景对照

| 题 | 知识要点 | 贴近场景 |
|----|----------|----------|
| 109 | proto3：message / 字段号 / 标量 / repeated / map / 枚举 / 嵌套 | 用户信息 `user.proto` |
| 110 | `protoc` + `protoc-gen-go`；`go generate` | 工具链 |
| 111 | `proto.Marshal` / `Unmarshal` | 二进制往返 |
| 112 | HTTP + `application/protobuf` Body | Protobuf API |
| 113 | JSON vs Protobuf 体积与契约演进 | 契约对照 |

## 常用命令

```bash
# 2.3（改完 .proto 后重新生成）
go generate ./pkg/part02/protobufpractice/
go test -v ./pkg/part02/protobufpractice/
go doc training_golang/pkg/part02/protobufpractice.Question111

# 2.2
go test -v ./pkg/part02/httppractice/

go run ./cmd/demo
```

实现完 2.3 后进入路线图 **2.4** 数据库（MySQL + GORM）。
