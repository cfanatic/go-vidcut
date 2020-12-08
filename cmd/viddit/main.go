package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cfanatic/go-viddit/pkg/video"
)

func main() {
	var path string
	var duration []string

	if len(os.Args[1:]) > 0 {
		path = os.Args[1:][0]
	} else {
		panic("No video given")
	}

	if len(os.Args[2:]) > 0 {
		duration = os.Args[2:]
	} else {
		log.Println("Provide at least two duration strings with unit suffix, e.g. '3m20s', and exit with 'q':")
	STDIN:
		for {
			var input string
			fmt.Scanf("%s", &input)
			switch input {
			case "q":
				break STDIN
			case "":
				continue
			default:
				duration = append(duration, input)
			}
		}
	}

	if len(duration)%2 != 0 || len(duration) == 0 {
		panic("Invalid duration strings given")
	}

	video, err := video.NewVideo(path, duration)
	if err == nil {
		switch len(duration) == 2 {
		case true:
			video.Trim()
		case false:
			video.Trim()
			video.Merge()
		}
	} else {
		panic(err)
	}

	log.Println("Done")
}
