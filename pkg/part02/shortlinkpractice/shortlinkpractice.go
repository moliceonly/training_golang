// Package shortlinkpractice 对齐路线图 Part 02 · 2.6「综合 IO 小项目（衔接 Web）」。
// 把文件、HTTP、MySQL、Redis、context 串成短链跳转链路。
// 题号：129 → 134
//
// 查看题干：
//
//	go doc training_golang/pkg/part02/shortlinkpractice.Question131
//
// 依赖本机 apt 安装的 MySQL + Redis（训练统一不用 Docker）。
// MySQL：trainer / Train2026Lib!（库 training_lib），可用 TRAINING_MYSQL_DSN 覆盖。
// Redis：127.0.0.1:6379，可用 TRAINING_REDIS_ADDR 覆盖。
//
// 连不上时 Question 演示可打印提示并 return；
// 单测可设 TRAINING_SKIP_MYSQL=1 或 TRAINING_SKIP_REDIS=1。
package shortlinkpractice

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"
	"bufio"
	"errors"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbUser           = "trainer"
	dbPass           = "Train2026Lib!"
	dbName           = "training_lib"
	dbAddr           = "127.0.0.1:3306"
	defaultRedisAddr = "127.0.0.1:6379"
)

// ShortLink 短链表：短码 → 长链，Hits 为累计访问（可由 Redis 汇总落库）。
type ShortLink struct {
	gorm.Model
	Code string `gorm:"uniqueIndex;size:16"`
	URL  string `gorm:"size:2048"`
	Hits int64
}

// Deps 短链服务依赖：MySQL + Redis。
type Deps struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func DefaultDSN() string {
	cfg := mysqldriver.Config{
		User:                 dbUser,
		Passwd:               dbPass,
		Net:                  "tcp",
		Addr:                 dbAddr,
		DBName:               dbName,
		AllowNativePasswords: true,
		Params: map[string]string{
			"charset":   "utf8mb4",
			"parseTime": "true",
			"loc":       "Local",
		},
	}
	return cfg.FormatDSN()
}

func dsnFromEnv() string {
	if v := os.Getenv("TRAINING_MYSQL_DSN"); v != "" {
		return v
	}
	return DefaultDSN()
}

func redisAddrFromEnv() string {
	if v := os.Getenv("TRAINING_REDIS_ADDR"); v != "" {
		return v
	}
	return defaultRedisAddr
}

func cacheKey(code string) string { return "short:" + code }
func hitsKey(code string) string  { return "hits:" + code }

