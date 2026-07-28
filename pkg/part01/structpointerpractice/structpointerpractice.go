package structpointerpractice

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   uint64
	Name string
	Age  int
}

// 41.定义结构体 User，字段：ID(uint64)、Name(string)、Age(int)
//
//	并实现 String() string，格式形如：User{id=1, name=Tom, age=20}
//	在 Question41 中构造一个 User 并打印
func (user User) String() string {
	return fmt.Sprintf("User{id=%d, name=%s, age=%d}", user.ID, user.Name, user.Age)
}
func Question41() {
	person := User{
		ID:   222510,
		Name: "小明",
		Age:  23,
	}
	fmt.Println(person)
	fmt.Println()
}

// 42.写函数 NewUser(id uint64, name string, age int) User，返回初始化后的 User
//
//	在 Question42 中调用 NewUser 并打印结果
func NewUser(id uint64, name string, age int) User {
	user := User{
		ID:   id,
		Name: name,
		Age:  age,
	}
	return user
}
func Question42(id uint64, name string, age int) {
	fmt.Println(NewUser(id, name, age))
	fmt.Println()
}

// 43.写函数 GrowUpValue(u User, years int) User：值传递，返回 Age 增加 years 后的新 User
//
//	不能修改传入的原变量；在 Question43 中对比修改前后原对象与返回值的 Age
func GrowUpValue(u User, years int) User {
	u.Age = u.Age + years
	return u
}
func Question43() {
	person := User{
		ID:   222510,
		Name: "小明",
		Age:  23,
	}
	fmt.Println(person)
	GrowUpValue(person, 12)
	fmt.Println(person)
	fmt.Println()
}

// 44.写函数 GrowUpPointer(u *User, years int)：指针传递，直接修改原对象的 Age
//
//	在 Question44 中验证调用后原对象 Age 已变化
func GrowUpPointer(u *User, years int) {
	u.Age = u.Age + years
	return
}
func Question44() {
	person := User{
		ID:   222510,
		Name: "小明",
		Age:  23,
	}
	fmt.Println(person)
	GrowUpPointer(&person, 12)
	fmt.Println(person)
	fmt.Println()
}

// 45.给 User 写「值接收者」方法 Birthday()：内部 Age+1
//
//	在 Question45 中调用后打印 Age，观察原对象是否变化
func (u User) Birthday() {
	u.Age = u.Age + 1
}
func Question45() {
	person := User{
		ID:   222510,
		Name: "小明",
		Age:  23,
	}
	fmt.Println(person)
	person.Birthday()
	fmt.Println(person)
	fmt.Println()
}

// 46.给 User 写「指针接收者」方法 BirthdayPtr()：内部 Age+1
//
//	在 Question46 中调用后打印 Age，观察原对象是否变化
func (u *User) BirthdayPtr() {
	u.Age = u.Age + 1
}
func Question46() {
	person := User{
		ID:   222510,
		Name: "小明",
		Age:  23,
	}
	fmt.Println(person)
	person.BirthdayPtr()
	fmt.Println(person)
	fmt.Println()
}

// 47.写函数 SwapUserAge(a, b *User)：交换两个用户的 Age
//
//	在 Question47 中构造两个 User，交换后打印各自 Age
func SwapUserAge(a, b *User) {
	a.Age, b.Age = b.Age, a.Age
}
func Question47() {
	persona := User{
		ID:   222510,
		Name: "小明",
		Age:  23,
	}
	personb := User{
		ID:   222321,
		Name: "小象",
		Age:  12,
	}
	fmt.Println(persona)
	fmt.Println(personb)
	SwapUserAge(&persona, &personb)
	fmt.Printf("交换年龄后: \n")
	fmt.Println(persona)
	fmt.Println(personb)
	fmt.Println()
}

// 48.写函数 CloneUser(u *User) *User：返回一份拷贝（修改拷贝不得影响原对象）
//
//	在 Question48 中修改拷贝的 Age，并对比原对象 Age
func CloneUser(u *User) *User {
	copy := *u
	return &copy
}
func Question48() {
	person := User{
		ID:   222321,
		Name: "小象",
		Age:  12,
	}
	copies := CloneUser(&person)
	copies.Age = copies.Age + 18
	fmt.Println(person)
	fmt.Println(copies)
	fmt.Println()
}

// 49.定义结构体 Account，字段 Balance int
//
//	写指针接收者方法：Deposit(amount int)、Withdraw(amount int) bool
//	余额不足时 Withdraw 返回 false；在 Question49 中演示存取款
type Account struct {
	Balance int
}

