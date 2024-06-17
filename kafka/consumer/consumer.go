package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	r    *kafka.Reader
	pool ants.Pool
}

type Config struct {
	Brokers  []string
	Topic    string
	GroupId  string
	MinBytes int
	MaxBytes int
	PoolSize int
}

func NewConsumer(conf Config) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  conf.Brokers,
		Topic:    conf.Topic,
		GroupID:  conf.GroupId,
		MinBytes: conf.MinBytes,
		MaxBytes: conf.MaxBytes,
	})
	return &Consumer{r: r}
}

func (p *Consumer) Consumer() {

	defer p.r.Close()
	for {
		msg, err := p.r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("读取消息出错：%v", err)
			continue
		}
		if err := p.handler(context.Background(), msg); err != nil {
			log.Printf("处理消息出错：%v", err)
			continue
		}
	}
}

func (p *Consumer) handler(ctx context.Context, msg kafka.Message) error {
	log.Printf("Message received: key = %s, value = %s, partaion: %d, offset: %d\n", string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)
	time.Sleep(time.Millisecond * 100)
	return nil
}

func main() {
	var group string
	var poolsize int
	flag.StringVar(&group, "g", "defaultConsumerGroup", "消费组名")
	flag.IntVar(&poolsize, "p", 50, "协程并发数")
	flag.Parse()
	topic := "hello"

	consumer := NewConsumer(Config{
		Topic:    topic,
		GroupId:  group,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		Brokers:  []string{"127.0.0.1:9092"},
		PoolSize: poolsize,
	})
	consumer.Consumer()
}
