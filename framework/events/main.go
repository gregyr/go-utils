package events

import (
	"math/rand/v2"
	"time"
)

func main() {

	p := NewProcessor()

	p.NewChannel("InitialEvents", 3)
	p.NewChannel("ProcessedEvents", 3)

	p.Process("InitialEvents", InitialProducer)
	p.ProcessTransitive("InitialEvents", "ProcessedEvents", ConvertEvent)
	p.Process("ProcessedEvents", DifferentProcessedEventProducer)
	p.Process("ProcessedEvents", ConsumeProcessedEvents)

	p.Run()
}

func InitialProducer(c *Channel) {
	time.Sleep(300 * time.Millisecond)
	// log.Println("New Initial Event")
	c.Send(SomeInitialEvent{Id: rand.Int(), SomeOtherValue: "Cool Message"})
}

func ConvertEvent(cFrom *Channel, cTo *Channel) {
	value := cFrom.Read().(SomeInitialEvent)
	time.Sleep(600 * time.Millisecond)
	// log.Println("Processed Event")
	cTo.Send(SomeProcessedEvent{Id: value.Id, NewValue: rand.Float32(), MessageValueOrSomething: value.SomeOtherValue})
}

func DifferentProcessedEventProducer(c *Channel) {
	time.Sleep(200 * time.Millisecond)
	// log.Println("New Processed Event from other Source")
	c.Send(SomeProcessedEvent{Id: rand.Int(), NewValue: rand.Float32(), MessageValueOrSomething: "other source message"})
}

func ConsumeProcessedEvents(c *Channel) {
	time.Sleep(190 * time.Millisecond)
	c.Read()
	c.Final()
	// log.Printf("Consumed Event - Id: %v, Value: %v, Message: %s", value.Id, value.NewValue, value.MessageValueOrSomething)
}

type SomeInitialEvent struct {
	Id             int
	SomeOtherValue string
}

type SomeProcessedEvent struct {
	NewValue                float32
	Id                      int
	MessageValueOrSomething string
}
