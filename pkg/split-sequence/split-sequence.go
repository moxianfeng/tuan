package split_sequence

import (
	"github.com/spf13/cast"
	"regexp"
	"tuan/pkg/tools"
)

type SplitSequence struct {
	regex *regexp.Regexp
}

func (t SplitSequence) Process(ctx tools.ProcessContext) (tools.ProcessContext, error) {
	lineCtx, ok := ctx.(tools.LinesContext)
	if !ok {
		return nil, tools.ErrInvalidContext("LinesContext", ctx.Kind())
	}

	var lines []tools.StructuredLine
	for _, line := range lineCtx.GetLines() {
		ret := t.regex.FindAllStringSubmatch(line, -1)

		if len(ret) == 0 {
			return nil, tools.ErrInvalidContent("SplitSequence", line)
		}

		lines = append(lines, tools.StructuredLine{
			Raw:      line,
			Sequence: cast.ToInt(ret[0][1]),
			Rest:     ret[0][2],
		})

	}

	return tools.NewStructuredLinesContext(lines), nil
}

func init() {
	tools.Register("split-sequence", &SplitSequence{
		regex: regexp.MustCompile("^(?:([0-9]+)[.])(.*)$"),
	})
}
