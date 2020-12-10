# go-viddit

This tool is used to trim and concatenate video files from the terminal in a convenient and concurrent way.

Each clip is processed in a dedicated thread.

## Configuration

Developed and tested on the following setup:

- macOS 11.0.1
- Go 1.15.5
- ffmpeg 4.3.1

## Installation

Run the build process:

```terminal
go get github.com/cfanatic/go-viddit
cd $GOPATH/src/github.com/cfanatic/go-viddit
go build -o $GOPATH/bin/viddit cmd/viddit/main.go
```

Run the unit test to see if your setup is working correctly:

```terminal
go test -v pkg/video/video_test.go pkg/video/video.go
```

## Usage

Assuming that `$GOPATH/bin` is added to your `PATH` variable, you can run:

```terminal
viddit misc/test.mp4 0m10s 0m12s 0m42s 0m48s
```

This will trim `test.mp4` between timestamps 00:10 and 00:12, as well as 00:42 and 00:48.

Both clips will then be concatenated.

The parameters are duration strings each with a [unit suffix](https://golang.org/pkg/time/#ParseDuration).

## Credits

The video `misc/test.mp4` is taken from [https://media.w3.org](https://media.w3.org/2010/05/sintel/trailer.mp4).
