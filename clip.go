package automovie

import (
	"strings"
)

type Clip struct {
	path    string
	filters []ClipFilter
}

type ClipFilter interface {
	String() string
}

func NewClip(fileName string) Clip {
	return Clip{
		path:    fileName,
		filters: make([]ClipFilter, 0),
	}
}

func (c *Clip) AddText(text string, x int, y int) *ClipText {
	ct := NewClipText(text, x, y)
	c.filters = append(c.filters, ct)
	return ct
}

func (c *Clip) prepare(outFile string) error {
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
