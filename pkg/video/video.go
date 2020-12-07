package video

import (
	"fmt"
	"log"
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

// Trim produces one or more clips based on Video.duration
func (v *Video) Trim() ([]string, error) {
	var path []string
	for i, j := 0, 1; i < len(v.duration); i, j = i+2, j+1 {
		log.Printf("Processing duration [%s %s]\n", v.duration[i], v.duration[i+1])
		start, _ := time.ParseDuration(v.duration[i])
		end, _ := time.ParseDuration(v.duration[i+1])
		file := strings.Split(v.path, ".mp4")
		dest := fmt.Sprintf("%s_%d.mp4", file[0], j)
		v.video.Trim(start, end)
		v.video.Render(dest)
		log.Println("Saved", dest)
		path = append(path, dest)
	}
	return path, nil
}
