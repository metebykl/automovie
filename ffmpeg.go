package automovie

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func convertJPGToMP4(inputJPG string, outputMP4 string, duration time.Duration) error {
	err := ffmpeg(
		"-loop", "1",
		"-i", inputJPG,
		"-c:v", "libx264",
		"-t", fmt.Sprintf("%.3f", duration.Seconds()),
		"-pix_fmt", "yuv420p",
		outputMP4,
	)
	if err != nil {
		return fmt.Errorf("failed to convert JPG to MP4: %w", err)
	}

	return nil
}

func ffmpeg(args ...string) error {
	cmd := exec.Command("ffmpeg", args...)
	return cmd.Run()
}

func ffmpegResult(args ...string) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("ffmpeg", args...)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}

	return stdout.String(), nil
}
