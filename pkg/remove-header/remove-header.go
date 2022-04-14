package remove_header

import (
	"regexp"
	"tuan/pkg/tools"
)

type RemoveHeader struct {
	regex *regexp.Regexp
}

func (t RemoveHeader) Process(ctx tools.ProcessContext) (tools.ProcessContext, error) {
	lineCtx, ok := ctx.(tools.LinesContext)
	if !ok {
		return nil, tools.ErrInvalidContext("LinesContext", ctx.Kind())
	}

	var lines []string
	for _, line := range lineCtx.GetLines() {
		if !t.regex.MatchString(line) {
			lines = append(lines, line)
		}
	}
	return tools.NewLinesContext(lines), nil
}

func init() {
	tools.Register("remove-header", &RemoveHeader{
		regex: regexp.MustCompile("^#.*$"),
	})
}
