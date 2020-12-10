# go-viddit

This tool is used to trim and concatenate video files from the terminal in a convenient way.

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

This is the expected output:

```terminal
=== RUN   TestTrim
2020/12/10 12:28:09 Loading /Users/cfanatic/Coding/Go/src/github.com/cfanatic/go-viddit/misc/test.mp4
2020/12/10 12:28:09 Trim start
2020/12/10 12:28:09 Video 1 between 10s to 12s
2020/12/10 12:28:09 Video 2 between 42s to 48s
2020/12/10 12:28:09 Video 1 done
2020/12/10 12:28:09 Video 2 done
2020/12/10 12:28:09 Trim done
    video_test.go:52: 102.4384765625 93.60000000000001 114.4
    video_test.go:52: 321.693359375 291.6 356.40000000000003
--- PASS: TestTrim (0.64s)
=== RUN   TestMerge
2020/12/10 12:28:09 Loading /Users/cfanatic/Coding/Go/src/github.com/cfanatic/go-viddit/misc/test.mp4
2020/12/10 12:28:09 Trim start
2020/12/10 12:28:09 Video 1 between 10s to 12s
2020/12/10 12:28:10 Video 2 between 42s to 48s
2020/12/10 12:28:10 Video 1 done
2020/12/10 12:28:10 Video 2 done
2020/12/10 12:28:10 Trim done
2020/12/10 12:28:10 Concatenate videos
    video_test.go:104: 422.4765625 381.6 466.40000000000003
--- PASS: TestMerge (0.68s)
PASS
ok      command-line-arguments  1.480s
```

The video `misc/test.mp4` is taken from [https://media.w3.org](https://media.w3.org/2010/05/sintel/trailer.mp4).

## Usage

Assuming that `$GOPATH/bin` is added to your `PATH` variable, you can run:

```terminal
viddit misc/test.mp4 0m10s 0m12s 0m42s 0m48s
```

This will trim `test.mp4` between timestamps 00:10 and 00:12, as well as 00:42 and 00:48.

Both clips will then be concatenated.

The parameters are duration strings each with a [unit suffix](https://golang.org/pkg/time/#ParseDuration).
