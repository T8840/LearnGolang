package main

import (
	"fmt"
)

var ( // 使用关键字 var 定义变量，自动初始化为零值。如果提供初始化值，可省略变量类型，由编译器自动推断。
	x int
	f float32 = 1.6
	s         = "abc"
)

func test() (int, string) {
	return 1, "abc"
}

func main() { // 编译器会将未使用的局部变量当做错误。

	n, q := 0x1234, "Hello, World!"
	fmt.Println(x, q, n)

	// 多变量赋值时，先计算所有相关值，然后再从左到右依次赋值。
	data, i := [3]int{0, 1, 2}, 0
	i, data[i] = 2, 100 // (i = 0) -> (i = 2), (data[0] = 100)

	// 特殊只写变量 "_"，用于忽略值占位。
	_, t := test()
	fmt.Println(t)

	// 注意重新赋值与定义新同名变量的区别。
	s := "abc"
	fmt.Println(&s)

	s, y := "hello", 20 // 重新赋值: 与前 s 在同一层次的代码块中，且有新的变量被定义。
	fmt.Println(&s, y)  // 通常函数多返回值 err 会被重复使用。

	{
		s, z := 1000, 30 // 定义新同名变量: 不在同一层次代码块。
		fmt.Println(&s, z)
	}


	// 引用类型：包括slice\map\channel
	// 内置函数 new 计算类型大小，为其分配零值内存，返回指针。而 make 会被编译器翻译成具体的创建函数，由其分配内存和初始化成员结构，返回对象而非指针
	a := []int{0, 0, 0} // 提供初始化表达式。
	a[1] = 10

	b := make([]int, 3) // makeslice
	b[1] = 10

	// c := new([]int)
	// c[1] = 10 // Error: invalid operation: c[1] (index of type *[]int)

 
	// 类型转换
	var g byte = 100
	// var n int = b // Error: cannot use b (type byte) as type int in assignment
	var h int = int(g) // 显式转换
	fmt.Println(h)
	// 使用括号避免优先级错误。
	
	// 同样不能将其他类型当 bool 值使用。
	// p := 100
	// if p { // Error: non-bool a (type int) used as if condition
	// 	println("true")
	// }


}
