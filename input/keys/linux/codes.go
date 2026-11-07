// not tested at all lol
package linuxkeys

type KeyCode uintptr

const (
	// Mouse Buttons (mapped to Linux BTN_LEFT, BTN_RIGHT, etc.)
	KEY_LBUTTON  KeyCode = 272 // BTN_LEFT
	KEY_RBUTTON  KeyCode = 273 // BTN_RIGHT
	KEY_MBUTTON  KeyCode = 274 // BTN_MIDDLE
	KEY_XBUTTON1 KeyCode = 275 // BTN_SIDE
	KEY_XBUTTON2 KeyCode = 276 // BTN_EXTRA

	// Modifier & Control Keys
	KEY_BACKSPACE KeyCode = 14
	KEY_TAB       KeyCode = 15
	KEY_ENTER     KeyCode = 28
	KEY_SHIFT     KeyCode = 42  // KEY_LEFTSHIFT
	KEY_CTRL      KeyCode = 29  // KEY_LEFTCTRL
	KEY_ALT       KeyCode = 56  // KEY_LEFTALT
	KEY_PAUSE     KeyCode = 119 // KEY_PAUSE
	KEY_CAPSLOCK  KeyCode = 58
	KEY_ESC       KeyCode = 1
	KEY_SPACE     KeyCode = 57

	// Navigation Keys
	KEY_PAGE_UP   KeyCode = 104 // KEY_PAGEUP
	KEY_PAGE_DOWN KeyCode = 109 // KEY_PAGEDOWN
	KEY_END       KeyCode = 107
	KEY_HOME      KeyCode = 102
	KEY_LEFT      KeyCode = 105
	KEY_UP        KeyCode = 103
	KEY_RIGHT     KeyCode = 106
	KEY_DOWN      KeyCode = 108
	KEY_INSERT    KeyCode = 110
	KEY_DELETE    KeyCode = 111

	// Number Keys (Top Row)
	KEY_0 KeyCode = 11
	KEY_1 KeyCode = 2
	KEY_2 KeyCode = 3
	KEY_3 KeyCode = 4
	KEY_4 KeyCode = 5
	KEY_5 KeyCode = 6
	KEY_6 KeyCode = 7
	KEY_7 KeyCode = 8
	KEY_8 KeyCode = 9
	KEY_9 KeyCode = 10

	// Letter Keys
	KEY_A KeyCode = 30
	KEY_B KeyCode = 48
	KEY_C KeyCode = 46
	KEY_D KeyCode = 32
	KEY_E KeyCode = 18
	KEY_F KeyCode = 33
	KEY_G KeyCode = 34
	KEY_H KeyCode = 35
	KEY_I KeyCode = 23
	KEY_J KeyCode = 36
	KEY_K KeyCode = 37
	KEY_L KeyCode = 38
	KEY_M KeyCode = 50
	KEY_N KeyCode = 49
	KEY_O KeyCode = 24
	KEY_P KeyCode = 25
	KEY_Q KeyCode = 16
	KEY_R KeyCode = 19
	KEY_S KeyCode = 31
	KEY_T KeyCode = 20
	KEY_U KeyCode = 22
	KEY_V KeyCode = 47
	KEY_W KeyCode = 17
	KEY_X KeyCode = 45
	KEY_Y KeyCode = 21
	KEY_Z KeyCode = 44

	// Windows/OS Keys
	KEY_LWIN KeyCode = 125 // KEY_LEFTMETA
	KEY_RWIN KeyCode = 126 // KEY_RIGHTMETA
	KEY_APPS KeyCode = 127 // KEY_COMPOSE

	// Numpad Keys
	KEY_NUMPAD0   KeyCode = 82  // KEY_KP0
	KEY_NUMPAD1   KeyCode = 79  // KEY_KP1
	KEY_NUMPAD2   KeyCode = 80  // KEY_KP2
	KEY_NUMPAD3   KeyCode = 81  // KEY_KP3
	KEY_NUMPAD4   KeyCode = 75  // KEY_KP4
	KEY_NUMPAD5   KeyCode = 76  // KEY_KP5
	KEY_NUMPAD6   KeyCode = 77  // KEY_KP6
	KEY_NUMPAD7   KeyCode = 71  // KEY_KP7
	KEY_NUMPAD8   KeyCode = 72  // KEY_KP8
	KEY_NUMPAD9   KeyCode = 73  // KEY_KP9
	KEY_MULTIPLY  KeyCode = 55  // KEY_KPASTERISK
	KEY_ADD       KeyCode = 78  // KEY_KPPLUS
	KEY_SEPARATOR KeyCode = 121 // KEY_KPCOMMA
	KEY_SUBTRACT  KeyCode = 74  // KEY_KPMINUS
	KEY_DECIMAL   KeyCode = 83  // KEY_KPDOT
	KEY_DIVIDE    KeyCode = 98  // KEY_KPSLASH

	// Function Keys
	KEY_F1  KeyCode = 59
	KEY_F2  KeyCode = 60
	KEY_F3  KeyCode = 61
	KEY_F4  KeyCode = 62
	KEY_F5  KeyCode = 63
	KEY_F6  KeyCode = 64
	KEY_F7  KeyCode = 65
	KEY_F8  KeyCode = 66
	KEY_F9  KeyCode = 67
	KEY_F10 KeyCode = 68
	KEY_F11 KeyCode = 87
	KEY_F12 KeyCode = 88

	// Lock Keys
	KEY_NUMLOCK    KeyCode = 69 // KEY_NUMLOCK
	KEY_SCROLLLOCK KeyCode = 70 // KEY_SCROLLLOCK

	// Punctuation / OEM Keys (Standard US Keyboard mapping to Linux evdev)
	KEY_SEMICOLON KeyCode = 39 // ; : (KEY_SEMICOLON)
	KEY_EQUAL     KeyCode = 13 // = + (KEY_EQUAL)
	KEY_COMMA     KeyCode = 51 // , < (KEY_COMMA)
	KEY_MINUS     KeyCode = 12 // - _ (KEY_MINUS)
	KEY_PERIOD    KeyCode = 52 // . > (KEY_DOT)
	KEY_SLASH     KeyCode = 53 // / ? (KEY_SLASH)
	KEY_BACKTICK  KeyCode = 41 // ` ~ (KEY_GRAVE)
	KEY_LBRACKET  KeyCode = 26 // [ { (KEY_LEFTBRACE)
	KEY_BACKSLASH KeyCode = 43 // \ | (KEY_BACKSLASH)
	KEY_RBRACKET  KeyCode = 27 // ] } (KEY_RIGHTBRACE)
	KEY_QUOTE     KeyCode = 40 // ' " (KEY_APOSTROPHE)
)

