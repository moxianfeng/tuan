package tools

import "fmt"

func ErrInvalidContext(expectContext, gotContext string) error {
	return fmt.Errorf("invalid context, expect is %v, got %v", expectContext, gotContext)
}

func ErrInvalidContent(plugin string, content string) error {
	return fmt.Errorf("[%s] invalid line content `%s`", plugin, content)
}
