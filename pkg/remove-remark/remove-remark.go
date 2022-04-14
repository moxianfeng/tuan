package remove_remark

import (
	"regexp"
	"tuan/pkg/tools"
)

type RemoveRemark struct {
	regex *regexp.Regexp
}

func (t RemoveRemark) Process(ctx tools.ProcessContext) (tools.ProcessContext, error) {
	structuredCtx, ok := ctx.(tools.StructuredLineContext)
	if !ok {
		return nil, tools.ErrInvalidContext("LinesContext", ctx.Kind())
	}

	var lines []tools.StructuredLine
	for _, line := range structuredCtx.GetLines() {
		rest := t.regex.ReplaceAllString(line.Rest, "")

		lines = append(lines, tools.StructuredLine{
			Raw:      line.Raw,
			Sequence: line.Sequence,
			Location: line.Location,
			Rest:     rest,
		})
	}

	return tools.NewStructuredLinesContext(lines), nil
}

func init() {
	tools.Register("remove-remark", &RemoveRemark{
		regex: regexp.MustCompile("[(][^)]*[)]"),
	})
}
