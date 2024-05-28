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
		log.Fatalf("should provide with 2 args, called by %v", os.Args)
	}
	startStr, endStr := os.Args[1], os.Args[2]
	start, err := time.Parse(TFMT, startStr)
	if err != nil {
		log.Fatalf("%s is not a valid RFC3339 timestamp", startStr)
	}
	_, err = time.Parse(TFMT, endStr)
	if err != nil {
		log.Fatalf("%s is not a valid RFC3339 timestamp", endStr)
	}
	return startStr, endStr, start
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
