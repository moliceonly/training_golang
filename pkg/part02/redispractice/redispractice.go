// Package redispractice 对齐路线图 Part 02 · 2.5「Redis」。
// 每个知识要点至少一题；题干尽量贴近商品 / 登录 / 防刷场景。
// 题号：122 → 128
//
// 查看题干：
//
//	go doc training_golang/pkg/part02/redispractice.Question125
//
// 依赖本机 apt 安装的 Redis（训练统一不用 Docker 跑 Redis）。
// 默认地址 127.0.0.1:6379，可被环境变量 TRAINING_REDIS_ADDR 覆盖。
//
// 未配置或连不上时，Question 演示可打印提示并 return；单测可设 TRAINING_SKIP_REDIS=1。
package redispractice

import (
	"context"
	"fmt"
	"os"
	"time"
	"encoding/json"
	"sync"
	"errors"

	"github.com/redis/go-redis/v9"
)

const defaultRedisAddr = "127.0.0.1:6379"

func addrFromEnv() string {
	if v := os.Getenv("TRAINING_REDIS_ADDR"); v != "" {
		return v
	}
	return defaultRedisAddr
}

// ---------------------------------------------------------------------------
// 122. go-redis 连接：NewClient、Ping
//
// Question122 练习打开 Redis 连接并 Ping。
//
// 写函数：
//   OpenRedis(addr string) (*redis.Client, error)
//     - redis.NewClient(&redis.Options{Addr: addr})
//     - context.Background() 下 Ping，失败返回 error
// 在 Question122 中：OpenRedis(addrFromEnv())，成功打印 "redis ok"，失败打印 err
func OpenRedis(addr string) (*redis.Client, error) {
	// TODO
	rClient := redis.NewClient(&redis.Options{Addr: addr})
	ctx := context.Background()
	if err := rClient.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rClient, nil
}

