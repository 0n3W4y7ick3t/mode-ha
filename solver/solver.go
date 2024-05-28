package solver

import (
	"bufio"
	"fmt"
	"net/http"
	"time"

	"mode-ha/parser"
)

// Solve uses a sized buffer to read from response, one line at a time,
// calculated the results and print them.
// it does not close the response body.
func Solve(resp *http.Response, start time.Time, TFMT string) {
	sum := 0.0                           // accumulates the numbers
	counter := 0                         // lines of data we accumulated in hbucket hour
	hbucket := start.Truncate(time.Hour) // current hourly bucket

	reader := bufio.NewReader(resp.Body) // default sized buffer
	for {
		line, ended := parser.ReadLine(reader)
		if ended {
			// if we still have data left
			if counter > 0 {
				fmt.Printf("%s %8.4f\n", hbucket.Format(TFMT), sum/float64(counter))
			}
			break
		}

		dataTime, dataNum := parser.ParseLine(line, TFMT)

		// log.Printf("[debug], %v, %v, %v\n", hbucket, dataTime, dataNum)

		if dataTime.Truncate(time.Hour) == hbucket {
			// still in current hourly bucket
			// accumulate and increase counter
			sum += dataNum
			counter += 1
		} else {
			// next hourly bucket
			// print current result, set to next bucket, reset counter and accumulates
			if counter > 0 {
				fmt.Printf("%s %8.4f\n", hbucket.Format(TFMT), sum/float64(counter))
			}
			hbucket = dataTime.Truncate(time.Hour)
			sum = dataNum
			counter = 1
		}
	}
}
