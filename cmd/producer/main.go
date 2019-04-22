package main

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

var (
	brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:19092").Strings()
	topic      = kingpin.Flag("topic", "Topic name").Default("important").String()
	maxRetry   = kingpin.Flag("maxRetry", "Retry limit").Default("5").Int()
)

func init() {
	// do something here to set environment depending on an environment variable
	// or command-line flag
	a := "production"
	fmt.Printf("(%s)あああああああ", a)
	fmt.Printf("(%s)あああああああ", *env)
	if *env == a {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		log.SetFormatter(&log.TextFormatter{})
	}
}
func main() {
	kingpin.Parse()

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = *maxRetry
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(*brokerList, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: *topic,
		Value: sarama.StringEncoder("Something Cool"),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", *topic, partition, offset)

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	log.WithFields(log.Fields{
		"animal": *env,
	}).Info("A walrus appears")
}
