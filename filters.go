package automovie

import "fmt"

type ClipText struct {
	x     int
	y     int
	size  int
	text  string
	color string
}

func NewClipText(text string, x int, y int) *ClipText {
	return &ClipText{
		x:     x,
		y:     y,
		size:  30,
		text:  text,
		color: "white",
	}
}

func (ct *ClipText) Color(color string) {
	ct.color = color
}

func (ct *ClipText) FontSize(size int) {
	ct.size = size
}

func (ct *ClipText) String() string {
	// "drawtext=text='{}':fontcolor={}:fontsize={}:x={}:y={}"
	return fmt.Sprintf("drawtext=text='%s':fontcolor=%s:fontsize=%d:x=%d:y=%d", ct.text, ct.color, ct.size, ct.x, ct.y)
}
