package video

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"testing"
)

// TestTrim calls video.Trim to check for correct clip outputs based on sha256 checksums
func TestTrim(t *testing.T) {
	var (
		path  string
		clips []string
		hash  []string
		video *Video
		err   error
	)

	path, _ = filepath.Abs("../../misc/test.mp4")
	duration := []string{
		"0m10s",
		"0m12s",
		"0m42s",
		"0m48s",
	}
	hash = []string{
		"7c8faa3190ccf08b75e595f89315f91d105dc1969495afbdd32aa57d8aff46ff",
		"56931c048adc020117bc5c7cec8fcf83d8d2ee96115ffb9579a724e5e78130a0",
	}

	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}

	if video, err = NewVideo(path, duration); err != nil {
		t.Fatal(err)
	}

	if clips, err = video.Trim(); err != nil {
		t.Fatal(err)
	}

	for i, clip := range clips {
		f, _ := os.Open(clip)
		defer f.Close()
		hasher := sha256.New()
		io.Copy(hasher, f)
		value := hex.EncodeToString(hasher.Sum(nil))
		if value != hash[i] {
			t.Fatal("Trim operation failed for", clips[i])
		}
	}
}
