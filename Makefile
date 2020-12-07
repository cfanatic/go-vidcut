build:
	go build -o bin/vidcut cmd/vidcut/main.go

run:
	@go build -o bin/vidcut cmd/vidcut/main.go
	@./bin/vidcut /Users/cfanatic/Coding/Go/src/github.com/cfanatic/go-vidcut/misc/test.mp4 0m10s 0m12s 0m42s 0m48s

clean:
	rm -f -r bin/vidcut
	rm -f -r misc/test_1.mp4
	rm -f -r misc/test_2.mp4
