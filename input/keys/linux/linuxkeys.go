package linuxkeys

import (
	"encoding/binary"
	"os"
	"sync"
)

type KeysInput struct {
	mu            sync.Mutex
	realTimeState map[KeyCode]bool
	currentState  map[KeyCode]bool
	previousState map[KeyCode]bool
	device        *os.File
}

// standard Linux input_event struct (64-bit architecture version)
type inputEvent struct {
	Sec   uint64 // timeval.tv_sec
	Usec  uint64 // timeval.tv_usec
	Type  uint16
	Code  uint16
	Value int32
}

// NewKeysInput requires the path to the keyboard device, e.g., "/dev/input/event0"
func NewKeysInput(devicePath string) (*KeysInput, error) {
	f, err := os.Open(devicePath)
	if err != nil {
		return nil, err
	}

	ki := &KeysInput{
		realTimeState: make(map[KeyCode]bool),
		currentState:  make(map[KeyCode]bool),
		previousState: make(map[KeyCode]bool),
		device:        f,
	}

	// Start a background thread to continuously read the event stream
	go ki.pollEvents()

	return ki, nil
}

func (k *KeysInput) pollEvents() {
	var event inputEvent
	for {
		// Read binary data directly from the device stream
		err := binary.Read(k.device, binary.LittleEndian, &event)
		if err != nil {
			break // Device closed, disconnected, or permission lost
		}

		// EV_KEY is type 1
		if event.Type == 1 {
			k.mu.Lock()
			// Value: 0 = Released, 1 = Pressed, 2 = Auto-Repeat
			switch event.Value {
			case 1, 2:
				k.realTimeState[KeyCode(event.Code)] = true
			case 0:
				k.realTimeState[KeyCode(event.Code)] = false
			}
			k.mu.Unlock()
		}
	}
}

func (k *KeysInput) Update(keysToTrack []KeyCode) {
	k.mu.Lock()
	defer k.mu.Unlock()

	// 1. Shift current state to previous state
	for key, isDown := range k.currentState {
		k.previousState[key] = isDown
	}

	// 2. Snap the real-time background state into the current frame's state
	for _, key := range keysToTrack {
		k.currentState[key] = k.realTimeState[key]
	}
}

func (k *KeysInput) KeyPressed(key KeyCode) bool {
	return k.currentState[key]
}

func (k *KeysInput) KeyJustPressed(key KeyCode) bool {
	return k.currentState[key] && !k.previousState[key]
}

func (k *KeysInput) KeyJustReleased(key KeyCode) bool {
	return !k.currentState[key] && k.previousState[key]
}
