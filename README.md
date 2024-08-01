# AutoMovie

AutoMovie is a Go library to create videos programmatically.

## Features
- Add single or multiple video clips.
- Add text to video clips with customizable color, size, and position.
- Generate a final video from the added clips.
- Simple API.
- Uses ffmpeg for video processing.


## Installation
To use AutoMovie in your project, you need to have ffmpeg installed on your system. You can download ffmpeg from [ffmpeg.org](https://ffmpeg.org).

Then, get the package:
```bash
go get github.com/fly2z/automovie
```
**Note: This library is currently under development and may undergo significant changes.**

## Usage
Here's a basic example demonstrating how to use the AutoMovie library:
```Go
package main

import "github.com/fly2z/automovie"

func main() {
	movie := automovie.NewMovie()

	clip1 := automovie.NewClip("./assets/1.mp4")
	t := clip1.AddText("Welcome to the movie!", 100, 100)
	t.Color("red")
	t.FontSize(100)

	clip2 := automovie.NewClip("./assets/2.mp4")
	clip2.AddText("Second Clip", 100, 100)

	clip3 := automovie.NewClip("./assets/3.mp4")
	clip3.AddText("Third Clip", 100, 100)

	err := movie.AddClips(clip1, clip2, clip3)
	if err != nil {
		panic(err)
	}

	err = movie.Generate("./movie.mp4")
	if err != nil {
		panic(err)
	}
}
```
