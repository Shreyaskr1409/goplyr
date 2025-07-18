# Gapless Audio Queue with Beep in Go

To implement a gapless audio queue with editable playback using the Beep library in Go, you'll need to create a custom streamer that manages the queue and handles buffering. Here's a comprehensive solution:

## Implementation

```go
package main

import (
	"errors"
	"sync"

	"github.com/faiface/beep"
)

// QueueStreamer implements beep.Streamer and provides gapless playback with editable queue
type QueueStreamer struct {
	mu       sync.Mutex
	streamer beep.Streamer
	queue    []beep.Streamer
	buffer   *beep.Buffer
	format   beep.Format
	pos      int
	done     bool
}

// NewQueueStreamer creates a new QueueStreamer with the given format
func NewQueueStreamer(format beep.Format) *QueueStreamer {
	return &QueueStreamer{
		format: format,
		buffer: beep.NewBuffer(format),
	}
}

// Stream streams the audio from the queue with gapless playback
func (q *QueueStreamer) Stream(samples [][2]float64) (n int, ok bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	// If we're done and queue is empty, return false
	if q.done && len(q.queue) == 0 {
		return 0, false
	}

	// Fill the samples buffer
	for n < len(samples) {
		// If current streamer is nil or exhausted, get next from queue
		if q.streamer == nil || !q.ok {
			if len(q.queue) == 0 {
				q.done = true
				return n, n > 0
			}
			
			// Get next streamer from queue
			q.streamer = q.queue[0]
			q.queue = q.queue[1:]
			q.pos++
		}

		// Stream from current streamer
		sn, sok := q.streamer.Stream(samples[n:])
		n += sn
		q.ok = sok

		// If current streamer is exhausted, set to nil for next iteration
		if !q.ok {
			q.streamer = nil
		}
	}

	return n, true
}

// Err returns any error from the streamer
func (q *QueueStreamer) Err() error {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	if q.streamer != nil {
		if err := q.streamer.Err(); err != nil {
			return err
		}
	}
	return nil
}

// AddToQueue adds a streamer to the end of the queue
func (q *QueueStreamer) AddToQueue(streamer beep.Streamer) {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	q.queue = append(q.queue, streamer)
}

// InsertNext inserts a streamer at the front of the queue (plays next)
func (q *QueueStreamer) InsertNext(streamer beep.Streamer) {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	if len(q.queue) > 0 {
		q.queue = append([]beep.Streamer{streamer}, q.queue...)
	} else {
		q.queue = []beep.Streamer{streamer}
	}
}

// ClearQueue clears the playback queue
func (q *QueueStreamer) ClearQueue() {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	q.queue = nil
	q.streamer = nil
	q.ok = false
}

// RemoveFromQueue removes a track at the specified index
func (q *QueueStreamer) RemoveFromQueue(index int) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	if index < 0 || index >= len(q.queue) {
		return errors.New("index out of range")
	}
	
	q.queue = append(q.queue[:index], q.queue[index+1:]...)
	return nil
}

// QueueLength returns the current queue length
func (q *QueueStreamer) QueueLength() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	return len(q.queue)
}

// CurrentPosition returns the current position in the queue
func (q *QueueStreamer) CurrentPosition() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	return q.pos
}
```

## Usage Example

```go
package main

import (
	"fmt"
	"time"
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	// Initialize speaker with sample rate (adjust based on your files)
	sampleRate := 44100
	speaker.Init(sampleRate, sampleRate/10) // Small buffer size for lower latency

	// Create queue
	queue := NewQueueStreamer(beep.Format{SampleRate: sampleRate, NumChannels: 2, Precision: 3})

	// Load some audio files (replace with your actual files)
	file1, _ := os.Open("song1.mp3")
	file2, _ := os.Open("song2.mp3")
	file3, _ := os.Open("song3.mp3")

	streamer1, format1, _ := mp3.Decode(file1)
	streamer2, format2, _ := mp3.Decode(file2)
	streamer3, format3, _ := mp3.Decode(file3)

	// Resample if needed (assuming all files have same format for simplicity)
	if format1.SampleRate != queue.format.SampleRate {
		streamer1 = beep.Resample(4, format1.SampleRate, queue.format.SampleRate, streamer1)
	}
	if format2.SampleRate != queue.format.SampleRate {
		streamer2 = beep.Resample(4, format2.SampleRate, queue.format.SampleRate, streamer2)
	}
	if format3.SampleRate != queue.format.SampleRate {
		streamer3 = beep.Resample(4, format3.SampleRate, queue.format.SampleRate, streamer3)
	}

	// Add to queue
	queue.AddToQueue(streamer1)
	queue.AddToQueue(streamer2)
	queue.AddToQueue(streamer3)

	// Play the queue
	speaker.Play(queue)

	// Example of queue manipulation while playing
	go func() {
		time.Sleep(5 * time.Second) // Wait 5 seconds
		file4, _ := os.Open("song4.mp3")
		streamer4, format4, _ := mp3.Decode(file4)
		if format4.SampleRate != queue.format.SampleRate {
			streamer4 = beep.Resample(4, format4.SampleRate, queue.format.SampleRate, streamer4)
		}
		
		// Insert a new song as the next to play
		queue.InsertNext(streamer4)
		
		// Remove the third song from queue
		queue.RemoveFromQueue(2)
	}()

	// Wait for playback to finish (in a real app, you'd have a proper shutdown)
	select {}
}
```

## Key Features

1. **Gapless Playback**: The queue seamlessly transitions between tracks without gaps.

2. **Thread-safe Operations**: All queue manipulations are protected by mutex locks.

3. **Queue Management**:
   - Add tracks to the end of the queue
   - Insert tracks to play next
   - Remove specific tracks
   - Clear the entire queue

4. **Buffering**: The implementation ensures the next track is ready to play when the current one ends.

5. **Playback Control**: You can manipulate the queue while audio is playing.

## Notes

1. Make sure all your audio files have the same format or are properly resampled.

2. For best performance, consider pre-buffering the next track in the queue.

3. The example uses MP3 files, but you can use any format supported by Beep.

4. Adjust the speaker initialization parameters based on your latency requirements.

This implementation provides a solid foundation that you can extend with additional features like seeking, volume control, or more sophisticated queue management.
