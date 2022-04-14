package tools

import (
	"fmt"
	"io"
)

type LinesContext interface {
	GetLines() []string
	SetLines(lines []string)

	ProcessContext
}

type linesContext struct {
	Lines []string
}

func NewLinesContext(lines []string) LinesContext {
	return &linesContext{
		Lines: lines,
	}
}

func (ctx linesContext) GetLines() []string {
	return ctx.Lines
}

func (ctx *linesContext) SetLines(lines []string) {
	ctx.Lines = lines
}

func (linesContext) Kind() string {
	return "LinesContext"
}

func (ctx linesContext) Dump(writer io.Writer, prefix string) {
	fmt.Fprint(writer, prefix)
	for _, l := range ctx.Lines {
		fmt.Fprintln(writer, l)
	}
}
