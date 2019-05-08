package main

import (
	"fmt"
	"time"
)

// package main

// import "fmt"

func sum(i int, c chan int) {
	c <- i // 把 sum 发送到通道 c
}

func main() {
	c := make(chan int)
	// go sum(1, c)
	// go sum(2, c)
	// go sum(3, c)
	// go sum(4, c)
	// go sum(5, c)
	// go sum(6, c)
	for i := 0; i < 10; i++ {
		go sum(i, c)
		time.Sleep(1000)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

// package main

// import (
// 	"fmt"
// )

// func f1(in chan int) {
// 	fmt.Println(<-in)
// }

// func main() {
// 	out := make(chan int)
// 	go f1(out)
// 	out <- 2
// 	out <- 3
// }
