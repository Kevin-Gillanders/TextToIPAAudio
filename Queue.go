package main

import (
	"log"

	"github.com/faiface/beep"
)


type Queue struct{
	streamers [] beep.Streamer
	started bool
}

func (q *Queue) Add(streamers ...beep.Streamer){
	q.started = true
	log.Println("len q streamers : ", len(q.streamers))
	log.Println("Streamer", streamers)
	q.streamers = append(q.streamers, streamers...)
	log.Println("len q streamers after: ", len(q.streamers))

}

func (q *Queue) Stream(samples [][2]float64) (n int, ok bool) {
	// We use the filled variable to track how many samples we've
	// successfully filled already. We loop until all samples are filled.
	filled := 0
	for filled < len(samples) {
		// There are no streamers in the queue, so we stream silence.
		if len(q.streamers) == 0 {
			// if ! q.started {
				for i := range samples[filled:] {
					samples[i][0] = 0
					samples[i][1] = 0
				}
				break
			} else {
			// 	return 0, false
			// }
		}

		// We stream from the first streamer in the queue.
		n, ok := q.streamers[0].Stream(samples[filled:])
		log.Println(len(q.streamers))
		log.Println(n, ok)
		// If it's drained, we pop it from the queue, thus continuing with
		// the next streamer.
		if !ok {
			log.Println("Popping drained stream : ", len(q.streamers))
			q.streamers = q.streamers[1:]
			log.Println("Popped drained stream : ", len(q.streamers))
		}
		// We update the number of filled samples.
		filled += n
	}
	return len(samples), true
}

func (q *Queue) Err() error {
	return nil
}