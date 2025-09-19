// Package ssestream provides utilities for Server-Sent Events (SSE) streaming.
package ssestream

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Stream represents a Server-Sent Events stream.
type Stream[T any] struct {
	decoder *Decoder
	err     error
	current *T
	hasNext bool
}

// Decoder handles reading and parsing SSE data from an HTTP response.
type Decoder struct {
	response *http.Response
	scanner  *bufio.Scanner
}

// NewStream creates a new SSE stream with the given decoder and error.
func NewStream[T any](decoder *Decoder, err error) *Stream[T] {
	return &Stream[T]{
		decoder: decoder,
		err:     err,
		hasNext: true,
	}
}

// NewDecoder creates a new decoder from an HTTP response.
func NewDecoder(response *http.Response) *Decoder {
	if response == nil {
		return nil
	}

	return &Decoder{
		response: response,
		scanner:  bufio.NewScanner(response.Body),
	}
}

// Next advances the stream to the next event and returns true if an event is available.
func (s *Stream[T]) Next() bool {
	if s.err != nil {
		return false
	}

	if s.decoder == nil {
		s.err = fmt.Errorf("decoder is nil")
		return false
	}

	event, err := s.decoder.readEvent()
	if err != nil {
		if err == io.EOF {
			s.hasNext = false
			return false
		}
		s.err = err
		return false
	}

	if event == nil {
		s.hasNext = false
		return false
	}

	var result T
	if err := json.Unmarshal([]byte(event.Data), &result); err != nil {
		s.err = fmt.Errorf("failed to unmarshal event data: %w", err)
		return false
	}

	s.current = &result
	return true
}

// Current returns the current event in the stream.
func (s *Stream[T]) Current() T {
	if s.current == nil {
		var zero T
		return zero
	}
	return *s.current
}

// Err returns any error that occurred during streaming.
func (s *Stream[T]) Err() error {
	return s.err
}

// Close closes the underlying HTTP response body.
func (s *Stream[T]) Close() error {
	if s.decoder != nil && s.decoder.response != nil && s.decoder.response.Body != nil {
		return s.decoder.response.Body.Close()
	}
	return nil
}

// Event represents a single Server-Sent Event.
type Event struct {
	ID    string
	Event string
	Data  string
	Retry int
}

// readEvent reads a single SSE event from the decoder.
func (d *Decoder) readEvent() (*Event, error) {
	if d.scanner == nil {
		return nil, fmt.Errorf("scanner is nil")
	}

	event := &Event{}

	for d.scanner.Scan() {
		line := strings.TrimSpace(d.scanner.Text())

		// Empty line indicates end of event
		if line == "" {
			if event.Data != "" || event.Event != "" || event.ID != "" {
				return event, nil
			}
			continue
		}

		// Skip comments
		if strings.HasPrefix(line, ":") {
			continue
		}

		// Parse field
		if colonIndex := strings.Index(line, ":"); colonIndex != -1 {
			field := strings.TrimSpace(line[:colonIndex])
			value := strings.TrimSpace(line[colonIndex+1:])

			switch field {
			case "data":
				if event.Data != "" {
					event.Data += "\n"
				}
				event.Data += value
			case "event":
				event.Event = value
			case "id":
				event.ID = value
			case "retry":
				// Retry field parsing is optional and may contain non-numeric values
				continue
			}
		}
	}

	if err := d.scanner.Err(); err != nil {
		return nil, err
	}

	// EOF reached
	return nil, io.EOF
}