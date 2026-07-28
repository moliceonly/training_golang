// Package stdlibpractice 对齐路线图 1.4「标准库 SDK（非网络 IO）」。
// 每个知识要点至少一题；题干尽量贴近参考模拟场景。
// 题号：75 → …（自行实现 TODO）
package stdlibpractice

import (
	"context"
	"time"
)

// ---------------------------------------------------------------------------
// 75. fmt / strings / strconv / bytes / unicode
// 场景引申：API 入参清洗（用户名、年龄字符串）
//
// 写函数 SanitizeProfile(rawName string, ageStr string) (name string, age int, err error)：
//   - 用 strings.TrimSpace 去掉首尾空白
//   - 用 unicode 判断：清洗后 name 不得为空，且每个 rune 应为字母或数字（可用 unicode.IsLetter/IsDigit）
//   - 用 strconv.Atoi 解析 ageStr；非法则返回明确 error
//   - 用 fmt.Errorf 包装错误信息（例如 "invalid age: %w"）
//   - 额外：用 bytes.ToUpper 把 name 的字节形式转大写再转回 string 作为返回 name
//     （体会 string/[]byte 转换；真实业务更常用 strings.ToUpper，此处为练 bytes）
// 在 Question75 中分别演示成功与失败各一次并打印
func SanitizeProfile(rawName string, ageStr string) (name string, age int, err error) {
	// TODO
	return "", 0, nil
}

func Question75() {
	// TODO
}

// ---------------------------------------------------------------------------
// 76. time：解析、格式化、Duration、时区、定时器直觉
// 场景：活动倒计时 — 按用户时区计算报名截止时间
//
// 写函数 DeadlineInLocation(deadlineUTC string, locName string) (remain time.Duration, localStr string, err error)：
//   - deadlineUTC 用 time.Parse 按 RFC3339 解析（例如 "2026-08-01T12:00:00Z"）
//   - locName 用 time.LoadLocation（例如 "Asia/Shanghai"）
//   - 将截止时刻转到该时区，localStr 用 Format("2006-01-02 15:04:05 MST") 输出
//   - remain = 截止时刻.Sub(当前时间)；若已过期 remain 可为负
// 在 Question76 中打印 remain 与 localStr；可用 time.NewTimer 演示「剩不到 1s 就等一下」的直觉（可选）
func DeadlineInLocation(deadlineUTC string, locName string) (remain time.Duration, localStr string, err error) {
	// TODO
	return 0, "", nil
}

func Question76() {
	// TODO
}

// ---------------------------------------------------------------------------
// 77. encoding/json：序列化 / 反序列化；tag、忽略字段、自定义 Marshal
// 场景：API 契约 — 结构体 ↔ JSON 往返；敏感字段忽略；非法 JSON 返回明确错误
//
// 定义：
//   type APIUser struct {
//     ID       uint64 `json:"id"`
//     Name     string `json:"name"`
//     Password string `json:"-"`                    // 响应中绝不出现
//     Role     Role   `json:"role"`
//   }
//   type Role string  // 例如 "admin" / "user"
//   为 Role 实现 MarshalJSON / UnmarshalJSON：
//     - 对外 JSON 必须是大写枚举名 "ADMIN"/"USER"（或你自定映射）
//     - 非法值 Unmarshal 时返回 error
//
// 写函数：
//   EncodeAPIUser(u APIUser) ([]byte, error)
//   DecodeAPIUser(data []byte) (APIUser, error)  // 非法 JSON 应返回明确 error
// 在 Question77 中：合法往返打印；再传入非法 JSON 打印 error；确认 Password 不会出现在 JSON 里
type Role string

type APIUser struct {
	// TODO: 字段与 json tag
}

func (r Role) MarshalJSON() ([]byte, error) {
	// TODO
	return nil, nil
}

func (r *Role) UnmarshalJSON(b []byte) error {
	// TODO
	return nil
}

func EncodeAPIUser(u APIUser) ([]byte, error) {
	// TODO
	return nil, nil
}

func DecodeAPIUser(data []byte) (APIUser, error) {
	// TODO
	return APIUser{}, nil
}

func Question77() {
	// TODO
}

// ---------------------------------------------------------------------------
// 78. encoding/base64：标准 / URL 编码与解码
// 场景：令牌传输 — 二进制密文放进 Header/URL，对端解码还原
//
// 写函数：
//   TokenToHeader(raw []byte) string   // 用 base64.StdEncoding
//   TokenFromHeader(s string) ([]byte, error)
//   TokenToURL(raw []byte) string      // 用 base64.URLEncoding（URL 安全）
//   TokenFromURL(s string) ([]byte, error)
// 在 Question78 中：对同一段 []byte 分别走 Header 与 URL 路径，解码后 bytes.Equal 应为 true；
// 并打印两种编码字符串，观察 '+' '/' 与 '-' '_' 的差异
func TokenToHeader(raw []byte) string {
	// TODO
	return ""
}

func TokenFromHeader(s string) ([]byte, error) {
	// TODO
	return nil, nil
}

func TokenToURL(raw []byte) string {
	// TODO
	return ""
}

func TokenFromURL(s string) ([]byte, error) {
	// TODO
	return nil, nil
}

func Question78() {
	// TODO
}

// ---------------------------------------------------------------------------
// 79. crypto/sha256：摘要；密码落库常见用法
// 场景：密码落库前 — 对 盐 + 密码 做 SHA256（了解生产应优先 bcrypt/argon2；此处练哈希 API）
//
// 写函数 HashPassword(salt, password string) string：
//   - 拼接 salt + password（或 password + salt，自定但要注释）
//   - sha256.Sum256，返回 hex 字符串（encoding/hex）
// 写函数 CheckPassword(salt, password, wantHex string) bool：对比哈希是否一致
// 在 Question79 中：同一盐密应 Check 成功；改密码应失败
func HashPassword(salt, password string) string {
	// TODO
	return ""
}

