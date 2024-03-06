package components

import (
	"time"

	"github.com/cactus/go-statsd-client/v5/statsd"
)

func sendTime(time time.Duration) {

	statsdDir := "graphite:8125"

	statsdClient, err := statsd.NewClient(statsdDir, "aes_go")

	if err != nil {
		println("Error initializing statsd client ", err)
		return
	}

	statsdClient.Timing("elapsed_time", time.Milliseconds(), 1)

}

func printTime(time time.Duration) {

	if( time.Hours() > 1) {
		println("Elapsed time: ", time.Hours(), " hours")
	} else if( time.Minutes() > 1) {
		println("Elapsed time: ", time.Minutes(), " minutes")
	} else if (time.Seconds() > 1) {
		println("Elapsed time: ", time.Seconds(), " seconds")
	} else {
		println("Elapsed time: ", time.Milliseconds(), " milliseconds")
	}

}

func RunAndMeasure(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	printTime(elapsed)
	sendTime(elapsed)
}