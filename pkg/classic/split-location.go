package classic

import (
	"github.com/spf13/cast"
	"regexp"
	"tuan/pkg/tools"
)

type Classic struct {
	regex *regexp.Regexp
}

func (t Classic) Process(ctx tools.ProcessContext) (tools.ProcessContext, error) {
	structuredCtx, ok := ctx.(tools.StructuredLineContext)
	if !ok {
		return nil, tools.ErrInvalidContext("LinesContext", ctx.Kind())
	}

	shelf := tools.NewShelf()
	var lines []tools.StructuredLine
	for _, line := range structuredCtx.GetLines() {

		ret := t.regex.FindAllStringSubmatch(tools.ToArabic(line.Rest), -1)
		if len(ret) == 0 {
			return nil, tools.ErrInvalidContent("Classic", line.Raw)
		}

		goodsList := map[string]int{}

		for _, item := range ret {
			goods := item[1]
			count := cast.ToInt(item[2])
			unit := item[3]

			goodsList[goods] = count

			shelf.AddGoods(*tools.NewGoods(goods, unit))
		}

		lines = append(lines, tools.StructuredLine{
			Raw:      line.Raw,
			Sequence: line.Sequence,
			Location: line.Location,
			Goods:    goodsList,
		})
	}

	returnCtx := tools.NewStructuredLinesContext(lines)
	returnCtx.SetShelf(*shelf)
	return returnCtx, nil
}

func init() {
	tools.Register("classic", &Classic{
		regex: regexp.MustCompile("(?:([^0-9]+)([0-9]+)([^0-9 \t,，:：;；.。、])[ \t,，:：;；.。、]?)"),
	})
}
