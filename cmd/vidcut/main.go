package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jtguibas/cinema"
)

// Video contains the input file and other processing information
type Video struct {
	video    *cinema.Video
	path     string
	duration []string
}

// NewVideo returns a Video object that can process trim and merge operations
func NewVideo(path string, duration []string) (*Video, error) {
	log.Println("Loading " + path)
	video, err := cinema.Load(path)
	return &Video{video: video, path: path, duration: duration}, err
}

// Trim produces a single clip or multiple clips based on Video.duration
func (v *Video) Trim() error {
	for i, j := 0, 1; i < len(v.duration); i, j = i+2, j+1 {
		log.Printf("Processing duration [%s %s]\n", v.duration[i], v.duration[i+1])
		start, _ := time.ParseDuration(v.duration[i])
		end, _ := time.ParseDuration(v.duration[i+1])
		file := strings.Split(v.path, ".mp4")
		dest := fmt.Sprintf("%s_%d.mp4", file[0], j)
		v.video.Trim(start, end)
		v.video.Render(dest)
		log.Println("Saved", dest)
	}
	return nil
}

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

	video, err := NewVideo(path, duration)
	if err == nil {
		video.Trim()
	} else {
		panic("Could not load video")
	}

	log.Println("Done")
}
