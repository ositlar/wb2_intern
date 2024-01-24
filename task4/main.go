package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func Worker(ch chan string, ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			println(id, ": ", <-ch)
		}
	}
}

func main() {
	ch := make(chan string)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	fmt.Println("Enter a number of workers: ")
	var num int
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		go Worker(ch, ctx, i)
	}
	for {
		select {
		case <-ctx.Done():
			cancel()
			log.Println("CANCEL INTERRUPT")
			return
		default:
			ch <- time.Now().GoString()
			time.Sleep(time.Second)
		}
	}
}
