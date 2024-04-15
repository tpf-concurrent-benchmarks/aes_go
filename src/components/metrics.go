package components

import (
	"fmt"
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

	statsdClient.Gauge("completion_time", time.Milliseconds(), 1)

}

func printTime(time time.Duration) {


	if time.Hours() > 1 {
		fmt.Printf("Elapsed time: %f hours\n", time.Hours())
	} else if time.Minutes() > 1 {
		fmt.Printf("Elapsed time: %f minutes\n", time.Minutes())
	} else if time.Seconds() > 1 {
		fmt.Printf("Elapsed time: %f seconds\n", time.Seconds())
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
