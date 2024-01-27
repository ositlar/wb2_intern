package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

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
