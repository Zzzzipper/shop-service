package merchant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	// Buffer size
	logBufSize = 30
)

type LogLine struct {
	line   string
	time   time.Time
	source string
}

type BufLogContainer struct {
	buffer []LogLine
	remote bool
	url    string
	source string
	mu     *sync.RWMutex
}

var logBuffer *BufLogContainer

//
// New - singleton constructor
//
func Log() *BufLogContainer {
	if logBuffer == nil {
		logBuffer = &BufLogContainer{
			mu: &sync.RWMutex{},
		}
		if os.Getenv("LOGBUFF_REMOTE") == "true" {
			logBuffer.remote = true
		}
		logBuffer.url = os.Getenv("LOGBUFF_URL")
		logBuffer.source = os.Getenv("LOG_SOURCE")
		if logBuffer.source == "" {
			logBuffer.source = "UNKNOWN"
		}
	}
	return logBuffer
}

//
// Send - send string line to srever
//
func (l *BufLogContainer) send(line LogLine) error {
	body, err := json.Marshal(line)

	if err != nil {
		fmt.Printf("BufLog: can't marshall line: %s\n", err.Error())
		return err
	}

	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, l.url, bodyReader)
	if err != nil {
		fmt.Printf("BufLog: could not create request: %s\n", err)
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	_, err = client.Do(req)
	if err != nil {
		fmt.Printf("BufLog: error making http request: %s\n", err)
		return err
	}

	return nil
}

//
// putString - put string to buffer with control size
//
func (l *BufLogContainer) putString(in LogLine) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if strings.Contains(in.line, "/GetLog") {
		return
	}

	fmt.Println(in.line)

	in.source = l.source

	if l.len() > logBufSize {
		l.buffer = append(l.buffer[:0], l.buffer[1:]...)
	}
	if l.remote {
		l.send(in)
	} else {
		l.buffer = append(l.buffer, in)
	}
}

//
// format - format print output
//
func (l *BufLogContainer) format(format string, a ...any) {
	l.putString(LogLine{
		line: fmt.Sprintf(format, a...),
		time: time.Now(),
	})
}

//
// line  - unformat print output
//
func (l *BufLogContainer) line(a ...any) {
	l.putString(LogLine{
		line: fmt.Sprintln(a...),
		time: time.Now(),
	})
}

//
// len - length of buffer (rows number)
//
func (l *BufLogContainer) len() int {
	return len(l.buffer)
}

//
// body - return string slice (array)
//
func (l *BufLogContainer) body() *[]LogLine {
	return &l.buffer
}
