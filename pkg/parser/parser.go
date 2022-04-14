package parser

import (
	"os"
	"strings"
	"tuan/pkg/tools"

	_ "tuan/pkg/classic"
	_ "tuan/pkg/pickup-goods"
	_ "tuan/pkg/remove-empty"
	_ "tuan/pkg/remove-enter"
	_ "tuan/pkg/remove-header"
	_ "tuan/pkg/remove-remark"
	_ "tuan/pkg/split-location"
	_ "tuan/pkg/split-sequence"
)

var DefaultPipeline = []string{
	"remove-empty",
	"remove-header",
	"remove-enter",
	"split-sequence",
	"split-location",
	"remove-remark",
	"pickup-goods",
	"classic",
}

type Parser struct {
	SourceFile string
	Lines      []string
	Pipeline   []string
	Debug      bool
}

type Option func(p *Parser)

func WithPipeline(pipeline []string) Option {
	return func(p *Parser) {
		p.Pipeline = pipeline
	}
}

func WithDebug(debug bool) Option {
	return func(p *Parser) {
		p.Debug = debug
	}
}

func NewParser(filename string, opts ...Option) *Parser {
	ret := &Parser{
		SourceFile: filename,
		Pipeline:   DefaultPipeline,
	}
	for _, o := range opts {
		o(ret)
	}
	return ret
}

func (p *Parser) ReadAll() error {
	data, err := os.ReadFile(p.SourceFile)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	p.Lines = lines
	return nil
}

func (p *Parser) Parse() error {
	if err := p.ReadAll(); err != nil {
		return err
	}

	var currentContext tools.ProcessContext = tools.NewLinesContext(p.Lines)
	if p.Debug {
		currentContext.Dump(os.Stderr, "init context:\n")
	}

	for _, toolName := range p.Pipeline {
		tool, err := tools.Get(toolName)
		if err != nil {
			return err
		}
		nextContext, err := tool.Process(currentContext)
		if err != nil {
			return err
		}
		if p.Debug {
			nextContext.Dump(os.Stderr, "output context:\n")
		}
		currentContext = nextContext
	}
	return nil
}
