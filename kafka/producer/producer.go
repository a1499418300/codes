package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	w *kafka.Writer
}

type Config struct {
	Brokers   []string
	BatchSize int
	// ... 更多配置
}

type Message struct {
	Key []byte
	Val []byte
}

func NewProducer(conf Config) *Producer {
	return &Producer{
		w: kafka.NewWriter(kafka.WriterConfig{
			Brokers:   conf.Brokers,
			BatchSize: conf.BatchSize,
		}),
	}
}

func (p *Producer) Producer(ctx context.Context, topic string, msgs []Message) error {
	tmp := make([]kafka.Message, 0, len(msgs))
	for _, v := range msgs {
		tmp = append(tmp, kafka.Message{
			Key:   v.Key,
			Value: v.Val,
			Topic: topic,
		})
	}
	return p.w.WriteMessages(ctx, tmp...)
}

const UnitW = 10000

func main() {
	producer := NewProducer(Config{
		Brokers:   []string{"127.0.0.1:9092"},
		BatchSize: 10000,
	})
	topic := "hello"
	n := UnitW * 100
	bathcSize := producer.w.BatchSize
	msgs := make([]Message, 0, bathcSize)
	for i := 0; i < n; i++ {
		msgs = append(msgs, Message{
			Key: []byte("key" + strconv.Itoa(i%6)),
			Val: []byte("hello world; id: " + strconv.Itoa(i)),
		})
		if i%bathcSize == 0 && i > 0 {
			begin := time.Now()
			err := producer.Producer(context.Background(), topic, msgs)
			log.Printf("第%d条消息生产耗时：%v", i, time.Since(begin))
			if err != nil {
				log.Printf("生产消息失败：" + err.Error())
			}
			msgs = make([]Message, 0, bathcSize)
		}
		if i%UnitW == 0 {
			log.Printf("生产消息，目标：%d, 当前进度：%d\n", i, i)
		}
	}
}
