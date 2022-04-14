package tools

import (
	"fmt"
	"io"
)

type StructuredLineContext interface {
	GetLines() []StructuredLine
	SetLines(lines []StructuredLine)

	SetShelf(shelf Shelf)
	GetShelf() Shelf

	ProcessContext
}

type structuredLinesContext struct {
	Lines []StructuredLine
	Shelf Shelf // 货架
}

type StructuredLine struct {
	Raw      string
	Sequence int
	Rest     string
	Location string
	Goods    map[string]int // 存储货品及购买数量

}

func (sl StructuredLine) Dump(writer io.Writer, prefix string) {
	fmt.Fprintf(writer, "%s", prefix)
	fmt.Fprintf(writer, "%s", sl.Rest)
}

func NewStructuredLinesContext(lines []StructuredLine) StructuredLineContext {
	return &structuredLinesContext{
		Lines: lines,
	}
}

func (ctx structuredLinesContext) GetLines() []StructuredLine {
	return ctx.Lines
}

func (ctx *structuredLinesContext) SetLines(lines []StructuredLine) {
	ctx.Lines = lines
}

func (ctx structuredLinesContext) GetShelf() Shelf {
	return ctx.Shelf
}

func (ctx *structuredLinesContext) SetShelf(shelf Shelf) {
	ctx.Shelf = shelf
}

func (structuredLinesContext) Kind() string {
	return "SequencedLinesContext"
}

func (ctx structuredLinesContext) Dump(writer io.Writer, prefix string) {
	fmt.Fprint(writer, prefix)
	for _, l := range ctx.Lines {
		fmt.Fprintf(writer, "%03d. ", l.Sequence)
		l.Dump(writer, "")
		fmt.Fprintf(writer, "\n")
	}
}
