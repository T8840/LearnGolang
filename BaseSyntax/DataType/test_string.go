package main

/*
字符串是不可变值类型，内部用指针指向 UTF-8 字节数组。
	默认值是空字符串 ""。
	用索引号访问某字节，如 s[i]。
	不能用序号获取字节元素指针，&s[i] 非法。
	不可变类型，无法修改字节数组。
	字节数组尾部不包含 NULL。
*/
import (
	"fmt"
)

func main() {
	// 使用索引号访问字符 (byte)
	s := "abc"
	fmt.Println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)

	// 使用 "`" 定义不做转义处理的原始字符串，支持跨行。
	a := `a
	b\r\n\x00
	c`
	fmt.Println(a)

	// 连接跨行字符串时，"+" 必须在上一行末尾，否则导致编译错误。
	b1 := "hello" +
		"world"
	fmt.Println(b1)
	// b2 := "hello"
	// + "world"
	// fmt.Println(b2) // Error: invalid operation: + untyped string

	// 支持用两个索引号返回子串。子串依然指向原字节数组，仅修改了指针和长度属性。
	c := "Hello, World!"
	s1 := c[:5]  // Hello
	s2 := c[7:]  // World!
	s3 := c[1:5] // ello
	fmt.Printf("%v,%v,%v,%v\n", c, s1, s2, s3)

	// 要修改字符串，可先将其转换成 []rune 或 []byte，完成后再转换为 string。无论哪种转换，都会重新分配内存，并复制字节数组。
	d := "abcd"
	bs := []byte(d)
	bs[1] = 'B'
	fmt.Println(string(bs))
	u := "电脑"
	us := []rune(u)
	us[1] = '话'
	fmt.Println(string(us))

	// 用for循环遍历字符串时，也需要用byte或rune的方式

	e := "abcdefg汉字"
	for i := 0; i < len(e); i++ { //byte
		fmt.Printf("%c,", e[i])
	}
	fmt.Println()
	for _, r := range e { // rune
		fmt.Printf("%c,", r)

	}
}