func CheckPassword(salt, password, wantHex string) bool {
	// TODO
	return false
}

func Question79() {
	// TODO
}

// ---------------------------------------------------------------------------
// 80. crypto/aes + cipher：GCM；密钥与 Nonce 管理意识
// 场景：配置脱敏 — AES-GCM 加密配置中的密钥字段，启动时解密注入内存
//
// 写函数：
//   EncryptSecret(key, plaintext []byte) (nonceAndCipher []byte, err error)
//   DecryptSecret(key, nonceAndCipher []byte) (plaintext []byte, err error)
// 要求：
//   - key 长度必须是 16/24/32 字节，否则返回 error
//   - 使用 aes.NewCipher + cipher.NewGCM
//   - Nonce 用 io.ReadFull(rand.Reader, nonce) 生成；把 nonce 前缀拼在密文前返回
//   - 解密时先拆出 nonce 再 Open
// 在 Question80 中：加密 "db-password-demo"，再解密打印；key 可用 32 字节常量（仅练习）
func EncryptSecret(key, plaintext []byte) (nonceAndCipher []byte, err error) {
	// TODO
	return nil, nil
}

func DecryptSecret(key, nonceAndCipher []byte) (plaintext []byte, err error) {
	// TODO
	return nil, nil
}

func Question80() {
	// TODO
}

// ---------------------------------------------------------------------------
// 81. encoding/csv（了解即可；xml 可在注释中对比一句）
// 场景引申：把用户列表导出为 CSV，供运营下载
//
// 写函数 UsersToCSV(users []APIUser) (string, error)：
//   - 表头：id,name,role
//   - 不要输出 Password
//   - 使用 encoding/csv.Writer 写到 strings.Builder 或 bytes.Buffer
// 在 Question81 中打印 CSV 文本
// 选做：用 encoding/xml 把同一个切片 Marshal 成 XML 打印，体会差异即可
func UsersToCSV(users []APIUser) (string, error) {
	// TODO
	return "", nil
}

func Question81() {
	// TODO
}

// ---------------------------------------------------------------------------
// 82. regexp 基础；sort / slices / maps
// 场景引申：校验手机号；对标签排序去重
//
// 写函数：
//   ValidCNMobile(s string) bool
//     使用 regexp.MustCompile 匹配中国大陆手机号粗规则：1 开头共 11 位数字
//   NormalizeTags(tags []string) []string
//     - 用 strings.TrimSpace；去掉空串
//     - 用 maps 或手动 map 去重
//     - 用 slices.Sort 排序后返回
// 在 Question82 中演示合法/非法手机号，以及标签规范化前后对比
func ValidCNMobile(s string) bool {
	// TODO
	return false
}

func NormalizeTags(tags []string) []string {
	// TODO
	return nil
}

func Question82() {
	// TODO
}

// ---------------------------------------------------------------------------
// 83. flag 入门（不必上 cobra）
// 场景：运维小工具 — 启动参数
//
// 写函数 RunFlagDemo(args []string) (configPath string, verbose bool, err error)：
//   - 使用 flag.NewFlagSet("ops", flag.ContinueOnError)
//   - 定义：-config 默认 "config.json"；-verbose 默认 false
//   - Parse(args)，返回解析结果
// 在 Question83 中分别传入 []string{"-config", "app.json", "-verbose"} 与空参数并打印
func RunFlagDemo(args []string) (configPath string, verbose bool, err error) {
	// TODO
	return "", false, nil
}

func Question83() {
	// TODO
}

// ---------------------------------------------------------------------------
// 84. log/slog：结构化日志基础
// 场景：运维小工具 — 关键操作打结构化日志
//
// 写函数 LogDeploy(service, version string, verbose bool)：
//   - 用 slog.Info 记录事件 "deploy"，带属性 service、version
//   - 若 verbose==true，再 slog.Debug 一条 "deploy details"
//   - 可先 slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))
// 在 Question84 中调用两次（verbose true/false）观察输出
func LogDeploy(service, version string, verbose bool) {
	// TODO
}

func Question84() {
	// TODO
}

// ---------------------------------------------------------------------------
// 85. context：WithCancel / WithTimeout / WithValue（先理解语义）
// 场景：活动倒计时后 — 超时用 context 取消后续处理
//
// 写函数：
//   ProcessWithTimeout(parent context.Context, d time.Duration, work func(ctx context.Context) error) error
//     - 用 context.WithTimeout(parent, d)；defer cancel
//     - 调用 work(ctx)；若 ctx 超时，work 应能通过 ctx.Done() 感知（见下方模拟）
//   SlowWork(ctx context.Context) error
//     - 用 select：case <-time.After(200*time.Millisecond) 成功返回 nil
//                  case <-ctx.Done() 返回 ctx.Err()
// 在 Question85 中：
//   1) timeout=1s 调用，应成功
//   2) timeout=50ms 调用，应得到 context.DeadlineExceeded
//   3) 用 context.WithValue 塞一个 request_id，在 SlowWork 或包装里读出并打印（练 WithValue 语义即可）
func ProcessWithTimeout(parent context.Context, d time.Duration, work func(ctx context.Context) error) error {
	// TODO
	return nil
}

func SlowWork(ctx context.Context) error {
	// TODO
	return nil
}

func Question85() {
	// TODO
}
