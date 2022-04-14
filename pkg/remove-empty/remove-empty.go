package remove_empty

import (
	"regexp"
	"tuan/pkg/tools"
)

type RemoveEmpty struct {
	regex *regexp.Regexp
}

func (t RemoveEmpty) Process(ctx tools.ProcessContext) (tools.ProcessContext, error) {
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
	tools.Register("remove-empty", &RemoveEmpty{
		regex: regexp.MustCompile("^[ \t]*$"),
	})
}
