package main

import (
	"fmt"
	"unsafe"
)

// ---------- 常量 ----------
const x, y int = 1, 2     // 多常量初始化
const s = "Hello, World!" // 类型推断

const ( // 常量组
	a, b      = 10, 100
	c    bool = false
)

const ( // 常量值还可以是 len、cap、unsafe.Sizeof 等编译器可确定结果的函数返回值。

	d = "abc"
	e = len(d)
	f = unsafe.Sizeof(e)
)

// ---------- 枚举 ----------
const (
	Sunday    = iota // 0
	Monday           // 1，通常省略后续⾏行表达式。
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

const ( // 在同一常量组中，可以提供多个 iota，它们各自增长。
	A, B = iota, iota << 10 // 0, 0 << 10
	C, D                    // 1, 1 << 10
)

const (
	_        = iota             // iota = 0
	KB int64 = 1 << (10 * iota) // iota = 1
	MB                          // 与 KB 表达式相同，但 iota = 2
	GB
	TB
)

const ( // 如果 iota 自增被打断，须显式恢复。
	M = iota // 0
	O        // 1
	P = "c"  // c
	Q        // c，与上⼀一⾏行相同。
	I = iota // 4，显式恢复。注意计数包含了 C、D 两⾏行。
	S        // 5
)

func main() {

	const x = "xxx" // 未使⽤用局部常量不会引发编译错误。
	fmt.Println(TB)
	fmt.Println(I)
}
