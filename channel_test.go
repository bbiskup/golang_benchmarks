package main

import (
	"fmt"
	"testing"
)

/* deadlock
func BenchmarkGoChannelSendReceive(b *testing.B) {
    var ch1 = make(chan int)
    var ch2 = make(chan int)

    go func() {
        fmt.Printf("goroutine: Read\n")
        tmp := <-ch1
        fmt.Printf("goroutine: Write\n")
        ch2 <- tmp
    }()

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        fmt.Printf("main: Write\n")
        ch1 <- i
        fmt.Printf("goroutine: Read\n")
        var tmp = <-ch2
        _ = tmp
    }
}*/

func produce(n int, msgs chan<- int, done chan<- bool) {
	for i := 0; i < n; i++ {
		msgs <- i
	}
	done <- true
}

func consume(msgs <-chan int) {
	var last int
	for msg := range msgs {
		//msg := <-msgs
		last = msg
		//fmt.Println(msg)
	}
	fmt.Printf("Last: %d\n", last)
}

func BenchmarkProducerConsumer(b *testing.B) {
	var done = make(chan bool)
	var msgs = make(chan int)
	b.ResetTimer()
	go produce(b.N, msgs, done)
	go consume(msgs)
	<-done
}
