package solver

import (
	"io"
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

func TestSolve(t *testing.T) {
	input := `2021-03-04T03:00:24Z 117.6289
2021-03-04T03:01:27Z 117.4429
2021-03-04T03:02:13Z 117.2574
2021-03-04T04:00:57Z 108.8629
2021-03-04T04:01:54Z 108.7147
2021-03-04T04:03:30Z 108.5674
2021-03-04T04:04:43Z 108.4208
2021-03-04T04:54:21Z 102.9821
2021-03-04T04:55:28Z 102.8686
2021-03-04T04:56:45Z 102.7558
2021-03-04T04:58:08Z 102.6437
2021-03-04T04:59:27Z 102.5322
2021-03-04T05:00:17Z 102.4213
`

	expected := `2021-03-04T03:00:00Z 117.4431
2021-03-04T04:00:00Z 105.3720
2021-03-04T05:00:00Z 102.4213
`

	reader := strings.NewReader(input)
	start, _ := time.Parse(time.RFC3339, "2021-03-04T03:00:00Z")
	TFMT := time.RFC3339

	output := captureOutput(func() {
		Solve(reader, start, TFMT)
	})

	if output != expected {
		t.Errorf("expected output:\n%s\ngot:\n%s", expected, output)
	}
}

// captureOutput captures stdout made by f() and pipes it to a buffer and returns its content as string
func captureOutput(f func()) string {
	oldStdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		os.Stdout = oldStdOut
	}()

	f()
	w.Close()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	r.Close()
	return buf.String()
}
