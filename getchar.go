// Original code by paulrademacher released under the Unlicense, Version not specified.
// While not required under the Unlicense if the code is used please leave reference
// to the original code repo. https://github.com/paulrademacher/climenu.git getchar.go.

package utilities

import "github.com/pkg/term"

// Returns either an ascii code, or (if input is an arrow) a Javascript key code.
// TODO: move to X11 keysymdef.h.
func getChar() (ascii int, keyCode int, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)

	bytes := make([]byte, 7)

	byteLength, err := t.Read(bytes)
	if err != nil {
		return
	}
	if byteLength == 3 && bytes[0] == 27 && bytes[1] == 91 {
		// Three-character control sequence, beginning with "ESC-[".

		// Since there are no ASCII codes for arrow keys, we use
		// Javascript key codes. TODO: switch to hexadecimal keycodes.
		if bytes[2] == 65 {
			// Up
			keyCode = 38
		} else if bytes[2] == 66 {
			// Down
			keyCode = 40
		} else if bytes[2] == 67 {
			// Right
			keyCode = 39
		} else if bytes[2] == 68 {
			// Left
			keyCode = 37
		}
	} else if byteLength == 1 {
		ascii = int(bytes[0])
	} else {
		// Two characters read??
	}
	t.Restore()
	t.Close()
	return
}
