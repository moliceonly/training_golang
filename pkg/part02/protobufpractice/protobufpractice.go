// Package protobufpractice 对齐路线图 Part 02 · 2.3「Protocol Buffers（proto3）」。
// 每个知识要点至少一题；题干尽量贴近参考模拟场景。
// 题号：109 → 113
//
// 查看题干：
//
//	go doc training_golang/pkg/part02/protobufpractice.Question111
//
// 生成 pb 代码（需已安装 protoc + protoc-gen-go）：
//
//	go generate ./pkg/part02/protobufpractice/...
package protobufpractice

//go:generate protoc --go_out=../../.. --go_opt=module=training_golang proto/user.proto

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"bytes"
	"io"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"

	"training_golang/pkg/part02/protobufpractice/userpb"
)

// ---------------------------------------------------------------------------
// 109. proto3 语法：message、字段编号、标量、repeated / map、枚举、嵌套
// 场景：用户信息报文 user.proto（id / name / balance + 扩展字段）
//
// Question109 阅读 proto/user.proto，用生成类型构造一个完整 User 并打印关键字段。
//
// 写函数：
//   NewDemoUser() *userpb.User
//     - Id=42, Name="alice", Balance=10050（分）
//     - Status=USER_STATUS_ACTIVE
//     - Tags=["vip","cn"]
//     - Metadata={"level":"gold"}
//     - Address.City="Beijing", Street="Wangfujing"
// 在 Question109 中打印 Id / Name / Balance / Status / Tags / City
func NewDemoUser() *userpb.User {
	// TODO
	return &userpb.User{
		Id:      42,
		Name:    "alice",
		Balance: 10050,
		Status:  userpb.UserStatus_USER_STATUS_ACTIVE,
		Tags:    []string{"vip", "cn"},
		Metadata: map[string]string{
			"level": "gold",
		},
		Address: &userpb.Address{
			City:   "Beijing",
			Street: "Wangfujing",
		},
	}
}

// Question109 演示 proto 生成结构体的字段语义。
func Question109() {
	u := NewDemoUser()
	fmt.Println("id:", u.Id)
	fmt.Println("name:", u.Name)
	fmt.Println("balance:", u.Balance)
	fmt.Println("status:", u.Status.String())
	fmt.Println("tags:", u.Tags)
	fmt.Println("city:", u.Address.City)
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 110. 工具链：protoc + protoc-gen-go；.proto → Go
//
// Question110 确认生成物可用，并打印 go generate / protoc 命令提示。
//
// 写函数：
//   ProtoToolchainHint() string
//     - 返回一段多行说明：如何 go generate、protoc 命令、go_package 含义
//   AssertUserTypeWorks() error
//     - NewDemoUser() 非 nil，且能 Marshal 一次
func ProtoToolchainHint() string {
	// TODO
	return `
	go generate:
	usage: go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]
	        $GOPACKAGE
                The name of the package of the file containing the directive.
	Generate runs commands described by directives within existing
	files. Those commands can run any process but the intent is to
	create or update Go source files.
	
	protoc: 
	Parse PROTO_FILES and generate output based on the options given:
	usage: protoc [options] file.proto
	`
}

func AssertUserTypeWorks() error {
	// TODO
	u := NewDemoUser()
	if u == nil {
		return fmt.Errorf("NewDemoUser returned nil")
	}
	b, err := proto.Marshal(u)
	if err != nil {
		return err
	}
	if len(b) == 0 {
		return fmt.Errorf("EncodeUser returned empty bytes")
	}
	return nil
}

// Question110 打印工具链说明并做一次冒烟检查。
func Question110() {
	fmt.Println(ProtoToolchainHint())
	if err := AssertUserTypeWorks(); err != nil {
		fmt.Println("assert:", err)
		return
	}
	fmt.Println("assert: ok")
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 111. 序列化：proto.Marshal / proto.Unmarshal
//
// Question111 练习二进制往返。
//
// 写函数：
//   EncodeUser(u *userpb.User) ([]byte, error)
//     - proto.Marshal(u)
//   DecodeUser(b []byte) (*userpb.User, error)
//     - proto.Unmarshal 到新 User
// 在 Question111 中：NewDemoUser → Encode → Decode → 打印 Name / Balance / 字节长度
func EncodeUser(u *userpb.User) ([]byte, error) {
	// TODO
	return proto.Marshal(u)
}

func DecodeUser(b []byte) (*userpb.User, error) {
	// TODO
	u := &userpb.User{}
	if err := proto.Unmarshal(b, u); err != nil {
		return nil, err
	}
	return u, nil
}

// Question111 演示 Marshal / Unmarshal 往返。
func Question111() {
	u := NewDemoUser()
	b, err := EncodeUser(u)
	if err != nil {
		fmt.Println("encode:", err)
		return
	}
	u2, err := DecodeUser(b)
	if err != nil {
		fmt.Println("decode:", err)
		return
	}
	fmt.Println("name:", u2.Name)
	fmt.Println("balance:", u2.Balance)
	fmt.Println("bytes:", len(b))
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 112. HTTP 承载：Content-Type application/protobuf 读写 Body
// 场景：Protobuf API — 读请求 Body 反序列化，处理后序列化写回
//
// Question112 实现纯标准库 Handler（可用 httptest 测）。
//
// 写函数：
//   WriteProtobuf(w http.ResponseWriter, status int, msg proto.Message)
//     - Header Content-Type: application/protobuf
//     - Marshal 后 WriteHeader + Write
//   UserEchoHandler(w, r)
//     - 仅 POST；否则 405
//     - io.ReadAll(r.Body) → Unmarshal 为 User
//     - 若 Id==0：400，可用 JSON {"error":"..."} 或空 body（自定并注释）
//     - 否则把 Balance += 100，再 WriteProtobuf 200 写回
// 在 Question112 中用 httptest 对 UserEchoHandler POST 一段 protobuf
func WriteProtobuf(w http.ResponseWriter, status int, msg proto.Message) {
	// TODO: msg 用 proto.Message
	w.Header().Set("Content-Type", "application/protobuf")
	mmsg, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	w.WriteHeader(status)
	w.Write(mmsg)
}

func UserEchoHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	} 
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return 
	}

	user := &userpb.User{}
	if err := proto.Unmarshal(body, user); err != nil {
		fmt.Println(err)
		return
	}

	if user.Id == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad id"}`))
		return
	}

	user.Balance += 100
	WriteProtobuf(w, 200, user)
	return
}

