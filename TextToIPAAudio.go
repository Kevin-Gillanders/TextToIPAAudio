package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func main() {
	// GetAudioSamples()
	log.SetFlags(log.Lshortfile)
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))

	var queue Queue
	var resampled *beep.Resampler

	resampled = ReadToMemory("./Audio/ɦ_5g4.wav", sr)
	// resampled = ReadToMemory("./Audio/PinkPanther30.wav", sr)
	AddToQueue(resampled, &queue)
	// resampled = ReadToMemory("./Audio/æ_050.wav", sr)
	// AddToQueue(resampled, &queue)
	resampled = ReadToMemory("./Audio/pʰ_300.wav", sr)
	AddToQueue(resampled, &queue)
	resampled = ReadToMemory("./Audio/i_000.wav", sr)
	AddToQueue(resampled, &queue)

	// resampled = ReadToMemory("./Audio/b_304.wav", sr)
	resampled = ReadToMemory("./Audio/b_204.wav", sr)
	AddToQueue(resampled, &queue)
	// resampled = ReadToMemory("./Audio/ɛ̈_042.wav", sr)
	// AddToQueue(resampled, &queue)
	resampled = ReadToMemory("./Audio/ɹ̠_654.wav", sr)
	AddToQueue(resampled, &queue)
	resampled = ReadToMemory("./Audio/d̠ʱ_354.wav", sr)
	AddToQueue(resampled, &queue)
	resampled = ReadToMemory("./Audio/e_020.wav", sr)
	AddToQueue(resampled, &queue)

	resampled = ReadToMemory("./Audio/z_544.wav", sr)
	AddToQueue(resampled, &queue)
	resampled = ReadToMemory("./Audio/o_029.wav", sr)
	AddToQueue(resampled, &queue)
	resampled = ReadToMemory("./Audio/i_000.wav", sr)
	AddToQueue(resampled, &queue)

	done := make(chan bool)

	speaker.Play(beep.Seq(&queue, beep.Callback(func() {
		done <- true
	})))
	fmt.Println("Playing ...")
	for {
		select {
		case <-done:
			return
		}
	}

}

func ReadToMemory(filename string, sr beep.SampleRate) *beep.Resampler {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)

	if err != nil {
		log.Fatal(err)
	}

	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	return resampled
}

func AddToQueue(resampled *beep.Resampler, queue *Queue) {

	speaker.Lock()
	queue.Add(resampled)
	speaker.Unlock()
}
