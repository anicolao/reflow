package ansi

import (
	"bytes"

	"github.com/mattn/go-runewidth"
)

// Buffer is a buffer aware of ANSI escape sequences.
type Buffer struct {
	bytes.Buffer
	state AnsiState
}

// PrintableRuneWidth returns the cell width of all printable runes in the
// buffer.
func (w Buffer) PrintableRuneWidth() int {
	return CalculateWidth(w.state, w.String());
}

func PrintableRuneWidth(s string) int {
	var state AnsiState
	return CalculateWidth(state, s);
}

// Calling CalculateWidth is to be avoided because it won't have 
// state sufficient to deal with escape codes with lots of data
// between the escape and the terminator. It was necessary only
// for truncate.go
func CalculateWidth(state AnsiState, s string) int {
	var n int
	var ansi bool

	for _, c := range s {
		if state.IsMarker(c) {
			// ANSI escape sequence
			ansi = true
		} else if ansi {
			if state.IsTerminator(c) {
				// ANSI sequence terminated
				ansi = false
			}
		} else {
			n += runewidth.RuneWidth(c)
		}
	}

	return n
}
