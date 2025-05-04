package playerv2

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func InitPlayer() {
	filename := "../../../samples/test_audio.mp3"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// NOT TO BE CALLED EVERYTIME
	sr := format.SampleRate
	speaker.Init(sr, sr.N(time.Second/10))
	// Later parameter is buffer size, should be added into config file

	resampled := beep.Resample(4, format.SampleRate, sr, streamer)
	// 4 is the quality rate (reasonable in present use-case)

	done := make(chan bool)
	speaker.Play(beep.Seq(resampled, beep.Callback(func() {
		done <- true // plays the song till done is triggered
	})))

	<-done
}
