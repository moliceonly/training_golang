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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
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
	w.Header().Set("Content-Type", "text/plain; charset=uft-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}

func SetSessionCookie(w http.ResponseWriter, sessionID string) {
	// TODO
	http.SetCookie(w, &http.Cookie{
		Name: "sesstion_id",
		Value: sessionID,
		Path: "/",
	})
}

func ReadMethodAndQuery(r *http.Request) (method, q string) {
	// TODO
	return r.Method, r.URL.Query().Get("q")
}

// Question102 演示方法、状态码、Header、Body、Cookie。
func Question102() {
	// TODO

		req := httptest.NewRequest(http.MethodGet, "/hello?q=world", nil)
		method, q := ReadMethodAndQuery(req)
		fmt.Println("method:", method, "q:", q)

		rr := httptest.NewRecorder()
		WritePlainOK(rr, "ok")
		SetSessionCookie(rr, "abc-123")

		fmt.Println("status:", rr.Code)
		fmt.Println("Content-Type:", rr.Header().Get("Content-Type"))
		fmt.Println("body:", rr.Body.String())
		fmt.Println("Set-Cookie:", rr.Header().Get("Set-Cookie"))
		fmt.Println()
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

var (
	todos []Todo
	mutex sync.Mutex
	todoID = 0
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	// TODO
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func NewTodoMux() http.Handler {
	// TODO: http.NewServeMux + HandleFunc
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			mutex.Lock()
			defer mutex.Unlock()
			if idStr := r.URL.Query().Get("id"); idStr != "" {
				id, err := strconv.Atoi(idStr)
				if err != nil {
					WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "bad id"})
					return
				}
				for _, t := range todos {
					if t.ID == id {
						WriteJSON(w, http.StatusOK, t)
						return
					}
				}
				WriteJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
				return 
			}
			WriteJSON(w, http.StatusOK, todos)

		case http.MethodPost:
			var title struct {
				Title string `json:"title"`
			}
			if err := json.NewDecoder(r.Body).Decode(&title); err != nil {
				WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "bad json"})
				return
			}
			mutex.Lock()
			defer mutex.Unlock()
			newTodo := Todo{
				ID: todoID,
				Title: title.Title,
				Done: false,
			}
			todoID++
			todos = append(todos, newTodo)
			WriteJSON(w, http.StatusOK, newTodo)
		
		default:
			WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method"})
		} 
	})
	return serveMux
}

// Question103 演示备忘录 Todo JSON API（httptest）。
func Question103() {
	h := NewTodoMux()

	body := strings.NewReader(`{"title":"buy apple"}`)
	req := httptest.NewRequest(http.MethodPost, "/todos", body)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	fmt.Println(req.Method, rr.Code, rr.Body.String())

	body = strings.NewReader(`{"title":"buy banana"}`)
	req = httptest.NewRequest(http.MethodPost, "/todos", body)
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	fmt.Println(req.Method, rr.Code, rr.Body.String())

	req = httptest.NewRequest(http.MethodGet, "/todos", nil)
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	fmt.Println(req.Method, rr.Code, rr.Body.String())

	req = httptest.NewRequest(http.MethodGet, "/todos?id=1", nil)
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	fmt.Println(req.Method, rr.Code, rr.Body.String())
	fmt.Println()
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("%s %s %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}

func WithRecover(next http.Handler) http.Handler {
	// TODO
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		defer func() {
			if rec := recover(); rec != nil {
				WriteJSON(w, http.StatusInternalServerError, map[string]string{"error":"internal"})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Chain(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {
	// TODO
	for i := len(mws) - 1; i>= 0; i-- {
		h = mws[i](h)
	}
	return h
}

// Question104 演示中间件链包裹 Todo API。
func Question104() {
	
	mux := http.NewServeMux()
	inner := NewTodoMux()
	mux.Handle("/", inner)
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r * http.Request) {
		panic("panic")
	})

	h := Chain(mux, WithRecover, WithLogging)

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	fmt.Println("todos", rr.Code, rr.Body.String())

	req = httptest.NewRequest(http.MethodGet, "/panic", nil)
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	fmt.Println("panic:", rr.Code, rr.Body.String())
	fmt.Println()
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
	return &http.Client{Timeout: d}
}

func FetchWeather(ctx context.Context, client *http.Client, baseURL, city string) (summary string, err error) {
	// TODO
	u := baseURL + "/weather?city=" + url.QueryEscape(city)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return "", err
	}

	rr, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer rr.Body.Close()

	var output struct {
		City string `json:"city"`
		Summary string `json:"summary"`
	}
	if err := json.NewDecoder(rr.Body).Decode(&output); err != nil {
		return "", err
	}
	return output.Summary, nil
}

func FetchWeatherCached(ctx context.Context, client *http.Client, baseURL, city, cachePath string) (string, error) {
	// TODO
	if cache, err := os.ReadFile(cachePath); err == nil {
		return string(cache), nil
	}
	summary, err := FetchWeather(ctx, client, baseURL, city)
	if err != nil {
		return "", nil
	}
	if err := os.WriteFile(cachePath, []byte(summary), 0o644); err != nil {
		return "", nil
	}
	return summary, nil
}

// Question105 演示带超时的天气客户端与文件缓存。
func Question105(tmpDir string) {
	
	client := NewTimeoutClient(3 * time.Second)

	serve := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]string{
			"city": "beijing",
			"summary": "rainy",
		})
	}))
	defer serve.Close()

	ctx := context.Background()
	cache := filepath.Join(tmpDir, "cache")

	ser1, err := FetchWeatherCached(ctx, client, serve.URL, "beijing", cache)
	fmt.Println(ser1, err)

	ser2, err := FetchWeatherCached(ctx, client, serve.URL, "beijing", cache)
	fmt.Println(ser2, err)

	timeout := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
	}))
	defer timeout.Close()
	_, err = FetchWeather(ctx, client, timeout.URL, "beijing")
	fmt.Println(err)
	fmt.Println()
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
	var lastErr error
	for count := 0; count < maxTries; count++ {
		if err := ctx.Err(); err != nil {
			fmt.Println("try", count+1, "ctx:", err)
			return 0, nil, err
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
		if err != nil {
			return 0, nil, err
		}

		rr, err := client.Do(req)
		if err != nil {
			lastErr = err
			if ctx.Err() != nil {
				fmt.Println("try", count+1, "do err:", err)
				return 0, nil, ctx.Err()
			}
			time.Sleep(10 * time.Millisecond)
			continue
		}

		b, err := io.ReadAll(rr.Body)
		rr.Body.Close()
		if err != nil {
			lastErr = err
			time.Sleep(10 * time.Millisecond)
			fmt.Println("try", count+1, "status:", rr.StatusCode, "body:", string(b))
			continue
		}

		if rr.StatusCode >= 500 {
			lastErr = fmt.Errorf("status %d", rr.StatusCode)
			if count == maxTries - 1 {
				fmt.Println("try", count+1, "status:", rr.StatusCode, "body:", string(b))
				return rr.StatusCode, b, lastErr
			}
			fmt.Println("try", count+1, "status:", rr.StatusCode, "body:", string(b))
			time.Sleep(10 * time.Millisecond)
			continue
		}

		return rr.StatusCode, b, nil
	}
	return 0, nil, lastErr
}

