package main

import (
	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

func main() {
	nsqd := "127.0.0.1:4150"
	producer, err := nsq.NewProducer(nsqd, nsq.NewConfig())
	producer.Publish("test", []byte("Hello Tinywan 0002"))
	if err != nil {
		panic(err)
	}
}
