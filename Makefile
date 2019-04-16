.PHONY: producer
producer:
	go run ./producer/main.go --topic test

.PHONY: consumer
consumer:
	go run ./consumer/main.go --topic test