// Question106 演示 context 取消与 GET 重试。
func Question106() {
	n := 0
	serve := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n++
		if n < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("fail"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer serve.Close()

	client := &http.Client{Timeout: 2 * time.Second}
	st, body, err := GetWithRetry(context.Background(), client, serve.URL, 5)
	fmt.Println(st, string(body), err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	slow := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
	}))
	defer slow.Close()
	_, _, err = GetWithRetry(ctx, client, slow.URL, 5)
	fmt.Println(err)
	fmt.Println()
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func WriteAPIError(w http.ResponseWriter, status int, msg string) {
	// TODO
	WriteJSON(w, status, map[string]string{"error": msg})
}

// Question107 演示 CORS 与错误状态码 JSON。
func Question107() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]string{"ok": "true"})
	})
	
	mux.HandleFunc("/missing", func(w http.ResponseWriter, r *http.Request) {
		WriteAPIError(w, http.StatusNotFound, "not found")
	})
	h := WithCORS(mux)

	req := httptest.NewRequest(http.MethodOptions, "/ping", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	fmt.Println("OPTIONS:", rr.Code, rr.Header().Get("Access-Control-Allow-Origin"))

	req = httptest.NewRequest(http.MethodGet, "/missing", nil)
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	fmt.Println("404:", rr.Code, rr.Body.String())
	fmt.Println()
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
	const prefix = "/users/"
	if !strings.HasPrefix(path, prefix) {
		return 0, false
	}
	rest := strings.TrimPrefix(path, prefix)
	if rest == "" {
		return 0, false
	}
	id, err := strconv.Atoi(rest)
	if err != nil || id <= 0 {
		return 0, false
	}
	return id, true
}

func UserDetailHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
	if r.Method != http.MethodGet {
		WriteAPIError(w, http.StatusMethodNotAllowed, "method")
		return
	}
	id, ok := ParseUserID(r.URL.Path)
	if !ok {
		WriteAPIError(w, http.StatusBadRequest, "bad id")
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{
		"id":   id,
		"name": fmt.Sprintf("user-%d", id),
	})
}

// Question108 演示手写 /users/:id 路由。
func Question108() {
	paths := []string{"/users/42", "/users/abc", "/users/0"}
	for _, p := range paths {
		req := httptest.NewRequest(http.MethodGet, p, nil)
		rr := httptest.NewRecorder()
		UserDetailHandler(rr, req)
		fmt.Println(p, rr.Code, rr.Body.String())
	}
	fmt.Println()
}
