package main

import (
	"context"
	"log"
	"testing"
	"time"
)

func Test_WaitN(t *testing.T) {
	log.Printf("启动")
	for i := 0; i < 10; i++ {
		WaitN(context.Background(), i)
	}
}

func Test_AllowN(t *testing.T) {
	log.Printf("启动")
	for i := 0; i < 10; i++ {
		AllowN(i)
	}
}

func Test_AllowN2(t *testing.T) {
	log.Printf("启动")
	time.Sleep(time.Second * 8)
	for i := 0; i < 10; i++ {
		AllowN(i)
	}
}
