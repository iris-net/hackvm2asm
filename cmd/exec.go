package cmd

import (
	codewriter "github.com/iris-net/hackvm2asm/code_writer"
	"github.com/spf13/cobra"
)

var (
	out string
	in  string
)

func init() {
	execCmd.Flags().StringVar(&in, "in", "i", "where to input an vm file")
	execCmd.Flags().StringVar(&out, "out", "o", "where to output a assembly file")

	rootCmd.AddCommand(execCmd)
}

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "analyze vm code and translate it into assembly code",
	RunE: func(cmd *cobra.Command, args []string) error {
		cw := codewriter.NewCodeWriter()
		err := cw.Execute(in, out)
		if err != nil {
			return err
		}

		return nil
	},
}
