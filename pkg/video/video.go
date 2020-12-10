package video

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/jtguibas/cinema"
)

// Video contains the input file and other processing information
type Video struct {
	path     string     // path to the video file
	duration []string   // list of duration strings with unit as in '0m48s, '3m20s', etc.
	list     []string   // list of videos to concatenate
	mu       sync.Mutex // update video list concurrently
}

// NewVideo returns a Video object that can process trim and merge operations
func NewVideo(path string, duration []string) (*Video, error) {
	list := make([]string, len(duration)/2, len(duration)/2)
	path, err := filepath.Abs(path)
	log.Println("Loading " + path)
	return &Video{path: path, duration: duration, list: list}, err
}

// Trim concurrently produces one or more videos based on Video.duration
func (v *Video) Trim() error {
	var wg sync.WaitGroup
	worker := func(j int, videoWorker *cinema.Video, start, end time.Duration, wg *sync.WaitGroup) {
		defer wg.Done()
		log.Printf("Video %d between %s to %s\n", j, start, end)
		file := strings.Split(v.path, ".mp4")
		dest := fmt.Sprintf("%s_%d.mp4", file[0], j)
		videoWorker.Trim(start, end)
		videoWorker.Render(dest)
		v.mu.Lock()
		v.list[j-1] = dest
		v.mu.Unlock()
		log.Printf("Video %d done\n", j)
	}
	log.Println("Trim start")
	for i, j := 0, 1; i < len(v.duration); i, j = i+2, j+1 {
		wg.Add(1)
		start, _ := time.ParseDuration(v.duration[i])
		end, _ := time.ParseDuration(v.duration[i+1])
		videoWorker, _ := cinema.Load(v.path)
		go worker(j, videoWorker, start, end, &wg)
	}
	wg.Wait()
	log.Println("Trim done")
	return nil
}

// Merge produces a single clip out of multiple videos based on Video.duration
func (v *Video) Merge() error {
	var clip *cinema.Clip
	var err error
	log.Println("Concatenate videos")
	if clip, err = cinema.NewClip(v.list); err != nil {
		return err
	}
	if err := clip.Concatenate("merged.mp4"); err != nil {
		return err
	}
	return nil
}
