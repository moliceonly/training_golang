# training_golang

对齐 `go-web-roadmap.html`。

- **Part 01**：1.1–1.5 在 `pkg/part01/`
- **Part 02**：2.1–2.3 已有练习；进行中 **2.4** MySQL + GORM（**统一用 apt 本机 MySQL，不用 Docker**）

## 目录

```text
training_golang/
├── pkg/part01/
├── pkg/part02/
│   ├── fileiopractice/      # 2.1
│   ├── httppractice/        # 2.2
│   ├── protobufpractice/    # 2.3
│   └── gormpractice/        # 2.4（114–121）
└── test_internal/
```

## 进度对照

| 路线 | 状态 | 落在哪 |
|------|------|--------|
| 1.1–1.5 | 练习库已齐 | `pkg/part01/...` |
| **2.1–2.3** | 练习中 | `pkg/part02/...` |
| **2.4** MySQL + GORM | 进行中 | `pkg/part02/gormpractice`（114–121） |

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

## 常用命令

```bash
# 2.4
go test -v ./pkg/part02/gormpractice/
go doc training_golang/pkg/part02/gormpractice.Question117

# 2.3
go generate ./pkg/part02/protobufpractice/
go test -v ./pkg/part02/protobufpractice/
```

实现完 2.4 后进入路线图 **2.5** Redis。
