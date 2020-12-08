build:
	go build -o bin/viddit cmd/viddit/main.go

run:
	@go build -o bin/viddit cmd/viddit/main.go
	@./bin/viddit misc/test.mp4 0m10s 0m12s 0m42s 0m48s

clean:
	rm -f -r bin/viddit
	rm -f -r misc/test_1.mp4
	rm -f -r misc/test_2.mp4
	rm -f -r misc/merged.mp4
	rm -f -r misc/concat.txt
