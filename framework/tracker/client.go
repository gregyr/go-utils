package tracker

import (
	"fmt"
	"net/http"
	"strings"
)

type Tracker struct {
	BackendUrl string
}

func (t *Tracker) EmitEmptyEvent(bucket string) error {
	event := CreateEvent(bucket, NewEmptyContent())
	return t.post(event)
}

func (t *Tracker) EmitCountEvent(bucket string, count int) error {
	event := CreateEvent(bucket, NewCountContent(count))
	return t.post(event)
}

func (t *Tracker) post(event Event) error {
	resp, err := http.Post(t.BackendUrl, "track", strings.NewReader(event.String()))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Received non 200 Status Code: %v", resp.StatusCode)
	}
	return nil
}
