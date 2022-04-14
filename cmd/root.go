package cmd

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"tuan/pkg/parser"
)

var rootOptions struct {
	SourceFile string `validate:"required,file"`
	Pipeline   []string
	Debug      bool
}

var validate = validator.New()

func errorReturnf(cmd *cobra.Command, format string, args ...any) error {
	if len(format) > 0 && format[len(format)-1] != '\n' {
		format += "\n"
	}

	fmt.Fprintf(cmd.OutOrStderr(), format, args...)
	return cmd.Usage()
}

func errorReturn(cmd *cobra.Command, err error) error {
	fmt.Fprintf(cmd.OutOrStderr(), "%v\n", err)
	return cmd.Usage()
}

var rootCmd = cobra.Command{
	Use: "tuan [flags] <source file>",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errorReturnf(cmd, "miss required argument <source file>")
		}
		rootOptions.SourceFile = args[0]
		if err := validate.Struct(&rootOptions); err != nil {
			return errorReturn(cmd, err)
		}

		p := parser.NewParser(rootOptions.SourceFile,
			parser.WithDebug(rootOptions.Debug),
			parser.WithPipeline(rootOptions.Pipeline),
		)

		err := p.Parse()
		if err != nil {
			return errorReturn(cmd, err)
		}
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolVarP(&rootOptions.Debug, "debug", "d", false, "enable/disable debug mode")
	rootCmd.Flags().StringSliceVarP(&rootOptions.Pipeline, "pipeline", "", parser.DefaultPipeline, "process pipeline")
}
