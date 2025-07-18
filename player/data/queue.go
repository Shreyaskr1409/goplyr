package data

import (
	"sync"

	"github.com/faiface/beep"
)

type Queue struct {
	Tracks       []*beep.Streamer
	CurrentTrack int

	mu            sync.Mutex
	queueStreamer beep.Streamer
	buffer        *beep.Buffer
	pos           int
}

func InitQueue() *Queue
func (q *Queue) AddTrack(track *Track) int // returns length of the queue
func (q *Queue) PlayNext(track *Track) int // returns kength of the queue
func (q *Queue) RemoveTrackNumber(int)     // searching each track one by one will be tedious
func (q *Queue) Clear()
func (q *Queue) GetPosition() int
func (q *Queue) GetLength() int
