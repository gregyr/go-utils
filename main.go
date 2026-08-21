package main

import (
	"fmt"
	"time"

	"github.com/gregyr/go-utils/tracing"
)

func DoSomething(span *tracing.Span) {
	fmt.Println("Doing Something")
	span.Trace("do something else", DoSomethingElse)
	time.Sleep(time.Second)
	fmt.Println("Done")
}

func DoSomethingElse(span *tracing.Span) {
	fmt.Println("Doing Something else")
	time.Sleep(2 * time.Second)
	fmt.Println("Done something else")
}

func main() {
	collector := tracing.NewCollector()
	collector.Trace("do something", DoSomething)
	tracing.Collect(collector)
}
