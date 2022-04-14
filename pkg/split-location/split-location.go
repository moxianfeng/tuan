package split_location

import (
	"fmt"
	"regexp"
	"strings"
	"tuan/pkg/tools"
)

type SplitLocation struct {
	regex *regexp.Regexp
}

func (t SplitLocation) Process(ctx tools.ProcessContext) (tools.ProcessContext, error) {
	structuredCtx, ok := ctx.(tools.StructuredLineContext)
	if !ok {
		return nil, tools.ErrInvalidContext("LinesContext", ctx.Kind())
	}

	var lines []tools.StructuredLine
	for _, line := range structuredCtx.GetLines() {
		ret := t.regex.FindAllStringSubmatch(line.Rest, 1)
		if len(ret) == 0 {
			return nil, tools.ErrInvalidContent("SplitLocation", line.Raw)
		}

		if len(ret[0][1]) > 20 {
			fmt.Println(len(ret[0][1]))
			err := tools.Confirm("位置", ret[0][1])
			if err != nil {
				return nil, tools.ErrInvalidContent("SplitLocation", line.Raw)
			}
		}
		rest := strings.TrimLeft(line.Rest, ret[0][0])

		lines = append(lines, tools.StructuredLine{
			Raw:      line.Raw,
			Sequence: line.Sequence,
			Location: ret[0][1],
			Rest:     rest,
		})
	}

	return tools.NewStructuredLinesContext(lines), nil
}

func init() {
	tools.Register("split-location", &SplitLocation{
		regex: regexp.MustCompile("^[ \t*]([^0-9]*[0-9]+[^ \t:：,，。]*)[ \t:：,，。]+"),
	})
}
