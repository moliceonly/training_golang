// Package httppractice 对齐路线图 Part 02 · 2.2「HTTP 客户端与服务端基础」。
// 每个知识要点至少一题；题干尽量贴近参考模拟场景。
// 题号：102 → 108
//
// 查看题干：
//
//	go doc training_golang/pkg/part02/httppractice.Question103
//
// 测试请优先用 net/http/httptest，避免在单测里永久 ListenAndServe。
package httppractice

import (
	"context"
	"net/http"
	"time"
)

// ---------------------------------------------------------------------------
// 102. HTTP 协议要点：方法、状态码、Header、Body、Cookie
//
// Question102 练习构造与解析一次「迷你」HTTP 交换的关键字段。
//
// 写函数：
//   WritePlainOK(w http.ResponseWriter, body string)
//     - 状态 200；Header Content-Type: text/plain; charset=utf-8
//     - 写 body
//   SetSessionCookie(w http.ResponseWriter, sessionID string)
//     - 设置 Cookie：Name=session_id，Value=sessionID，Path=/
//   ReadMethodAndQuery(r *http.Request) (method, q string)
//     - 返回 r.Method 与 r.URL.Query().Get("q")
// 在 Question102 中用 httptest.NewRecorder + httptest.NewRequest 演示并打印
func WritePlainOK(w http.ResponseWriter, body string) {
	// TODO
}

func SetSessionCookie(w http.ResponseWriter, sessionID string) {
	// TODO
}

func ReadMethodAndQuery(r *http.Request) (method, q string) {
	// TODO
	return "", ""
}

// Question102 演示方法、状态码、Header、Body、Cookie。
func Question102() {
	// TODO
}

// ---------------------------------------------------------------------------
// 103. 服务端：ServeMux、Handler；读 Query/JSON；统一写 JSON
// 场景：备忘录 API — 纯标准库 Todo CRUD，统一 JSON
//
// Question103 实现内存版备忘录路由（不必真 Listen；返回 http.Handler 即可）。
//
// 定义：
//   type Todo struct { ID int `json:"id"`; Title string `json:"title"`; Done bool `json:"done"` }
//
// 写函数：
//   WriteJSON(w http.ResponseWriter, status int, v any)
//     - Content-Type: application/json；json.NewEncoder(w).Encode(v)
//   NewTodoMux() http.Handler
//     - GET  /todos           → 返回全部 Todo 列表 JSON
//     - POST /todos           → Body JSON {title}，分配自增 ID，Done=false，返回新建对象
//     - GET  /todos?id=1      → 按 id 查询；不存在 404 JSON {"error":"..."}
//     （可用包级变量存 slice/map + Mutex；GET 列表与按 id 查询可用同一路径靠 Query 区分，或分路径，自定但文档写清）
// 在 Question102/测试中用 httptest 调 NewTodoMux 演示
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	// TODO
}

func NewTodoMux() http.Handler {
	// TODO: http.NewServeMux + HandleFunc
	return http.NewServeMux()
}

// Question103 演示备忘录 Todo JSON API（httptest）。
func Question103() {
	// TODO
}

// ---------------------------------------------------------------------------
// 104. 中间件链：日志、耗时、recover
// 场景：备忘录 API 请求进出打访问日志
//
// Question104 练习中间件：func(http.Handler) http.Handler。
//
// 写函数：
//   Middleware = func(http.Handler) http.Handler  （可用 type 别名）
//   WithLogging(next http.Handler) http.Handler
//     - 记录 method、path、耗时（time.Since）；可用 log/slog 或 fmt
//   WithRecover(next http.Handler) http.Handler
//     - defer recover；panic 时写 500 JSON {"error":"internal"}
//   Chain(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler
//     - 从外到内或从内到外包装（注释写清顺序）；例如先 Recover 再 Logging
// 在 Question104 中：Chain(NewTodoMux(), WithRecover, WithLogging)，用 httptest 请求并制造一次 panic 路由（可选 HandleFunc /panic）
func WithLogging(next http.Handler) http.Handler {
	// TODO
	return next
}

func WithRecover(next http.Handler) http.Handler {
	// TODO
	return next
}

