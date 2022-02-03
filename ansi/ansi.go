package ansi

const Marker = '\x1B'

type AnsiState struct {
	isAnsi bool
	oneSemi bool
	urlChars bool
}

func (a *AnsiState) IsMarker(c rune) bool {
	return c == Marker
}

func (a *AnsiState) IsTerminator(c rune) bool {
	if (a.urlChars) {
		a.urlChars = (c != 0x5c)
		return !a.urlChars
	} else if (c == 0x3b) {
		if (!a.oneSemi) {
			a.oneSemi = true
		} else {
			a.urlChars = true
			a.oneSemi = false
		}
	} else {
		a.oneSemi = false
	}

	return (c >= 0x40 && c <= 0x5a) || (c >= 0x61 && c <= 0x7a)
}
