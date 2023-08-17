package main
import (
	"time"
	//"sync"
	"fmt"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
	   ch <- i + 1
	   time.Sleep(time.Second)
	}
	close(ch)
   }
   func consume(ch <-chan int) {
	for n := range ch {
	println(n)
	}
   }
   func main() {
	/*
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
	produce(ch)
	wg.Done()
	}()
	go func() {
	consume(ch)
	 wg.Done()
	}()
	wg.Wait()
	*/
	var m = map[string]int{
		"tony": 21,
		"tom": 22,
		"jim": 23,
	   }
	   counter := 0
	   for k, v := range m {
		if counter == 0 {
		delete(m, "tony")
		}
		counter++
		fmt.Println(k, v)
	   }
	   fmt.Println("counter is ", counter)
	   
   }