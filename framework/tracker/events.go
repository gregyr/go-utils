package tracker

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gregyr/go-utils/util/core/fun"
)

const (
	Version1_0 string = "TRK/1.0"

	ContentTypeEmpty string = "Empty"
	ContentTypeCount string = "Count"
)

type Content interface {
	// content type identifier
	Type() string
	// String Packaged Content
	Package() string

	Unpackage(string)
}

type Event struct {
	Bucket    string
	Content   Content
	Timestamp string
}

// Creates a new Event with the Timestamp being the creation time
func CreateEvent(bucket string, content Content) Event {
	return Event{Bucket: bucket, Content: content, Timestamp: time.Now().Format(time.RFC3339)}
}

func NewEvent(bucket string, content Content, timestamp string) Event {
	return Event{Bucket: bucket, Content: content, Timestamp: timestamp}
}

func (e *Event) String() string {
	var b strings.Builder
	b.WriteString(Version1_0)
	b.WriteRune(' ')
	fmt.Fprintf(&b, "%s %s %s %s", e.Content.Type(), e.Bucket, e.Timestamp, e.Content.Package())
	return b.String()
}

func ParseEvent(event string) (Event, error) {
	components := strings.Split(event, " ")
	if len(components) < 4 {
		return Event{}, NewInvalidEventError()
	}
	version := components[0]
	switch version {
	case Version1_0:
		return HandleVersion1_0(components)
	}
	return Event{}, fmt.Errorf("Could not parse event %s", event)
}

func HandleVersion1_0(eventComponents []string) (Event, error) {
	cType := eventComponents[1]
	bucket := eventComponents[2]
	timestamp := eventComponents[3]

	var content Content
	if cType == ContentTypeEmpty {
		content = NewEmptyContent()
	}

	strcontent := ""
	if len(eventComponents) != 5 {
		return Event{}, fmt.Errorf("Length of components is expected to be 5")
	} else {
		strcontent = strings.Trim(eventComponents[4], "\x00")
	}
	switch cType {
	case ContentTypeCount:
		content = &CountContent{}
	}

	content.Unpackage(strcontent)
	return NewEvent(bucket, content, timestamp), nil

}

type EmptyContent struct{}

func NewEmptyContent() *EmptyContent {
	return &EmptyContent{}
}

func (c *EmptyContent) Unpackage(s string) {}
func (c *EmptyContent) Type() string {
	return ContentTypeEmpty
}
func (c *EmptyContent) Package() string {
	return ""
}

type CountContent struct {
	count int
}

func NewCountContent(count int) *CountContent {
	return &CountContent{count: count}
}
func (c *CountContent) Unpackage(s string) {
	c.count = fun.AtoiOrZero(s)
}
func (c *CountContent) Type() string {
	return ContentTypeCount
}
func (c *CountContent) Package() string {
	return strconv.Itoa(c.count)
}