// LetterKeys returns a fixed-length array of the 26 standard alphabet keys.
func LetterKeys() []KeyCode {
	return []KeyCode{
		KEY_A, KEY_B, KEY_C, KEY_D, KEY_E, KEY_F, KEY_G, KEY_H, KEY_I, KEY_J,
		KEY_K, KEY_L, KEY_M, KEY_N, KEY_O, KEY_P, KEY_Q, KEY_R, KEY_S, KEY_T,
		KEY_U, KEY_V, KEY_W, KEY_X, KEY_Y, KEY_Z,
	}
}

// NumberKeys returns a fixed-length array of the 10 top-row number keys.
func NumberKeys() []KeyCode {
	return []KeyCode{
		KEY_0, KEY_1, KEY_2, KEY_3, KEY_4, KEY_5, KEY_6, KEY_7, KEY_8, KEY_9,
	}
}

// FunctionKeys returns a fixed-length array of the 12 standard function keys.
func FunctionKeys() []KeyCode {
	return []KeyCode{
		KEY_F1, KEY_F2, KEY_F3, KEY_F4, KEY_F5, KEY_F6,
		KEY_F7, KEY_F8, KEY_F9, KEY_F10, KEY_F11, KEY_F12,
	}
}

// NumpadKeys returns a fixed-length array of the 16 numeric keypad keys (numbers and operators).
func NumpadKeys() []KeyCode {
	return []KeyCode{
		KEY_NUMPAD0, KEY_NUMPAD1, KEY_NUMPAD2, KEY_NUMPAD3, KEY_NUMPAD4,
		KEY_NUMPAD5, KEY_NUMPAD6, KEY_NUMPAD7, KEY_NUMPAD8, KEY_NUMPAD9,
		KEY_MULTIPLY, KEY_ADD, KEY_SEPARATOR, KEY_SUBTRACT, KEY_DECIMAL, KEY_DIVIDE,
	}
}

