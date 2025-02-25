PORT=3000

.PHONY: build run

build:
	docker build -t supply .

run: build
	docker run -p $(PORT):3000 supply

clean:
	docker rmi -f supply