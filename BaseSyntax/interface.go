package main
import (  
    "fmt"
)

type I interface {
    M()
    }
    
type T struct {
}

func (T) M() {
}
func case1() int {
    println("eval case1 expr")
    return 1
   }
   func case2() int {
    println("eval case2 expr")
    return 2
   }
   func switchexpr() int {
    println("eval switch expr")
    return 1
   }
  
func main() {  
    a := [...]float64{67.7, 89.8, 21, 78,99}
    sum := float64(0)
    for i, v := range a {//range returns both the index and value
        fmt.Printf("%d the element of a is %.2f\n", i, v)
        sum += v
    }
    b:= [...][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
    }
    for i,v:=range b{
        fmt.Printf("%d the element of a is %s\n", i, v)
    }

    c := [5]int{76, 77, 78, 79, 80}
    var d []int = c[1:4] // creates a slice from a[1] to a[3]
    fmt.Println(d)
    for i:= range(d) {
        d[i]++
    }
    fmt.Println(cap(d))

    i := make([]int, 5, 5)
    fmt.Println(i)

    type MyByte = byte
    var m MyByte
    fmt.Printf("%T",m)
    type MyInt int
    var n MyInt
    fmt.Printf("%T",n)
    fmt.Println("分割线---------------------------------")
   
    
           fmt.Println("分割线---------------------------------")

           var t T
 var l I = t
 switch l.(type) {
 case T:
 println("it is type T")
 case int:
 println("it is type int")
 case string:
 println("it is type string")
 }
}