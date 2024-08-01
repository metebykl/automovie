package automovie

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
)

type Movie struct {
	clips []Clip
}

func NewMovie() *Movie {
	return &Movie{
		clips: make([]Clip, 0),
	}
}

func (m *Movie) AddClip(clip Clip) error {
	if _, err := os.Stat(clip.path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("clip '%s' does not exist", clip.path)
		}

		return fmt.Errorf("failed to open clip '%s'", clip.path)
	}

	m.clips = append(m.clips, clip)

	return nil
}

func (m *Movie) AddClips(clips ...Clip) error {
	for _, c := range clips {
		err := m.AddClip(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Movie) Generate(fileName string) error {
	listPath := "__input.txt"
	lf, err := os.Create(listPath)
	if err != nil {
		return fmt.Errorf("failed to create input list file: %w", err)
	}

	defer func() {
		lf.Close()
		if err := os.Remove(listPath); err != nil {
			fmt.Printf("failed to remove list file: %v\n", err)
		}
	}()

	autoMovieDir := "./.automovie"
	clipDir := filepath.Join(autoMovieDir, "clips")
	_ = os.MkdirAll(clipDir, 0700)
	defer func() {
		if err := os.RemoveAll(autoMovieDir); err != nil {
			fmt.Printf("failed to remove clip directory: %v\n", err)
		}
	}()

	for id, clip := range m.clips {
		clipPath := filepath.Join(clipDir, fmt.Sprintf("clip_%d.mp4", id))

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = fmt.Sprintf(" Preparing clip 'clip_%d.mp4'", id)
		s.Start()

		if err := clip.prepare(clipPath); err != nil {
			// break the loop if clip fails to prepare
			s.Stop()
			fmt.Printf("failed to prepare clip: %v\n", err)
			break
		}

		s.Stop()
		fmt.Fprintf(lf, "file '%s'\n", clipPath)
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Preparing the movie"

	s.Start()
	err = ffmpeg(
		"-y",
		"-f", "concat",
		"-safe", "0",
		"-i", listPath,
		"-c", "copy",
		fileName,
	)
	s.Stop()

	return err
}
