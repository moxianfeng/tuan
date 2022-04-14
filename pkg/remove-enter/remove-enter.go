package remove_enter

import (
	"regexp"
	"tuan/pkg/tools"
)

type RemoveEnter struct {
	regex *regexp.Regexp
}

func (t RemoveEnter) Process(ctx tools.ProcessContext) (tools.ProcessContext, error) {
	lineCtx, ok := ctx.(tools.LinesContext)
	if !ok {
		return nil, tools.ErrInvalidContext("LinesContext", ctx.Kind())
	}

	var pendingLine string
	var lines []string
	for _, line := range lineCtx.GetLines() {
		if pendingLine == "" {
			pendingLine = line
			continue
		}

		if t.regex.MatchString(line) {
			lines = append(lines, pendingLine)
			pendingLine = line
		} else {
			pendingLine += " " + line
		}
	}

	lines = append(lines, pendingLine)
	return tools.NewLinesContext(lines), nil
}

func init() {
	tools.Register("remove-enter", &RemoveEnter{
		regex: regexp.MustCompile("^[0-9]+.*$"),
	})
}
