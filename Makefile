default: compile

compile:
	go build -o build/main examples/main.go

clean:
	rm -rf build

fmt:
	gofmt -w -l .

