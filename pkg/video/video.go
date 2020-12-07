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
	clips    []string
}

// NewVideo returns a Video object that can process trim and merge operations
func NewVideo(path string, duration []string) (*Video, error) {
	log.Println("Loading " + path)
	video, err := cinema.Load(path)
	return &Video{video: video, path: path, duration: duration}, err
}

// Trim produces one or more clips based on Video.duration
func (v *Video) Trim() error {
	for i, j := 0, 1; i < len(v.duration); i, j = i+2, j+1 {
		log.Printf("Trim video between %s to %s\n", v.duration[i], v.duration[i+1])
		start, _ := time.ParseDuration(v.duration[i])
		end, _ := time.ParseDuration(v.duration[i+1])
		file := strings.Split(v.path, ".mp4")
		dest := fmt.Sprintf("%s_%d.mp4", file[0], j)
		v.video.Trim(start, end)
		if err := v.video.Render(dest); err != nil {
			return err
		}
		v.clips = append(v.clips, dest)
	}
	return nil
}

// Merge produces a single movie out of multiple clips based on Video.duration
func (v *Video) Merge() error {
	log.Println("Merge video clips")
	if err := v.video.Concatenate(v.clips, "merged.mp4"); err != nil {
		return err
	}
	return nil
}
