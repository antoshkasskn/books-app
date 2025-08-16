# Makefile

TEST_FLAGS=-v -race -covermode atomic

# Цель по умолчанию
.PHONY: all
all: test

# Цель для запуска тестов
.PHONY: test
test:
	go test $(TEST_FLAGS) ./...

# Цель для запуска тестов с покрытием
.PHONY: test-cover
test-cover:
	go test $(TEST_FLAGS) -cover ./...

# Цель для запуска тестов с анализом
.PHONY: lint
lint:
	go vet ./...