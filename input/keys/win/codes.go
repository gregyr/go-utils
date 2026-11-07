package winkeys

type KeyCode uintptr

func (c KeyCode) value() uintptr {
	return uintptr(c)
}

const (
	// Mouse Buttons
	KEY_LBUTTON  KeyCode = 1
	KEY_RBUTTON  KeyCode = 2
	KEY_MBUTTON  KeyCode = 4
	KEY_XBUTTON1 KeyCode = 5
	KEY_XBUTTON2 KeyCode = 6

	// Modifier & Control Keys
	KEY_BACKSPACE KeyCode = 8
	KEY_TAB       KeyCode = 9
	KEY_ENTER     KeyCode = 13
	KEY_SHIFT     KeyCode = 16
	KEY_CTRL      KeyCode = 17
	KEY_ALT       KeyCode = 18
	KEY_PAUSE     KeyCode = 19
	KEY_CAPSLOCK  KeyCode = 20
	KEY_ESC       KeyCode = 27
	KEY_SPACE     KeyCode = 32

	// Navigation Keys
	KEY_PAGE_UP   KeyCode = 33
	KEY_PAGE_DOWN KeyCode = 34
	KEY_END       KeyCode = 35
	KEY_HOME      KeyCode = 36
	KEY_LEFT      KeyCode = 37
	KEY_UP        KeyCode = 38
	KEY_RIGHT     KeyCode = 39
	KEY_DOWN      KeyCode = 40
	KEY_INSERT    KeyCode = 45
	KEY_DELETE    KeyCode = 46

	// Number Keys (Top Row)
	KEY_0 KeyCode = 48
	KEY_1 KeyCode = 49
	KEY_2 KeyCode = 50
	KEY_3 KeyCode = 51
	KEY_4 KeyCode = 52
	KEY_5 KeyCode = 53
	KEY_6 KeyCode = 54
	KEY_7 KeyCode = 55
	KEY_8 KeyCode = 56
	KEY_9 KeyCode = 57

	// Letter Keys
	KEY_A KeyCode = 65
	KEY_B KeyCode = 66
	KEY_C KeyCode = 67
	KEY_D KeyCode = 68
	KEY_E KeyCode = 69
	KEY_F KeyCode = 70
	KEY_G KeyCode = 71
	KEY_H KeyCode = 72
	KEY_I KeyCode = 73
	KEY_J KeyCode = 74
	KEY_K KeyCode = 75
	KEY_L KeyCode = 76
	KEY_M KeyCode = 77
	KEY_N KeyCode = 78
	KEY_O KeyCode = 79
	KEY_P KeyCode = 80
	KEY_Q KeyCode = 81
	KEY_R KeyCode = 82
	KEY_S KeyCode = 83
	KEY_T KeyCode = 84
	KEY_U KeyCode = 85
	KEY_V KeyCode = 86
	KEY_W KeyCode = 87
	KEY_X KeyCode = 88
	KEY_Y KeyCode = 89
	KEY_Z KeyCode = 90

	// Windows/OS Keys
	KEY_LWIN KeyCode = 91 // Left Windows key
	KEY_RWIN KeyCode = 92 // Right Windows key
	KEY_APPS KeyCode = 93 // Context Menu key

	// Numpad Keys
	KEY_NUMPAD0   KeyCode = 96
	KEY_NUMPAD1   KeyCode = 97
	KEY_NUMPAD2   KeyCode = 98
	KEY_NUMPAD3   KeyCode = 99
	KEY_NUMPAD4   KeyCode = 100
	KEY_NUMPAD5   KeyCode = 101
	KEY_NUMPAD6   KeyCode = 102
	KEY_NUMPAD7   KeyCode = 103
	KEY_NUMPAD8   KeyCode = 104
	KEY_NUMPAD9   KeyCode = 105
	KEY_MULTIPLY  KeyCode = 106
	KEY_ADD       KeyCode = 107
	KEY_SEPARATOR KeyCode = 108
	KEY_SUBTRACT  KeyCode = 109
	KEY_DECIMAL   KeyCode = 110
	KEY_DIVIDE    KeyCode = 111

	// Function Keys
	KEY_F1  KeyCode = 112
	KEY_F2  KeyCode = 113
	KEY_F3  KeyCode = 114
	KEY_F4  KeyCode = 115
	KEY_F5  KeyCode = 116
	KEY_F6  KeyCode = 117
	KEY_F7  KeyCode = 118
	KEY_F8  KeyCode = 119
	KEY_F9  KeyCode = 120
	KEY_F10 KeyCode = 121
	KEY_F11 KeyCode = 122
	KEY_F12 KeyCode = 123

	// Lock Keys
	KEY_NUMLOCK    KeyCode = 144
	KEY_SCROLLLOCK KeyCode = 145

	// Punctuation / OEM Keys (Standard US Keyboard mapping)
	KEY_SEMICOLON KeyCode = 186 // ; :
	KEY_EQUAL     KeyCode = 187 // = +
	KEY_COMMA     KeyCode = 188 // , <
	KEY_MINUS     KeyCode = 189 // - _
	KEY_PERIOD    KeyCode = 190 // . >
	KEY_SLASH     KeyCode = 191 // / ?
	KEY_BACKTICK  KeyCode = 192 // ` ~
	KEY_LBRACKET  KeyCode = 219 // [ {
	KEY_BACKSLASH KeyCode = 220 // \ |
	KEY_RBRACKET  KeyCode = 221 // ] }
	KEY_QUOTE     KeyCode = 222 // ' "
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

// AllKeys returns a fixed-length array containing all 105 defined standard keys and mouse buttons.
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