// NavigationKeys returns a fixed-length array of the 10 directional and page navigation keys.
func NavigationKeys() []KeyCode {
	return []KeyCode{
		KEY_PAGE_UP, KEY_PAGE_DOWN, KEY_END, KEY_HOME, KEY_LEFT,
		KEY_UP, KEY_RIGHT, KEY_DOWN, KEY_INSERT, KEY_DELETE,
	}
}

// ModifierAndControlKeys returns a fixed-length array of the 10 primary control keys.
func ModifierAndControlKeys() []KeyCode {
	return []KeyCode{
		KEY_BACKSPACE, KEY_TAB, KEY_ENTER, KEY_SHIFT, KEY_CTRL,
		KEY_ALT, KEY_PAUSE, KEY_CAPSLOCK, KEY_ESC, KEY_SPACE,
	}
}

// PunctuationKeys returns a fixed-length array of the 11 standard OEM punctuation keys.
func PunctuationKeys() []KeyCode {
	return []KeyCode{
		KEY_SEMICOLON, KEY_EQUAL, KEY_COMMA, KEY_MINUS, KEY_PERIOD,
		KEY_SLASH, KEY_BACKTICK, KEY_LBRACKET, KEY_BACKSLASH, KEY_RBRACKET, KEY_QUOTE,
	}
}

// SystemAndLockKeys returns a fixed-length array of the 5 OS and state-lock keys.
func SystemAndLockKeys() []KeyCode {
	return []KeyCode{
		KEY_LWIN, KEY_RWIN, KEY_APPS, KEY_NUMLOCK, KEY_SCROLLLOCK,
	}
}

// MouseButtons returns a fixed-length array of the 5 standard mouse buttons.
func MouseButtons() []KeyCode {
	return []KeyCode{
		KEY_LBUTTON, KEY_RBUTTON, KEY_MBUTTON, KEY_XBUTTON1, KEY_XBUTTON2,
	}
}

// AllKeys returns a fixed-length array containing all 105 defined standard keys.
func AllKeys() []KeyCode {
	return []KeyCode{
		// Letters (26)
		KEY_A, KEY_B, KEY_C, KEY_D, KEY_E, KEY_F, KEY_G, KEY_H, KEY_I, KEY_J,
		KEY_K, KEY_L, KEY_M, KEY_N, KEY_O, KEY_P, KEY_Q, KEY_R, KEY_S, KEY_T,
		KEY_U, KEY_V, KEY_W, KEY_X, KEY_Y, KEY_Z,
		// Numbers (10)
		KEY_0, KEY_1, KEY_2, KEY_3, KEY_4, KEY_5, KEY_6, KEY_7, KEY_8, KEY_9,
		// Function (12)
		KEY_F1, KEY_F2, KEY_F3, KEY_F4, KEY_F5, KEY_F6, KEY_F7, KEY_F8, KEY_F9, KEY_F10, KEY_F11, KEY_F12,
		// Numpad (16)
		KEY_NUMPAD0, KEY_NUMPAD1, KEY_NUMPAD2, KEY_NUMPAD3, KEY_NUMPAD4,
		KEY_NUMPAD5, KEY_NUMPAD6, KEY_NUMPAD7, KEY_NUMPAD8, KEY_NUMPAD9,
		KEY_MULTIPLY, KEY_ADD, KEY_SEPARATOR, KEY_SUBTRACT, KEY_DECIMAL, KEY_DIVIDE,
		// Navigation (10)
		KEY_PAGE_UP, KEY_PAGE_DOWN, KEY_END, KEY_HOME, KEY_LEFT, KEY_UP, KEY_RIGHT, KEY_DOWN, KEY_INSERT, KEY_DELETE,
		// Modifiers & Control (10)
		KEY_BACKSPACE, KEY_TAB, KEY_ENTER, KEY_SHIFT, KEY_CTRL, KEY_ALT, KEY_PAUSE, KEY_CAPSLOCK, KEY_ESC, KEY_SPACE,
		// Punctuation (11)
		KEY_SEMICOLON, KEY_EQUAL, KEY_COMMA, KEY_MINUS, KEY_PERIOD, KEY_SLASH, KEY_BACKTICK, KEY_LBRACKET, KEY_BACKSLASH, KEY_RBRACKET, KEY_QUOTE,
		// System & Locks (5)
		KEY_LWIN, KEY_RWIN, KEY_APPS, KEY_NUMLOCK, KEY_SCROLLLOCK,
	}
}

