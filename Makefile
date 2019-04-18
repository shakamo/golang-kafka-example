.PHONY: producer
producer:
	go run ./cmd/producer/main.go --topic test

.PHONY: consumer
consumer:
	go run ./cmd/consumer/main.go --topic test

.PHONY: up
up:
	cd build && docker-compose up -d

.PHONY: down
down:
	cd build && docker-compose down

.PHONY: ps
ps:
	cd build && docker-compose ps
