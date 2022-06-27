package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

var sr beep.SampleRate
var ipaFileLoc string
var engToIPA [][] string

func init(){
	sr = beep.SampleRate(44100)
	ipaFileLoc = "./en_UK.tsv"


	file, err := os.Open(ipaFileLoc)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
    	listing := strings.Split(scanner.Text(), "\t")
        engToIPA = append(engToIPA, listing)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

}

func main() {
	// GetAudioSamples()
	log.SetFlags(log.Lshortfile)
	speaker.Init(sr, sr.N(time.Second/10))

	// var queue Queue
	// var resampled *beep.Resampler

	fmt.Println("Please input a line of text you wish to be transcribed...")

	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()

    line := scanner.Text()

    translatedText := CreateTranslatedText(line)
    fmt.Println(translatedText)
	// SayHappyBirthDayZoe(&queue)

	// done := make(chan bool)

	// speaker.Play(beep.Seq(&queue, beep.Callback(func() {
	// 	done <- true
	// })))
	// fmt.Println("Playing ...")
	// for {
	// 	select {
	// 	case <-done:
	// 		fmt.Println("Channel cleared")
	// 		return
	// 	}
	// }

}

func SayHappyBirthDayZoe(queue *Queue) {
	resampled := ReadToMemory("./Audio/ɦ_5g4.wav", sr)
	// resampled = ReadToMemory("./Audio/PinkPanther30.wav", sr)
	AddToQueue(resampled, queue)
	// resampled = ReadToMemory("./Audio/æ_050.wav", sr)
	// AddToQueue(resampled, &queue)
	resampled = ReadToMemory("./Audio/pʰ_300.wav", sr)
	AddToQueue(resampled, queue)
	resampled = ReadToMemory("./Audio/i_000.wav", sr)
	AddToQueue(resampled, queue)

	// resampled = ReadToMemory("./Audio/b_304.wav", sr)
	resampled = ReadToMemory("./Audio/b_204.wav", sr)
	AddToQueue(resampled, queue)
	// resampled = ReadToMemory("./Audio/ɛ̈_042.wav", sr)
	// AddToQueue(resampled, &queue)
	resampled = ReadToMemory("./Audio/ɹ̠_654.wav", sr)
	AddToQueue(resampled, queue)
	resampled = ReadToMemory("./Audio/d̠ʱ_354.wav", sr)
	AddToQueue(resampled, queue)
	resampled = ReadToMemory("./Audio/e_020.wav", sr)
	AddToQueue(resampled, queue)

	resampled = ReadToMemory("./Audio/z_544.wav", sr)
	AddToQueue(resampled, queue)
	resampled = ReadToMemory("./Audio/o_029.wav", sr)
	AddToQueue(resampled, queue)
	resampled = ReadToMemory("./Audio/i_000.wav", sr)
	AddToQueue(resampled, queue)
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
