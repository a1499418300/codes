package main

import (
	"log"
	"time"
)

func main() {
	test()
}

func test() {
	ticker := time.NewTicker(time.Second)
Loop:
	i := 0
	for {
		time.Sleep(time.Second)
		i++
		log.Printf("hello: %d", i)
		if i%10 == 0 {
			log.Printf("触发goto: %d", i)
			goto Loop
		}
	}
}
