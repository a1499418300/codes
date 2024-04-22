package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/time/rate"
)

var limit = rate.Every(time.Second * 2)
var r = rate.NewLimiter(limit, 10)

func WaitN(ctx context.Context, i int) {
	if err := r.WaitN(ctx, 2); err != nil {
		log.Fatalf("等待发生错误：%v", err)
	}
	log.Printf("hello %d", i)
}

func AllowN(i int) {
	if !r.AllowN(time.Now(), 2) {
		log.Printf("令牌桶不足，丢弃请求 %d", i)
		return
	}
	log.Printf("hello %d", i)
}
