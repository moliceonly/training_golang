package errorinterfacepractice

import (
	"fmt"
	"errors"
)
//57.定义包级错误变量 ErrNotFound（使用 errors.New），表示「资源未找到」
//   在后续题目的 Get 等方法中复用它
//58.写函数 Div(a, b int) (int, error)：b==0 时返回错误，否则返回商
//   在 Question57and58 中分别演示除零与正常除法
var ErrNotFound = errors.New("资源未找到")
func Div(a, b int) (int, error) {
	if b==0 {
		return 0, errors.New("除数为0, 不计算\n")
	} else {
		return a / b, nil
	}
}
func Question57and58(a, b int) {
	divnum, err := Div(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d除以%d为%d\n", a, b, divnum)
	}
	fmt.Println()
}

//59.定义接口 UserStore，包含方法：
//   Save(u User) error
//   Get(id uint64) (User, error)
//60.用 map 实现 MemoryUserStore（可存取 User）
//   Get 不到时返回 ErrNotFound
//   提供 NewMemoryUserStore() *MemoryUserStore
//   在 Question59and60 中：Save 一个用户，Get 成功与失败各演示一次

type User struct {
	id uint64
	name string
}

type MemoryUserStore map[uint64]string

type UserStore interface {
	Save(u User) error
	Get(id uint64) (User, error)
}

func NewMemoryUserStore() *MemoryUserStore {
	store := make(MemoryUserStore)
	return &store
}

func (mus MemoryUserStore)Save(u User) error {
	if _, ok := mus[u.id]; ok {
		return errors.New("用户已存在")
	} else {
		mus[u.id] = u.name
		return nil
	}
}

func (mus MemoryUserStore)Get(id uint64) (User, error) {
	if name, ok := mus[id]; ok {
		getuser := User{ id: id, name: name,}
		return getuser, nil
	} else {
		return User{}, ErrNotFound
	}
}

func Question59and60() {
	u1 := User{
		id: 123,
		name: "小明",
	}
	u2 := User{
		id: 124,
		name: "小美",
	}
	store := NewMemoryUserStore()
	//正常存储以及重复存储
	store.Save(u1)
	store.Save(u2)
	fmt.Println(store.Save(u1))
	fmt.Println(store)
	//正常查询和异常查询
	var i uint64 = 124
	for i<126{
		fmt.Printf("正在查询id为%d的用户\t", i)
		queryuser, err := store.Get(i)
		if err == nil {
			fmt.Printf("查询成功\t")
			fmt.Println(queryuser)
		} else {
			fmt.Errorf("%w", err)
			fmt.Println(err)
		}
		i++
	}
	fmt.Println()
}

//61.自定义错误类型 AgeError，字段 Age int，实现 Error() string
//   文案需包含非法年龄，例如：invalid age: -1
//62.写函数 ValidateUser(u User) error：
//   Age < 0 时返回 AgeError；Name 为空时返回 errors.New("empty name")；否则返回 nil
//   在 Question61and62 中用非法 User 调用并打印 error
type person struct {
	Age int
	Name string
}
type AgeError struct {
	Age int
}
func (ageerr *AgeError) Error() string {
	return fmt.Sprintf("invalid age: %d", ageerr.Age)
}

func ValidateUser(p person) error {
	switch {
		case p.Age < 0:
			return &AgeError {
				Age: p.Age,
			}
		case p.Name == "":
			err := errors.New("empty name")
			return err
		default:
			return nil
	}
}
func Question61and62() {
	p1 := person{
		Age: -2,
		Name: "小花",
	}
	p2 := person{
		Age: 18,
		Name: "",
	}
	p3 := person{
		Age: 12,
		Name: "小红",
	}
	fmt.Println(ValidateUser(p1))
	fmt.Println(ValidateUser(p2))
	fmt.Println(ValidateUser(p3))
	fmt.Println()
}

