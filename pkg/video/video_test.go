package video

import (
	"os"
	"path/filepath"
	"testing"
)

// TestTrim calls video.Trim to check for correct clip outputs based on the expected file size
func TestTrim(t *testing.T) {
	var (
		video    *Video
		path     string
		duration []string
		size     []float64
		err      error
	)

	path, _ = filepath.Abs("../../misc/test.mp4")
	duration = []string{
		"0m10s",
		"0m12s",
		"0m42s",
		"0m48s",
	}
	size = []float64{
		104.0, // size of misc/test_1.mp4 in kilobyte
		324.0, // size of misc/test_2.mp4 in kilobyte
	}

	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}

	if video, err = NewVideo(path, duration); err != nil {
		t.Fatal(err)
	}

	if err = video.Trim(); err != nil {
		t.Fatal(err)
	}

	if len(video.list) != 2 {
		t.Fatal("Expected two clips, but found only", len(video.list))
	}

	for i, clip := range video.list {
		f, _ := os.Stat(clip)
		value := float64(f.Size()) / 1024.0 // convert to kilobytes
		thres1 := size[i] * 0.90            // -10% margin of difference across different systems
		thres2 := size[i] * 1.10            // +10% margin of difference across different systems
		t.Log(value, thres1, thres2)
		if value < thres1 || value > thres2 {
			t.Fatal("Trim operation failed for", video.list[i])
		}
	}
}

// TestMerge calls video.Merge to check for a correct video concatenate operation based on the expected file size
func TestMerge(t *testing.T) {
	var (
		video    *Video
		path     string
		duration []string
		size     float64
		err      error
	)

	path, _ = filepath.Abs("../../misc/test.mp4")
	duration = []string{
		"0m10s",
		"0m12s",
		"0m42s",
		"0m48s",
	}
	size = 424.0 // size of misc/merged.mp4 in kilobyte

	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}

	if video, err = NewVideo(path, duration); err != nil {
		t.Fatal(err)
	}

	if err = video.Trim(); err != nil {
		t.Fatal(err)
	}

	if len(video.list) != 2 {
		t.Fatal("Expected two clips, but found only", len(video.list))
	}

	if err = video.Merge(); err != nil {
		t.Fatal(err)
	}

	output := filepath.Join(filepath.Dir(path), "merged.mp4")

	f, _ := os.Stat(output)
	value := float64(f.Size()) / 1024.0 // convert to kilobytes
	thres1 := size * 0.90               // -10% margin of difference across different systems
	thres2 := size * 1.10               // +10% margin of difference across different systems
	t.Log(value, thres1, thres2)
	if value < thres1 || value > thres2 {
		t.Fatal("Merge operation failed for")
	}
}
