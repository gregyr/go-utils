package events

import (
	"fmt"
	"log"
	"strings"
	"time"

	. "github.com/gregyr/go-utils/std"
)

type CoreProcessor struct {
	channels           map[string]chan any
	channelNames       []string
	finalConsumedCount int
}

func NewProcessor() CoreProcessor {
	return CoreProcessor{channels: make(map[string]chan any), channelNames: make([]string, 0)}
}

func (p *CoreProcessor) NewChannel(name string, buffersize int) {
	buffersize = ClampMin(buffersize, 0)
	if p.channels == nil {
		p.channels = make(map[string]chan any)
	}
	p.channels[name] = make(chan any, buffersize)
	p.channelNames = append(p.channelNames, name)
}

func (p *CoreProcessor) SendTo(name string, data any) error {
	channel, ok := p.channels[name]
	if !ok {
		log.Panicf("Could not find channel %s", name)
	}
	channel <- data
	return nil
}

func (p *CoreProcessor) ReadFrom(name string) any {
	channel, ok := p.channels[name]
	if !ok {
		log.Panicf("Could not find channel %s", name)
	}
	return <-channel
}

func (p *CoreProcessor) RemoveChannel(name string) error {
	channel, ok := p.channels[name]
	if !ok {
		log.Panicf("Could not find channel %s", name)
	}
	close(channel)
	delete(p.channels, name)
	return nil
}

func (p *CoreProcessor) Run() {
	for {
		time.Sleep(10 * time.Millisecond)
		var b strings.Builder
		for _, name := range p.channelNames {
			fmt.Fprintf(&b, "%s: %2v | ", name, len(p.channels[name]))
		}
		fmt.Fprintf(&b, "Consumed: %4v", p.finalConsumedCount)
		b.WriteString("\r")
		fmt.Print(b.String())
	}
}

func (p *CoreProcessor) Process(name string, procFunc func(*Channel)) {
	channel := Channel{name: name, ep: p}
	go func() {
		for {
			procFunc(&channel)
		}
	}()
}

func (p *CoreProcessor) ProcessTransitive(from, to string, procFunc func(*Channel, *Channel)) {
	fromChannel := Channel{name: from, ep: p}
	toChannel := Channel{name: to, ep: p}
	go func() {
		for {
			procFunc(&fromChannel, &toChannel)
		}
	}()
}

type Channel struct {
	name string
	ep   *CoreProcessor
}

func (p *Channel) Send(data any) {
	p.ep.SendTo(p.name, data)
}

func (p *Channel) Read() any {
	return p.ep.ReadFrom(p.name)
}

func (p *Channel) Final() {
	p.ep.finalConsumedCount++
}
