package merchant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// Buffer size
	logBufSize = 30
)

type LogLine struct {
	Code   uint32    `json:"code"`
	Line   string    `json:"line"`
	Time   time.Time `json:"createdSince"`
	Source string    `json:"source"`
}

type BufLogContainer struct {
	buffer       []LogLine
	remote       bool
	url          string
	source       string
	mu           *sync.RWMutex
	logLoadSince int
}

var container *BufLogContainer

//
// New - singleton constructor
//
func Log() *BufLogContainer {
	if container == nil {
		container = &BufLogContainer{
			mu: &sync.RWMutex{},
		}

		i, err := strconv.Atoi(os.Getenv("LOG_LOAD_SINCE"))
		if err != nil {
			container.logLoadSince = 60
		} else {
			container.logLoadSince = i
		}

		container.url = os.Getenv("LOGBUFF_URL")
		if container.url != "" {
			container.remote = true
		}

		container.source = os.Getenv("LOG_SOURCE")
		if container.source == "" {
			container.source = "UNKNOWN"
		}

	}
	return container
}

//
// Send - send string line to srever
//
func (l *BufLogContainer) send(line LogLine) error {

	fmt.Printf("LogLine to BufLog send: %v\n", line)

	body, err := json.Marshal(line)

	if err != nil {
		fmt.Printf("BufLog: can't marshall line: %s\n", err.Error())
		return err
	}

	fmt.Printf("Body to BufLog send: %v\n", string(body))

	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, strings.TrimSpace(l.url), bodyReader)
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

	fmt.Printf("LogLine in putString: %v\n", in)

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
		Line:   fmt.Sprintf(format, a...),
		Time:   time.Now(),
		Source: l.source,
	})
}

//
// line  - unformat print output
//
func (l *BufLogContainer) line(a ...any) {
	l.putString(LogLine{
		Line:   fmt.Sprintln(a...),
		Time:   time.Now(),
		Source: l.source,
	})
}

//
// errorf  - printf error and return it
//
func (l *BufLogContainer) errorf(format string, a ...any) error {
	l.putString(LogLine{
		Line:   fmt.Sprintf(format, a...),
		Time:   time.Now(),
		Source: l.source,
	})
	return fmt.Errorf(format, a...)
}

//
// statusErrorf  - printf error with return code and return it
//
func (l *BufLogContainer) statusErrorf(code codes.Code, format string, a ...any) error {
	l.putString(LogLine{
		Code:   uint32(code),
		Line:   fmt.Sprintf(format, a...),
		Time:   time.Now(),
		Source: l.source,
	})
	return status.Errorf(code, format, a...)
}

//
// statusError  - printf error with return code and return it
//
func (l *BufLogContainer) statusError(code codes.Code, a string) error {
	l.putString(LogLine{
		Code:   uint32(code),
		Line:   fmt.Sprintln(a),
		Time:   time.Now(),
		Source: l.source,
	})
	return status.Error(code, a)
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