// AllKeysAndMouse returns a fixed-length array containing all 105 defined standard keys and mouse buttons.
func AllKeysAndMouse() []KeyCode {
	return []KeyCode{
		// Mouse Buttons (5)
		KEY_LBUTTON, KEY_RBUTTON, KEY_MBUTTON, KEY_XBUTTON1, KEY_XBUTTON2,
		// Letters (26)
		KEY_A, KEY_B, KEY_C, KEY_D, KEY_E, KEY_F, KEY_G, KEY_H, KEY_I, KEY_J,
		KEY_K, KEY_L, KEY_M, KEY_N, KEY_O, KEY_P, KEY_Q, KEY_R, KEY_S, KEY_T,
		KEY_U, KEY_V, KEY_W, KEY_X, KEY_Y, KEY_Z,
		// Numbers (10)
		KEY_0, KEY_1, KEY_2, KEY_3, KEY_4, KEY_5, KEY_6, KEY_7, KEY_8, KEY_9,
		// Function (12)
		KEY_F1, KEY_F2, KEY_F3, KEY_F4, KEY_F5, KEY_F6, KEY_F7, KEY_F8, KEY_F9, KEY_F10, KEY_F11, KEY_F12,
		// Numpad (16)
		KEY_NUMPAD0, KEY_NUMPAD1, KEY_NUMPAD2, KEY_NUMPAD3, KEY_NUMPAD4,
		KEY_NUMPAD5, KEY_NUMPAD6, KEY_NUMPAD7, KEY_NUMPAD8, KEY_NUMPAD9,
		KEY_MULTIPLY, KEY_ADD, KEY_SEPARATOR, KEY_SUBTRACT, KEY_DECIMAL, KEY_DIVIDE,
		// Navigation (10)
		KEY_PAGE_UP, KEY_PAGE_DOWN, KEY_END, KEY_HOME, KEY_LEFT, KEY_UP, KEY_RIGHT, KEY_DOWN, KEY_INSERT, KEY_DELETE,
		// Modifiers & Control (10)
		KEY_BACKSPACE, KEY_TAB, KEY_ENTER, KEY_SHIFT, KEY_CTRL, KEY_ALT, KEY_PAUSE, KEY_CAPSLOCK, KEY_ESC, KEY_SPACE,
		// Punctuation (11)
		KEY_SEMICOLON, KEY_EQUAL, KEY_COMMA, KEY_MINUS, KEY_PERIOD, KEY_SLASH, KEY_BACKTICK, KEY_LBRACKET, KEY_BACKSLASH, KEY_RBRACKET, KEY_QUOTE,
		// System & Locks (5)
		KEY_LWIN, KEY_RWIN, KEY_APPS, KEY_NUMLOCK, KEY_SCROLLLOCK,
	}
}
