package video

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/jtguibas/cinema"
)

// Video contains the input file and other processing information
type Video struct {
	video    *cinema.Video // object that enables video processing
	path     string        // path to the video file
	duration []string      // list of duration strings with unit as in '0m48s, '3m20s', etc.
	list     []string      // list of videos to concatenate
}

// NewVideo returns a Video object that can process trim and merge operations
func NewVideo(path string, duration []string) (*Video, error) {
	path, _ = filepath.Abs(path)
	log.Println("Loading " + path)
	video, err := cinema.Load(path)
	return &Video{video: video, path: path, duration: duration}, err
}

// Trim produces one or more videos based on Video.duration
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
		v.list = append(v.list, dest)
	}
	return nil
}

// Merge produces a single clip out of multiple videos based on Video.duration
func (v *Video) Merge() error {
	var clip *cinema.Clip
	var err error
	log.Println("Concatenate video files")
	if clip, err = cinema.NewClip(v.list); err != nil {
		return err
	}
	if err := clip.Concatenate("merged.mp4"); err != nil {
		return err
	}
	return nil
}
