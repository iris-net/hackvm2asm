package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vm2asm",
	Short: "A translation tool that translates the vm into the assembly for HACK platform",
}

func Execute() error {
	rootCmd.SetOut(os.Stdout)

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
