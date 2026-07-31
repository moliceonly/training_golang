# training_golang

对齐 [`go-web-roadmap.html`](../go-web-roadmap.html) 的 **Part 01–02** 练习库（题干 + `QuestionN` 演示 + `_test.go` 串跑）。

- **Part 01**：语言基础与标准库直觉 → `pkg/part01/`
- **Part 02**：文件 / HTTP / Protobuf / MySQL(GORM) / Redis / 短链综合 → `pkg/part02/`


## 目录

```text
training_golang/
├── pkg/part01/                 # 1.1–1.5
├── pkg/part02/
│   ├── fileiopractice/         # 2.1
│   ├── httppractice/           # 2.2
│   ├── protobufpractice/       # 2.3
│   ├── gormpractice/           # 2.4（114–121）
│   ├── redispractice/          # 2.5（122–128）
│   └── shortlinkpractice/      # 2.6（129–134）
└── test_internal/
```

## 进度对照

| 路线　　　　　　　　 | 状态　　　 | 落在哪　　　　　　　　　　　　　　　　　　|
| ----------------------| ------------| -------------------------------------------|
| **1.1–1.5**　　　　　| 练习库已齐 | `pkg/part01/...`　　　　　　　　　　　　　|
| **2.1** 文件 IO　　　| 练习库已齐 | `pkg/part02/fileiopractice`　　　　　　　 |
| **2.2** HTTP　　　　 | 练习库已齐 | `pkg/part02/httppractice`　　　　　　　　 |
| **2.3** Protobuf　　 | 练习库已齐 | `pkg/part02/protobufpractice`　　　　　　 |
| **2.4** MySQL + GORM | 练习库已齐 | `pkg/part02/gormpractice`（114–121）　　　|
| **2.5** Redis　　　　| 练习库已齐 | `pkg/part02/redispractice`（122–128）　　 |
| **2.6** 综合短链　　 | 练习库已齐 | `pkg/part02/shortlinkpractice`（129–134） |

## 2.4 题目与场景对照

| 题 | 知识要点 | 贴近场景 |
|----|----------|----------|
| 114 | DSN、`gorm.Open`、连接池 | 连上 MySQL |
| 115 | Model / `gorm.Model` / AutoMigrate | 图书·读者表 |
| 116 | CRUD + 分页 | 图书馆后台 |
| 117 | `Transaction` 借书 | 扣库存 + 流水，失败回滚 |
| 118 | 悲观锁 `FOR UPDATE` | 锁行再扣库存 |
| 119 | 乐观锁 `version` | 冲突重试 |
| 120 | Belongs To + Preload | 借阅详情（书+读者） |
| 121 | EXPLAIN、占位符、迁移策略 | 安全与排障 |

## MySQL（apt 本机安装）

| 项 | 值 |
|----|-----|
| 用户 | `trainer` |
| 密码 | `Train2026Lib!` |
| 库名 | `training_lib` |
| 主机 | `127.0.0.1:3306` |
| 覆盖 DSN | `TRAINING_MYSQL_DSN` |
| 跳过单测 | `TRAINING_SKIP_MYSQL=1` |

```bash
sudo apt update && sudo apt install -y mysql-server
sudo service mysql start

# 若尚未建库/用户（sudo mysql）：
# CREATE DATABASE IF NOT EXISTS training_lib DEFAULT CHARSET utf8mb4;
# CREATE USER 'trainer'@'localhost' IDENTIFIED BY 'Train2026Lib!';
# CREATE USER 'trainer'@'127.0.0.1' IDENTIFIED BY 'Train2026Lib!';
# GRANT ALL PRIVILEGES ON training_lib.* TO 'trainer'@'localhost';
# GRANT ALL PRIVILEGES ON training_lib.* TO 'trainer'@'127.0.0.1';
# FLUSH PRIVILEGES;

mysql -u trainer -p -h 127.0.0.1 -e "SELECT 1;"
```

```bash
export GOPROXY=https://goproxy.cn,direct
go get gorm.io/gorm gorm.io/driver/mysql
```

## 2.5 题目与场景对照

| 题 | 知识要点 | 贴近场景 |
|----|----------|----------|
| 122 | go-redis `NewClient`、Ping | 连上 Redis |
| 123 | String / Hash / List / Set / ZSet | 五类命令各练一遍 |
| 124 | TTL；穿透 / 击穿 / 雪崩笔记 | 过期与缓存问题 |
| 125 | Cache-Aside + 主动失效 | 商品详情先缓存后 DB |
| 126 | Token 存 Redis，退出 Del | 登录态立即失效 |
| 127 | `SET NX EX` 分布式锁 | 加锁 / 校验 value 解锁 |
| 128 | 固定窗口限流 | 按 IP，超限对应 429 |

## Redis（apt 本机安装）

| 项 | 值 |
|----|-----|
| 地址 | `127.0.0.1:6379` |
| 覆盖 | `TRAINING_REDIS_ADDR` |
| 跳过单测 | `TRAINING_SKIP_REDIS=1` |

```bash
sudo apt install -y redis-server
sudo service redis-server start
redis-cli ping   # PONG

export GOPROXY=https://goproxy.cn,direct
go get github.com/redis/go-redis/v9
```

## 2.6 题目与场景对照

| 题 | 知识要点 | 贴近场景 |
|----|----------|----------|
| 129 | MySQL + Redis 依赖、`AutoMigrate` | 短链表就绪 |
| 130 | 创建短链写 MySQL | 生成短码入库 |
| 131 | Cache-Aside + HTTP 302 | 访问优先 Redis，未命中回填 |
| 132 | CSV / `io.Reader` 批量导入 | 运营批量建链 |
| 133 | Redis 计数 + 汇总落库 | 跳转统计 |
| 134 | `context` 超时 + 冒烟 | 稳定性约束 |

依赖同 2.4 / 2.5。跳过：`TRAINING_SKIP_MYSQL=1` 或 `TRAINING_SKIP_REDIS=1`。

## 常用命令

```bash
cd training_golang
export GOPROXY=https://goproxy.cn,direct

# Part 02
go test -v -count=1 ./pkg/part02/shortlinkpractice/
go test -v -count=1 ./pkg/part02/redispractice/
go test -v -count=1 ./pkg/part02/gormpractice/
go generate ./pkg/part02/protobufpractice/...
go test -v -count=1 ./pkg/part02/protobufpractice/
go test -v -count=1 ./pkg/part02/httppractice/
go test -v -count=1 ./pkg/part02/fileiopractice/

# 看题干
go doc training_golang/pkg/part02/shortlinkpractice.Question131
go doc training_golang/pkg/part02/gormpractice.Question117
```

