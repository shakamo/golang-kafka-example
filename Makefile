
BROKER_LIST := localhost:19092
ifeq ($(ENV),staging)
  BROKER_LIST := kafka:9092
endif
ifeq ($(ENV),production)
  BROKER_LIST := localhost:19092
endif




.PHONY: producer
producer:
	go run ./cmd/producer/main.go \
	--topic test \
	--brokerList $(BROKER_LIST)

.PHONY: consumer
consumer:
	go run ./cmd/consumer/main.go \
	--topic test \
	--brokerList $(BROKER_LIST)

.PHONY: up
up:
	cd build && docker-compose up -d --build

.PHONY: down
down:
	cd build && docker-compose down

.PHONY: down-all
down-all:
	cd build && docker-compose down --rmi all

.PHONY: ps
ps:
	cd build && docker-compose ps
