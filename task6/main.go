package main

import (
	"context"
	"fmt"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

var stopFlag int32

func CLosedByChan(stop <-chan bool) {
	for {
		select {
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		case <-stop:
			fmt.Println("Stoped")
			return
		}
	}
}

func ClosedOnEnd() {
	go func(x int) {
		time.Sleep(time.Second)
		fmt.Printf("%d \n", x)
	}(10)
}

func ClosedByTimeOutInMain(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "c OK"
}

func ClosedByFlag() {
	for {
		if stopFlag == 1 {
			fmt.Println("Stoped")
			return
		}
		fmt.Println("Working")
		time.Sleep(500 * time.Millisecond)
	}
}

func ClosedByContext(ctx context.Context) {
	for {
		select {
		default:
			fmt.Println("Working")
			time.Sleep(500 * time.Millisecond)
		case <-ctx.Done():
			fmt.Println("Done")
			return
		}
	}
}

func main() {
	fmt.Println("\nGoroutine is closed by chan")
	time.Sleep(500 * time.Millisecond)
	stop := make(chan bool)
	go CLosedByChan(stop)
	time.Sleep(2 * time.Second)
	stop <- true

	fmt.Println("\nGoroutine finished its job")
	time.Sleep(500 * time.Millisecond)
	go ClosedOnEnd()
	time.Sleep(2 * time.Second)

	fmt.Println("\nTimeout in main")
	time.Sleep(500 * time.Millisecond)
	c1 := make(chan string)
	go ClosedByTimeOutInMain(c1)
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeout c")
	}

	fmt.Println("\nGlobal flag")
	go ClosedByFlag()
	time.Sleep(1 * time.Second)
	stopFlag = 1
	time.Sleep(time.Second)

	fmt.Println("\nContext")
	ctx, cancel := context.WithCancel(context.Background())
	go ClosedByContext(ctx)
	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
}
