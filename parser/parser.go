package parser

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	log.SetOutput(os.Stderr)
}

// ParseCMD reads from os.Args, returns the start and end string, as well as the start time,
// @TFMT is the time format string used
// it fails and exits the program when the args are invalid.
func ParseCMD(TFMT string) (string, string, time.Time) {
	if len(os.Args) != 3 {
		log.Fatalf("should provide with 2 args, calling args: %v\n", os.Args)
	}
	startStr, endStr := os.Args[1], os.Args[2]
	start, err := time.Parse(TFMT, startStr)
	if err != nil {
		log.Fatalf("input invalid: start(%s) is not a valid RFC3339 timestamp\n", startStr)
	}
	end, err := time.Parse(TFMT, endStr)
	if err != nil {
		log.Fatalf("input invalid: end(%s) is not a valid RFC3339 timestamp\n", endStr)
	}

	if start.After(end) {
		log.Fatalf("input invalid: start(%s) timestamp is after end(%s)\n", startStr, endStr)
	}

	if start.Minute() != 0 || start.Second() != 0 || end.Minute() != 0 || end.Second() != 0 {
		log.Fatalf("input invalid: start(%s) and end(%s) should be hourly, their minutes and seconds must all be 0\n", startStr, endStr)
	}
	end = end.Add(59*time.Second + 59*time.Minute) // add 00:59:59
	return startStr, end.Format(TFMT), start
}

// ReadLine reads a line of data, returns it(without the \n) and a bool value whether the buffer is empty,
// true indicates end of buffer.
// it fails and exit the program if some error occurrs during reading.
func ReadLine(reader *bufio.Reader) (string, bool) {
	line, err := reader.ReadString('\n')
	if err != nil {
		if err.Error() != "EOF" {
			log.Fatalf("failed to read line: %v\n", err)
		}
		return "", true
	}
	return strings.TrimSuffix(line, "\n"), false
}

// ParseLine parses a line of data, return the time and float number,
// @TFMT is the time format string used
// it fails and exits the program if this line is invalid.
func ParseLine(line, TFMT string) (time.Time, float64) {
	lineSplits := strings.Fields(line)
	if len(lineSplits) != 2 {
		log.Fatalf("reponse line invalid: %s\n", line)
	}
	dataTime, err := time.Parse(TFMT, lineSplits[0])
	if err != nil {
		log.Fatalf("invalid timestamp: %s\n", lineSplits[0])
	}

	num, err := strconv.ParseFloat(lineSplits[1], 64)
	if err != nil {
		log.Fatalf("invalid number: %v\n", lineSplits[1])
	}
	return dataTime, num
}
