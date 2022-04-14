package tools

import "strings"

func ToArabic(input string) string {
	var mapping = []struct {
		Chinese string
		Arabic  string
	}{
		{Chinese: "零", Arabic: "0"},
		{Chinese: "一", Arabic: "1"},
		{Chinese: "二", Arabic: "2"},
		{Chinese: "三", Arabic: "3"},
		{Chinese: "四", Arabic: "4"},
		{Chinese: "五", Arabic: "5"},
		{Chinese: "六", Arabic: "6"},
		{Chinese: "七", Arabic: "7"},
		{Chinese: "八", Arabic: "8"},
		{Chinese: "九", Arabic: "9"},
		{Chinese: "十", Arabic: "10"},
	}

	ret := input
	for _, m := range mapping {
		ret = strings.ReplaceAll(ret, m.Chinese, m.Arabic)
	}
	return ret
}
