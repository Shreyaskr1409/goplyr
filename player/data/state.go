package data

import "time"

const (
	NONE PlaybackModeOpts = iota
	REPEAT
	REPEAT_SONG
)

const (
	PAUSE PlaybackStatusOpts = iota
	PLAY
)

type (
	PlaybackModeOpts   int
	PlaybackStatusOpts int
	PlayerStatus       struct {
		PlaybackMode   *PlaybackModeOpts
		PlaybackStatus *PlaybackStatusOpts
		timestamp      *time.Duration
	}
)
