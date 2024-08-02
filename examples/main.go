package main

import "github.com/fly2z/automovie"

func main() {
	movie := automovie.NewMovie()

	clip1 := automovie.NewVideoClip("./assets/1.mp4")
	t := clip1.AddText("Welcome to the movie!", 100, 100)
	t.Color("red")
	t.FontSize(100)

	clip2 := automovie.NewVideoClip("./assets/2.mp4")
	clip2.AddText("Second Clip", 100, 100)

	clip3 := automovie.NewVideoClip("./assets/3.mp4")
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
