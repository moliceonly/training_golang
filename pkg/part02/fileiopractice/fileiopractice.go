// Package fileiopractice 对齐路线图 Part 02 · 2.1「文件与本地 IO」。
// 每个知识要点至少一题；题干尽量贴近参考模拟场景。
// 题号：94 → 101
//
// 查看题干示例：
//
//	go doc training_golang/pkg/part02/fileiopractice.Question94
package fileiopractice

import (
	"io"
	"os"
	"fmt"
	"path/filepath"
	"errors"
	"time"
	"strings"
	"bufio"
	"io/fs"
	// "encoding/csv"
)

func EnsureDir(path string) error {
	// TODO: os.MkdirAll，已存在不报错
	return os.MkdirAll(path, 0o777)
}

func WriteHelloFile(dir, name string) (fullPath string, err error) {
	// TODO: filepath.Join + os.Create 写 "hello\n"，defer Close
	fullPath = filepath.Join(dir, name)
	file, createErr := os.Create(fullPath)
	if createErr != nil {
		return "", createErr
	}
	defer file.Close()
	_, writeErr := file.WriteString("hello\n")
	err = errors.Join(createErr, writeErr)
	return fullPath, err
}

func RemoveFile(path string) error {
	// TODO: os.Remove
	return os.Remove(path)
}

func DemoEnvAndWD() (wd string, home string, err error) {
	// TODO: os.Getwd；os.Getenv("HOME")
	wd, err = os.Getwd()
	return wd, os.Getenv("HOME"), err
}

// Question94 练习 os：打开/创建/删除；环境变量与工作目录。
//
// 实现 EnsureDir / WriteHelloFile / RemoveFile / DemoEnvAndWD，然后在本函数中：
//   - 在 tmpDir 下确保子目录、写入 hello 文件、再删除
//   - 打印工作目录与 HOME
func Question94(tmpDir string) {
	dir, home, err :=DemoEnvAndWD()
	fmt.Println(err)
	err = EnsureDir(tmpDir)
	
	fmt.Println(err)
	filePath, err := WriteHelloFile(tmpDir, "hello")
	fmt.Println(err)
	time.After( 2 * time.Second)
	err = RemoveFile(filePath)
	fmt.Println(dir, home, err)
	fmt.Println()
}

func SafeJoin(root, userPath string) (string, error) {
	// TODO: Clean(Join)；Rel 后拒绝 ".." 逃逸
	full := filepath.Clean(filepath.Join(root, userPath))
	rel, err := filepath.Rel(root, full)
	if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", errors.New("path escapes root")
	}
	return full, err
}

// Question95 练习 path/filepath：拼接与清理，防止上传路径 .. 穿越。
//
// 实现 SafeJoin(root, userPath)：
//   - full := filepath.Clean(filepath.Join(root, userPath))
//   - rel, err := filepath.Rel(root, full)
//   - 若 rel==".." 或以 ".."+分隔符开头，返回 error
//
// 本函数中：对 "a/b.txt" 应成功；对 "../../etc/passwd" 应失败并打印
func Question95(tmpDir string) {
	p1, err1 := SafeJoin(tmpDir, "a/b.txt")
	fmt.Println("ok:", p1, err1)
	p2, err2 := SafeJoin(tmpDir, "../../etc/passwd")
	fmt.Println("bad:", p2, err2)
	fmt.Println()
}

func CopyFile(dst, src string) (written int64, err error) {
	// TODO: Open/Create + io.Copy，defer Close
	copyFile, _ := os.Create(dst)
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, errors.New("source file not exists")
	}
	return io.Copy(copyFile, srcFile)
}

func CountBytes(r io.Reader) (n int64, err error) {
	// TODO: io.Copy(io.Discard, r)
	return io.Copy(io.Discard, r)
}

// Question96 练习 io：Reader/Writer、Copy、接口组合。
//
// 实现 CopyFile、CountBytes；本函数中写小文件、复制、再统计字节数并打印。
func Question96(tmpDir string) {
	_ = EnsureDir(tmpDir)
	full, _ := WriteHelloFile(tmpDir, "a.txt")
	copyFilePath := filepath.Join(tmpDir, "b.txt")
	writtenNum, err := CopyFile(copyFilePath, full)
	fmt.Println(writtenNum, err)
	copyFile, _ := os.Open(copyFilePath)
	defer copyFile.Close()
	copyFileWC, err := CountBytes(copyFile)
	fmt.Println(copyFileWC, err)
	fmt.Println()
}

func WriteLines(path string, lines []string) error {
	// TODO: bufio.NewWriter 按行写并 Flush
	writtenFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	defer writtenFile.Close()
	bw := bufio.NewWriter(writtenFile)
	for _, s := range lines {
		bw.WriteString(s)
	} 
	if err := bw.Flush(); err != nil {
		return err
	}
	return nil
}

func ReadLines(path string) ([]string, error) {
	// TODO: bufio.Scanner 按行读
	readLines := []string{}
	readFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()
	ns := bufio.NewScanner(readFile)
	for ns.Scan(){
		readLines = append(readLines, ns.Text())
	}
	if err := ns.Err(); err != nil {
		return nil, err
	}
	return readLines, nil
}

// Question97 练习 bufio：按行读、缓冲写（访问日志一行一条）。
//
// 实现 WriteLines / ReadLines；本函数写入若干行再读回打印。
func Question97(tmpDir string) {
	lines := []string{" Question97 练习 bufio：按行读、缓冲写（访问日志一行一条）",
	 "实现 WriteLines / ReadLines；本函数写入若干行再读回打印。",
	 "// TODO: bufio.Scanner 按行读",
	}
	filePath := filepath.Join(tmpDir, "b.txt")
	fmt.Println(WriteLines(filePath, lines))
	readLines, err := ReadLines(filePath)
	fmt.Println(readLines, err)
	fmt.Println()
}

