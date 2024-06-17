package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	test(nil)
	timer := time.NewTicker(3 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				log.Printf("程序停止")
				return
			case <-timer.C:
				log.Printf("执行任务")
			}
		}
	}()
	time.Sleep(time.Minute)
	timer.Stop()
	done <- true
}

func test(a []string) {
	fmt.Println(len(a))
}