// ---------------------------------------------------------------------------
// 129. 依赖就绪：连 MySQL + Redis；ShortLink AutoMigrate
//
// Question129 打开双依赖并迁移短链表。
//
// 写函数：
//   OpenDB(dsn string) (*gorm.DB, error)
//     - gorm.Open + Ping（池参数可简配）
//   OpenRedis(addr string) (*redis.Client, error)
//     - NewClient + Ping
//   OpenDeps() (*Deps, error)
//     - 组合上面两个；任一失败返回 error
//   AutoMigrateShortLink(db *gorm.DB) error
// 在 Question129 中：OpenDeps + AutoMigrate，成功打印 "deps ok"
func OpenDB(dsn string) (*gorm.DB, error) {
	// TODO: gorm.Open(mysql.Open(dsn), &gorm.Config{}) + Ping
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sql, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sql.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func OpenRedis(addr string) (*redis.Client, error) {
	// TODO: redis.NewClient + Ping
	rdb := redis.NewClient(&redis.Options{Addr: addr})
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}

func OpenDeps() (*Deps, error) {
	// TODO
	var dep Deps
	var err error

	dep.DB, err = OpenDB(dsnFromEnv())
	if err != nil {
		return nil, err
	}

	dep.RDB, err = OpenRedis(redisAddrFromEnv())
	if err != nil {
		return nil, err
	}

	return &dep, nil
}

func AutoMigrateShortLink(db *gorm.DB) error {
	// TODO
	return db.AutoMigrate(&ShortLink{})
}

// Question129 演示依赖与迁移。
func Question129() {
	// TODO
	dep, err := OpenDeps()
	if err != nil {
		fmt.Println("Open DB, redis: ", err)
		return
	}
	fmt.Println(dep)

	if err := AutoMigrateShortLink(dep.DB); err != nil {
		fmt.Println("AutoMigrate: ", err)
		return
	}
	if dep.DB.Migrator().HasTable(&ShortLink{}) {
		fmt.Println("table ok: short_links")
	} else {
		fmt.Println("table missing")
	}
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 130. 创建短链：长链写入 MySQL，生成短码
//
// Question130 练习创建一条短链。
//
// 写函数：
//   genCode() string
//     - 可用时间戳/随机串生成短码（长度建议 ≤16，注意 unique）
//   CreateShortLink(ctx context.Context, db *gorm.DB, longURL string) (code string, err error)
//     - 生成 code，插入 ShortLink{Code, URL}；失败返回 error
//     - 应用 WithTimeout（如 3s）亦可在 Question 外层包一层
// 在 Question130 中：创建一条，打印 code 与 longURL；演示结束可删掉该行
func genCode() string {
	// TODO
	// return strconv.FormatInt(time.Now().UnixNano(), 36)
	return fmt.Sprintf("%x", time.Now().UnixNano())
}

func CreateShortLink(ctx context.Context, db *gorm.DB, longURL string) (string, error) {
	// TODO
	shortLink := &ShortLink{
		Code: genCode(),
		URL: longURL,
	}

	if err := db.WithContext(ctx).Create(shortLink).Error; err != nil {
		return "", err
	}

	return shortLink.Code, nil
}

// Question130 演示创建短链。
func Question130() {
	deps, err := OpenDeps()
	if err != nil {
		fmt.Println("db/rdb init failed: ", err )
		return 
	}
	
	if err := AutoMigrateShortLink(deps.DB); err != nil {
		fmt.Println("create table failed: ", err )
	} 

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	
	code, err := CreateShortLink(ctx, deps.DB, "https://example.com")
	if err != nil {
		fmt.Println("insert shortlink failed: ", err)
		return
	}
	fmt.Println("code:", code, "url: https://example.com")
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 131. 访问解析：先 Redis 后 DB；回填；HTTP 302 跳转
// 场景：Cache-Aside；未命中查库再 SET 缓存（建议带 TTL）
//
// 写函数：
//   ResolveURL(ctx context.Context, rdb *redis.Client, db *gorm.DB, code string) (longURL string, err error)
//     - GET cacheKey(code)；命中直接返回
//     - miss：查 ShortLink by code；没有则 error
//     - 回填 Redis（如 TTL 5m）
//   NewShortLinkMux(deps *Deps) http.Handler
//     - POST /shorten  body: {"url":"..."} → JSON {"code":"..."}
//     - GET  /r/{code} → ResolveURL 后 302 Location；可选顺带 Incr hitsKey
// 在 Question131 中：httptest 创建 + 跳转，打印状态码与 Location
func ResolveURL(ctx context.Context, rdb *redis.Client, db *gorm.DB, code string) (string, error) {
	// TODO

	// 拿key
	key := cacheKey(code)

	// 查询redis
	result, err := rdb.Get(ctx, key).Result()

	if err == nil { // 查询成功

		return result, nil
	}
	if err != redis.Nil {
		return "", err
	}

	//未命中进入DB查询
	var shortLink ShortLink
	if err := db.WithContext(ctx).Where("code = ?", code).First(&shortLink).Error; err != nil {
		return "", err
	}

	if err := rdb.Set(ctx, key, shortLink.URL, 5*time.Minute).Err(); err != nil {
		fmt.Println("回填失败:", err)
	}

	return shortLink.URL, nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func NewShortLinkMux(deps *Deps) http.Handler {
	// TODO
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil || strings.TrimSpace(body.URL) == "" {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "bad url"})
			return
		}

		ctx := r.Context()
		code, err := CreateShortLink(ctx, deps.DB, body.URL)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		WriteJSON(w, http.StatusOK, map[string]string{"code": code})
	})

	mux.HandleFunc("GET /r/{code}", func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("code")
		if code == "" {
			http.Error(w, "missing code", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		longURL, err := ResolveURL(ctx, deps.RDB, deps.DB, code)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		// 可选：访问计数（Q133 的 IncrHit；这里直接 INCR，避免依赖未完成的 stub）
		_ = deps.RDB.Incr(ctx, hitsKey(code)).Err()

		http.Redirect(w, r, longURL, http.StatusFound)
	})

	return mux
}

// Question131 演示解析与 302。
func Question131() {

	deps, err := OpenDeps()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := AutoMigrateShortLink(deps.DB); err != nil {
		fmt.Println(err)
		return
	}

	h := NewShortLinkMux(deps)
	longURL := "https://example.com/shortlink-demo"

	// Post
	body := strings.NewReader(`{"url":"` + longURL + `"}`)
	req := httptest.NewRequest(http.MethodPost, "/shorten", body)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	fmt.Println("POST /shorten:", rr.Code, rr.Body.String())

	var resp struct {
		Code string `json:"code"`
	}
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil || resp.Code == "" {
		fmt.Println("bad shorten resp")
		return
	}
	defer deps.DB.Unscoped().Where("code = ?", resp.Code).Delete(&ShortLink{})
	defer deps.RDB.Del(context.Background(), cacheKey(resp.Code), hitsKey(resp.Code))

	req2 := httptest.NewRequest(http.MethodGet, "/r/"+resp.Code, nil)
	rr2 := httptest.NewRecorder()
	h.ServeHTTP(rr2, req2)
	fmt.Println("GET /r/{code}:", rr2.Code, "Location:", rr2.Header().Get("Location"))
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 132. 运营批量建链：读 CSV（每行一个长链），写入 DB
//
// Question132 练习从 io.Reader 导入。
//
// 写函数：
//   ImportCSV(ctx context.Context, db *gorm.DB, r io.Reader) (n int, err error)
//     - 按行读取（可 encoding/csv 或 bufio.Scanner）
//     - 跳过空行；每行调用 CreateShortLink（或批量 Create）
//     - 返回成功条数
// 在 Question132 中：用 strings.NewReader 喂几行 URL，打印 n
func ImportCSV(ctx context.Context, db *gorm.DB, r io.Reader) (int, error) {
	// TODO

	scanner := bufio.NewScanner(r)
	count := 0
	for scanner.Scan() {
		
		url := scanner.Text()
		if strings.TrimSpace(url) == "" {
			continue
		}
		
		shortLink := &ShortLink{
			URL: url,
			Code: "",
		}
		_, err := CreateShortLink(ctx, db, shortLink.URL)
		
		if err != nil {
			return count, err
		}
		count++
	}

	return count, scanner.Err()
}

// Question132 演示 CSV 批量导入。
func Question132() {
	deps, err := OpenDeps()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := AutoMigrateShortLink(deps.DB); err != nil {
		fmt.Println(err)
		return
	}
	ctx := context.Background()
	csv := strings.NewReader(
		`https://example.com/a
		https://example.com/b
		https://example.com/c

		https://example.com/c
		`)
	n, err := ImportCSV(ctx, deps.DB, csv)
	if err != nil {
		fmt.Println("import:", err)
		return
	}
	fmt.Println("imported:", n) // 空行跳过，应为 3
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 133. 访问统计：跳转时 Redis 计数；按需汇总落库
//
// Question133 练习 hits 计数与 Flush。
//
// 写函数：
//   IncrHit(ctx context.Context, rdb *redis.Client, code string) error
//     - INCR hitsKey(code)
//   FlushHits(ctx context.Context, rdb *redis.Client, db *gorm.DB, code string) error
//     - GET hits 计数（没有则 0）
//     - 将 ShortLink.Hits 更新为 DB 原值 + redis 增量，或直接写成累计策略（题中自定并注释）
//     - 可选：Flush 后 DEL / 置零 Redis 计数，避免重复累加
// 建议：Resolve 或 GET /r/{code} 成功时调用 IncrHit
// 在 Question133 中：创建 → 模拟几次命中 → FlushHits → 打印 DB 中 Hits
func IncrHit(ctx context.Context, rdb *redis.Client, code string) error {
	_, err := rdb.Incr(ctx, hitsKey(code)).Result()
	return err
}

func FlushHits(ctx context.Context, rdb *redis.Client, db *gorm.DB, code string) error {

	redisHitCount, err := rdb.Get(ctx, hitsKey(code)).Int64()
	if err == redis.Nil {
		redisHitCount = 0
	} else if err != nil {
		return err
	}

	var shortLink ShortLink
	if err := db.WithContext(ctx).Where("code = ?", code).First(&shortLink).Error; err != nil {
		return err
	}
	shortLink.Hits += redisHitCount
	if err := db.WithContext(ctx).Model(&shortLink).Update("hits", shortLink.Hits).Error; err != nil {
		return err
	}

	_ = rdb.Del(ctx, hitsKey(code)).Err()

	return nil
}

// Question133 演示访问计数与落库。
func Question133() {

	deps, err := OpenDeps()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := AutoMigrateShortLink(deps.DB); err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	code, err := CreateShortLink(ctx, deps.DB, "https://example.com/hits-demo")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer deps.DB.Unscoped().Where("code = ?", code).Delete(&ShortLink{})
	defer deps.RDB.Del(ctx, cacheKey(code), hitsKey(code))

	for i := 0; i < 3; i++ {
		if err := IncrHit(ctx, deps.RDB, code); err != nil {
			fmt.Println("incr:", err)
			return
		}
	}

	if err := FlushHits(ctx, deps.RDB, deps.DB, code); err != nil {
		fmt.Println("flush:", err)
		return
	}

	var row ShortLink
	if err := deps.DB.Where("code = ?", code).First(&row).Error; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("code:", code, "hits:", row.Hits)
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 134. 稳定性：关键路径 context 超时；核心路径基本测试
//
// Question134 给创建 / 解析加上超时，并跑一轮冒烟。
//
// 写函数：
//   CreateShortLinkTimeout(deps *Deps, longURL string, timeout time.Duration) (code string, err error)
//     - context.WithTimeout → CreateShortLink
//   ResolveURLTimeout(deps *Deps, code string, timeout time.Duration) (longURL string, err error)
//     - context.WithTimeout → ResolveURL
//   SmokeShortLink(deps *Deps) error
//     - 创建 → 解析 → 校验 URL 一致；失败返回 error
// 在 Question134 中：OpenDeps 后 SmokeShortLink，成功打印 "smoke ok"
// 另：shortlinkpractice_test.go 可对 Smoke / Resolve 再加独立 Test（可选）
func CreateShortLinkTimeout(deps *Deps, longURL string, timeout time.Duration) (string, error) {
	// TODO
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	return CreateShortLink(ctx, deps.DB, longURL)
}

func ResolveURLTimeout(deps *Deps, code string, timeout time.Duration) (string, error) {
	// TODO
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	return ResolveURL(ctx, deps.RDB, deps.DB, code)
}

func SmokeShortLink(deps *Deps) error {
	// TODO
	code, err := CreateShortLinkTimeout(deps, "https://example.com/a", time.Second) 
	if err != nil {
		return err
	}
	resolveUrl, err := ResolveURLTimeout(deps, code, time.Second)
	if err != nil {
		return err
	}
	if resolveUrl != "https://example.com/a" {
		return errors.New("Resolve url wrong")
	}
	return nil
}

// Question134 演示超时包装与冒烟。
func Question134() {

	deps, err := OpenDeps()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := AutoMigrateShortLink(deps.DB); err != nil {
		fmt.Println(err)
		return
	}

	if err := SmokeShortLink(deps); err != nil {
		fmt.Println("Smoke failes")
	}
	fmt.Println("Smoke ok")
}
