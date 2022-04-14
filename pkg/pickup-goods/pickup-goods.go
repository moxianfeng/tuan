package pickup_goods

import (
	"regexp"
	"tuan/pkg/tools"
)

type PickupGoods struct {
	regex         *regexp.Regexp
	regexSplitter *regexp.Regexp
}

func (t PickupGoods) Process(ctx tools.ProcessContext) (tools.ProcessContext, error) {
	structuredCtx, ok := ctx.(tools.StructuredLineContext)
	if !ok {
		return nil, tools.ErrInvalidContext("LinesContext", ctx.Kind())
	}

	shelf := tools.NewShelf()
	var lines []tools.StructuredLine
	for _, line := range structuredCtx.GetLines() {
		fields := t.regexSplitter.Split(line.Rest, -1)
		if len(fields) == 0 {
			return nil, tools.ErrInvalidContent("PickupGoods", line.Raw)
		}
		for _, f := range fields {
			if len(f) == 0 {
				continue
			}
			ret := t.regex.FindAllStringSubmatch(f, -1)
			if len(ret) == 0 {
				return nil, tools.ErrInvalidContent("PickupGoods", line.Raw)
			}

			for _, item := range ret {
				goods := item[1]
				unit := item[3]

				shelf.AddGoods(*tools.NewGoods(goods, unit))
			}
		}

		//lines = append(lines, tools.StructuredLine{
		//	Raw:      line.Raw,
		//	Sequence: line.Sequence,
		//	Location: line.Location,
		//})
	}

	returnCtx := tools.NewStructuredLinesContext(lines)
	returnCtx.SetShelf(*shelf)
	return returnCtx, nil
}

func init() {
	tools.Register("pickup-goods", &PickupGoods{
		regexSplitter: regexp.MustCompile(" \t,，:：;；.。、"),
		regex:         regexp.MustCompile("(?:(\\pL+)([0-9一二两俩三四五六七八九十]+)([^0-9一二两俩三四五六七八九十 \t,，:：;；.。、])[ \t,，:：;；.。、]?)"),
	})
}