//63.写函数 PrintStoreUser(store UserStore, id uint64)
//   使用 errors.Is 判断是否为 ErrNotFound，对「未找到」和「其他错误/成功」打印不同提示
//   在 Question63 中演示未找到的情况
func Question63() {
	//沿用Q59&60
	u1 := User{
		id: 123,
		name: "小明",
	}
	u2 := User{
		id: 124,
		name: "小美",
	}
	store := NewMemoryUserStore()
	store.Save(u1)
	store.Save(u2)
	var i uint64 = 124
	for i<126{
		fmt.Printf("正在查询id为%d的用户\t", i)
		queryuser, err := store.Get(i)
		if errors.Is(err, ErrNotFound) {
			fmt.Printf("未找到相应用户\n")
		} else {
			fmt.Printf("查找成功\t")
			fmt.Println(queryuser)
		}
		i++
	}
	fmt.Println()
}

//64.定义接口 Validator，包含方法 Validate() error
//   让 User 实现 Validator（可复用 ValidateUser 的逻辑）
//   在 Question64 中把 User 赋给 Validator 变量并调用 Validate
type Validator interface {
	Validate() error
}
func (p person) Validate() error {
	return ValidateUser(p)
}
func Question64() {
	var v Validator = person{
		Age: 18,
		Name: "",
	}
	fmt.Println(v.Validate())
	fmt.Println()
}

//65.写函数 RunValidators(vs ...Validator) []error
//   依次调用 Validate，收集所有非 nil 的 error 并返回
//   在 Question65 中传入多个 User（合法与非法），打印收集到的 errors
func RunValidators(vs ...Validator) []error {
	errs := make([]error, 0, 0)
	for _, val := range vs {
		err := val.Validate()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
func Question65() {
	p1 := person{
		Age: -2,
		Name: "小花",
	}
	p2 := person{
		Age: 18,
		Name: "",
	}
	p3 := person{
		Age: 12,
		Name: "小红",
	}
	for index, err := range RunValidators(p1, p2, p3){
		fmt.Printf("第%d个错误: %s\n",index, err)
	}
	fmt.Println()
}


// ========== 对齐 go-web-roadmap Part 1.2 补强题（仅题干）==========

//66.写函数 WrapNotFound(id uint64) error：
//   使用 fmt.Errorf("user %d: %w", id, ErrNotFound) 包装 ErrNotFound
//   在 Question66 中：对返回值用 errors.Is(..., ErrNotFound) 验证仍能识别
func WrapNotFound(id uint64) error {
	return fmt.Errorf("user %d: %w", id, ErrNotFound)
}
//沿用Q59&60
func Question66() {
	err := WrapNotFound(122)
	fmt.Println("error: ",err)

	if errors.Is(err, ErrNotFound) {
		fmt.Println("函数对返回值仍能识别出 ErrNotFound")
	} else {
		fmt.Println("函数对返回值未能识别出 ErrNotFound")
	}
	fmt.Println()
}

//67.写函数 DescribeValidateError(err error) string：
//   使用 errors.As 判断是否为 *AgeError（或 AgeError）
//   若是：返回含具体 Age 的说明；若不是：返回 err.Error()
//   在 Question67 中分别传入 AgeError 与普通 errors.New 演示
func DescribeValidateError(err error) string {
	var ageerror *AgeError
	if errors.As(err, &ageerror) {
		return fmt.Sprintf("是AgeError错误，错误Age为%d", ageerror.Age)
	} else {
		return err.Error()
	}
}
func Question67() {
	p1 := person{
		Age: -2,
		Name: "小花",
	}
	p2 := person{
		Age: 18,
		Name: "",
	}
	fmt.Println(DescribeValidateError(ValidateUser(p1)))
	fmt.Println(DescribeValidateError(ValidateUser(p2)))
	fmt.Println()
}

//68.写函数 DescribeValue(v any) string，使用 type switch：
//   区分 int / string / User / 其他，返回不同描述字符串
//   在 Question68 中对多种类型各调用一次并打印
func DescribeValue(v any) string {
	switch t := v.(type) {
	case int:
		return fmt.Sprintf("输入为整型: %d", t)
	case string:
		return fmt.Sprintf("输入为字符串: %s", t)
	case User:
		return fmt.Sprintf("输入为User: %v", t)
	case float64:
		return fmt.Sprintf("输入为浮点数: %.4f", t)
	default:
		return fmt.Sprintf("输入为其他类型: %T", t)
	}
}
func Question68() {
	fmt.Println(DescribeValue(10))
	fmt.Println(DescribeValue( "你好啊"))
	fmt.Println(DescribeValue(User{id: 67, name: "徐华",}))
	fmt.Println(DescribeValue(2.1114))
	fmt.Println(DescribeValue(person{Age: 12, Name:"王强"}))
	fmt.Println()
}

//69.写函数 SafeRun(fn func()) (panicMsg string)：
//   在内部用 defer + recover 捕获 fn 中的 panic
//   若发生 panic：返回 panic 信息字符串；否则返回 ""
//   在 Question69 中：一次传入会 panic 的 fn，一次传入正常 fn，分别打印结果
func SafeRun(fn func()) (panicMsg string) { 
	defer func() {
		if err := recover(); err != nil {
			panicMsg = fmt.Sprintf("%v", err)
		}
	}()

	fn()
	panicMsg = "SafeRun运行正常"
	return panicMsg
}
func Question69() {
	msg1 := SafeRun(func() {
		panic("panic happened")
	})
	fmt.Println(msg1)
	msg2 := SafeRun(func() {
		fmt.Println("")
	})
	fmt.Println(msg2)
	fmt.Println()
}

//73.单独的类型断言（不用 type switch）：
//   写函数 AsString(v any) (string, bool)：
//   使用 v.(string) 的「ok 形式」；成功返回 (s, true)，失败返回 ("", false)
//   再写函数 MustAsInt(v any) int：使用无 ok 的 v.(int)；断言失败会 panic
//   在 Question73 中：
//   1) 对 string / int / nil 调用 AsString 并打印
//   2) 对正确 int 调用 MustAsInt；对错误类型用 SafeRun 包一层演示 panic
func AsString(v any) (string, bool) {
	if s, ok := v.(string); ok {
		return s, true
	} else {
		return "", false
	}
}

func MustAsInt(v any) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("发生panic: %s\n", err)
		}
	}()
	return v.(int)
}