// Question112 演示 HTTP + protobuf Body。
func Question112() {
	u := NewDemoUser()
	b, err := EncodeUser(u)
	if err != nil {
		fmt.Println("encode:", err)
		return
	}

	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(b))
	rr := httptest.NewRecorder()

	UserEchoHandler(rr, req)

	fmt.Println("status:", rr.Code)
	fmt.Println("Content-Type:", rr.Header().Get("Content-Type"))

	out, err := DecodeUser(rr.Body.Bytes())
	if err != nil {
		fmt.Println("decode:", err)
		return
	}

	fmt.Println("name:", out.Name, "balance:", out.Balance) // 期望 balance = 10150
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 113. 与 JSON 对比：体积、解析代码、字段号兼容直觉
// 场景：契约对照 — 同一用户数据 JSON vs Protobuf 往返
//
// Question113 对比同一份用户数据两种编码。
//
// 写函数：
//   UserToJSON(u *userpb.User) ([]byte, error)
//     - encoding/json 序列化（可用中间 struct 或 protojson；练习用简单 struct 也可）
//   CompareEncodings(u *userpb.User) (jsonLen, protoLen int, err error)
//     - 分别算 JSON 与 Protobuf 字节长度并返回
// 在 Question113 中打印两种长度，并注释：为何改字段名不影响二进制兼容、改字段号会怎样
func UserToJSON(u *userpb.User) ([]byte, error) {
	// TODO
	return protojson.Marshal(u)
}

func CompareEncodings(u *userpb.User) (jsonLen, protoLen int, err error) {
	// TODO
	jb, err := UserToJSON(u)
	if err != nil {
		return 0, 0, err
	}
	pb, err := EncodeUser(u)
	if err != nil {
		return 0, 0, err
	}
	return len(jb), len(pb), nil
}

// Question113 对比 JSON 与 Protobuf 体积。
func Question113() {
	u := NewDemoUser()
	jsonLen, protoLen, err := CompareEncodings(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("json bytes:", jsonLen)
	fmt.Println("proto bytes:", protoLen)
	fmt.Println()
}