// Question122 演示 Redis 连接。
func Question122() {
	if _, err := OpenRedis(addrFromEnv()); err != nil {
		fmt.Println(err)
		return 
	} 
	fmt.Println("redis ok")
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 123. 数据类型：String / Hash / List / Set / ZSet
// 场景：各写一条典型命令，体会用途差异
//
// Question123 用同一 client 演示五种类型。
//
// 写函数：
//   DemoRedisTypes(ctx context.Context, rdb *redis.Client, prefix string) error
//     - String：SET/GET  prefix+":str"
//     - Hash：HSET/HGET  prefix+":hash"  field name
//     - List：LPUSH/LRANGE prefix+":list"
//     - Set：SADD/SMEMBERS prefix+":set"
//     - ZSet：ZADD/ZRANGE WITHSCORES prefix+":zset"
//     - 可用 Pipeline 批量写入（可选）
// 在 Question123 中：连接后调用 DemoRedisTypes，打印各类型读出结果，defer 删掉练习 key
func DemoRedisTypes(ctx context.Context, rdb *redis.Client, prefix string) error {
	// TODO

	strKey := prefix + ":str"
	if err := rdb.Set(ctx, strKey, "Hello", 0).Err(); err != nil {
		return err
	}
	s, err := rdb.Get(ctx, strKey).Result()
	if err != nil {
		return err
	}
	fmt.Println(s)

	hashKey := prefix + ":hash"
	if err := rdb.HSet(ctx, hashKey, "name", "world").Err(); err != nil {
		return err
	}
	h, err := rdb.HGet(ctx, hashKey, "name").Result()
	if err != nil {
		return err
	}
	fmt.Println(h)

	listKey := prefix + ":list"
	if err := rdb.LPush(ctx, listKey, "list", "world").Err(); err != nil {
		return err
	}
	list, err := rdb.LRange(ctx, listKey, 0, 2).Result()
	if err != nil {
		return err
	}
	fmt.Println(list)

	setKey := prefix + ":set"
	if err := rdb.SAdd(ctx, setKey, "set", "world").Err(); err != nil {
		return err
	}
	set, err := rdb.SMembers(ctx, setKey).Result()
	if err != nil {
		return err
	}
	fmt.Println(set)

	zsetKey := prefix + ":zset"
	if err := rdb.ZAdd(ctx, zsetKey, 
		redis.Z{Score: 12.23, Member: "小明"},
		redis.Z{Score: 100.1, Member: "肖华"},
		redis.Z{Score: 82.3, Member: "小二"},
	).Err(); err != nil {
		return err
	}
	zset, err := rdb.ZRangeWithScores(ctx, zsetKey, 0, -1).Result()
	if err != nil {
		return err
	}
	fmt.Println(zset)

	return nil
}

// Question123 演示五种 Redis 数据类型。
func Question123() {
	rdb, err := OpenRedis(addrFromEnv())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rdb.Close()

	ctx := context.Background()
	prefix := fmt.Sprintf("demo:%d", time.Now().UnixNano())

	if err := DemoRedisTypes(ctx, rdb, prefix); err != nil {
		fmt.Println(err)
		fmt.Println()
	}

}

// ---------------------------------------------------------------------------
// 124. 过期 TTL；缓存穿透 / 击穿 / 雪崩思路
//
// Question124 练习带 TTL 写入，并写下缓存问题笔记。
//
// 写函数：
//   SetWithTTL(ctx context.Context, rdb *redis.Client, key, val string, ttl time.Duration) error
//     - SET key val EX（或 Set + expiration）
//   TTLSeconds(ctx context.Context, rdb *redis.Client, key string) (int64, error)
//     - TTL 命令，返回剩余秒（-1/-2 按 redis 语义原样返回亦可）
//   CacheProblemNotes() string
//     - 多行说明：穿透、击穿、雪崩各是什么 + 一种常见应对
// 在 Question124 中：SetWithTTL 短 TTL，打印 TTLSeconds 与 CacheProblemNotes
func SetWithTTL(ctx context.Context, rdb *redis.Client, key, val string, ttl time.Duration) error {
	// TODO
	return rdb.Set(ctx, key, val, ttl).Err()
}

func TTLSeconds(ctx context.Context, rdb *redis.Client, key string) (int64, error) {
	// TODO
	ttl, err := rdb.TTL(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	if ttl < 0 {
		return int64(ttl / time.Second), nil
	}
	return int64(ttl.Seconds()), nil
}

func CacheProblemNotes() string {
	// TODO
	return `穿透：查不存在的数据，缓存没有、DB 也没有，每次打到 DB。应对：缓存空值短 TTL，或布隆过滤。
	击穿：热点 key 突然过期，大量请求同时打 DB。应对：互斥锁重建缓存，或热点永不过期+异步刷新。
	雪崩：大量 key 同一时刻过期，DB 被打满。应对：TTL 加随机抖动，多级缓存，限流降级。`
}

// Question124 演示 TTL 与缓存问题笔记。
func Question124() {

	rdb, err := OpenRedis(addrFromEnv())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rdb.Close()

	ctx := context.Background()
	key := fmt.Sprintf("ttl:%d", time.Now().UnixNano())
	defer rdb.Del(ctx, key)

	if err := SetWithTTL(ctx, rdb, key, "v", 10*time.Second); err != nil {
		fmt.Println(err)
		return
	}
	sec, err := TTLSeconds(ctx, rdb, key)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sec)
	fmt.Println(CacheProblemNotes())
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 125. 商品详情页：先 Redis 后 DB；回填 TTL；更新时主动删缓存
// 场景：Cache-Aside（旁路缓存）；DB 用内存 map 模拟即可
//
// Product 表示商品。
type Product struct {
	ID    string
	Name  string
	Price int // 分
}

var (
	productMutex sync.Mutex
	productDB = map[string]Product{}
)

// 写函数：
//   可用包级 var productDB = map[string]Product{} 模拟数据库（注意演示里加锁或单测串行）
//   SaveProductDB(p Product)                         // 写入「库」
//   GetProductDB(id string) (Product, bool)
//   productCacheKey(id string) string                // 如 "product:"+id
//   GetProduct(ctx, rdb, id) (Product, error)
//     - 先 GET 缓存；命中则返回
//     - 未命中查 DB；没有则返回 error（可选：短 TTL 空值防穿透）
//     - 有则 JSON/简单字符串回填 Redis 并设 TTL（如 60s）
//   InvalidateProduct(ctx, rdb, id) error            // DEL 缓存
//   UpdateProduct(ctx, rdb, p Product) error         // 改 DB 后 Invalidate
// 在 Question125 中：Save → Get（应回填）→ 再 Get（应命中）→ Update → Get 再回填
func SaveProductDB(p Product) {
	// TODO
	productMutex.Lock()
	defer productMutex.Unlock()
	productDB[p.ID] = p
}

func GetProductDB(id string) (Product, bool) {
	// TODO
	productMutex.Lock()
	defer productMutex.Unlock()
	prod, ok := productDB[id]
	return prod, ok
}

func productCacheKey(id string) string {
	return "product:" + id
}

func GetProduct(ctx context.Context, rdb *redis.Client, id string) (Product, error) {
	// TODO

	key := productCacheKey(id)
	result, err := rdb.Get(ctx, key).Result()
	if err == nil {
		var p Product
		if err := json.Unmarshal([]byte(result), &p); err != nil {
			return Product{}, err
		}
		fmt.Println("hit cache", id)
		return p, nil
	}
	if err != redis.Nil {
		return Product{}, err
	}

	fmt.Println("cache miss", id)
	p, ok := GetProductDB(id)
	if !ok {
		return Product{}, fmt.Errorf("product not found: %s", id)
	}

	cache, err := json.Marshal(p)
	if err != nil {
		return Product{}, err
	}
	if err := rdb.Set(ctx, key, cache, 10 * time.Second).Err(); err != nil {
		return Product{}, err
	}

	return p, nil
}

func InvalidateProduct(ctx context.Context, rdb *redis.Client, id string) error {
	// TODO
	return rdb.Del(ctx, productCacheKey(id)).Err()
}

func UpdateProduct(ctx context.Context, rdb *redis.Client, p Product) error {
	// TODO
	SaveProductDB(p)
	return InvalidateProduct(ctx, rdb, p.ID)
}

// Question125 演示商品缓存旁路。
func Question125() {
	rdb, err := OpenRedis(addrFromEnv())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rdb.Close()
	ctx := context.Background()

	p := Product{ID: "p1", Name: "Go Book", Price: 9900}
	SaveProductDB(p)
	defer func() {
		productMutex.Lock()
		delete(productDB, p.ID)
		productMutex.Unlock()
		_ = InvalidateProduct(ctx, rdb, p.ID)
	}()

	p1, err := GetProduct(ctx, rdb, p.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p1)
	p2, err := GetProduct(ctx, rdb, p.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2)

	
	p.Price = 8800
	if err := UpdateProduct(ctx, rdb, p); err != nil {
		fmt.Println(err)
		return
	}
	p3, err := GetProduct(ctx, rdb, p.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p3)
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 126. 登录态：Token 存 Redis；退出立即失效
// 场景：不能等到自然过期
//
// 写函数：
//   tokenKey(token string) string  // 如 "sess:"+token
//   PutToken(ctx, rdb, token, userID string, ttl time.Duration) error
//   UserIDByToken(ctx, rdb, token string) (userID string, err error)
//     - key 不存在返回 error（或 "", err）
//   RevokeToken(ctx, rdb, token string) error  // Del，模拟退出登录
// 在 Question126 中：Put → 查到 userID → Revoke → 再查应失败
func tokenKey(token string) string {
	return "sess:" + token
}

func PutToken(ctx context.Context, rdb *redis.Client, token, userID string, ttl time.Duration) error {
	// TODO

	key := tokenKey(token)
	if err := rdb.Set(ctx, key, userID, ttl).Err(); err != nil {
		return err
	}
	return nil
}

func UserIDByToken(ctx context.Context, rdb *redis.Client, token string) (string, error) {
	// TODO

	result, err := rdb.Get(ctx, tokenKey(token)).Result()
	if err != nil {
		return "", err
	}
	return result, err
}

func RevokeToken(ctx context.Context, rdb *redis.Client, token string) error {
	// TODO
	return rdb.Del(ctx, tokenKey(token)).Err()
}

// Question126 演示 Token 存储与立即失效。
func Question126() {

	rdb, err := OpenRedis(addrFromEnv())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rdb.Close()
	ctx := context.Background()

	token := fmt.Sprintf("tok-%d", time.Now().UnixNano())
	defer rdb.Del(ctx, tokenKey(token))

	if err := PutToken(ctx, rdb, token, "user-42", time.Hour); err != nil {
		fmt.Println(err)
		return
	}
	uid, err := UserIDByToken(ctx, rdb, token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("user:", uid)

	if err := RevokeToken(ctx, rdb, token); err != nil {
		fmt.Println(err)
		return
	}
	_, err = UserIDByToken(ctx, rdb, token)
	if err != nil {
		fmt.Println("after revoke:", err)
		return
	}
	fmt.Println("unexpected: token still valid")
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 127. 分布式锁入门：SET NX EX；释放时校验 value
// 场景：避免误删别人的锁
//
// 写函数：
//   lockKey(name string) string  // 如 "lock:"+name
//   TryLock(ctx, rdb, name, token string, ttl time.Duration) (ok bool, err error)
//     - SET key token NX EX ttl
//   Unlock(ctx, rdb, name, token string) error
//     - 仅当 GET 的值 == token 时再 DEL（可用 Lua，或 Get+Del 练习版）
// 在 Question127 中：TryLock 成功 → 第二次同名应失败 → Unlock → 再 TryLock 成功
func lockKey(name string) string {
	return "lock:" + name
}

func TryLock(ctx context.Context, rdb *redis.Client, name, token string, ttl time.Duration) (bool, error) {
	// TODO

	key := lockKey(name)
	if _, err := rdb.Get(ctx, key).Result(); err == nil {
		return false, errors.New("已存在锁")
	}

	if err := rdb.Set(ctx, key, token, ttl).Err(); err != nil {
		return false, err
	}

	return true, nil
}

func Unlock(ctx context.Context, rdb *redis.Client, name, token string) error {
	// TODO

	key := lockKey(name)
	result, err := rdb.Get(ctx, key).Result()
	if err == nil {
		if result == token {
			return rdb.Del(ctx, key).Err()
		} else {
			return errors.New("token wrong")
		}
	}
	return err
}

// Question127 演示简易分布式锁。
func Question127() {

	rdb, err := OpenRedis(addrFromEnv())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rdb.Close()
	ctx := context.Background()

	name := "order:demo"
	defer rdb.Del(ctx, lockKey(name))

	token1 := fmt.Sprintf("t1-%d", time.Now().UnixNano())
	token2 := fmt.Sprintf("t2-%d", time.Now().UnixNano())

	ok, err := TryLock(ctx, rdb, name, token1, 10*time.Second)
	fmt.Println("lock1:", ok, err) 

	ok, err = TryLock(ctx, rdb, name, token2, 10*time.Second)
	fmt.Println("lock2:", ok, err) 

	if err := Unlock(ctx, rdb, name, token1); err != nil {
		fmt.Println("unlock:", err)
		return
	}
	ok, err = TryLock(ctx, rdb, name, token2, 10*time.Second)
	fmt.Println("lock3:", ok, err)
	fmt.Println()

}

// ---------------------------------------------------------------------------
// 128. 限流：固定窗口；按 IP 计数，超限返回 false（对应 HTTP 429）
// 场景：1 分钟内超过阈值拒绝
//
// 写函数：
//   rateKey(ip string, windowStart int64) string
//     - 如 fmt.Sprintf("rate:%s:%d", ip, windowStart)
//   AllowIP(ctx, rdb, ip string, limit int, window time.Duration) (allowed bool, err error)
//     - 用当前时间对齐到窗口起点（如 unix/window 秒）
//     - INCR key；若为 1 则 EXPIRE window
//     - count <= limit 则 allowed=true
// 在 Question128 中：limit=3，连续调 4 次，打印每次 allowed

func rateKey(ip string, windowStart int64) string {
	return fmt.Sprintf("rate:%s:%d", ip, windowStart)
}

func AllowIP(ctx context.Context, rdb *redis.Client, ip string, limit int, window time.Duration) (bool, error) {
	// TODO
	ttl := int64(window.Seconds())
	if ttl <= 0 {
		return false, errors.New("Invalid ttl")
	}
	now := time.Now().Unix()
	windowStart := now - (now % ttl)

	key := rateKey(ip, windowStart)
	count, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if count == 1 {
		if err := rdb.Expire(ctx, key, window).Err(); err != nil {
			return false, err
		}
	} 
	if count <= int64(limit) {
		return true, nil
	}
	return false, nil
}

// Question128 演示固定窗口限流。
func Question128() {

	rdb, err := OpenRedis(addrFromEnv())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rdb.Close()
	ctx := context.Background()

	ip := fmt.Sprintf("demo-%d", time.Now().UnixNano())
	window := time.Minute
	for i := 1; i <= 4; i++ {
		ok, err := AllowIP(ctx, rdb, ip, 3, window)
		fmt.Println(i, "allowed:", ok, "err:", err)
	}

}
