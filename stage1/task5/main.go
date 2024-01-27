package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

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
	fmt.Print("Enter a lifetime of workers: ")
	var duration int
	fmt.Scan(&duration)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*time.Duration(duration)))
	for i := 0; i < 5; i++ {
		go Worker(ch, ctx, i)
	}
	log.Println("Start!")
	for {
		select {
		case <-ctx.Done():
			cancel()
			log.Println("Time's up!")
			return
		default:
			ch <- time.Now().GoString()
			time.Sleep(time.Second)
		}
	}
}
