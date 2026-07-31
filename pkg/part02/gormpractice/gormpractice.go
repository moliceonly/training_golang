// Package gormpractice 对齐路线图 Part 02 · 2.4「数据库（MySQL + GORM）」。
// 每个知识要点至少一题；题干尽量贴近图书馆模拟场景。
// 题号：114 → 121
//
// 查看题干：
//
//	go doc training_golang/pkg/part02/gormpractice.Question117
//
// 依赖本机 apt 安装的 MySQL（训练统一不用 Docker 跑库）。
// 练习账号：trainer / Train2026Lib!（库 training_lib）
// 默认用 mysql.Config 生成 DSN，避免密码特殊字符手写编码出错。
// 也可被环境变量 TRAINING_MYSQL_DSN 覆盖。
//
// 未配置或连不上时，Question 演示可打印提示并 return；单测可设 TRAINING_SKIP_MYSQL=1。
package gormpractice

import (
	"fmt"
	"os"
	"time"
	"errors"
	"database/sql"

	mysqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/gorm/clause"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 练习库账号（与本机 MySQL 中 trainer 一致）
const (
	dbUser = "trainer"
	dbPass = "Train2026Lib!"
	dbName = "training_lib"
	dbAddr = "127.0.0.1:3306"
)

// DefaultDSN 由 Config 生成，含正确转义。
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

// ---------------------------------------------------------------------------
// 114. GORM 连接 MySQL：DSN、gorm.Open、连接池
//
// Question114 练习打开连接并配置池参数。
//
// 写函数：
//   OpenDB(dsn string) (*gorm.DB, error)
//     - gorm.Open(mysql.Open(dsn), &gorm.Config{})
//     - 取出 sql.DB：SetMaxOpenConns(10)、SetMaxIdleConns(5)、SetConnMaxLifetime(time.Hour)
//     - Ping 失败则返回 error
// 在 Question114 中：OpenDB(dsnFromEnv())，成功打印 "db ok"，失败打印 err
func OpenDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Question114 演示连接与连接池配置。
func Question114() {
	db, err := OpenDB(dsnFromEnv())
	if err != nil {
		fmt.Println("db err:", err)
		return
	}
	_ = db
	fmt.Println("db ok")
}

// ---------------------------------------------------------------------------
// 115. Model：gorm.Model、标签、主键/索引；迁移 AutoMigrate
// 场景：图书馆 — 图书 / 读者表
//
// Question115 定义 Model 并 AutoMigrate（练习环境可用；生产慎用）。
//
// 定义（可微调字段，但需能支撑后续借还书）：
//   type Book struct {
//     gorm.Model
//     ISBN   string `gorm:"size:32;uniqueIndex"`
//     Title  string `gorm:"size:128;index"`
//     Stock  int
//     Version int  `gorm:"default:0"` // 乐观锁用，可先占位
//   }
//   type Reader struct {
//     gorm.Model
//     Name string `gorm:"size:64;index"`
//   }
// 写函数：
//   AutoMigrateLibrary(db *gorm.DB) error
//     - db.AutoMigrate(&Book{}, &Reader{}, &BorrowRecord{})
//   BorrowRecord 可在 117/120 再完善；此处可先定义空壳或一并写好：
//     BookID, ReaderID uint；BorrowedAt time.Time；Returned bool

type Book struct {
	gorm.Model
	ISBN    string `gorm:"size:32;uniqueIndex"`
	Title   string `gorm:"size:128;index"`
	Stock   int
	Version int    `gorm:"default:0"`
}

type Reader struct {
	gorm.Model
	Name string `gorm:"size:64;index"`
}

type BorrowRecord struct {
	gorm.Model
	BookID     uint
	ReaderID   uint
	Book       Book      `gorm:"foreignKey:BookID"`
	Reader     Reader    `gorm:"foreignKey:ReaderID"`
	BorrowedAt time.Time
	Returned   bool
}

func AutoMigrateLibrary(db *gorm.DB) error {
	return db.AutoMigrate(&Book{}, &Reader{}, &BorrowRecord{})
}

// Question115 演示 AutoMigrate。
func Question115() {
	db, err := OpenDB(dsnFromEnv())
	if err != nil {
		fmt.Println("db err:", err)
		return
	}
	if err := AutoMigrateLibrary(db); err != nil {
		fmt.Println("migrate:", err)
		return
	}
	fmt.Println("migrate ok")
	fmt.Println(0)
}

// ---------------------------------------------------------------------------
// 116. CRUD + 条件分页
// 场景：图书增删改查与分页列表
//
// Question116 练习 Create / First / Find / Updates / Delete 与 Offset/Limit。
//
// 写函数：
//   CreateBook(db, isbn, title string, stock int) (*Book, error)
//   GetBookByISBN(db, isbn string) (*Book, error)
//   ListBooks(db, titleLike string, page, pageSize int) (books []Book, total int64, err error)
//     - titleLike 非空时 Where("title LIKE ?", "%"+titleLike+"%")
//     - Count + Offset((page-1)*pageSize).Limit(pageSize).Find
//   UpdateBookStock(db, id uint, stock int) error
//   SoftDeleteBook(db, id uint) error  // Delete，依赖 gorm.Model 软删
// 在 Question116 中串一次增查改删并打印
func CreateBook(db *gorm.DB, isbn, title string, stock int) (*Book, error) {
	// TODO
	book := &Book{
		ISBN: isbn,
		Title: title,
		Stock: stock,
	}
	if err := db.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func GetBookByISBN(db *gorm.DB, isbn string) (*Book, error) {
	// TODO
	var book Book
	if err := db.Where("isbn = ?", isbn).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func ListBooks(db *gorm.DB, titleLike string, page, pageSize int) (books []Book, total int64, err error) {
	// TODO
	q := db.Model(&Book{})
	if titleLike != "" {
		q = q.Where("title LIKE ?", "%"+titleLike+"%")
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	if err := q.Offset(offset).Limit(pageSize).Find(&books).Error; err != nil {
		return nil, 0, err
	}
	return books, total, nil
}

func UpdateBookStock(db *gorm.DB, id uint, stock int) error {
	// TODO
	return db.Model(&Book{}).Where("id = ?", id).Update("stock", stock).Error
}

func SoftDeleteBook(db *gorm.DB, id uint) error {
	// TODO
	return db.Where("id = ?", id).Delete(&Book{}).Error
}

// Question116 演示图书 CRUD 与分页。
func Question116() {
	db, err := OpenDB(dsnFromEnv())

	if err != nil {
		fmt.Println("db err:", err)
		return
	}
	_ = AutoMigrateLibrary(db)

	db.Unscoped().Where("isbn IN ?", []string{"abc-123", "abc-456"}).Delete(&Book{})

	book1, err := CreateBook(db, "abc-123", "Alphabet", 5)
	if err != nil {
		fmt.Println("create1:", err)
		return
	}
	defer db.Unscoped().Delete(&Book{}, book1.ID)
	fmt.Println("create1:", book1.ID, book1.ISBN, book1.Stock)

	book2, err := CreateBook(db, "abc-456", "Alphabet2", 2)
	if err != nil {
		fmt.Println("create2:", err)
		return
	}
	defer db.Unscoped().Delete(&Book{}, book2.ID)
	fmt.Println("create2:", book2.ID, book2.ISBN, book2.Stock)

	got, err := GetBookByISBN(db, "abc-123")
	if err != nil {
		fmt.Println("get:", err)
		return
	}
	fmt.Println("get:", got.Title, got.Stock)
	if err := UpdateBookStock(db, book1.ID, 3); err != nil {
		fmt.Println("update:", err)
		return
	}
	got, _ = GetBookByISBN(db, "abc-123")
	fmt.Println("after update stock:", got.Stock)

	books, total, err := ListBooks(db, "Alpha", 1, 10)
	if err != nil {
		fmt.Println("list:", err)
		return
	}
	fmt.Println("list total:", total, "page len:", len(books))

	if err := SoftDeleteBook(db, book2.ID); err != nil {
		fmt.Println("delete:", err)
		return
	}
	_, err = GetBookByISBN(db, "abc-456")
	fmt.Println("after soft delete get err:", err)
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 117. 事务：借书 — 扣库存 + 写借阅流水
// 场景：db.Transaction；库存不足或写流水失败则全部回滚
//
// Question117 实现事务借书。
//
// 写函数：
//   BorrowBookTx(db *gorm.DB, bookID, readerID uint) error
//     - db.Transaction(func(tx *gorm.DB) error {
//         查书；Stock<=0 则返回错误
//         Stock-- ；Updates
//         创建 BorrowRecord{BookID, ReaderID, BorrowedAt: time.Now()}
//         任一步失败 return err → 自动回滚
//       })
// 在 Question117 中准备一本库存 1 的书、一个读者，借两次：第二次应失败且库存仍为 0
func BorrowBookTx(db *gorm.DB, bookID, readerID uint) error {
	// TODO
	return db.Transaction(func(tx *gorm.DB) error {
		var book Book
		tx.Where("id = ?", bookID).First(&book)
		if book.Stock <= 0 {
			return errors.New("Book's stock zero")
		}

		if err := UpdateBookStock(tx, bookID, book.Stock-1); err != nil {
			return err 
		}

		borrowRecord := &BorrowRecord{
			BookID: bookID,
			ReaderID: readerID,
			BorrowedAt: time.Now(),
			Returned: false,
		}

		if err := tx.Create(borrowRecord).Error; err != nil {
			return err
		}
		return nil
	})
}

// Question117 演示事务借书与回滚。
func Question117() {

	db, err := OpenDB(dsnFromEnv())

	if err != nil {
		fmt.Println(err)
		return
	}
	_ = AutoMigrateLibrary(db)

	isbn := fmt.Sprintf("borrow-%d", time.Now().UnixNano())
	book, err := CreateBook(db, isbn, "Borrow Demo", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Unscoped().Delete(&Book{}, book.ID)
	fmt.Println(book)

	reader := &Reader{Name: "alice"}
	if err := db.Create(reader).Error; err != nil {
		fmt.Println(err)
		return
	}
	defer db.Unscoped().Delete(&Reader{}, reader.ID)
	
	if err := BorrowBookTx(db, book.ID, reader.ID); err != nil {
		fmt.Println(err)
		return
	}
	b1, _ := GetBookByISBN(db, isbn)
	fmt.Println("after borrow1 stock:", b1.Stock) 
	
	err = BorrowBookTx(db, book.ID, reader.ID)
	fmt.Println(err) 
	b2, _ := GetBookByISBN(db, isbn)
	fmt.Println("after borrow2 stock:", b2.Stock) 
	db.First(&book, book.ID)
	fmt.Println(book)
	fmt.Println()

}

// ---------------------------------------------------------------------------
// 118. 悲观锁：SELECT ... FOR UPDATE 再改库存
// 场景：事务内锁行再扣库存
//
// Question118 练习 Clauses(clause.Locking{Strength: "UPDATE"})。
//
// 写函数：
//   DeductStockPessimistic(db *gorm.DB, bookID uint, n int) error
//     - Transaction 内：
//       tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&book, bookID)
//       库存不足则 err；否则 Stock -= n；Save/Updates
// 在 Question118 中对一本书扣 1 本并打印剩余库存
func DeductStockPessimistic(db *gorm.DB, bookID uint, n int) error {
	// TODO
	return db.Transaction(func(tx *gorm.DB) error {
		var book Book
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&book, bookID).Error; err != nil {
			return err
		}
		left := book.Stock - n
		if left < 0 {
			return errors.New("Not enough stock")
		}
		if err := UpdateBookStock(tx, bookID, left); err != nil {
			return err
		}
		return nil
	})
}

// Question118 演示悲观锁扣库存。
func Question118() {
	db, err := OpenDB(dsnFromEnv())

	if err != nil {
		fmt.Println(err)
		return
	}
	_ = AutoMigrateLibrary(db)

	isbn := fmt.Sprintf("borrow-%d", time.Now().UnixNano())
	book, err := CreateBook(db, isbn, "Borrow Demo", 78)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Unscoped().Delete(&Book{}, book.ID)
	fmt.Println(book)

	if err := DeductStockPessimistic(db, book.ID, 2); err != nil {
		fmt.Println(err)
	}
	db.First(&book, book.ID)
	fmt.Println(book)
	fmt.Println()
}

// ---------------------------------------------------------------------------
// 119. 乐观锁：version 字段；更新失败则重试或返回错误
// 场景：Model 带 version；WHERE version=? ，RowsAffected=0 视为冲突
//
// Question119 练习乐观锁扣库存。
//
// 写函数：
//   DeductStockOptimistic(db *gorm.DB, bookID uint, n int, maxRetry int) error
//     - 循环最多 maxRetry 次：
//       First 取书；库存不足 return err
//       result := db.Model(&Book{}).Where("id = ? AND version = ?", book.ID, book.Version).
//         Updates(map[string]any{"stock": book.Stock - n, "version": book.Version + 1})
//       若 result.RowsAffected == 1 成功 return nil；否则重试
//       用尽重试则返回冲突错误
// 在 Question119 中扣库存并打印；可注释对比悲观锁：重试成本 vs 锁等待/死锁
func DeductStockOptimistic(db *gorm.DB, bookID uint, n int, maxRetry int) error {
	// TODO

	for i := 0; i < maxRetry; i++ {
		var book Book
		if err := db.First(&book, bookID).Error; err != nil {
			return err
		}

		if book.Stock < n {
			return fmt.Errorf("stock not enough: have %d, need %d", book.Stock, n)
		}

		result := db.Model(&Book{}).
			Where("id = ? AND version = ?", book.ID, book.Version).
			Updates(map[string]any{
				"stock":   book.Stock - n,
				"version": book.Version + 1,
			})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 1 {
			return nil 
		}
		
	}

	return fmt.Errorf("optimistic lock conflict after %d retries", maxRetry)
}

// Question119 演示乐观锁扣库存。
func Question119() {

	db, err := OpenDB(dsnFromEnv())
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = AutoMigrateLibrary(db)

	isbn := fmt.Sprintf("borrow-%d", time.Now().UnixNano())
	book, err := CreateBook(db, isbn, "Borrow Demo", 78)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Unscoped().Delete(&Book{}, book.ID)

	if err := DeductStockOptimistic(db, book.ID, 1, 3); err != nil {
		fmt.Println(err)
		return
	}
	db.First(&book, book.ID)
	fmt.Println(book)
	fmt.Println()

}

// ---------------------------------------------------------------------------
// 120. 关联 + Preload：借阅记录查出「书 + 读者」
// 场景：Belongs To；Preload 避免 N+1
//
// Question120 完善 BorrowRecord 关联并 Preload 查询。
//
// BorrowRecord：
//   BookID, ReaderID uint
//   Book Book `gorm:"foreignKey:BookID"`
//   Reader Reader `gorm:"foreignKey:ReaderID"`
//   BorrowedAt time.Time
//   Returned bool
// 写函数：
//   ListBorrowDetails(db *gorm.DB, limit int) ([]BorrowRecord, error)
//     - db.Preload("Book").Preload("Reader").Limit(limit).Find(&rows)
// 在 Question120 中打印每条：书名、读者名、借出时间
func ListBorrowDetails(db *gorm.DB, limit int) ([]BorrowRecord, error) {
	// TODO
	var rows []BorrowRecord
	err := db.Preload("Book").Preload("Reader").
		Order("id desc").
		Limit(limit).
		Find(&rows).Error
	return rows, err
}

// Question120 演示 Preload 关联查询。
func Question120() {

	db, err := OpenDB(dsnFromEnv())
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = AutoMigrateLibrary(db)

	isbn := fmt.Sprintf("abc-%d", time.Now().UnixNano())
	book, err := CreateBook(db, isbn, "abc-123", 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Unscoped().Delete(&Book{}, book.ID)

	reader := &Reader{Name: "小明"}
	if err := db.Create(reader).Error; err != nil {
		fmt.Println(err)
		return
	}
	defer db.Unscoped().Delete(&Reader{}, reader.ID)

	if err := BorrowBookTx(db, book.ID, reader.ID); err != nil {
		fmt.Println("borrow:", err)
		return
	}

	rows, err := ListBorrowDetails(db, 5)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range rows {
		fmt.Println(r.Book.Title, r.Reader.Name, r.BorrowedAt)
	}

}

// ---------------------------------------------------------------------------
// 121. SQL 直觉 + 安全：EXPLAIN；参数化；慎用生产 AutoMigrate
//
// Question121 练习一眼 EXPLAIN。
//
// 写函数：
//   ExplainBookByISBN(db *gorm.DB, isbn string) (plan string, err error)
//     - 可用 db.Raw("EXPLAIN SELECT * FROM books WHERE isbn = ?", isbn).Scan / Rows
//       把关键列拼成字符串返回（实现方式不限，能跑通即可）
// 排障注意点:
//   1. 用 ? 占位符传参，不要把用户输入拼进 SQL（防注入）
//   2. 慢查询看 slow_query_log / EXPLAIN，确认是否走索引（如 isbn）
//   3. 关联用 Preload/Joins，避免循环里逐条查造成 N+1
//   4. AutoMigrate 适合练习/开发；生产用版本化迁移（golang-migrate 等），先审再上
// 在 Question121 中打印 plan 与 SafetyNotes
func ExplainBookByISBN(db *gorm.DB, isbn string) (plan string, err error) {
	// TODO
	rows, err := db.Raw("EXPLAIN SELECT * FROM books WHERE isbn = ?", isbn).Rows()
	if err != nil {
		return "", err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return "", err
	}
	var out string
	for rows.Next() {
		vals := make([]sql.NullString, len(cols))
		ptrs := make([]any, len(cols))
		for i := range vals {
			ptrs[i] = &vals[i]
		}
		if err := rows.Scan(ptrs...); err != nil {
			return "", err
		}
		for i, c := range cols {
			if i > 0 {
				out += " | "
			}
			v := "NULL"
			if vals[i].Valid {
				v = vals[i].String
			}
			out += c + "=" + v
		}
		out += "\n"
	}
	return out, rows.Err()
}

func SafetyNotes() string {
	return `1. 用 ? 占位符传参，不要把用户输入拼进 SQL（防注入）
2. 慢查询看 slow_query_log / EXPLAIN，确认是否走索引（如 isbn）
3. 关联用 Preload/Joins，避免循环里逐条查造成 N+1
4. AutoMigrate 适合练习/开发；生产用版本化迁移，先审再上`
}

// Question121 演示 EXPLAIN。
func Question121() {

	db, err := OpenDB(dsnFromEnv())
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = AutoMigrateLibrary(db)

	isbn := "demo-isbn"
	book, err := CreateBook(db, isbn, "abc-123", 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Unscoped().Delete(&Book{}, book.ID)

	plan, err := ExplainBookByISBN(db, "demo-isbn")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(plan)
	fmt.Println(SafetyNotes())
}