func Question73() {
	if val, ok := AsString("小明"); ok {
		fmt.Println(val)
	}
	if val, ok := AsString(1); !ok {
		fmt.Println(val)
	}
	if val, ok := AsString(nil); !ok {
		fmt.Println(val)
	}
	fmt.Println(MustAsInt(2))
	fmt.Println(MustAsInt("2"))
	fmt.Println()
}

//74.空接口 any 与泛型边界直觉：
//   写泛型函数 First[T any](xs []T) (T, bool)：返回首元素；空切片返回零值与 false
//   写非泛型函数 FirstAny(xs []any) (any, bool)：同样语义，但元素类型被擦成 any
//   在 Question74 中对比：
//   1) First([]int{1,2,3}) 得到的是 int，可直接参与算术
//   2) FirstAny([]any{1,2,3}) 得到的是 any，需再断言才能当 int 用
//   打印两边结果，体会：any 灵活但丢类型；泛型保留具体类型
func First[T any](xs []T) (T, bool) {
	var zero T
	if len(xs) == 0 {
		return zero, false
	}
	return xs[0], true
}

func FirstAny(xs []any) (any, bool) {
	if len(xs) == 0 {
		return nil, false
	}
	return xs[0], true
}

func Question74() {
	a, _ := First([]int{1, 2, 3})
	b, _ := FirstAny([]any{1, 2, 3})
	// sum := a + b
	sum := a + b.(int)
	fmt.Println(sum)
}
