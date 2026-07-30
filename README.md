# training_golang

对齐 `go-web-roadmap.html`。

- **Part 01**：1.1–1.5 在 `pkg/part01/`
- **Part 02**：2.1–2.4 已有练习；进行中 **2.5** Redis（**统一用 apt 本机 Redis，不用 Docker**）

## 目录

```text
training_golang/
├── pkg/part01/
├── pkg/part02/
│   ├── fileiopractice/      # 2.1
│   ├── httppractice/        # 2.2
│   ├── protobufpractice/    # 2.3
│   ├── gormpractice/        # 2.4（114–121）
│   └── redispractice/       # 2.5（122–128）
└── test_internal/
```

## 进度对照

| 路线 | 状态 | 落在哪 |
|------|------|--------|
| 1.1–1.5 | 练习库已齐 | `pkg/part01/...` |
| **2.1–2.3** | 练习中 | `pkg/part02/...` |
| **2.4** MySQL + GORM | 练习库已齐 | `pkg/part02/gormpractice`（114–121） |
| **2.5** Redis | 进行中 | `pkg/part02/redispractice`（122–128） |

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

## MySQL（apt 本机安装，训练统一方案）

本仓库 **2.4 及后续需要 MySQL 的练习一律用 apt 安装的本机服务**，不使用 Docker 跑 MySQL。

练习账号（已写入 `gormpractice` 默认连接）：

| 项 | 值 |
|----|-----|
| 用户 | `trainer` |
| 密码 | `Train2026Lib!` |
| 库名 | `training_lib` |
| 主机 | `127.0.0.1:3306` |

代码里用 `mysql.Config` + `FormatDSN()` 生成连接串，避免手写 `%40` / `%21` 出错。

```bash
# 安装与启动（Ubuntu / WSL）
sudo apt update
sudo apt install -y mysql-server
sudo service mysql start

# 用 root（sudo mysql）建库与练习用户（若尚未创建）
# CREATE DATABASE IF NOT EXISTS training_lib DEFAULT CHARSET utf8mb4;
# CREATE USER 'trainer'@'localhost' IDENTIFIED BY 'Train2026Lib!';
# CREATE USER 'trainer'@'127.0.0.1' IDENTIFIED BY 'Train2026Lib!';
# GRANT ALL PRIVILEGES ON training_lib.* TO 'trainer'@'localhost';
# GRANT ALL PRIVILEGES ON training_lib.* TO 'trainer'@'127.0.0.1';
# FLUSH PRIVILEGES;

# 自检
mysql -u trainer -p -h 127.0.0.1 -e "SELECT 1; SHOW DATABASES;"
# 密码：Train2026Lib!
```

实现 GORM 代码时再装驱动：

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

## Redis（apt 本机安装，训练统一方案）

本仓库 **2.5 练习一律用 apt 安装的本机 Redis**，不使用 Docker 跑 Redis。

| 项 | 值 |
|----|-----|
| 地址 | `127.0.0.1:6379` |
| 覆盖 | 环境变量 `TRAINING_REDIS_ADDR` |
| 跳过单测 | `TRAINING_SKIP_REDIS=1` |

```bash
# 安装与启动（Ubuntu / WSL）
sudo apt update
sudo apt install -y redis-server
sudo service redis-server start   # 或: sudo systemctl start redis-server

# 自检
redis-cli ping   # 应返回 PONG
```

实现 Redis 代码时再装客户端：

```bash
export GOPROXY=https://goproxy.cn,direct
go get github.com/redis/go-redis/v9
```

## 常用命令

```bash
# 2.5
go test -v ./pkg/part02/redispractice/
go doc training_golang/pkg/part02/redispractice.Question125

# 2.4
go test -v ./pkg/part02/gormpractice/
go doc training_golang/pkg/part02/gormpractice.Question117

# 2.3
go generate ./pkg/part02/protobufpractice/
go test -v ./pkg/part02/protobufpractice/
```

实现完 2.5 后进入路线图 **2.6** 综合 IO 小项目。
