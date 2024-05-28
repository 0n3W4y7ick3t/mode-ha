package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"mode-ha/parser"
	"mode-ha/solver"
)

const (
	TIME_FMT = time.RFC3339
)

func main() {
	log.SetOutput(os.Stderr)

	startStr, endStr, start := parser.ParseCMD(TIME_FMT)

	url := fmt.Sprintf("https://tsserv.tinkermode.dev/data?begin=%s&end=%s", startStr, endStr)
	log.Printf("getting %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to perform GET request: %v", err)
	}
	defer resp.Body.Close()

	solver.Solve(resp, start, TIME_FMT)
}