func (account Account) String() string {
	return fmt.Sprintf("用户当前存款为%d", account.Balance)
}
func (account *Account) Deposit(amount int) {
	fmt.Println()
	account.Balance = account.Balance + amount
	fmt.Printf("已存款%d元\n", amount)
	fmt.Println(account)
}
func (account *Account) Withdraw(amount int) bool {
	fmt.Println()
	fmt.Println(account)
	if account.Balance < amount {
		fmt.Printf("余额不足，无法取款%d\n", amount)
		return false
	} else {
		account.Balance = account.Balance - amount
		fmt.Printf("余额充足，已取款%d\n", amount)
		fmt.Println(account)
		return true
	}
}
func Question49() {
	paccount := Account{Balance: 200}
	fmt.Println(paccount)
	paccount.Deposit(100)
	paccount.Deposit(225)
	paccount.Withdraw(355)
	paccount.Withdraw(1000)
}

// 50.写函数 Describe(u *User)：若 u == nil 打印 "nil user"，否则打印 u.String()
//
//	在 Question50 中分别传入 nil 和非 nil 进行演示
func Describe(u *User) string {
	if u == nil {
		return fmt.Sprintf("nil user")
	} else {
		return u.String()
	}
}
func Question50() {
	var persona User
	personb := User{
		ID:   222510,
		Name: "小明",
		Age:  23,
	}
	fmt.Println(Describe(nil))
	fmt.Println(Describe(&persona))
	fmt.Println(Describe(&personb))
	fmt.Println()
}

// 51.为 Account 再写「值接收者」方法 DepositByValue(amount int)
//
//	在 Question51 中：初始余额 100，调用 DepositByValue(50) 后打印余额
//	预期：余额仍为 100（证明值接收者无法更新原对象）
//	建议另写单元测试：调用前后 Balance 相等
func (account Account) DepositByValue(amount int) {
	account.Balance = account.Balance + amount
	fmt.Printf("尝试存款%d元\n", amount)
}
func Question51() {
	paccount := Account{Balance: 100}
	fmt.Println(paccount)
	paccount.DepositByValue(100)
	fmt.Println(paccount)
	fmt.Println()
}

// 52.扩展（或新建）带切片字段的结构体，例如：
//
//	type Bill struct { Owner string; Items []int }
//	在 Question52 中：
//	1) 构造 bill，Items 为若干元素
//	2) 浅拷贝：copyBill := bill（或 *ptr 解引用赋值）
//	3) 修改 copyBill.Items[0]
//	4) 打印原 bill.Items[0]，观察是否被连带修改（浅拷贝共享底层数组）
type Bill struct {
	Owner string
	Items []int
}

func (bill Bill) String() string {
	return fmt.Sprintf("%s的账单为%v", bill.Owner, bill.Items)
}
func Question52() {
	bill := Bill{
		Owner: "小明",
		Items: []int{23, 421, 21, 48},
	}
	fmt.Println(bill)
	copyBill := bill
	copyBill.Items[0], copyBill.Items[3] = 11, 87
	fmt.Println(bill)
	fmt.Println()
	//原对象切片被修改，结构体中有引用类型的字段会被浅拷贝变量影响
}

// 53.为上一题的 Bill（或同等结构）手写 DeepCopyBill(b Bill) Bill
//
//	要求：拷贝后修改新对象的 Items 元素，不得影响原对象
//	在 Question53 中演示对比浅拷贝与深拷贝的差异
func DeepCopyBill(b Bill) Bill {
	copiedbill := Bill{
		Owner: b.Owner,
		Items: make([]int, len(b.Items), cap(b.Items)),
	}
	copy(copiedbill.Items, b.Items)
	return copiedbill
}
func Question53() {
	bill := Bill{
		Owner: "小明",
		Items: []int{23, 421, 21, 48},
	}
	copybill := DeepCopyBill(bill)
	copybill.Items[1] = 1111
	fmt.Println(bill)
	fmt.Println(copybill)
	fmt.Println()
	//引用类型切忌用=拷贝
}

// 54.定义嵌套结构体并加 json tag，例如：
//
//	Address{City, Street string `json:"..."`}
//	Profile{Name string; Addr Address `json:"addr"`}
//	在 Question54 中：构造 Profile → encoding/json Marshal 打印 JSON
//	再把该 JSON Unmarshal 回结构体并打印，完成一次往返
type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}
type Profile struct {
	Name string
	Addr Address `json:"addr"`
}

func Question54() {
	profile := Profile{
		Name: "小明",
		Addr: Address{
			City:   "北京",
			Street: "北京路1号",
		},
	}
	jsondata, err := json.Marshal(profile)
	fmt.Println(err)
	fmt.Println(string(jsondata))
	var profiledata Profile
	json.Unmarshal(jsondata, &profiledata)
	fmt.Println(profiledata)
	fmt.Println()
}