func ImportProductsCSV(path string, handle func(id, name string) error) (rows int, err error) {
	// TODO: Open+Scanner 流式读；跳过表头；Split 后调 handle。禁止 ReadFile 整文件
	readFile, err := os.Open(path)
	if err != nil {
		return 0, nil
	}
	defer readFile.Close()

	ns := bufio.NewScanner(readFile)
	for ns.Scan() {
		parts := strings.SplitN(ns.Text(), ",", 2)
		if len(parts) < 2 {
			continue
		}
		if err := handle(parts[0], parts[1]); err != nil {
			return rows, err
		}
		rows++
	}
	return rows, ns.Err()
}

func GenerateDemoCSV(path string, n int) error {
	// TODO: 流式写表头 + n 行 "i,product-i"

	writeFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	defer writeFile.Close()

	bw := bufio.NewWriter(writeFile)
	bw.WriteString("id,name\n")
	for i := 0; i < n; i++ {
		fmt.Fprintf(bw, "%d,product-%d\n", i, i)
	}
	return bw.Flush()
}

// Question98 运营导表：流式导入 CSV，不能一次性读进内存。
//
// 场景：商品 CSV（练习可用 1000 行代替 10 万行）。
// 实现 GenerateDemoCSV、ImportProductsCSV；本函数生成后导入并打印行数。
func Question98(tmpDir string) {
	mkFile := filepath.Join(tmpDir, "test.csv")
	fmt.Println(GenerateDemoCSV(mkFile, 100000))
	rows, err := ImportProductsCSV(mkFile, func(id, name string) error {
		fmt.Println(id, name)
		return nil
	}) 
	fmt.Println(rows, err)
	fmt.Println()
}

// RotatingWriter 按大小滚动的日志写入器。
type RotatingWriter struct {
	// TODO: path, maxBytes, f *os.File, size int64
	path string
	maxBytes int64
	f *os.File
	size int64
}

func OpenRotating(path string, maxBytes int64) (*RotatingWriter, error) {
	// TODO

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return nil, nil
	}
	defer file.Close()

	rWriter := RotatingWriter{
		path: path,
		maxBytes: maxBytes,
		f: file,
		size: 0,
	}

	for i := int64(0); i < maxBytes + maxBytes; i++ {
		rWriter.Write([]byte("i"))
		rWriter.size++
		if i >=199 {
			fmt.Println(rWriter)
		}
		if i >=210 {
			break
		}
	}

	return &rWriter, nil
}

func (w *RotatingWriter) Write(p []byte) (int, error) {
	// TODO: 超限则 Rename 为 path+".1" 后新建

	if w.size == 200 {
		w.Close()
	}

	return 	w.f.Write(p)
}

func (w *RotatingWriter) Close() error {
	
	w.path = w.path + ".1"
	w.size = 0

	return w.f.Close()
}

// Question99 服务日志滚动：单文件超过阈值自动切分。
//
// 场景：避免日志撑满磁盘（练习 maxBytes 可用 200 字节，不必真 100MB）。
// 实现 RotatingWriter；本函数循环写入，观察是否生成 .1 文件。
func Question99(tmpDir string) {
	file := filepath.Join(tmpDir, "U")
	OpenRotating(file, 200)
	fmt.Println()
}

func ScanUploadDir(root string) (byExt map[string]int64, largest string, largestSize int64, err error) {
	// TODO: filepath.WalkDir 统计后缀占用与最大文件
	byExt = make(map[string]int64)
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		size := info.Size()
		ext := filepath.Ext(path)

		byExt[ext] += size
		if size > largestSize {
			largestSize = size
			largest = path
		}
		return nil
	})
	return byExt, largest, largestSize, err
}

// Question100 磁盘巡检：扫描上传目录，统计后缀占用，找出异常大文件。
//
// 实现 ScanUploadDir；本函数在 tmpDir 建不同后缀文件后打印统计。
func Question100(tmpDir string) {
	_ = os.WriteFile(filepath.Join(tmpDir, "a.txt"), []byte("hi"), 0o644)
	_ = os.WriteFile(filepath.Join(tmpDir, "b.jpg"), []byte("1234567890"), 0o644)
	_ = os.WriteFile(filepath.Join(tmpDir, "c.txt"), []byte("hello world"), 0o644)
	byExt, largest, largestSize, err := ScanUploadDir(tmpDir)
	fmt.Println("byExt:", byExt)
	fmt.Println("largest:", largest, "size:", largestSize)
	fmt.Println("err:", err)
	fmt.Println()

}

func WriteTempAndRead(content string) (read string, err error) {
	// TODO: CreateTemp；defer Remove；注意 Close 错误与命名返回值
	tmpFile, err := os.CreateTemp("", "Test-*.txt")
	if err != nil {
		return "", nil
	}
	name := tmpFile.Name()
	defer os.Remove(name)

	tmpFile.Write([]byte(content))
	if err = tmpFile.Close(); err != nil {
		return "", err
	}

	data, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Question101 临时文件、defer Close 与错误处理顺序。
//
// 实现 WriteTempAndRead；本函数打印读写结果。
func Question101() {
	got, err := WriteTempAndRead("hello temp file\n")
	fmt.Println("read:", got)
	fmt.Println("err:", err)
	fmt.Println()
}

var (
	_ = io.Discard
	_ = os.Stderr
)