func Chain(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {
	// TODO
	return h
}

// Question104 演示中间件链包裹 Todo API。
func Question104() {
	// TODO
}

// ---------------------------------------------------------------------------
// 105. 客户端：http.Client、Timeout；天气聚合场景
// 场景：带 2 秒超时调用第三方天气 API，解析后缓存到本地文件
//
// Question105 用可注入的 baseURL（测试里 httptest.Server.URL）模拟第三方。
//
// 写函数：
//   FetchWeather(ctx context.Context, client *http.Client, baseURL, city string) (summary string, err error)
//     - GET baseURL + "/weather?city=" + url.QueryEscape(city)
//     - 解析 JSON：{"city":"...","summary":"sunny"}，返回 summary
//   FetchWeatherCached(ctx context.Context, client *http.Client, baseURL, city, cachePath string) (string, error)
//     - 若 cachePath 存在且可读，直接返回文件内容
//     - 否则 FetchWeather，成功后 os.WriteFile 缓存，再返回
//   NewTimeoutClient(d time.Duration) *http.Client
//     - &http.Client{Timeout: d}（练习用 2*time.Second）
// 在 Question105 中：httptest 假天气服务 + 临时缓存文件演示；可再演示超时（服务端 Sleep > client Timeout）
func NewTimeoutClient(d time.Duration) *http.Client {
	// TODO
	return http.DefaultClient
}

func FetchWeather(ctx context.Context, client *http.Client, baseURL, city string) (summary string, err error) {
	// TODO
	return "", nil
}

func FetchWeatherCached(ctx context.Context, client *http.Client, baseURL, city, cachePath string) (string, error) {
	// TODO
	return "", nil
}

// Question105 演示带超时的天气客户端与文件缓存。
func Question105(tmpDir string) {
	// TODO
}

// ---------------------------------------------------------------------------
// 106. context 取消请求；重试与幂等意识
//
// Question106 练习带 ctx 的请求与简单重试（仅对幂等 GET）。
//
// 写函数：
//   GetWithRetry(ctx context.Context, client *http.Client, rawURL string, maxTries int) (status int, body []byte, err error)
//     - 使用 http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
//     - 失败（网络错误或 5xx）则重试，最多 maxTries 次；中间可 time.Sleep 很短
//     - ctx 取消则立即返回 ctx.Err()
//     - 注释：为何不对非幂等 POST 随意重试
// 在 Question106 中：httptest 前两次 500、第三次 200；以及短 timeout ctx 取消
func GetWithRetry(ctx context.Context, client *http.Client, rawURL string, maxTries int) (status int, body []byte, err error) {
	// TODO
	return 0, nil, nil
}

// Question106 演示 context 取消与 GET 重试。
func Question106() {
	// TODO
}

// ---------------------------------------------------------------------------
// 107. CORS、Content-Type、状态码语义（了解即可）
//
// Question107 练习给响应加上基础 CORS，并区分 400/404/500。
//
// 写函数：
//   WithCORS(next http.Handler) http.Handler
//     - 响应头：Access-Control-Allow-Origin: *
//     - 若 Method == OPTIONS：写 204 并 return（预检）
//   WriteAPIError(w http.ResponseWriter, status int, msg string)
//     - 用 WriteJSON 写 {"error": msg}；status 为 400/404/500 等
// 在 Question107 中用 httptest 打 OPTIONS 与一次 404 演示
func WithCORS(next http.Handler) http.Handler {
	// TODO
	return next
}

func WriteAPIError(w http.ResponseWriter, status int, msg string) {
	// TODO
}

// Question107 演示 CORS 与错误状态码 JSON。
func Question107() {
	// TODO
}

// ---------------------------------------------------------------------------
// 108. 用户详情路由：手写匹配 /users/:id，非法 id 返回 400
// 场景：用户详情路由
//
// Question108 不用第三方路由库，手写路径解析。
//
// 写函数：
//   ParseUserID(path string) (id int, ok bool)
//     - 匹配前缀 "/users/"，其余部分 Atoi 为 id；id<=0 或非数字则 ok=false
//   UserDetailHandler(w http.ResponseWriter, r *http.Request)
//     - 仅处理 GET；否则 405
//     - ParseUserID(r.URL.Path)；失败 WriteAPIError 400
//     - 成功 WriteJSON 200：{"id":id,"name":"user-<id>"}
// 在 Question108 中 httptest 请求 /users/42、/users/abc、/users/0
func ParseUserID(path string) (id int, ok bool) {
	// TODO
	return 0, false
}

func UserDetailHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// Question108 演示手写 /users/:id 路由。
func Question108() {
	// TODO
}
