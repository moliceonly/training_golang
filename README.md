# Go基础练习题

目录已按 `part01/` 拆分，对齐 `train_hub/go-web-roadmap.html` 的 Part 01。

题号连续：**1 → 74**。

## part01 练习一览

| 包 | 题号 | 对应路线 | 说明 |
|----|------|----------|------|
| `part01/basegrammer` | 1–10 | 起点 | 基础语法 |
| `part01/controlpractice` | 11–20 | 起点 | 流程控制 |
| `part01/arrslicepractice` | 21–30 | 起点 / 1.1 切片 | 数组与切片 |
| `part01/functionmethodpractice` | 31–40 | 起点 | 函数与方法 |
| `part01/structpointerpractice` | 41–50 | **1.1** | 结构体与指针 |
| `part01/structpointerpractice` | 51–56 | **1.1 补强** | 值接收者账户、深浅拷贝、JSON、map 标签、iota |
| `part01/structpointerpractice` | 70–72 | **1.1 补强** | 匿名嵌入、nil map、切片截取与内存滞留（仅题干） |
| `part01/errorinterfacepractice` | 57–65 | **1.2** | 错误与接口 |
| `part01/errorinterfacepractice` | 66–69 | **1.2 补强** | `%w` / `As` / type switch / recover |
| `part01/errorinterfacepractice` | 73–74 | **1.2 补强** | 类型断言 ok 形式、any 与泛型边界（仅题干） |

## 如何跑测试

```bash
cd part01/structpointerpractice && go test -v ./
cd ../errorinterfacepractice && go test -v ./
```

或在模块根目录：

```bash
go test -v ./part01/structpointerpractice/
go test -v ./part01/errorinterfacepractice/
```

## 建议顺序

1. 41–50（结构体/指针）
2. 51–56、70–72（1.1 补强）
3. 57–69、73–74（1.2 错误/接口及补强）
4. 再进入路线图 1.3 包结构 / 1.4 标准库