// 55.用 map[string]struct{} 实现「文章标签集合」
//
//	写函数：AddTag / RemoveTag / HasTag / ListTags
//	在 Question55 中演示：添加重复标签应去重、删除、查询、列出全部标签
func AddTag(m map[string]struct{}, tag string) {
	m[tag] = struct{}{}
}
func RemoveTag(m map[string]struct{}, tag string) {
	delete(m, tag)
}
func HasTag(m map[string]struct{}, tag string) bool {
	_, ok := m[tag]
	if !ok {
		return false
	} else {
		return true
	}
}
func ListTags(m map[string]struct{}) {
	fmt.Printf("现有标签：\t")
	for key, _ := range m {
		fmt.Printf("%s\t", key)
	}
	fmt.Println()
}

func Question55() {
	Tags := map[string]struct{}{}
	AddTag(Tags, "a")
	ListTags(Tags)
	AddTag(Tags, "a")
	ListTags(Tags)
	AddTag(Tags, "b")
	ListTags(Tags)
	if HasTag(Tags, "b") {
		RemoveTag(Tags, "b")
	}
	ListTags(Tags)
}

// 56.用 iota 定义订单状态枚举，例如：
//
//	type OrderStatus int
//	const ( OrderPending OrderStatus = iota; OrderPaid; OrderShipped; OrderDone )
//	为 OrderStatus 实现 String() string
//	在 Question56 中打印各状态的数值与字符串
type OrderStatus int

const (
	OrderPending OrderStatus = iota
	OrderPaid
	OrderShipped
	OrderDone
)

func (orderstatus OrderStatus) String() string {
	switch orderstatus {
	case OrderPending:
		return fmt.Sprintf("OrderPending: %d", orderstatus)
	case OrderPaid:
		return fmt.Sprintf("OrderPaid: %d", orderstatus)
	case OrderShipped:
		return fmt.Sprintf("OrderShipped: %d", orderstatus)
	case OrderDone:
		return fmt.Sprintf("OrderDone: %d", orderstatus)
	default:
		return ""
	}
}
func Question56() {
	for i := 1; i <= 4; i++ {
		fmt.Println(OrderStatus(i))
	}

}

// 70.匿名嵌入：复用上面的 Address，定义 PersonProfile
//
//	PersonProfile 匿名嵌入 Address，并自有字段 Name string
//	（对比 Q54 的 Profile：那里是「具名字段 Addr Address」，这里是匿名嵌入）
//	写函数 DescribePerson(p PersonProfile) string：通过提升字段直接读 City/Street
//	在 Question70 中构造实例、打印 DescribePerson 结果，并打印 p.Address.City 对比
type PersonProfile struct {
	Address
	Name string
}

func DescribePerson(p PersonProfile) string {
	return fmt.Sprintf("姓名: %s, 城市: %s, 街道: %s", p.Name, p.City, p.Street)
}
func Question70() {
	information := PersonProfile{
		Address: Address{
			City:   "上海",
			Street: "不知道什么路",
		},
		Name: "肖华",
	}
	fmt.Println(DescribePerson(information))
	fmt.Println()
}

// 71.nil map 写入会 panic：
//
//	写函数 SafeSet(m map[string]int, key string, val int) (ok bool)
//	- 若 m == nil：不写入，返回 false（不要 panic）
//	- 否则写入并返回 true
//	在 Question71 中：
//	1) 对 var m map[string]int 直接 m["a"]=1，用 defer+recover 捕获并打印 panic
//	2) 用 SafeSet(nil, ...) 演示安全失败
//	3) 用 make 后的 map 调用 SafeSet，打印成功结果
func SafeSet(m map[string]int, key string, val int) (ok bool) {
	if m == nil {
		return false
	} else {
		return true
	}
}

func Question71() {
	n := make(map[string]int)
	if SafeSet(n, "c", 2) {
		fmt.Println("make创建映射表运行成功")
	}
	defer SafeSet(nil, "a", 1)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var m map[string]int
	m["a"] = 1

}

// 72.切片截取与「内存泄漏」注意点：
//
//	写函数 KeepFirst3(src []byte) []byte：
//	- 错误示范思路：直接 return src[:3]（小切片仍引用整段底层大数组，大数组无法被 GC）
//	- 正确做法：copy 到新切片再返回，断开对原底层数组的引用
//	在 Question72 中：
//	1) 构造较大 []byte（例如 1<<20 长度），截取前 3 字节的「错误版」与「正确版」
//	2) 打印两种结果的 len/cap，说明为何错误版 cap 仍很大
func KeepFirst3(src []byte) []byte {
	// n := 3
	// if len(src) < 3 {
	// 	n = len(src)
	// }
	// cloneslice := make([]byte, n)
	// copy(cloneslice, src)
	// return cloneslice
	return src[:3]
}

func Question72() {
	var a []byte = []byte{1, 2, 3, 4, 5, 2}
	var b []byte = []byte{10, 2, 4}
	c := KeepFirst3(a)
	c[0] = 10
	fmt.Println(a)
	fmt.Println(c)
	d := KeepFirst3(b)
	d[1] = 3
	fmt.Println(b)
	fmt.Println(d)
	e := a
	fmt.Println(cap(a))
	fmt.Println(cap(e))
}
