run:
	go run cmd/quranapi/main.go
build:
	go build -o quranapi cmd/quranapi/main.go
prod: build
	./quranapi -d=true
	