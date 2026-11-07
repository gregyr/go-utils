package winkeys

import (
	"golang.org/x/sys/windows"
)

type KeysInput struct {
	getKeyState   *windows.LazyProc
	currentState  map[KeyCode]bool
	previousState map[KeyCode]bool
}

func NewKeysInput() *KeysInput {
	user32_dll := windows.NewLazySystemDLL("user32.dll")
	return &KeysInput{
		getKeyState:   user32_dll.NewProc("GetKeyState"),
		currentState:  make(map[KeyCode]bool),
		previousState: make(map[KeyCode]bool),
	}
}

func (k *KeysInput) Update(keysToTrack []KeyCode) {
	for key, isDown := range k.currentState {
		k.previousState[key] = isDown
	}

	for _, key := range keysToTrack {
		r1, _, _ := k.getKeyState.Call(key.value())
		k.currentState[key] = (r1&0x8000 != 0)
	}
}

func (k *KeysInput) KeyPressed(key KeyCode) bool {
	return k.currentState[key]
}

// Was the key just pressed this exact frame? (Down now, but Up last frame)
func (k *KeysInput) KeyJustPressed(key KeyCode) bool {
	return k.currentState[key] && !k.previousState[key]
}

// Was the key just released this exact frame? (Up now, but Down last frame)
func (k *KeysInput) KeyJustReleased(key KeyCode) bool {
	return !k.currentState[key] && k.previousState[key]
}
