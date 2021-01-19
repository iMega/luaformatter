package formatter

import (
	"io"
)

type newLineWriter interface {
	NewLine() error
}

type cursorPositioner interface {
	Cursor() cursorPosition
}

type writer struct {
	Writer io.Writer
	CurPos cursorPosition
}

type cursorPosition struct {
	Col uint64
	Ln  uint64
}

func (w *writer) Write(p []byte) (int, error) {
	b, err := w.Writer.Write(p)
	// fmt.Printf("%s\n", string(p))
	w.CurPos.Col += uint64(b)

	return b, err
}

func (w *writer) NewLine() error {
	_, err := w.Write([]byte(newLineSymbol))
	w.CurPos.Col = 0
	w.CurPos.Ln++

	return err
}

func (w *writer) Cursor() cursorPosition {
	return w.CurPos
}
