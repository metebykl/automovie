package automovie

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

type Clip interface {
	prepare(outFile string) error
	getPath() string
}

type VideoClip struct {
	path    string
	filters []ClipFilter
}

type ClipFilter interface {
	String() string
}

func NewVideoClip(fileName string) VideoClip {
	return VideoClip{
		path:    fileName,
		filters: make([]ClipFilter, 0),
	}
}

func (c *VideoClip) AddText(text string, x int, y int) *ClipText {
	ct := NewClipText(text, x, y)
	c.filters = append(c.filters, ct)
	return ct
}

func (c VideoClip) getPath() string {
	return c.path
}

func (c VideoClip) prepare(outFile string) error {
	var filterArr []string
	for _, filter := range c.filters {
		filterArr = append(filterArr, filter.String())
	}
	filterStr := strings.Join(filterArr[:], ",")

	return ffmpeg(
		"-y",
		"-i", c.path,
		"-vf", filterStr,
		"-c:a", "copy",
		outFile,
	)
}

type ImageClip struct {
	path     string
	duration time.Duration
	filters  []ClipFilter
}

func NewImageClip(fileName string, duration time.Duration) ImageClip {
	return ImageClip{
		path:     fileName,
		duration: duration,
		filters:  make([]ClipFilter, 0),
	}
}

func (c *ImageClip) AddText(text string, x int, y int) *ClipText {
	ct := NewClipText(text, x, y)
	c.filters = append(c.filters, ct)
	return ct
}

func (c ImageClip) getPath() string {
	return c.path
}

func (c ImageClip) prepare(outFile string) error {
	filename := strings.TrimSuffix(outFile, filepath.Ext(outFile))
	vidPath := fmt.Sprintf("%s_vid.mp4", filename)

	err := convertJPGToMP4(c.path, vidPath, c.duration)
	if err != nil {
		return fmt.Errorf("failed to conver image clip to mp4: %v", err)
	}

	var filterArr []string
	for _, filter := range c.filters {
		filterArr = append(filterArr, filter.String())
	}
	filterStr := strings.Join(filterArr[:], ",")

	return ffmpeg(
		"-y",
		"-i", vidPath,
		"-vf", filterStr,
		"-c:a", "copy",
		outFile,
	)
}
