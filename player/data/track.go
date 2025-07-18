package data

import (
	"time"

	"github.com/faiface/beep"
)

type Track struct {
	S           beep.Streamer
	Name        string
	Album       string
	Artist      string
	AlbumArtist string
	AlbumArtURI string
	Duration    time.Duration
	FileURI     string
}

func InitTrack(fileURI string) (*Track, error)
