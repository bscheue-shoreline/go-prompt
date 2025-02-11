package prompt

// GoLineEnd Go to the End of the line
func GoLineEnd(buf *Buffer, hist *History) {
	x := []rune(buf.Document().TextAfterCursor())
	buf.CursorRight(len(x))
}

// GoLineBeginning Go to the beginning of the line
func GoLineBeginning(buf *Buffer, hist *History) {
	x := []rune(buf.Document().TextBeforeCursor())
	buf.CursorLeft(len(x))
}

// DeleteChar Delete character under the cursor
func DeleteChar(buf *Buffer, hist *History) {
	buf.Delete(1)
}

// DeleteWord Delete word before the cursor
func DeleteWord(buf *Buffer, hist *History) {
	buf.DeleteBeforeCursor(len([]rune(buf.Document().TextBeforeCursor())) - buf.Document().FindStartOfPreviousWordWithSpace())
}

// DeleteBeforeChar Go to Backspace
func DeleteBeforeChar(buf *Buffer, hist *History) {
	buf.DeleteBeforeCursor(1)
}

// GoRightChar Forward one character
func GoRightChar(buf *Buffer, hist *History) {
	buf.CursorRight(1)
}

// GoRightChar Forward one character
func GoRightCharOrComplete(buf *Buffer, hist *History) {
	if fishCompleteOpt:= hist.GetFishOpt(buf.Text()); fishCompleteOpt != "" {
		buf.InsertText(fishCompleteOpt, false, true)
	}
	buf.CursorRight(1)
}


// GoLeftChar Backward one character
func GoLeftChar(buf *Buffer, hist *History) {
	buf.CursorLeft(1)
}

// GoRightWord Forward one word
func GoRightWord(buf *Buffer, hist *History) {
	buf.CursorRight(buf.Document().FindEndOfCurrentWordWithSpace())
}

// GoLeftWord Backward one word
func GoLeftWord(buf *Buffer, hist *History) {
	buf.CursorLeft(len([]rune(buf.Document().TextBeforeCursor())) - buf.Document().FindStartOfPreviousWordWithSpace())
}
