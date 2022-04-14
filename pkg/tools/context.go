package tools

import "io"

type ProcessContext interface {
	Kind() string
	Dump(writer io.Writer, prefix string)
}
