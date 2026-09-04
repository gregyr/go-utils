package tracker

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type EventWriter interface {
	WriteEvent(event Event) error
}

type Handler struct {
	eventWriter EventWriter
}

func NewHandler(eventWriter EventWriter) Handler {
	return Handler{eventWriter: eventWriter}
}

func (h *Handler) Run(port string) {
	log.Printf("Serving on port %s\n", port)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%s", port), h)
	if err != nil {
		panic(err)
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	data := make([]byte, 512)
	req.Body.Read(data)
	fmt.Println(string(data))
	event, err := ParseEvent(string(data))
	if err != nil {
		log.Println("REJECTED - Invalid Body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Event Format"))
	}
	log.Printf("RECEIVED %s @ %s -> %v\n", event.Bucket, event.Timestamp, event.Content.Package())
	err = h.eventWriter.WriteEvent(event)
	if err != nil {
		log.Println("FAIL - Failed to write Event")
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error, failed to collect Event"))
		return
	}
	w.WriteHeader(200)
}

type EventFileWriter struct {
	rootPath string
}

func (w *EventFileWriter) WriteEvent(event Event) error {
	file, err := os.OpenFile(
		filepath.Join(w.rootPath, fmt.Sprintf("%s.txt", event.Bucket)),
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s - %v\n", event.Timestamp, event.Content.Package())
	return err
}

func NewEventFileWriter() EventFileWriter {
	return EventFileWriter{rootPath: ""}
}
