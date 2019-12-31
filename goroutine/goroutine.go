package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int){
			for{
				//fmt.Printf("hello form " + "goroutine %d\n",i)   //fmt.printf会做io切换
				a[i]++
				runtime.Gosched()  //交出控制权
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
